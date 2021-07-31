/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"../giti"
)

// Multiplexer implement giti.LinkMultiplexer interface
type Multiplexer struct {
	physicalConnection giti.NetworkPhysicalConnection

	// Ports store all available link port to other physical or logical devices!
	// 256 is max ports that Chapar protocol support directly in one hop!!
	ports [256]port
	// UpperHandlers store all upper layer handlers
	// 256 is max next header ID that Chapar protocol support!
	upperHandlers [256]giti.NetworkTransportOSMultiplexer

	connections connections
}

// Init initializes new Multiplexer object otherwise panic will occur on un-registered port or handler call!
func (mux *Multiplexer) Init(pConnection giti.NetworkPhysicalConnection) {
	mux.physicalConnection = pConnection
	pConnection.RegisterLinkMultiplexer(mux)

	var i byte
	for i = 0; i <= 255; i++ {
		var pne = PortNonExist{portNumber: i}
		mux.ports.RegisterPort(ports)

		var nonUH = UpperHandlerNonExist{headerID: i}
		mux.upperHandlers.RegisterHandler(nonUH)
	}

	mux.connections.init()
}

// RegisterTransportHandler registers new port on given ports pool!
func (mux *Multiplexer) RegisterTransportHandler(tm giti.NetworkTransportOSMultiplexer) {
	// TODO::: check handler exist already and warn user??
	mux.upperHandlers[tm.HeaderID()] = tm
}

// UnRegisterTransportHandler delete the port by given port number on given ports pool!
func (mux *Multiplexer) UnRegisterTransportHandler(tm giti.NetworkTransportOSMultiplexer) {
	var nonUH = UpperHandlerNonExist{headerID: tm.HeaderID()}
	mux.upperHandlers.RegisterHandler(nonUH)
}

// removeTransportHandler delete the port by given port number on given ports pool!
func (mux *Multiplexer) removeTransportHandler(handlerID byte) {
	var nonUH = UpperHandlerNonExist{headerID: handlerID}
	mux.upperHandlers.RegisterHandler(nonUH)
}

func (mux *Multiplexer) getTransportHandler(id byte) giti.NetworkTransportOSMultiplexer {
	return mux.upperHandlers[id]
}

// SendBroadcastAsync send the payload to all ports async!
func (mux *Multiplexer) SendBroadcastAsync(nexHeaderID giti.NetworkLinkNextHeaderID, payload giti.Codec) (err giti.Error) {
	var payloadLen int = payload.Len()
	if payloadLen > maxBroadcastPayloadLen {
		return ErrMTU
	}

	var pathLen byte = maxHopCount
	var payloadLoc int = 3 + int(pathLen)
	var frameLength int = payloadLoc + payloadLen
	var frame = make([]byte, frameLength)

	SetHopCount(frame, broadcastHopCount)
	SetNextHeader(frame, byte(nexHeaderID))
	payload.MarshalTo(frame[payloadLoc:])

	// send the frame to all ports as BroadcastFrame!
	var i byte
	for i = 0; i <= 255; i++ {
		err = mux.getPort(i).SendAsync(frame)
	}
	return
}

// Receive handles income frame to ports!
func (mux *Multiplexer) Receive(frame []byte) {
	var err = CheckFrame(frame)
	if err != nil {
		// TODO::: ???
		return
	}
	mux.ports[GetNextHop(frame)].Receive(frame)
}

// Shutdown ready the connection pools to shutdown!!
func (mux *Multiplexer) Shutdown() {
	mux.connections.shutdown()
}

func (mux *Multiplexer) getPort(id byte) port { return mux.ports[id] }

func (mux *Multiplexer) registerPort(port port) {
	// TODO::: check port exist already and warn user??
	mux.ports[p.PortNumber()] = port
	port.RegisterMultiplexer(mux)
}

func (mux *Multiplexer) unRegisterPort(port port) {
	var pne = PortNonExist{portNumber: port.PortNumber()}
	mux.registerPort(pne)
	pne.RegisterMultiplexer(mux)
}

func (mux *Multiplexer) removePort(portNumber byte) {
	var pne = PortNonExist{portNumber: portNumber}
	mux.registerPort(pne)
	pne.RegisterMultiplexer(mux)
}
