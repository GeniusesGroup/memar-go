/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"sync"

	er "../error"
	"../protocol"
)

// connections store pools of connection to retrieve in many ways!
// TODO::: add analytic methods
type connections struct {
	mutex           sync.Mutex
	poolByPath      map[string]protocol.NetworkLinkConnection // It is optimize to convert byte slice to string as key in map: https://github.com/golang/go/commit/f5f5a8b6209f84961687d993b93ea0d397f5d5bf
	poolByThingID   map[[32]byte]protocol.NetworkLinkConnection
	poolByBlackList map[[32]byte]protocol.NetworkLinkConnection
}

func (c *connections) init() {
	if c.poolByPath == nil {
		c.poolByPath = make(map[string]protocol.NetworkLinkConnection, 16384)
	}
	if c.poolByThingID == nil {
		c.poolByThingID = make(map[[32]byte]protocol.NetworkLinkConnection, 16384)
	}
	if c.poolByBlackList == nil {
		c.poolByBlackList = make(map[[32]byte]protocol.NetworkLinkConnection)
	}
}

func (c *connections) newConnection(port *port, frame []byte) (conn protocol.NetworkLinkConnection) {
	conn = &Connection{
		port: port,
	}
	conn.init(frame)

	// TODO::: get ThingID from peer or func args??

	c.registerConnection(conn)
	return
}

func (c *connections) establishConnectionByPath(path []byte) (conn *Connection, err protocol.Error) {
	return
}

func (c *connections) establishConnectionByThingID(thingID [32]byte) (conn *Connection, err protocol.Error) {
	return
}

// GetConnectionByPath get a connection by its path from connections pool!!
func (c *connections) getConnectionByPath(path []byte) (conn protocol.NetworkLinkConnection, err protocol.Error) {
	conn = c.poolByPath[string(path)]
	if conn == nil {
		conn = &Connection{
			path: path,
		}
		err = conn.GetLastByPath()
	}
	return
}

func (c *connections) getConnectionsByThingID(thingID [32]byte) (conn protocol.NetworkLinkConnection, err protocol.Error) {
	conn = c.poolByThingID[thingID]
	if conn == nil {
		conn = &Connection{
			thingID: thingID,
		}
		err = conn.GetLastByThingID()
	}
	return
}

func (c *connections) registerConnection(conn protocol.NetworkLinkConnection) {
	c.mutex.Lock()
	c.poolByPath[conn.Path.GetAsString()] = conn
	c.poolByThingID[conn.ThingID] = conn
	c.mutex.Unlock()
}

func (c *connections) registerNewPathForConnection(conn protocol.NetworkLinkConnection, alternativePath []byte) {
	conn.setAlternativePath(alternativePath)

	c.mutex.Lock()
	c.poolByPath[string(alternativePath)] = conn
	c.mutex.Unlock()
}

func (c *connections) closeConnection(conn protocol.NetworkLinkConnection) {
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
	protocol.App.LogInfo("Chapar - ShutDown - Connections saving proccess begin...")
	protocol.App.LogInfo("Chapar - ShutDown - Number of active connections:", len(c.poolByThingID))
	for _, conn := range c.poolByThingID {
		conn.saveConn()
	}
	protocol.App.LogInfo("Chapar - ShutDown - Connections saving proccess end now!")
}
