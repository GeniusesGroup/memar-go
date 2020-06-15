/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import "../crypto"

// connections store pools of connection to retrieve in many ways!
type connections struct {
	poolByPeerAdd   map[[16]byte]*Connection
	poolByUserID    map[[16]byte][]*Connection
	poolByDomainID  map[[16]byte]*Connection
	poolByBlackList map[[16]byte]*Connection
}

func (c *connections) init() {
	if c.poolByPeerAdd == nil {
		c.poolByPeerAdd = make(map[[16]byte]*Connection)
	}
	if c.poolByUserID == nil {
		c.poolByUserID = make(map[[16]byte][]*Connection)
	}
	if c.poolByDomainID == nil {
		c.poolByDomainID = make(map[[16]byte]*Connection)
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

	return
}

// MakeNewConnectionByPeerAdd use to make new connection by peer GP and initialize it!
func (c *connections) MakeNewConnectionByPeerAdd(peerGP [16]byte) (conn *Connection, err error) {
	// TODO::: Make connection by ask peer public key from peer GP router!

	conn = &Connection{
		GPAddress:  peerGP,
		Cipher:     crypto.NewGCM(crypto.NewAES256([32]byte{})),
		StreamPool: make(map[uint32]*Stream),
	}

	c.RegisterConnection(conn)

	return
}

// RegisterConnection use to register new connection in server connection pool!!
func (c *connections) RegisterConnection(conn *Connection) {
	c.poolByPeerAdd[conn.GPAddress] = conn
	c.poolByUserID[conn.UserID] = append(c.poolByUserID[conn.UserID], conn)
	if conn.DomainID != [16]byte{} {
		c.poolByDomainID[conn.DomainID] = conn
	}
}

// GetConnectionByPeerAdd use to get a connection by peer GP from connections pool!!
func (c *connections) GetConnectionByPeerAdd(peerAddress [16]byte) *Connection {
	return c.poolByPeerAdd[peerAddress]
}

// GetConnectionByDomainID use to get a connection by peer domain ID from connections pool!!
func (c *connections) GetConnectionByDomainID(domainID [16]byte) (conn *Connection, ok bool) {
	conn, ok = c.poolByDomainID[domainID]
	// check if connection is in not ready status
	return
}

// CloseConnection use to un-register exiting connection in server connection pool!!
func (c *connections) CloseConnection(conn *Connection) {
	// TODO::: Don't delete connection just reset it and send it to pool of unused connection due to GC!
	delete(c.poolByPeerAdd, conn.GPAddress)
	delete(c.poolByUserID, conn.UserID)

	// Let unfinished stream handled!!
}

// RevokeConnection use to un-register exiting connection in server connection pool without given any time to finish any stream!!
func (c *connections) RevokeConnection(conn *Connection) {
	// Remove all unfinished stream first!!

	delete(c.poolByPeerAdd, conn.GPAddress)
	delete(c.poolByUserID, conn.UserID)
}
