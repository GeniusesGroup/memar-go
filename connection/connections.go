/* For license and copyright information please see LEGAL file in repository */

package connection

import (
	"strconv"
	"sync"
	"time"

	"../log"
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
	shutdownSignal             chan struct{}
}

func (c *Connections) Init() {
	c.poolByPeerAddr = make(map[[16]byte]protocol.Connection, 16384)
	c.poolByUserIDDelegateUserID = make(map[[32]byte]protocol.Connection, 16384)
	c.poolByUserID = make(map[[16]byte][]protocol.Connection, 16384)
	c.poolByDomain = make(map[string]protocol.Connection, 16384)
	c.poolByBlackList = make(map[[16]byte]protocol.Connection, 256)
	c.shutdownSignal = make(chan struct{})

	go c.connectionIdleTimeoutSaveAndFree()
}

// GetConnectionByPeerAddr get a connection by peer GP from connections pool.
func (c *Connections) GetConnectionByPeerAddr(addr [16]byte) (conn protocol.Connection, err protocol.Error) {
	conn = c.poolByPeerAddr[addr]
	if conn == nil {
		err = ErrNoConnection
	}
	return
}

// GetConnectionByUserIDDelegateUserID return the connection from pool or app storage.
// A connection can use just by single app node, so user can't use same connection to connect other node before close connection on usage node.
func (c *Connections) GetConnectionByUserIDDelegateUserID(userID, delegateUserID [16]byte) (conn protocol.Connection, err protocol.Error) {
	conn = c.poolByUserIDDelegateUserID[c.userIDDelegateUserID(userID, delegateUserID)]
	if conn == nil {
		err = ErrNoConnection
	}
	return
}

// GetConnectionsByUserID get the connections by peer userID||domainID from connections pool.
func (c *Connections) GetConnectionsByUserID(userID [16]byte) (conn []protocol.Connection, err protocol.Error) {
	conn = c.poolByUserID[userID]
	if conn == nil {
		err = ErrNoConnection
	}
	return
}

// GetConnectionByDomain return the connection from pool or app storage if any exist.
func (c *Connections) GetConnectionByDomain(domain string) (conn protocol.Connection, err protocol.Error) {
	conn = c.poolByDomain[domain]
	if conn == nil {
		err = ErrNoConnection
	}
	return
}

// RegisterConnection register new connection in server connection pool
func (c *Connections) RegisterConnection(conn protocol.Connection) (err protocol.Error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var userID = conn.UserID().UUID()
	c.poolByUserIDDelegateUserID[c.userIDDelegateUserID(userID, conn.DelegateUserID().UUID())] = conn
	if conn.Addr() != [16]byte{} {
		c.poolByPeerAddr[conn.Addr()] = conn
	}

	var connDomain = conn.DomainName()
	if connDomain != "" {
		c.poolByDomain[connDomain] = conn
	}

	if userID != [16]byte{} {
		c.poolByUserID[userID] = append(c.poolByUserID[userID], conn)
	}

	if conn.UserID().Type() == protocol.UserType_Unset {
		c.guestConnectionCount++
	}
	return
}

// DeregisterConnection delete the given connection in connections pool
func (c *Connections) DeregisterConnection(conn protocol.Connection) (err protocol.Error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	err = c.deregisterConnection(conn)
	return
}

func (c *Connections) deregisterConnection(conn protocol.Connection) (err protocol.Error) {
	if conn.UserID().Type() == protocol.UserType_Unset {
		c.guestConnectionCount--
	}

	var userID = conn.UserID().UUID()
	delete(c.poolByUserIDDelegateUserID, c.userIDDelegateUserID(userID, conn.DelegateUserID().UUID()))
	delete(c.poolByPeerAddr, conn.Addr())
	delete(c.poolByDomain, conn.DomainName())

	var userConnections = c.poolByUserID[userID]
	var userConnectionsLen =len(userConnections)
	for i := 0; i < userConnectionsLen; i++ {
		if userConnections[i].UserID().UUID() == userID {
			var newUserConnectionsLen = userConnectionsLen - 1
			// copy(userConnections[i:], userConnections[i+1:])
			// userConnections = userConnections[:newUserConnectionsLen]
			userConnections[i] = userConnections[newUserConnectionsLen]
			userConnections = userConnections[:newUserConnectionsLen]
			break
		}
	}
	c.poolByUserID[userID] = userConnections
	return
}

func (c *Connections) Shutdown() {
	c.shutdownSignal <- struct{}{}

	protocol.App.Log(log.InfoEvent(domainEnglish, "ShutDown - Saving proccess begin ...\n"+
		"Number of active connections:"+c.activeConnectionsNumber()))
	for _, conn := range c.poolByUserIDDelegateUserID {
		go conn.Close()
	}
	protocol.App.Log(log.InfoEvent(domainEnglish, "ShutDown - Saving proccess end now"))
}

func (c *Connections) userIDDelegateUserID(userID, delegateUserID [16]byte) (userIDDelegateUserID [32]byte) {
	copy(userIDDelegateUserID[:], userID[:])
	copy(userIDDelegateUserID[16:], delegateUserID[:])
	return
}

func (c *Connections) connectionIdleTimeoutSaveAndFree() {
	var timer = time.NewTimer(ConnectionIdleTimeout)
	var timerTime time.Time
	for {
		select {
		case timerTime = <-timer.C:
			protocol.App.Log(log.InfoEvent(domainEnglish, "Cron - Idle connections timeout save and free proccess begin ...\n"+
				"Number of active connections: "+c.activeConnectionsNumber()))

			var timeNowMilli = timerTime.UnixMilli()

			c.mutex.Lock()
			defer c.mutex.Unlock()
			for _, conn := range c.poolByUserIDDelegateUserID {
				var lastUsage = int64(conn.LastUsage())
				if lastUsage+ConnectionIdleTimeout < timeNowMilli {
					continue
				}

				c.deregisterConnection(conn)
				go conn.Close()
			}

			timer.Reset(ConnectionIdleTimeout)
			protocol.App.Log(log.InfoEvent(domainEnglish, "Cron - Number of active connections after:"+c.activeConnectionsNumber()+
				"\nIdle connections timeout save and free proccess end now"))
		case <-c.shutdownSignal:
			return
		}
	}
}

func (c *Connections) activeConnectionsNumber() string {
	return strconv.FormatInt(int64(len(c.poolByUserIDDelegateUserID)), 10)
}
