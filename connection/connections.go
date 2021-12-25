/* For license and copyright information please see LEGAL file in repository */

package connection

import (
	"sync"
	"time"

	"../authorization"
	etime "../earth-time"
	"../protocol"
)

// Connections store pools of connection to retrieve in many ways
type Connections struct {
	mutex                      sync.Mutex
	poolByPeerAddr             map[[16]byte]protocol.Connection
	poolByUserIDDelegateUserID map[[32]byte]protocol.Connection
	poolByUserID               map[[16]byte][]protocol.Connection
	poolByDomain               map[string]protocol.Connection
	poolByBlackList            map[[16]byte]protocol.Connection // key is PeerAddr
	guestConnectionCount       uint64
}

func (c *Connections) Init() {
	c.poolByPeerAddr = make(map[[16]byte]protocol.Connection, 16384)
	c.poolByUserIDDelegateUserID = make(map[[32]byte]protocol.Connection, 16384)
	c.poolByUserID = make(map[[16]byte][]protocol.Connection, 16384)
	c.poolByDomain = make(map[string]protocol.Connection, 16384)
	c.poolByBlackList = make(map[[32]byte]protocol.Connection, 256)

	go c.connectionIdleTimeoutSaveAndFree()
}

// GetConnectionByPeerAddr get a connection by peer GP from connections pool.
func (c *connections) GetConnectionByPeerAddr(addr [16]byte) (conn protocol.Connection, err Error) {
	return c.poolByPeerAddr[addr]
}

// GetConnectionByUserIDDelegateUserID return the connection from pool or app storage.
// A connection can use just by single app node, so user can't use same connection to connect other node before close connection on usage node.
func (c *Connections) GetConnectionByUserIDDelegateUserID(userID, delegateUserID [16]byte) (conn protocol.Connection, err Error) {
	conn = c.poolByID[connID]
	if conn == nil {
		conn = c.getConnectionByID(connID)
	}
	return
}

// GetConnectionsByUserID get the connections by peer userID||domainID from connections pool.
func (c *Connections) GetConnectionsByUserID(userID [32]byte) (conn []protocol.Connection, err Error) {
	conn = c.poolByUserID[userID]
	if conn == nil {
		// TODO::: check storage
	}
	// TODO::: check if connection is in not ready status
	return
}

// GetConnectionByDomain return the connection from pool or app storage if any exist.
func (c *Connections) GetConnectionByDomain(domain string) (conn protocol.Connection, err Error) {
	conn = c.poolByDomain[domain]
	if conn == nil {
		// TODO::: check storage
	}
	// TODO::: check if connection is in not ready status
	return
}

// RegisterConnection register new connection in server connection pool
func (c *connections) RegisterConnection(conn protocol.Connection, err Error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.poolByConnID[conn.ID] = conn
	if conn.Addr() != AddrNil {
		c.poolByPeerAddr[conn.Addr()] = conn
	}
	if conn.UserID != [32]byte{} {
		c.poolByUserID[conn.UserID()] = append(c.poolByUserID[conn.UserID()], conn)
	}
	if conn.UserType == protocol.UserTypeGuest {
		c.GuestConnectionCount++
	}
}

// CloseConnection un-register exiting connection in server connection pool
func (c *connections) CloseConnection(conn *Connection, err Error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if conn.UserType == protocol.UserTypeGuest {
		c.GuestConnectionCount--
	}

	// TODO::: Is it worth to don't delete connection just reset it and send it to pool of unused connection due to GC
	delete(c.poolByConnID, conn.ID)
	delete(c.poolByPeerAddr, conn.Addr())
	// delete(c.poolByUserID, conn.UserID)

	// Let unfinished stream handled
}

// RevokeConnection un-register exiting connection in server connection pool without given any time to finish any stream
func (c *connections) RevokeConnection(conn *Connection, err Error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if conn.UserType == protocol.UserTypeGuest {
		c.GuestConnectionCount--
	}

	// Remove all unfinished stream first

	delete(c.poolByConnID, conn.ID)
	delete(c.poolByPeerAddr, conn.Addr())
	// delete(c.poolByUserID, conn.UserID)
}

func (c *Connections) Shutdown() {
	protocol.App.Log(protocol.Log_Information, "Connections - ShutDown - Saving proccess begin ...")
	protocol.App.Log(protocol.Log_Information, "Connections - ShutDown - Number of active connections:", len(c.poolByUserIDDelegateUserID))
	for _, conn := range c.poolByUserIDDelegateUserID {
		go c.SaveConnection(conn)
	}
	protocol.App.Log(protocol.Log_Information, "Connections - ShutDown - Saving proccess end now")
}

func (c *Connections) connectionIdleTimeoutSaveAndFree() {
	var timer = time.NewTimer(Server.Manifest.TechnicalInfo.ConnectionIdleTimeout.ConvertToTimeDuration())
	for {
		select {
		case <-timer.C:
			protocol.App.Log(protocol.Log_Information, "Connections - Cron - Idle connections timeout save and free proccess begin ...")
			protocol.App.Log(protocol.Log_Information, "Connections - Cron - Number of active connections:", len(c.poolByUserIDDelegateUserID))
			// It is conccurent proccess and not take more than one or two second, so set timeNow here
			var timeNow = etime.Now()

			c.mutex.Lock()
			defer c.mutex.Unlock()
			for _, conn := range c.poolByUserIDDelegateUserID {
				if !conn.LastUsage.AddDuration(Server.Manifest.TechnicalInfo.ConnectionIdleTimeout).Pass(timeNow) {
					continue
				}

				go c.SaveConnection(conn)

				if conn.UserType == authorization.UserTypeGuest {
					c.GuestConnectionCount--
				}
				// Remove all unfinished stream first
				delete(c.poolByUserIDDelegateUserID, conn.ID)
				delete(c.poolByPeerAddr, conn.Addr())
			}

			timer.Reset(Server.Manifest.TechnicalInfo.ConnectionIdleTimeout.ConvertToTimeDuration())
			protocol.App.Log(protocol.Log_Information, "Connections - Cron - Number of active connections after:", len(c.poolByUserIDDelegateUserID))
			protocol.App.Log(protocol.Log_Information, "Connections - Cron - Idle connections timeout save and free proccess end now")
		}
	}
}
