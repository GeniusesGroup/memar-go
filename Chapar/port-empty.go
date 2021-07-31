/* For license and copyright information please see LEGAL file in repository */

package chapar

import "../giti"

// PortNonExist use for default and empty switch port due to non of ports can be nil!
type PortNonExist struct {
	portNumber byte
}

// Send send frame sync that block sender until frame send and sure received successfully!
func (pne *PortNonExist) Send(frame []byte) {}

// SendAsync send frame async that will not block sender but frame might not send successfully
// if port occur problem after port queued frame and caller can't notify about this situation!
func (pne *PortNonExist) SendAsync(frame []byte) {}

// Receive
func (pne *PortNonExist) Receive(frame []byte) {}
