/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

import (
	"hash/crc32"

	"../chapar"
	"../crypto"
	"../uuid"
)

// userConnections store pools of connection to retrieve in many ways!
type userConnections struct {
	old            *userConnections // use just in expanding
	pathsPool      [254][]*UserConnection
	usersGPPoolLen uint32
	usersGPPool    []userGPBucket
	usersPoolLen   uint32
	usersPool      []userBucket
}

type userGPBucket struct {
	FreeID   uint8
	UserGPID [8]uint32
	Conn     [8]*UserConnection
}

type userBucket struct {
	FreeID uint8
	UserID [8][16]byte
	Conn   [8][]*UserConnection
}

func (uc *userConnections) init() {
	// TotalHop = 1 or One hop frame use just in p2p networks and not pass here!!
	// TotalHop = 2 or Two hop implement by just one Chapar switch device!
	uc.pathsPool[0] = make([]*UserConnection, 256)
	// TotalHop > 2 or More than two hop implement by more than one Chapar switch device!
	var i uint8
	for i = 1; i < 255; i++ {
		uc.pathsPool[i] = make([]*UserConnection, 2048)
	}

	uc.usersGPPool = make([]userGPBucket, 2048)
	uc.usersGPPoolLen = 2048

	uc.usersPool = make([]userBucket, 2048)
	uc.usersPoolLen = 2048
}

func (uc *userConnections) expandPathPool(hopNum uint8) (newPoolLen uint32) {
	var poolLen uint32 = uint32(len(uc.pathsPool[hopNum]))
	newPoolLen = poolLen * 2

	var new = make([]*UserConnection, newPoolLen)

	// Copy old paths
	var existingConn *UserConnection
	var loc uint32
	var i uint32
	for i = 0; i < poolLen; i++ {
		existingConn = uc.pathsPool[hopNum][i]
		if existingConn != nil {
			loc = crc32.ChecksumIEEE(existingConn.Path) % newPoolLen
			new[loc] = existingConn
		}
	}
	uc.pathsPool[hopNum] = new

	return
}

func (uc *userConnections) expandUsersGPPool() {
	var newPoolLen uint32 = uc.usersGPPoolLen * 2
	var new = make([]userGPBucket, newPoolLen)

	// Copy old records
	var loc uint32
	var i uint32
	var j uint8
	var f uint8
	for i = 0; i < uc.usersGPPoolLen; i++ {
		f = uc.usersGPPool[i].FreeID
		for j = 0; j < f; j++ {
			loc = uc.usersGPPool[i].Conn[j].UserGPID % newPoolLen
			new[loc].UserGPID[f] = uc.usersGPPool[i].Conn[j].UserGPID
			new[loc].Conn[f] = uc.usersGPPool[i].Conn[j]
			uc.usersGPPool[i].FreeID++
		}
	}

	uc.usersGPPool = new
	uc.usersGPPoolLen = newPoolLen
}

func (uc *userConnections) expandUsersPool() {
	var newPoolLen uint32 = uc.usersPoolLen * 2
	var newPool = make([]userBucket, newPoolLen)

	// Copy old records
	var loc uint32
	var i uint32
	var j uint8
	for i = 0; i < uc.usersGPPoolLen; i++ {
		for j = 0; j < uc.usersPool[i].FreeID; j++ {
			// TODO::: change finding loc algorithm!!
			loc = uuid.GetFirstUint32(uc.usersPool[i].UserID[j]) % newPoolLen
			newPool[loc].UserID[uc.usersPool[i].FreeID] = uc.usersPool[i].UserID[j]
			newPool[loc].Conn[uc.usersPool[i].FreeID] = uc.usersPool[i].Conn[j]
			uc.usersPool[i].FreeID++
		}
	}

	uc.usersPool = newPool
	uc.usersPoolLen = newPoolLen
}

// registerNewPath use to register or overwrite a path in the related pool!
func (uc *userConnections) registerNewPath(conn *UserConnection, path []byte) {
	var pathLen uint8 = uint8(len(conn.Path))
	var poolLen uint32 = uint32(len(uc.pathsPool[pathLen]))
	var loc = crc32.ChecksumIEEE(conn.Path) % poolLen

	// Check collision!!
	var existingConn = uc.pathsPool[pathLen][loc]
	if existingConn != nil && existingConn.UserID != conn.UserID {
		poolLen = uc.expandPathPool(pathLen)
		loc = crc32.ChecksumIEEE(conn.Path) % poolLen
	}

	uc.pathsPool[pathLen][loc] = conn
}

// registerNewUser use to register or overwrite a user in the related pool!
func (uc *userConnections) registerNewUser(conn *UserConnection) {
	// TODO:::
}

// MakeNewUserConnectionReq is the request structure of MakeNewUserConnection()
type MakeNewUserConnectionReq struct {
	UserID        [16]byte
	ThingID       [16]byte
	UserPublicKey [32]byte
	Path          []byte
	MaxBandwidth  uint64
	Signature     [256]byte
}

// MakeNewUserConnection use to make new connection!
func (uc *userConnections) MakeNewUserConnection(req *MakeNewUserConnectionReq) {
	// TODO::: Check signature first!

	var conns []*UserConnection
	var conn *UserConnection
	conns = uc.GetConnectionsByUserID(req.UserID)
	// Check if just new way to exiting user in specific thing!!
	if conns != nil {
		var i int
		var ln int = len(conns)
		for i = 0; i < ln; i++ {
			if conns[i].UserID == req.UserID && conns[i].ThingID == req.ThingID {
				conn = conns[i]
				break
			}
		}
	}

	if conn != nil {
		conn.AlternativePath = append(conn.AlternativePath, req.Path)
		uc.registerNewPath(conn, req.Path)
	} else {
		// TODO::: Check user can make new connection

		conn = &UserConnection{
			UserID:        req.UserID,
			Path:          req.Path,
			ReversePath:   chapar.ReversePath(req.Path),
			ThingID:       req.ThingID,
			Status:        routerConnectionStateOpen,
			MaxBandwidth:  req.MaxBandwidth,
			PeerPublicKey: req.UserPublicKey,
			Cipher:        crypto.NewGCM(crypto.NewAES256([32]byte{})),
		}
		uc.RegisterConnection(conn)
	}
}

// RegisterConnection use to register new connection in server connection pool!!
func (uc *userConnections) RegisterConnection(conn *UserConnection) {
	uc.registerNewPath(conn, conn.Path)
	uc.registerNewUser(conn)
}

// GetConnectionByPath use to get a connection by peer GP from connections pool!!
func (uc *userConnections) GetConnectionByPath(path []byte) (conn *UserConnection) {
	var pathLen = len(path)
	var poolLen uint32 = uint32(len(uc.pathsPool[pathLen]))
	var loc = crc32.ChecksumIEEE(path) % poolLen

	conn = uc.pathsPool[pathLen][loc]
	return
}

// GetConnectionsByUserID use to get connections by peer UserID from connections pool!!
func (uc *userConnections) GetConnectionsByUserID(userID [16]byte) (conns []*UserConnection) {
	// TODO::: change finding loc algorithm!!
	var loc = uuid.GetFirstUint32(userID) % uc.usersPoolLen

	var i uint8
	for i = 0; i < uc.usersPool[loc].FreeID; i++ {
		if uc.usersPool[loc].UserID[i] == userID {
			conns = uc.usersPool[loc].Conn[i]
			break
		}
		// check if connection is in not ready status
	}
	return
}

// GetConnectionByUserGPID use to get a connection by user GP ID part from connections pool!!
func (uc *userConnections) GetConnectionByUserGPID(userGPID uint32) (conn *UserConnection) {
	// TODO::: change finding loc algorithm!!
	var loc = userGPID % uc.usersGPPoolLen

	var i uint8
	for i = 0; i < uc.usersGPPool[loc].FreeID; i++ {
		if uc.usersGPPool[loc].UserGPID[i] == userGPID {
			conn = uc.usersGPPool[loc].Conn[i]
			break
		}
		// check if connection is in not ready status
	}
	return
}

// CloseConnection use to un-register exiting connection in server connection pool!!
func (uc *userConnections) CloseConnection(conn *UserConnection) {
	// TODO::: Check GC performance : delete connection vs just reset it and send it to pool of unused connection!!
	var pathLen = len(conn.Path)
	var poolLen uint32 = uint32(len(uc.pathsPool[pathLen]))
	var loc = crc32.ChecksumIEEE(conn.Path) % poolLen

	uc.pathsPool[pathLen][loc] = nil

	loc = conn.UserGPID % uc.usersGPPoolLen
	var eGPPool = uc.usersGPPool[loc]
	var f = eGPPool.FreeID
	var i uint8
	for i = 0; i < f; i++ {
		if eGPPool.UserGPID[i] == conn.UserGPID {
			eGPPool.UserGPID[i] = 0
			eGPPool.Conn[i] = nil
			break
		}
	}

	loc = uuid.GetFirstUint32(conn.UserID) % uc.usersPoolLen
	var eUserPool = uc.usersPool[loc]
	f = eUserPool.FreeID
	for i = 0; i < f; i++ {
		if eUserPool.UserID[i] == conn.UserID {
			var conns = eUserPool.Conn[i]
			var j = len(conns) - 1
			for ; j >= 0; j-- {
				if conns[j].UserGPID == conn.UserGPID {
					// TODO::: is it worth to delete an item in middle of array!!?
					conns[j] = nil
					break
				}
			}

			break
		}
	}
}
