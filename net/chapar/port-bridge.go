/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"libgo/protocol"
)

// BridgePort use by physical port as its physicalConnection(protocol.NetworkInterface)
// It will use only when two switcher can't wire on same port number
type BridgePort struct {
	sidePortNumber byte
	portNumber     byte
	physicalMux    *Multiplexer
	protocol.NetworkInterface
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (p *BridgePort) Init(sidePortNumber, portNumber byte, physicalMux *Multiplexer, physicalConnection protocol.NetworkInterface) (err protocol.Error) {
	p.sidePortNumber = sidePortNumber
	p.portNumber = portNumber
	p.physicalMux = physicalMux
	p.NetworkInterface = physicalConnection
	return
}
func (p *BridgePort) Reinit() (err protocol.Error) { return }
func (p *BridgePort) Deinit() (err protocol.Error) { return }

//libgo:impl libgo/protocol.Quiddity
func (p *BridgePort) Name() string         { return "bridge" }
func (p *BridgePort) Abbreviation() string { return "" }
func (p *BridgePort) Aliases() []string    { return nil }

//libgo:impl libgo/protocol.NetworkInterface
func (p *BridgePort) Send(frame []byte) (err protocol.Error) {
	var f = Frame(frame)
	var lastHop = f.IncrementNextHop(p.portNumber)
	if lastHop {
		// err = &
		return
	}
	err = p.NetworkInterface.Send(frame)
	return
}
func (p *BridgePort) Receive(frame []byte) (err protocol.Error) {
	var f = Frame(frame)
	var lastHop = f.IncrementNextHop(p.sidePortNumber)
	if lastHop {
		// err = &
		return
	}
	p.physicalMux.Receive(nil, frame)
	return
}
