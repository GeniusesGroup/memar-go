/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"libgo/protocol"
)

// BridgePort use by physical port as its physicalConnection(protocol.NetworkPhysical_Connection)
// It will use only when two switcher can't wire on same port number
type BridgePort struct {
	sidePortNumber     byte
	portNumber         byte
	physicalMux        protocol.NetworkLink_Multiplexer
	physicalConnection protocol.NetworkPhysical_Connection
}

//libgo:impl /libgo/protocol.NetworkPhysical_Connection
func (p *BridgePort) Init(sidePortNumber, portNumber byte, physicalMux protocol.NetworkLink_Multiplexer, physicalConnection protocol.NetworkPhysical_Connection) {
	p.sidePortNumber = sidePortNumber
	p.portNumber = portNumber
	p.physicalMux = physicalMux
	p.physicalConnection = physicalConnection
}
func (p *BridgePort) Deinit()                                                            {}
func (p *BridgePort) RegisterLinkMultiplexer(linkMux protocol.NetworkLink_Multiplexer)   {}
func (p *BridgePort) UnRegisterLinkMultiplexer(linkMux protocol.NetworkLink_Multiplexer) {}
func (p *BridgePort) Send(frame []byte) (err protocol.Error) {
	var f = Frame(frame)
	var lastHop = f.IncrementNextHop(p.portNumber)
	if lastHop {
		// err = &
		return
	}
	err = p.physicalConnection.Send(frame)
	return
}
func (p *BridgePort) Receive(frame []byte) (err protocol.Error) {
	var f = Frame(frame)
	var lastHop = f.IncrementNextHop(p.sidePortNumber)
	if lastHop {
		// err = &
		return
	}
	p.physicalMux.Receive(p.physicalConnection, frame)
	return
}
