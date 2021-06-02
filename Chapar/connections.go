/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"sync"

	er "../error"
	"../giti"
	"../log"
)

// Connections store pools of connection to retrieve in many ways!
// TODO::: add analytic methods
type Connections struct {
	mutex           sync.Mutex
	poolByPath      map[string]giti.LinkConnection // It is optimize to convert byte slice to string as key in map: https://github.com/golang/go/commit/f5f5a8b6209f84961687d993b93ea0d397f5d5bf
	poolByThingID   map[[32]byte]giti.LinkConnection
	poolByBlackList map[[32]byte]giti.LinkConnection
}

// Init initialize it!
func (c *Connections) Init() {
	if c.poolByPath == nil {
		c.poolByPath = make(map[string]giti.LinkConnection, 16384)
	}
	if c.poolByThingID == nil {
		c.poolByThingID = make(map[[32]byte]giti.LinkConnection, 16384)
	}
	if c.poolByBlackList == nil {
		c.poolByBlackList = make(map[[32]byte]giti.LinkConnection)
	}
}

// NewConnection ...
// TODO::: get ThingID from peer or func args??
func (c *Connections) NewConnection(port giti.LinkPort, path []byte) (conn giti.LinkConnection) {
	conn = &Connection{
		port: port,
	}
	conn.setPath(path)

	c.RegisterConnection(conn)
	return
}

// GetConnectionByPath get a connection by its path from Connections pool!!
func (c *Connections) GetConnectionByPath(path []byte) (conn giti.LinkConnection, err *er.Error) {
	conn = c.poolByPath[string(path)]
	if conn == nil {
		conn = &Connection{
			path: path,
		}
		err = conn.GetLastByPath()
	}
	return
}

// GetConnectionsByThingID get the Connections by peer thingID from Connections pool!!
func (c *Connections) GetConnectionsByThingID(thingID [32]byte) (conn giti.LinkConnection, err *er.Error) {
	conn = c.poolByThingID[thingID]
	if conn == nil {
		conn = &Connection{
			thingID: thingID,
		}
		err = conn.GetLastByThingID()
	}
	return
}

// RegisterConnection register new connection in the connection pool!!
func (c *Connections) RegisterConnection(conn giti.LinkConnection) {
	c.mutex.Lock()
	c.poolByPath[string(conn.Path)] = conn
	c.poolByThingID[conn.ThingID] = conn
	c.mutex.Unlock()
}

// RegisterNewPathForConnection register new alternative path for connection and save it in the connection pool!!
func (c *Connections) RegisterNewPathForConnection(conn giti.LinkConnection, alternativePath []byte) {
	conn.setAlternativePath(alternativePath)

	c.mutex.Lock()
	c.poolByPath[string(alternativePath)] = conn
	c.mutex.Unlock()
}

// CloseConnection un-register exiting connection in the connection pool!!
func (c *Connections) CloseConnection(conn giti.LinkConnection) {
	c.mutex.Lock()
	// TODO::: Is it worth to don't delete connection just reset it and send it to pool of unused connection due to GC!
	delete(c.poolByPath, conn.Path)
	delete(c.poolByThingID, conn.ThingID)
	for path := range conn.conn.AlternativePath {
		delete(c.poolByPath, path)
	}
	c.mutex.Unlock()
}

// Shutdown ready the connection pools to shutdown!!
func (c *Connections) Shutdown() {
	log.Info("Chapar - ShutDown - Connections saving proccess begin...")
	log.Info("Chapar - ShutDown - Number of active connections:", len(c.poolByThingID))
	for _, conn := range c.poolByThingID {
		go conn.saveConn()
	}
	log.Info("Chapar - ShutDown - Connections saving proccess end now!")
}
