/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"bytes"
	"memar/protocol"
)

// Multiplexer Hardware implementation has some difference from Software(this) implementation:
// - multiplexers send frames on the connection to other mux not call other functionality
// - they must provide some congestion mechanism like cache to prevent sender frame.
// - mux must have some mechanism to drop frames on destination port unavailability (congestion, ...)
type Multiplexer struct {
	portNumber byte

	// Ports store all available link port to other physical or logical devices.
	ports [defaultPortNumber]port

	connections Connections
}

//memar:impl memar/protocol.Network_Framer
func (mux *Multiplexer) FrameType() protocol.Network_FrameType {
	return protocol.Network_FrameType_Chapar
}

// Init initializes new Multiplexer object otherwise panic will occur on un-registered port or handler call.
//
//memar:impl memar/protocol.ObjectLifeCycle
func (mux *Multiplexer) Init(portNumber byte, pConnection protocol.OSI_Physical, connections Connections) {
	mux.portNumber = portNumber
	mux.connections = connections

	var i byte
	for i = 0; i <= 255; i++ {
		if i == portNumber {
			mux.ports[i].Init(portNumber, mux, pConnection)
		} else {
			// TODO::: get port info
			// mux.ports[i].Init(i, ?, ?)
		}
	}
}

// Deinit ready the connection pools to de-allocated.
func (mux *Multiplexer) Deinit() (err protocol.Error) {
	err = mux.connections.Deinit()
	return
}

// Send send the payload to all ports async.
func (mux *Multiplexer) Send(packet []byte) (err protocol.Error) {
	var f = Frame(packet)

	// err = f.CheckFrame()
	// if err != nil {
	// 	// TODO::: ???
	// 	return
	// }

	if f.IsBroadcastFrame() {
		err = mux.sendBroadcast(packet)
	} else {
		var portNum byte = f.NextPortNum()
		var port = mux.getPort(portNum)
		err = port.Send(packet)
	}
	return
}

// send the packet to all ports as BroadcastFrame!
func (mux *Multiplexer) sendBroadcast(packet []byte) (err protocol.Error) {
	// send the frame to all ports as BroadcastFrame!
	var portNum byte
	for portNum = 0; portNum <= 255; portNum++ {
		err = mux.getPort(portNum).Send(packet)
	}
	return
}

// Receive handles income frame to ports.
func (mux *Multiplexer) Receive(soc protocol.Socket, packet []byte) (err protocol.Error) {
	var f = Frame(packet)

	err = f.CheckFrame()
	if err != nil {
		// TODO::: ???
		return
	}

	var lastHop = f.IncrementNextHop(mux.portNumber)
	if lastHop {
		if !AcceptLastHop {
			err = &ErrNotAcceptLastHop
		} else {
			var path = f.Path()

			var conn *Connection
			conn, _ = mux.connections.GetConnectionByPath(path)
			if conn == nil {
				var newConn Connection
				newConn.Init(f, &mux.ports[mux.portNumber])
				conn = &newConn
				_ = mux.connections.RegisterConnection(conn)
			} else if !bytes.Equal(conn.pathFromPeer.Get(), path) {
				// TODO::: receive frame on alternative path, Any action needed??
			}

			// TODO::: Set conn to given soc(protocol.Socket) for others and response logic
		}
		return
	}

	if f.IsBroadcastFrame() {
		err = mux.sendBroadcast(packet)
	} else {
		var portNum byte = f.NextPortNum()
		err = mux.getPort(portNum).Receive(packet)
	}
	return
}

func (mux *Multiplexer) getPort(id byte) *port { return &mux.ports[id] }

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
