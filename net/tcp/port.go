/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/protocol"
)

type port struct {
	sourcePort      PortNumber // local
	destinationPort PortNumber // remote
}

// Init use to initialize the stream after allocation in both server or client
//
//memar:impl memar/protocol.ObjectLifeCycle
func (p *port) Init(sourcePort, destinationPort PortNumber) (err protocol.Error) {
	p.sourcePort = sourcePort
	p.destinationPort = destinationPort
	return
}
func (p *port) Reinit(sourcePort, destinationPort PortNumber) (err protocol.Error) {
	p.sourcePort = sourcePort
	p.destinationPort = destinationPort
	return
}
func (p *port) Deinit() (err protocol.Error) { return }

func (p *port) SourcePort() PortNumber      { return p.sourcePort }
func (p *port) DestinationPort() PortNumber { return p.destinationPort }
