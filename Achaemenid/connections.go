/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import "../crypto"

// Connections store pools of connection to retrieve in many ways!
type Connections struct {
	PoolByDomainID  map[[16]byte]*Connection
	PoolByPeerGP    map[[16]byte]*Connection
	PoolByUserID    map[[16]byte][]*Connection
	PoolByBlackList map[[16]byte]*Connection
}

func (c *Connections) init() {
	if c.PoolByPeerGP == nil {
		c.PoolByPeerGP = make(map[[16]byte]*Connection)
	}
	if c.PoolByUserID == nil {
		c.PoolByUserID = make(map[[16]byte][]*Connection)
	}
}

// MakeNewConnectionByDomainID use to
func (c *Connections) MakeNewConnectionByDomainID(domainID [16]byte) (conn *Connection, err error) {
	// TODO::: Get closest domain GP
	var domainGP = [16]byte{}
	conn, err = c.MakeNewConnectionByPeerGP(domainGP)

	return
}

// MakeNewConnectionByPeerGP use to make new connection by peer GP and initialize it!
func (c *Connections) MakeNewConnectionByPeerGP(peerGP [16]byte) (conn *Connection, err error) {
	// TODO::: Make connection by ask peer public key from peer GP router!

	conn = &Connection{
		Cipher:     crypto.NewGCM(crypto.NewAES256([32]byte{})),
		StreamPool: make(map[uint32]*Stream),
	}

	c.RegisterConnection(conn)

	return
}

// RegisterConnection use to register new connection in server connection pool!!
func (c *Connections) RegisterConnection(conn *Connection) {
	c.PoolByPeerGP[conn.GPAddress] = conn
	// c.PoolByUserID[conn.OwnerUserID] = conn
}

// GetConnectionByPeerGP use to get a connection by peer GP from connections pool!!
func (c *Connections) GetConnectionByPeerGP(peerGPAddress [16]byte) (conn *Connection, ok bool) {
	conn, ok = c.PoolByPeerGP[peerGPAddress]
	return
}

// GetConnectionByDomainID use to get a connection by peer domain ID from connections pool!!
func (c *Connections) GetConnectionByDomainID(domainID [16]byte) (conn *Connection, ok bool) {
	conn, ok = c.PoolByDomainID[domainID]
	// check if connection is in not ready status
	return
}

// CloseConnection use to un-register exiting connection in server connection pool!!
func (c *Connections) CloseConnection(conn *Connection) {
	// TODO : Don't delete connection just reset it and send it to pool of unused connection due to GC!
	delete(c.PoolByPeerGP, conn.GPAddress)
	delete(c.PoolByUserID, conn.UserID)

	// Let unfinished stream handled!!
}

// RevokeConnection use to un-register exiting connection in server connection pool without given any time to finish any stream!!
func (c *Connections) RevokeConnection(conn *Connection) {
	// Remove all unfinished stream first!!

	delete(c.PoolByPeerGP, conn.GPAddress)
	delete(c.PoolByUserID, conn.UserID)
}
