/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"net"
	"sync"

	"../crypto"
	"../uuid"
)

// connections store pools of connection to retrieve in many ways!
type connections struct {
	server               *Server
	mutex                sync.Mutex
	poolByPeerGPAdd      map[[14]byte]*Connection
	poolByConnID         map[[16]byte]*Connection
	poolByUserID         map[[16]byte][]*Connection
	poolByDomainID       map[[16]byte][]*Connection
	poolByBlackList      map[[16]byte]*Connection
	GuestConnectionCount uint64
	GetConnByID          func(connID [16]byte) (conn *Connection)
	GetConnByUserID      func(userID, appID, thingID [16]byte) (conn *Connection)
	SaveConn             func(conn *Connection) (err error)
}

func (c *connections) init(s *Server) {
	c.server = s

	if c.poolByPeerGPAdd == nil {
		c.poolByPeerGPAdd = make(map[[14]byte]*Connection)
	}
	if c.poolByConnID == nil {
		c.poolByConnID = make(map[[16]byte]*Connection)
	}
	if c.poolByUserID == nil {
		c.poolByUserID = make(map[[16]byte][]*Connection)
	}
	if c.poolByDomainID == nil {
		c.poolByDomainID = make(map[[16]byte][]*Connection)
	}
	if c.poolByBlackList == nil {
		c.poolByBlackList = make(map[[16]byte]*Connection)
	}
}

// MakeNewConnectionByDomainID use to
func (c *connections) MakeNewConnectionByDomainID(domainID [16]byte) (conn *Connection, err error) {
	// TODO::: Get closest domain GP add
	var domainGPAddr = [14]byte{}
	conn, err = c.MakeNewConnectionByPeerAdd(domainGPAddr)
	conn.DomainID = domainID
	return
}

// MakeNewConnectionByPeerAdd make new connection by peer GP and initialize it!
func (c *connections) MakeNewConnectionByPeerAdd(gpAddr [14]byte) (conn *Connection, err error) {
	var userID, appID, thingID [16]byte
	// TODO::: Get peer publickey & userID & appID & thingID from peer GP router!

	if userID != [16]byte{} {
		conn = c.GetConnByUserID(userID, appID, thingID)
		conn.StreamPool.Init()
	}

	// If conn not exist means guest connection.
	if conn == nil {
		conn, err = c.MakeNewGuestConnection()
		if err == nil {
			conn.GPAddr = gpAddr
			conn.Cipher = crypto.NewGCM(crypto.NewAES256([32]byte{}))
		}
	}
	return
}

// MakeNewGuestConnection make new connection and register on given stream due to it is first attempt connect to server!
func (c *connections) MakeNewGuestConnection() (conn *Connection, err error) {
	if c.server.Manifest.TechnicalInfo.GuestMaxConnections == 0 {
		return nil, ErrAchaemenidGuestConnectionNotAllow
	} else if c.server.Manifest.TechnicalInfo.GuestMaxConnections > 0 && c.server.Connections.GuestConnectionCount > c.server.Manifest.TechnicalInfo.GuestMaxConnections {
		return nil, ErrAchaemenidGuestConnectionMaxReached
	}

	conn = &Connection{
		Server:   c.server,
		ID:       uuid.NewV4(),
		UserType: 0, // guest
	}
	conn.StreamPool.Init()
	return
}

// MakeNewIPConnection make new connection and return to caller.
// TODO::: due to IPv4&&IPv6 support we need this! Remove it when remove those support.
func (c *connections) MakeNewIPConnection(ipAddr net.IPAddr) (conn *Connection, err error) {
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
func (c *connections) MakeNewUDPConnection(udpAddr *net.UDPAddr) (conn *Connection, err error) {
	// TODO::: check new connection can create

	conn = &Connection{
		Server: c.server,
	}
	conn.StreamPool.Init()

	c.RegisterConnection(conn)
	return
}

// GetConnectionByPeerGPAdd get a connection by peer GP from connections pool!!
func (c *connections) GetConnectionByPeerGPAdd(gpAddr [14]byte) *Connection {
	return c.poolByPeerGPAdd[gpAddr]
}

// GetConnectionByID get a connection by its ID from connections pool!!
func (c *connections) GetConnectionByID(connID [16]byte) (conn *Connection) {
	conn = c.poolByConnID[connID]
	if conn == nil {
		// get from platfrom resources
		conn = c.GetConnByID(connID)
	}
	return
}

// GetConnectionsByDomainID get the connections by peer domain ID from connections pool!!
func (c *connections) GetConnectionsByDomainID(domainID [16]byte) (conn []*Connection) {
	conn = c.poolByDomainID[domainID]
	// TODO::: check if connection is in not ready status
	return
}

// RegisterConnection register new connection in server connection pool!!
func (c *connections) RegisterConnection(conn *Connection) {
	c.mutex.Lock()
	c.poolByConnID[conn.ID] = conn
	if conn.GPAddr != [14]byte{} {
		c.poolByPeerGPAdd[conn.GPAddr] = conn
	}
	if conn.UserID != [16]byte{} {
		c.poolByUserID[conn.UserID] = append(c.poolByUserID[conn.UserID], conn)
	}
	if conn.DomainID == [16]byte{} {
		c.poolByDomainID[conn.DomainID] = append(c.poolByDomainID[conn.DomainID], conn)
	}
	if conn.UserType == 0 {
		c.GuestConnectionCount++
	}
	c.mutex.Unlock()
}

// CloseConnection un-register exiting connection in server connection pool!!
func (c *connections) CloseConnection(conn *Connection) {
	c.mutex.Lock()
	if conn.UserType == 0 {
		c.GuestConnectionCount--
	}

	// TODO::: Don't delete connection just reset it and send it to pool of unused connection due to GC!
	delete(c.poolByPeerGPAdd, conn.GPAddr)
	delete(c.poolByUserID, conn.UserID)

	// Let unfinished stream handled!!
	c.mutex.Unlock()
}

// RevokeConnection un-register exiting connection in server connection pool without given any time to finish any stream!!
func (c *connections) RevokeConnection(conn *Connection) {
	c.mutex.Lock()
	if conn.UserType == 0 {
		c.GuestConnectionCount--
	}

	// Remove all unfinished stream first!!

	delete(c.poolByPeerGPAdd, conn.GPAddr)
	delete(c.poolByUserID, conn.UserID)
	c.mutex.Unlock()
}
