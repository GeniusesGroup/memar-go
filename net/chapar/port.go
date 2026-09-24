/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
)

type port struct {
	portNumber         byte
	mux                *Multiplexer
	physicalConnection net_p.OSI_Physical
}

func (p *port) Init(portNumber byte, mux *Multiplexer, physicalConnection net_p.OSI_Physical) {
	p.portNumber = portNumber
	p.mux = mux
	p.physicalConnection = physicalConnection
}

func (p *port) PortNumber() byte { return p.portNumber }

// Send send packet sync that block sender until packet send and sure received successfully.
func (p *port) Send(packet net_p.Packet) (err error_p.Error) {
	err = p.physicalConnection.Send(packet)
	return
}

// Receive read the frame and call upper layer handler.
func (p *port) Receive(packet net_p.Packet) (err error_p.Error) {
	err = p.mux.Receive(nil, packet)
	return
}
