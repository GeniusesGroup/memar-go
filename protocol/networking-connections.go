/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type Connections interface {
	GetConnectionByPeerAddr(addr [16]byte) (conn Connection, err Error)
	// A connection can use just by single app node, so user can't use same connection to connect other node before close connection on usage node.
	GetConnectionByUserIDDelegateUserID(userID, delegateUserID [16]byte) (conn Connection, err Error)
	GetConnectionsByUserID(userID [16]byte) (conns []Connection, err Error)
	GetConnectionByDomain(domain string) (conn Connection, err Error)

	// state=unregistered -> 'register' -> state=registered -> 'deregister' -> state=unregistered.
	RegisterConnection(conn Connection) (err Error)
	DeregisterConnection(conn Connection) (err Error)

	ConnectionsMetrics
}

// ConnectionMetrics
type ConnectionsMetrics interface {
	LastUsage() Time   // Last use of the connection
	OpenCount() int64  // number of opened and pending open connections
	InUseCount() int64 // The number of connections currently in use.
	IdleCount() int64  // The number of idle connections.

	GuestConnectionCount() int64
	ClosedCount() int64         // ClosedCount is an atomic counter which represents a total number of closed connections.
	WaitCount() int64           // Total number of connections waited for.
	IdleClosedCount() int64     // Total number of connections closed due to idle count.
	IdleTimeClosedCount() int64 // Total number of connections closed due to idle time.
	LifetimeClosedCount() int64 // Total number of connections closed due to max connection lifetime limit.

	// Rate() uint // Byte/Second
}
