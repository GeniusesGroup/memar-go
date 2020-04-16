/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// Connections store pools of connection to retrieve in many ways!
type Connections struct {
	PoolByDomainID    map[[16]byte]*Connection
	PoolByPeerUIP     map[[16]byte]*Connection
	PoolByOwnerUserID map[[16]byte][]*Connection
	PoolByBlackList   map[[16]byte]*Connection
}

func (c *Connections) init() {
	if c.PoolByPeerUIP == nil {
		c.PoolByPeerUIP = make(map[[16]byte]*Connection)
	}
	if c.PoolByOwnerUserID == nil {
		c.PoolByOwnerUserID = make(map[[16]byte][]*Connection)
	}
}

// RegisterConnection use to register new connection in server connection pool!!
func (c *Connections) RegisterConnection(conn *Connection) {
	c.PoolByPeerUIP[conn.PeerUIPAddress] = conn
	// c.PoolByOwnerUserID[conn.OwnerUserID] = conn
}

// GetConnectionByPeerUIP use to get a connection by peer UIP from connections pool!!
func (c *Connections) GetConnectionByPeerUIP(peerUIPAddress [16]byte) (conn *Connection, ok bool) {
	conn, ok = c.PoolByPeerUIP[peerUIPAddress]
	return
}

// GetConnectionByDomainID use to get a connection by peer domain ID from connections pool!!
func (c *Connections) GetConnectionByDomainID(domainID [16]byte) (conn *Connection, ok bool) {
	conn, ok = c.PoolByDomainID[domainID]
	return
}

// CloseConnection use to un-register exiting connection in server connection pool!!
func (c *Connections) CloseConnection(conn *Connection) {
	// TODO : Don't delete connection just reset it and send it to pool of unused connection due to GC!
	delete(c.PoolByPeerUIP, conn.PeerUIPAddress)
	delete(c.PoolByOwnerUserID, conn.OwnerUserID)

	// Let unfinished stream handled!!
}

// RevokeConnection use to un-register exiting connection in server connection pool without given any time to finish any stream!!
func (c *Connections) RevokeConnection(conn *Connection) {
	// Remove all unfinished stream first!!

	delete(c.PoolByPeerUIP, conn.PeerUIPAddress)
	delete(c.PoolByOwnerUserID, conn.OwnerUserID)
}
