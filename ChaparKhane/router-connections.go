/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

import (
	"hash/crc32"

	"../chapar"
	"../crypto"
)

// routerConnections store pools of connection to retrieve in many ways!
type routerConnections struct {
	old            *routerConnection        // use just in expanding
	pathsPool      [254][]*routerConnection // use 254 due to Chapar frame with TotalHop = 0||1 handled before pass it here!!
	routersPoolLen uint32
	routersPool    []*routerConnection // It will work due to routers IDs is in increment manner!
}

// TODO::: convert []*routerConnection to below structure
type pathPool struct {
	Len uint32
	Add uintptr
}

func (rc *routerConnections) init() {
	// TotalHop = 1 or One hop frame use just in p2p networks and not pass here!!
	// TotalHop = 2 or Two hop implement by just one Chapar switch device!
	rc.pathsPool[0] = make([]*routerConnection, 256)
	// TotalHop > 2 or More than two hop implement by more than one Chapar switch device!
	var i uint8
	for i = 1; i < 255; i++ {
		rc.pathsPool[i] = make([]*routerConnection, 2048)
	}

	// TODO::: use exiting router number instead of below 1024 number
	rc.routersPool = make([]*routerConnection, 1024)
	rc.routersPoolLen = 1024
}

func (rc *routerConnections) expandPathPool(hopNum uint8) (newPoolLen uint32) {
	var poolLen uint32 = uint32(len(rc.pathsPool[hopNum]))
	newPoolLen = poolLen * 2

	var new = make([]*routerConnection, newPoolLen)

	// Copy old paths
	var exitingConn *routerConnection
	var loc uint32
	var i uint32
	for i = 0; i < poolLen; i++ {
		exitingConn = rc.pathsPool[hopNum][i]
		if exitingConn != nil {
			loc = crc32.ChecksumIEEE(exitingConn.Path) % newPoolLen
			new[loc] = exitingConn
		}
	}
	rc.pathsPool[hopNum] = new

	return
}

func (rc *routerConnections) expandRoutersPool() {
	rc.routersPool = append(rc.routersPool, make([]*routerConnection, rc.routersPoolLen)...)
	rc.routersPoolLen *= 2
}

// registerNewPath use to register or overwrite a path in the related pool!
func (rc *routerConnections) registerNewPath(conn *routerConnection, path []byte) {
	var pathLen uint8 = uint8(len(conn.Path))
	var poolLen uint32 = uint32(len(rc.pathsPool[pathLen]))
	var loc = crc32.ChecksumIEEE(conn.Path) % poolLen

	// Check collision!!
	var exitingConn = rc.pathsPool[pathLen][loc]
	if exitingConn != nil && exitingConn.RouterID != conn.RouterID {
		poolLen = rc.expandPathPool(pathLen)
		loc = crc32.ChecksumIEEE(conn.Path) % poolLen
	}

	rc.pathsPool[pathLen][loc] = conn
}

// registerNewRouter use to register or overwrite the given router in the related pool!
func (rc *routerConnections) registerNewRouter(conn *routerConnection) {
	if conn.RouterID > rc.routersPoolLen {
		rc.expandRoutersPool()
	}

	rc.routersPool[conn.RouterID] = conn
}

// MakeNewRouterConnectionReq is the request structure of MakeNewRouterConnection()
type MakeNewRouterConnectionReq struct {
	RouterID     uint32
	Path         []byte
	MaxBandwidth uint64
	Signature    [256]byte
}

// MakeNewRouterConnection use to make connection from router advertisement!
func (rc *routerConnections) MakeNewRouterConnection(req *MakeNewRouterConnectionReq) {
	// TODO::: Check signature first!

	var conn *routerConnection
	conn = rc.GetConnectionByRouterID(req.RouterID)
	// Check if just new way to exiting router!!
	if conn != nil && conn.RouterID == req.RouterID {
		conn.AlternativePath = append(conn.AlternativePath, req.Path)
		rc.registerNewPath(conn, req.Path)
	} else {
		conn = &routerConnection{
			RouterID:     req.RouterID,
			Path:         req.Path,
			ReversePath:  chapar.ReversePath(req.Path),
			Status:       routerConnectionStateOpen,
			MaxBandwidth: req.MaxBandwidth,
			Cipher:       crypto.NewGCM(crypto.NewAES256([32]byte{})),
		}
		rc.RegisterConnection(conn)
	}
}

// RegisterConnection use to register new connection in routerConnections pools!!
func (rc *routerConnections) RegisterConnection(conn *routerConnection) {
	rc.registerNewPath(conn, conn.Path)
	rc.registerNewRouter(conn)
}

// GetConnectionByPath use to get a connection by path from routerConnections pool!!
func (rc *routerConnections) GetConnectionByPath(path []byte) (conn *routerConnection) {
	var pathLen = len(path)
	var poolLen uint32 = uint32(len(rc.pathsPool[pathLen]))
	var loc = crc32.ChecksumIEEE(path) % poolLen

	conn = rc.pathsPool[pathLen][loc]
	return
}

// GetConnectionByRouterID use to get a connection by routerID from routerConnections pool!!
func (rc *routerConnections) GetConnectionByRouterID(routerID uint32) (conn *routerConnection) {
	// respect routersPool length otherwise panic will occur!
	if routerID <= rc.routersPoolLen {
		conn = rc.routersPool[routerID]
	}
	return
}

// CloseConnection use to un-register exiting connection in server connection pool!!
func (rc *routerConnections) CloseConnection(conn *routerConnection) {
	// TODO::: Check GC performance : delete connection vs just reset it and send it to pool of unused connection!!
	var pathLen = len(conn.Path)
	var poolLen uint32 = uint32(len(rc.pathsPool[pathLen]))
	var loc = crc32.ChecksumIEEE(conn.Path) % poolLen

	rc.pathsPool[pathLen][loc] = nil
	rc.routersPool[conn.RouterID] = nil
}
