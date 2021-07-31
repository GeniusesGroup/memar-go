/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"sync"

	er "../error"
	"../giti"
	"../log"
)

// connections store pools of connection to retrieve in many ways!
// TODO::: add analytic methods
type connections struct {
	mutex           sync.Mutex
	poolByPath      map[string]giti.NetworkLinkConnection // It is optimize to convert byte slice to string as key in map: https://github.com/golang/go/commit/f5f5a8b6209f84961687d993b93ea0d397f5d5bf
	poolByThingID   map[[32]byte]giti.NetworkLinkConnection
	poolByBlackList map[[32]byte]giti.NetworkLinkConnection
}

func (c *connections) init() {
	if c.poolByPath == nil {
		c.poolByPath = make(map[string]giti.NetworkLinkConnection, 16384)
	}
	if c.poolByThingID == nil {
		c.poolByThingID = make(map[[32]byte]giti.NetworkLinkConnection, 16384)
	}
	if c.poolByBlackList == nil {
		c.poolByBlackList = make(map[[32]byte]giti.NetworkLinkConnection)
	}
}

func (c *connections) newConnection(port *portEndPoint, frame []byte) (conn giti.NetworkLinkConnection) {
	conn = &Connection{
		port: port,
	}
	conn.init(frame)

	// TODO::: get ThingID from peer or func args??

	c.registerConnection(conn)
	return
}

func (c *connections) establishConnectionByPath(path []byte) (conn *Connection, err giti.Error) {
	return
}

func (c *connections) establishConnectionByThingID(thingID [32]byte) (conn *Connection, err giti.Error) {
	return
}

// GetConnectionByPath get a connection by its path from connections pool!!
func (c *connections) getConnectionByPath(path []byte) (conn giti.NetworkLinkConnection, err giti.Error) {
	conn = c.poolByPath[string(path)]
	if conn == nil {
		conn = &Connection{
			path: path,
		}
		err = conn.GetLastByPath()
	}
	return
}

func (c *connections) getConnectionsByThingID(thingID [32]byte) (conn giti.NetworkLinkConnection, err giti.Error) {
	conn = c.poolByThingID[thingID]
	if conn == nil {
		conn = &Connection{
			thingID: thingID,
		}
		err = conn.GetLastByThingID()
	}
	return
}

func (c *connections) registerConnection(conn giti.NetworkLinkConnection) {
	c.mutex.Lock()
	c.poolByPath[conn.Path.GetAsString()] = conn
	c.poolByThingID[conn.ThingID] = conn
	c.mutex.Unlock()
}

func (c *connections) registerNewPathForConnection(conn giti.NetworkLinkConnection, alternativePath []byte) {
	conn.setAlternativePath(alternativePath)

	c.mutex.Lock()
	c.poolByPath[string(alternativePath)] = conn
	c.mutex.Unlock()
}

func (c *connections) closeConnection(conn giti.NetworkLinkConnection) {
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
func (c *connections) shutdown() {
	log.Info("Chapar - ShutDown - Connections saving proccess begin...")
	log.Info("Chapar - ShutDown - Number of active connections:", len(c.poolByThingID))
	for _, conn := range c.poolByThingID {
		go conn.saveConn()
	}
	log.Info("Chapar - ShutDown - Connections saving proccess end now!")
}
