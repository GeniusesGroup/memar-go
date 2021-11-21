/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"../protocol"
)

// Multiplexer implement protocol.LinkMultiplexer interface
// Hardware implementation has one difference from Software(this) implementation:
// Software: Receive method call Receive method of desire port
// Hardware: Receive method will call Send method of desire port
type Multiplexer struct {
	// Ports store all available link port to other physical or logical devices!
	// 256 is max ports that Chapar protocol support directly in one hop!!
	ports [256]port
	// UpperHandlers store all upper layer handlers
	// 256 is max next header ID that Chapar protocol support!
	upperHandlers [256]protocol.NetworkTransportOSMultiplexer

	connections connections
}

// Init initializes new Multiplexer object otherwise panic will occur on un-registered port or handler call!
func (mux *Multiplexer) Init(pConnection protocol.NetworkPhysicalConnection) {
	// mux.physicalConnection = pConnection
	pConnection.RegisterLinkMultiplexer(mux)

	var i byte
	for i = 0; i <= 255; i++ {
		var p = port{
			portNumber: i,
			mux: mux,
			// physicalConnection: TODO:::
		}
		mux.registerPort(p)

		var nonUH = UpperHandlerNonExist{headerID: i}
		mux.upperHandlers.RegisterHandler(nonUH)
	}

	mux.connections.init()
}

// RegisterTransportHandler registers new port on given ports pool!
func (mux *Multiplexer) RegisterTransportHandler(tm protocol.NetworkTransportOSMultiplexer) {
	// TODO::: check handler exist already and warn user??
	mux.upperHandlers[tm.HeaderID()] = tm
}

// UnRegisterTransportHandler delete the port by given port number on given ports pool!
func (mux *Multiplexer) UnRegisterTransportHandler(tm protocol.NetworkTransportOSMultiplexer) {
	var nonUH = UpperHandlerNonExist{headerID: tm.HeaderID()}
	mux.upperHandlers.RegisterHandler(nonUH)
}

// removeTransportHandler delete the port by given port number on given ports pool!
func (mux *Multiplexer) removeTransportHandler(handlerID byte) {
	var nonUH = UpperHandlerNonExist{headerID: handlerID}
	mux.upperHandlers.RegisterHandler(nonUH)
}

func (mux *Multiplexer) getTransportHandler(id byte) protocol.NetworkTransportOSMultiplexer {
	return mux.upperHandlers[id]
}

// Send send the payload to all ports async!
func (mux *Multiplexer) Send(frame []byte) (err protocol.Error) {
	err = CheckFrame(frame)
	if err != nil {
		// TODO::: ???
		return
	}

	if IsBroadcastFrame(frame) {
		err = mux.sendBroadcast(frame)
	} else {
		var portNum byte = GetNextPortNum(frame)
		err = mux.getPort(portNum).Send(frame)
	}
	return
}

// SendBroadcast send the payload to all ports async!
func (mux *Multiplexer) SendBroadcast(nexHeaderID protocol.NetworkLinkNextHeaderID, payload protocol.Codec) (err protocol.Error) {
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
	err = mux.sendBroadcast(frame)
	return
}

// send the frame to all ports as BroadcastFrame!
func (mux *Multiplexer) sendBroadcast(frame []byte) (err protocol.Error) {
	// send the frame to all ports as BroadcastFrame!
	var portNum byte
	for portNum = 0; portNum <= 255; portNum++ {
		err = mux.getPort(portNum).Send(frame)
	}
}

// Receive handles income frame to ports!
func (mux *Multiplexer) Receive(frame []byte) {
	var err = CheckFrame(frame)
	if err != nil {
		// TODO::: ???
		return
	}

	IncrementNextHop(frame, pm.portNumber)

	if IsBroadcastFrame(frame) {
		err = mux.sendBroadcast(frame)
	} else {
		var portNum byte = GetNextPortNum(frame)
		mux.getPort(portNum).Receive(frame)
	}
}

// Shutdown ready the connection pools to shutdown!!
func (mux *Multiplexer) Shutdown() {
	mux.connections.shutdown()
}

func (mux *Multiplexer) getPort(id byte) port { return mux.ports[id] }

func (mux *Multiplexer) registerPort(p port) {
	// TODO::: check port exist already and warn user??
	mux.ports[p.PortNumber()] = p
}

func (mux *Multiplexer) unRegisterPort(p port) {
	mux.removePort(p.PortNumber())
}

func (mux *Multiplexer) removePort(portNumber byte) {
	// mux.ports[portNumber].physicalConnection = TODO:::
}
