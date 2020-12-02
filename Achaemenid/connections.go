/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"net"
	"sync"
	"time"

	gp "../GP"
	"../authorization"
	etime "../earth-time"
	er "../error"
	"../log"
	"../uuid"
)

// connections store pools of connection to retrieve in many ways!
type connections struct {
	server                 *Server
	mutex                  sync.Mutex
	poolByPeerGPAddr       map[gp.Addr]*Connection
	poolByConnID           map[[32]byte]*Connection
	poolByUserID           map[[32]byte][]*Connection
	poolByBlackList        map[[32]byte]*Connection
	GuestConnectionCount   uint64
	GetConnByID            func(connID [32]byte) (conn *Connection)
	GetConnByUserIDThingID func(userID, thingID [32]byte) (conn *Connection)
	SaveConn               func(conn *Connection)
}

func (c *connections) init(s *Server) {
	c.server = s

	if c.poolByPeerGPAddr == nil {
		c.poolByPeerGPAddr = make(map[gp.Addr]*Connection, 16384)
	}
	if c.poolByConnID == nil {
		c.poolByConnID = make(map[[32]byte]*Connection, 16384)
	}
	if c.poolByUserID == nil {
		c.poolByUserID = make(map[[32]byte][]*Connection, 16384)
	}
	if c.poolByBlackList == nil {
		c.poolByBlackList = make(map[[32]byte]*Connection)
	}

	go c.connectionIdleTimeoutSaveAndFree()
}

// MakeNewConnectionByDomainID make new connection by peer domain ID and initialize it!
func (c *connections) MakeNewConnectionByDomainID(domainID [32]byte) (conn *Connection, err *er.Error) {
	// TODO::: Get closest domain GP add
	var domainGPAddr = gp.Addr{}
	conn, err = c.MakeNewConnectionByPeerAdd(domainGPAddr)
	if err != nil {
		return
	}
	conn.UserID = domainID
	conn.UserType = authorization.UserTypeApp
	return
}

// MakeNewConnectionByPeerAdd make new connection by peer GP and initialize it!
func (c *connections) MakeNewConnectionByPeerAdd(gpAddr gp.Addr) (conn *Connection, err *er.Error) {
	var userID, thingID [32]byte
	// TODO::: Get peer publickey & userID & thingID from peer GP router!

	if userID != [32]byte{} {
		conn = c.GetConnByUserIDThingID(userID, thingID)
	}

	// If conn not exist means guest connection.
	if conn == nil {
		conn, err = c.MakeNewGuestConnection()
		if err == nil {
			conn.GPAddr = gpAddr
			// conn.Cipher = crypto.NewGCM(crypto.NewAES256([32]byte{}))
		}
	}
	return
}

// MakeNewGuestConnection make new connection and register on given stream due to it is first attempt connect to server!
func (c *connections) MakeNewGuestConnection() (conn *Connection, err *er.Error) {
	if c.server.Manifest.TechnicalInfo.GuestMaxConnections == 0 {
		return nil, ErrGuestConnectionNotAllow
	} else if c.server.Manifest.TechnicalInfo.GuestMaxConnections > 0 && c.server.Connections.GuestConnectionCount > c.server.Manifest.TechnicalInfo.GuestMaxConnections {
		return nil, ErrGuestConnectionMaxReached
	}

	conn = &Connection{
		Server:   c.server,
		ID:       uuid.Random32Byte(),
		State:    StateNew,
		UserType: authorization.UserTypeGuest,
	}
	conn.AccessControl.GiveFullAccess()
	conn.StreamPool.Init()
	return
}

// MakeNewIPConnection make new connection and return to caller.
// TODO::: due to IPv4&&IPv6 support we need this! Remove it when remove those support.
func (c *connections) MakeNewIPConnection(ipAddr [16]byte) (conn *Connection, err *er.Error) {
	// TODO::: check new connection can create

	conn = &Connection{
		Server: c.server,
		IPAddr: ipAddr,
	}
	conn.StreamPool.Init()

	c.RegisterConnection(conn)
	return
}

// MakeNewUDPConnection make new connection and return to caller.
// TODO::: due to IPv4&&IPv6 support we need this! Remove it when remove those support.
func (c *connections) MakeNewUDPConnection(udpAddr *net.UDPAddr) (conn *Connection, err *er.Error) {
	// TODO::: check new connection can create

	conn = &Connection{
		Server: c.server,
	}
	conn.StreamPool.Init()

	c.RegisterConnection(conn)
	return
}

// GetConnectionByPeerGPAdd get a connection by peer GP from connections pool!!
func (c *connections) GetConnectionByPeerGPAdd(gpAddr gp.Addr) *Connection {
	return c.poolByPeerGPAddr[gpAddr]
}

// GetConnectionByID get a connection by its ID from connections pool!!
func (c *connections) GetConnectionByID(connID [32]byte) (conn *Connection) {
	conn = c.poolByConnID[connID]
	if conn == nil {
		// get from platfrom resources
		conn = c.GetConnByID(connID)
	}
	return
}

// GetConnectionsByUserID get the connections by peer userID||domainID from connections pool!!
func (c *connections) GetConnectionsByUserID(userID [32]byte) (conn []*Connection) {
	conn = c.poolByUserID[userID]
	// TODO::: check if connection is in not ready status
	return
}

// RegisterConnection register new connection in server connection pool!!
func (c *connections) RegisterConnection(conn *Connection) {
	c.mutex.Lock()
	c.poolByConnID[conn.ID] = conn
	if conn.GPAddr != gp.AddrNil {
		c.poolByPeerGPAddr[conn.GPAddr] = conn
	}
	if conn.UserID != [32]byte{} {
		c.poolByUserID[conn.UserID] = append(c.poolByUserID[conn.UserID], conn)
	}
	if conn.UserType == authorization.UserTypeGuest {
		c.GuestConnectionCount++
	}
	c.mutex.Unlock()
}

// CloseConnection un-register exiting connection in server connection pool!!
func (c *connections) CloseConnection(conn *Connection) {
	c.mutex.Lock()
	if conn.UserType == authorization.UserTypeGuest {
		c.GuestConnectionCount--
	}

	// TODO::: Is it worth to don't delete connection just reset it and send it to pool of unused connection due to GC!
	delete(c.poolByConnID, conn.ID)
	delete(c.poolByPeerGPAddr, conn.GPAddr)
	// delete(c.poolByUserID, conn.UserID)

	// Let unfinished stream handled!!
	c.mutex.Unlock()
}

// RevokeConnection un-register exiting connection in server connection pool without given any time to finish any stream!!
func (c *connections) RevokeConnection(conn *Connection) {
	c.mutex.Lock()
	if conn.UserType == authorization.UserTypeGuest {
		c.GuestConnectionCount--
	}

	// Remove all unfinished stream first!!

	delete(c.poolByConnID, conn.ID)
	delete(c.poolByPeerGPAddr, conn.GPAddr)
	// delete(c.poolByUserID, conn.UserID)
	c.mutex.Unlock()
}

func (c *connections) shutdown() {
	log.Info("Achaemenid - ShutDown - connections saving proccess begin...")
	log.Info("Achaemenid - ShutDown - Number of active connections:", len(c.poolByConnID))
	for _, conn := range c.poolByConnID {
		go c.SaveConn(conn)
	}
	log.Info("Achaemenid - ShutDown - connections saving proccess end now!")
}

func (c *connections) connectionIdleTimeoutSaveAndFree() {
	var timer = time.NewTimer(c.server.Manifest.TechnicalInfo.ConnectionIdleTimeout.ConvertToTimeDuration())
	for {
		select {
		case <-timer.C:
			log.Info("Achaemenid - Cron - Idle connections timeout save and free proccess begin...")
			log.Info("Achaemenid - Cron - Number of active connections:", len(c.poolByConnID))
			// It is conccurent proccess and not take more than one or two second! so set timeNow here
			var timeNow = etime.Now()

			c.mutex.Lock()
			for _, conn := range c.poolByConnID {
				if !conn.LastUsage.AddDuration(c.server.Manifest.TechnicalInfo.ConnectionIdleTimeout).Pass(timeNow) {
					continue
				}

				go c.SaveConn(conn)

				if conn.UserType == authorization.UserTypeGuest {
					c.GuestConnectionCount--
				}
				// Remove all unfinished stream first!!
				delete(c.poolByConnID, conn.ID)
				delete(c.poolByPeerGPAddr, conn.GPAddr)
			}
			c.mutex.Unlock()

			timer.Reset(c.server.Manifest.TechnicalInfo.ConnectionIdleTimeout.ConvertToTimeDuration())
			log.Info("Achaemenid - Cron - Number of active connections after:", len(c.poolByConnID))
			log.Info("Achaemenid - Cron - Idle connections timeout save and free proccess end now!")
		}
	}
}
