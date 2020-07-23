/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"net"

	"../crypto"
	"../errors"
)

// connections store pools of connection to retrieve in many ways!
type connections struct {
	server               *Server
	poolByPeerAdd        map[[16]byte]*Connection
	poolByConnID         map[[16]byte]*Connection
	poolByUserID         map[[16]byte][]*Connection
	poolByDomainID       map[[16]byte][]*Connection
	poolByBlackList      map[[16]byte]*Connection
	GuestConnectionCount uint64
	GetConnByID          func(connID [16]byte) (conn *Connection)
	GetConnByUserID      func(userID, appID, thingID [16]byte) (conn *Connection)
}

// Errors
var (
	ErrGuestUsersConnectionNotAllow    = errors.New("GuestUserConnectionNotAllow", "Guest users don't allow to make new connection")
	ErrMaxOpenedGuestConnectionReached = errors.New("MaxOpenedGuestConnectionReached", "This server not have enough resource to make new guest connection, register or try other server")
)

func (c *connections) init(s *Server) {
	c.server = s

	if c.poolByPeerAdd == nil {
		c.poolByPeerAdd = make(map[[16]byte]*Connection)
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
	// TODO::: Get closest domain GP
	var domainGP = [16]byte{}
	conn, err = c.MakeNewConnectionByPeerAdd(domainGP)
	c.poolByDomainID[conn.DomainID] = append(c.poolByDomainID[conn.DomainID], conn)
	return
}

// MakeNewConnectionByPeerAdd use to make new connection by peer GP and initialize it!
func (c *connections) MakeNewConnectionByPeerAdd(peerGP [16]byte) (conn *Connection, err error) {
	var userID, appID, thingID [16]byte
	// TODO::: Get peer publickey & userID & appID & thingID from peer GP router!

	if userID != [16]byte{} {
		conn = c.GetConnByUserID(userID, appID, thingID)
	}

	// If conn not exist means guest connection.
	if conn == nil {
		conn = &Connection{
			Server:     c.server,
			GPAddr:     peerGP,
			Cipher:     crypto.NewGCM(crypto.NewAES256([32]byte{})),
			StreamPool: make(map[uint32]*Stream),
		}
	}

	c.RegisterConnection(conn)

	return
}

// RegisterConnection use to register new connection in server connection pool!!
func (c *connections) RegisterConnection(conn *Connection) {
	c.poolByPeerAdd[conn.GPAddr] = conn
	c.poolByConnID[conn.ID] = conn
	if conn.UserID != [16]byte{} {
		c.poolByUserID[conn.UserID] = append(c.poolByUserID[conn.UserID], conn)
	}
}

// GetConnectionByPeerAdd use to get a connection by peer GP from connections pool!!
func (c *connections) GetConnectionByPeerAdd(peerAddress [16]byte) *Connection {
	return c.poolByPeerAdd[peerAddress]
}

// GetConnectionByID use to get a connection by its ID from connections pool!!
func (c *connections) GetConnectionByID(connID [16]byte) (conn *Connection) {
	conn = c.poolByConnID[connID]
	if conn == nil {
		// get from platfrom resources
		conn = c.GetConnByID(connID)
	}
	return
}

// GetConnectionsByDomainID use to get a connection by peer domain ID from connections pool!!
func (c *connections) GetConnectionsByDomainID(domainID [16]byte) (conn []*Connection) {
	conn = c.poolByDomainID[domainID]
	// TODO::: check if connection is in not ready status
	return
}

// GetConnectionByDomainID use to get a connection by peer domain ID from connections pool!!
func (c *connections) GetConnectionByDomainID(domainID [16]byte) (conn *Connection) {
	conn = c.poolByDomainID[domainID][0]
	// TODO::: check if connection is in not ready status
	return
}

// CloseConnection use to un-register exiting connection in server connection pool!!
func (c *connections) CloseConnection(conn *Connection) {
	if conn.UserType == 0 {
		c.GuestConnectionCount--
	}

	// TODO::: Don't delete connection just reset it and send it to pool of unused connection due to GC!
	delete(c.poolByPeerAdd, conn.GPAddr)
	delete(c.poolByUserID, conn.UserID)

	// Let unfinished stream handled!!
}

// RevokeConnection use to un-register exiting connection in server connection pool without given any time to finish any stream!!
func (c *connections) RevokeConnection(conn *Connection) {
	if conn.UserType == 0 {
		c.GuestConnectionCount--
	}

	// Remove all unfinished stream first!!

	delete(c.poolByPeerAdd, conn.GPAddr)
	delete(c.poolByUserID, conn.UserID)
}

// MakeNewIPConnection make new connection and return to caller.
// TODO::: due to IPv4&&IPv6 support we need this! Remove it when remove those support.
func (c *connections) MakeNewIPConnection(ipAddr *net.IPAddr) (conn *Connection, err error) {
	// TODO::: check new connection can create

	conn = &Connection{
		Server:     c.server,
		StreamPool: make(map[uint32]*Stream),
		IPAddr:     ipAddr,
	}

	c.RegisterConnection(conn)
	return
}

// MakeNewUDPConnection make new connection and return to caller.
// TODO::: due to IPv4&&IPv6 support we need this! Remove it when remove those support.
func (c *connections) MakeNewUDPConnection(udpAddr *net.UDPAddr) (conn *Connection, err error) {
	// TODO::: check new connection can create

	conn = &Connection{
		Server:     c.server,
		StreamPool: make(map[uint32]*Stream),
		UDPAddr:    udpAddr,
	}

	c.RegisterConnection(conn)
	return
}
