/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type LoadConnection interface {
	GetConnectionByPeerAddr(addr [16]byte) (conn Connection, err Error)
	// A connection can use just by single app node, so user can't use same connection to connect other node before close connection on usage node.
	GetConnectionByUserIDDelegateUserID(userID, delegateUserID [16]byte) (conn Connection, err Error)
	GetConnectionsByUserID(userID [16]byte) (conns []Connection, err Error)
	GetConnectionByDomain(domain string) (conn Connection, err Error)

	RegisterConnection(conn Connection) (err Error)
	CloseConnection(conn Connection) (err Error)
	RevokeConnection(conn Connection) (err Error)
}

/* Security data */
// Cipher() Cipher
