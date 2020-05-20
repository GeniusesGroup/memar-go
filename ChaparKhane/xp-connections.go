/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

import (
	"hash/crc32"

	"../chapar"
	"../crypto"
)

// xpConnections store pools of connection to retrieve in many ways!
type xpConnections struct {
	old        *xpConnections // use just in expanding
	pathsPool  [254][]*routerConnection
	XPsPoolLen uint32
	XPsPool    []*xpPool
}

type xpPool struct {
	XPID     uint32
	FreeID   uint8
	RouterID [8]uint32
	Conn     [8]*routerConnection
}

func (xp *xpConnections) init() {
	// TotalHop = 1 or One hop frame use just in p2p networks and not pass here!!
	// TotalHop = 2 or Two hop implement by just one Chapar switch device!
	xp.pathsPool[0] = make([]*routerConnection, 256)
	// TotalHop > 2 or More than two hop implement by more than one Chapar switch device!
	var i uint8
	for i = 1; i < 255; i++ {
		xp.pathsPool[i] = make([]*routerConnection, 2048)
	}

	// TODO::: use existing XP number instead of below 1024 number
	xp.XPsPool = make([]*xpPool, 1024)
	xp.XPsPoolLen = 1024
}

func (xp *xpConnections) expandPathPool(hopNum uint8) (newPoolLen uint32) {
	var poolLen uint32 = uint32(len(xp.pathsPool[hopNum]))
	newPoolLen = poolLen * 2

	var new = make([]*routerConnection, newPoolLen)

	// Copy old paths
	var existingConn *routerConnection
	var loc uint32
	var i uint32
	for i = 0; i < poolLen; i++ {
		existingConn = xp.pathsPool[hopNum][i]
		if existingConn != nil {
			loc = crc32.ChecksumIEEE(existingConn.Path) % newPoolLen
			new[loc] = existingConn
		}
	}
	xp.pathsPool[hopNum] = new

	return
}

func (xp *xpConnections) expandXPsPool() {
	var newPoolLen uint32 = xp.XPsPoolLen * 2
	var new = make([]*xpPool, newPoolLen)

	// Copy old XPs
	var eXPPool *xpPool
	var loc uint32
	var i uint32
	for i = 0; i < xp.XPsPoolLen; i++ {
		eXPPool = xp.XPsPool[i]
		if eXPPool != nil {
			loc = eXPPool.XPID % newPoolLen
			new[loc] = eXPPool
		}
	}

	xp.XPsPool = new
	xp.XPsPoolLen = newPoolLen
}

// registerNewPath use to register or overwrite a path in the related pool!
func (xp *xpConnections) registerNewPath(conn *routerConnection, path []byte) {
	var pathLen uint8 = uint8(len(conn.Path))
	var poolLen uint32 = uint32(len(xp.pathsPool[pathLen]))
	var loc = crc32.ChecksumIEEE(conn.Path) % poolLen

	// Check collision!!
	var existingConn = xp.pathsPool[pathLen][loc]
	if existingConn != nil && existingConn.RouterID != conn.RouterID {
		poolLen = xp.expandPathPool(pathLen)
		loc = crc32.ChecksumIEEE(conn.Path) % poolLen
	}

	xp.pathsPool[pathLen][loc] = conn
}

// registerNewRouter use to register or overwrite a router in the related pool!
func (xp *xpConnections) registerNewRouter(conn *routerConnection) {
	var eXPPool = xp.GetXPPool(conn.XPID)
	if eXPPool != nil {
		if eXPPool.XPID == conn.XPID {
			for i := 0; i < 8; i++ {
				// Check for existing connection
				if eXPPool.RouterID[i] == conn.RouterID {
					eXPPool.Conn[i].AlternativePath = append(eXPPool.Conn[i].AlternativePath, conn.Path)
					xp.registerNewPath(eXPPool.Conn[i], conn.Path)
				} else if eXPPool.RouterID[i] == 0 {
					eXPPool.FreeID++
					eXPPool.RouterID[i] = conn.RouterID
					eXPPool.Conn[i] = conn
					return
				}
			}
		} else {
			// Collision occurred!!
			xp.expandXPsPool()
			xp.registerNewRouter(conn)
		}
	} else {
		eXPPool = &xpPool{
			XPID:     conn.XPID,
			FreeID:   1,
			RouterID: [8]uint32{conn.RouterID, 0, 0, 0, 0, 0, 0, 0},
			Conn:     [8]*routerConnection{conn, nil, nil, nil, nil, nil, nil, nil},
		}
		var xpLoc = conn.XPID % xp.XPsPoolLen
		xp.XPsPool[xpLoc] = eXPPool
	}
}

// MakeNewXPConnectionReq is the request structure of MakeNewXPConnection()
type MakeNewXPConnectionReq struct {
	XPID         uint32
	RouterID     uint32
	Path         []byte
	MaxBandwidth uint64
	Signature    [256]byte
}

// MakeNewXPConnection use to make new connection!
func (xp *xpConnections) MakeNewXPConnection(req *MakeNewXPConnectionReq) {
	// TODO::: Check signature first!

	var conn = routerConnection{
		RouterID:     req.RouterID,
		Path:         req.Path,
		ReversePath:  chapar.ReversePath(req.Path),
		Status:       routerConnectionStateOpen,
		MaxBandwidth: req.MaxBandwidth,
		Cipher:       crypto.NewGCM(crypto.NewAES256([32]byte{})),
	}
	xp.RegisterConnection(&conn)
}

// RegisterConnection use to register new connection in server xpConnections pool!!
func (xp *xpConnections) RegisterConnection(conn *routerConnection) {
	xp.registerNewPath(conn, conn.Path)
	xp.registerNewRouter(conn)
}

// GetConnectionByPath use to get a connection by peer GP from xpConnections pool!!
func (xp *xpConnections) GetConnectionByPath(path []byte) (conn *routerConnection) {
	var pathLen = len(path)
	var poolLen uint32 = uint32(len(xp.pathsPool[pathLen]))
	var loc = crc32.ChecksumIEEE(path) % poolLen

	conn = xp.pathsPool[pathLen][loc]
	return
}

// GetConnectionByXPID use to get a connection by given XP ID from xpConnections pool!!
func (xp *xpConnections) GetXPPool(xpID uint32) (eXPPool *xpPool) {
	var xpLoc = xpID % xp.XPsPoolLen
	eXPPool = xp.XPsPool[xpLoc]
	return
}

// GetConnectionByXPID use to get a connection by given XP ID from xpConnections pool!!
func (xp *xpConnections) GetConnectionByXPID(xpID uint32) (conn *routerConnection) {
	var xpLoc = xpID % xp.XPsPoolLen
	conn = xp.XPsPool[xpLoc].Conn[0]

	// check if connection is in not ready status

	return
}

// CloseConnection use to un-register existing connection in xpConnections pool!!
func (xp *xpConnections) CloseConnection(conn *routerConnection) {
	// TODO::: Check GC performance : delete connection vs just reset it and send it to pool of unused connection!!
	var pathLen = len(conn.Path)
	var poolLen uint32 = uint32(len(xp.pathsPool[pathLen]))
	var loc = crc32.ChecksumIEEE(conn.Path) % poolLen
	var xpLoc = conn.XPID % xp.XPsPoolLen

	xp.pathsPool[pathLen][loc] = nil

	var i uint8
	var f = xp.XPsPool[xpLoc].FreeID
	var eXPPool = xp.XPsPool[xpLoc]
	for ; i < f; i++ {
		if eXPPool.RouterID[i] == conn.RouterID {
			eXPPool.RouterID[i] = 0
			eXPPool.Conn[i] = nil
			break
		}
	}
}
