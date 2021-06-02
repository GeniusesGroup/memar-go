/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	er "../error"
	"../giti"
)

// Multiplexer implement giti.LinkMultiplexer interface
type Multiplexer struct {
	// Ports store all available link port to other physical or logical devices!
	// 256 is max ports that Chapar protocol support directly!!
	ports [256]giti.LinkPort
	// UpperHandlers store all upper layer handlers
	// 256 is max next header ID that Chapar protocol support!
	upperHandlers [256]giti.TransportMultiplexer

	Connections
}

// Init initializes new Multiplexer object otherwise panic will occur on un-registered port or handler call!
func (mux *Multiplexer) Init() {
	var i byte
	for i = 0; i <= 255; i++ {
		var pne = PortNonExist{portNumber: i}
		mux.ports.RegisterPort(ports)

		var nonUH = UpperHandlerNonExist{headerID: i}
		mux.upperHandlers.RegisterHandler(nonUH)
	}
}

// Receive handles income frame to ports!
func (mux *Multiplexer) Receive(frame []byte) {
	mux.ports[GetNextHop(frame)].Receive(frame)
}

// EstablishNewConnectionByPath make new connection by peer path and initialize it!
func (mux *Multiplexer) EstablishNewConnectionByPath(path []byte) (conn *Connection, err *er.Error) {
	return
}

// EstablishNewConnectionByThingID make new connection by peer thing ID and initialize it!
func (mux *Multiplexer) EstablishNewConnectionByThingID(thingID [32]byte) (conn *Connection, err *er.Error) {
	return
}

/*
*******************
Ports as giti.LinkPort
*******************
*/

// Receive handles income frame to ports!
func (mux *Multiplexer) GetPort(id byte) giti.LinkPort {
	return mux.ports[id]
}

// RegisterPort registers new port on given ports pool!
func (mux *Multiplexer) RegisterPort(port giti.LinkPort) {
	// TODO::: check port exist already and warn user??
	mux.ports[p.PortNumber()] = port
	port.RegisterMultiplexer(mux)
}

// UnRegisterPort delete the port by given port number on given ports pool!
func (mux *Multiplexer) UnRegisterPort(port giti.LinkPort) {
	var pne = PortNonExist{portNumber: port.PortNumber()}
	mux.RegisterPort(pne)
	pne.RegisterMultiplexer(mux)
}

// RemovePort delete the port by given port number on given ports pool!
func (mux *Multiplexer) RemovePort(portNumber byte) {
	var pne = PortNonExist{portNumber: portNumber}
	mux.RegisterPort(pne)
	pne.RegisterMultiplexer(mux)
}

/*
*******************
UpperHandlers as giti.TransportMultiplexer
*******************
*/

// GetUpperHandler return upper layer handler
func (mux *Multiplexer) GetUpperHandler(id byte) giti.TransportMultiplexer {
	return mux.upperHandlers[id]
}

// RegisterHandler registers new port on given ports pool!
func (mux *Multiplexer) RegisterHandler(tm giti.TransportMultiplexer) {
	// TODO::: check handler exist already and warn user??
	mux.upperHandlers[tm.HeaderID()] = tm
}

// UnRegisterHandler delete the port by given port number on given ports pool!
func (mux *Multiplexer) UnRegisterHandler(tm giti.TransportMultiplexer) {
	var nonUH = UpperHandlerNonExist{headerID: tm.HeaderID()}
	mux.upperHandlers.RegisterHandler(nonUH)
}

// RemovePort delete the port by given port number on given ports pool!
func (mux *Multiplexer) RemoveHandler(handlerID byte) {
	var nonUH = UpperHandlerNonExist{headerID: handlerID}
	mux.upperHandlers.RegisterHandler(nonUH)
}
