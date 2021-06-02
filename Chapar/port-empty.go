/* For license and copyright information please see LEGAL file in repository */

package chapar

import "../giti"

// PortNonExist use for default and empty switch port due to non of ports can be nil!
type PortNonExist struct {
	portNumber byte
}

// RegisterMultiplexer register given multiplexer to the port for further usage!
func (pne *PortNonExist) RegisterMultiplexer(lm giti.LinkMultiplexer) {}

// SetPortNumber set port number for given port
func (pne *PortNonExist) SetPortNumber(num byte) { pne.portNumber = num }

// PortNumber return port number of exiting port
func (pne *PortNonExist) PortNumber() (num byte) { return pne.portNumber }

// Send send frame sync that block sender until frame send and sure received successfully!
func (pne *PortNonExist) Send(frame []byte) {}

// Receive
func (pne *PortNonExist) Receive(frame []byte) {}
