/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
)

// BridgePort use by physical port as its physicalConnection(protocol.OSI_Physical)
// It will use only when two switcher can't wire on same port number
type BridgePort struct {
	sidePortNumber byte
	portNumber     byte
	physicalMux    *Multiplexer
	net_p.OSI_Physical
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (p *BridgePort) Init(sidePortNumber, portNumber byte, physicalMux *Multiplexer, physicalConnection net_p.OSI_Physical) (err error_p.Error) {
	p.sidePortNumber = sidePortNumber
	p.portNumber = portNumber
	p.physicalMux = physicalMux
	p.OSI_Physical = physicalConnection
	return
}
func (p *BridgePort) Reinit() (err error_p.Error) { return }
func (p *BridgePort) Deinit() (err error_p.Error) { return }

//memar:impl memar/protocol.Quiddity
func (p *BridgePort) Name() string         { return "bridge" }
func (p *BridgePort) Abbreviation() string { return "" }
func (p *BridgePort) Aliases() []string    { return nil }

//memar:impl memar/protocol.OSI_Physical
func (p *BridgePort) Send(packet net_p.Packet) (err error_p.Error) {
	// TODO::: first find chapar Frame in the given packet

	var f = Frame(packet)
	var lastHop = f.IncrementNextHop(p.portNumber)
	if lastHop {
		// err = &
		return
	}
	err = p.OSI_Physical.Send(packet)
	return
}
func (p *BridgePort) Receive(frame net_p.Packet) (err error_p.Error) {
	// TODO::: Need to check FrameType first??

	var f = Frame(frame)
	var lastHop = f.IncrementNextHop(p.sidePortNumber)
	if lastHop {
		// err = &
		return
	}
	p.physicalMux.Receive(nil, frame)
	return
}
