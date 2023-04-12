/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"libgo/protocol"
)

type port struct {
	portNumber         byte
	mux                protocol.NetworkLink_Multiplexer
	physicalConnection protocol.NetworkPhysical_Connection
}

func (p *port) Init(portNumber byte, mux protocol.NetworkLink_Multiplexer, physicalConnection protocol.NetworkPhysical_Connection) {
	p.portNumber = portNumber
	p.mux = mux
	p.physicalConnection = physicalConnection
}

func (p *port) PortNumber() byte { return p.portNumber }

// Send send frame sync that block sender until frame send and sure received successfully.
func (p *port) Send(frame []byte) (err protocol.Error) {
	err = p.physicalConnection.Send(frame)
	return
}

// Receive read the frame and call upper layer handler.
func (p *port) Receive(frame []byte) (err protocol.Error) {
	err = p.mux.Receive(p.physicalConnection, frame)
	return
}
