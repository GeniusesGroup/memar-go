/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"memar/protocol"
)

type port struct {
	portNumber         byte
	mux                *Multiplexer
	physicalConnection protocol.OSI_Physical
}

func (p *port) Init(portNumber byte, mux *Multiplexer, physicalConnection protocol.OSI_Physical) {
	p.portNumber = portNumber
	p.mux = mux
	p.physicalConnection = physicalConnection
}

func (p *port) PortNumber() byte { return p.portNumber }

// Send send packet sync that block sender until packet send and sure received successfully.
func (p *port) Send(packet []byte) (err protocol.Error) {
	err = p.physicalConnection.Send(packet)
	return
}

// Receive read the frame and call upper layer handler.
func (p *port) Receive(packet []byte) (err protocol.Error) {
	err = p.mux.Receive(nil, packet)
	return
}
