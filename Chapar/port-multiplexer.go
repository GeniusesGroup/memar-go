/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"../giti"
)

// PortMultiplexer implement port interface as multiplexer port
// Usually implement as hardware circuit! But also can use to make virtual network in OS virtualization!
type PortMultiplexer struct {
	portNumber byte
	mux        *Multiplexer
}

// Send send frame sync that block sender until frame send and sure received successfully!
func (pm *PortMultiplexer) Send(frame []byte) (err giti.Error) {
	err = pm.mux.physicalConnection.Send(frame)
	return
}

// SendAsync send frame async that will not block sender but frame might not send successfully
// if port occur problem after port queued frame and caller can't notify about this situation!
func (pm *PortMultiplexer) SendAsync(frame []byte) (err giti.Error) {
	err = pm.mux.physicalConnection.SendAsync(frame)
	return
}

// Receive
func (pm *PortMultiplexer) Receive(frame []byte) {
	IncrementNextHop(frame, pm.portNumber)

	if IsBroadcastFrame(frame) {
		// send the frame to all ports as BroadcastFrame!
		var i byte
		for i = 0; i <= 255; i++ {
			// Send frame to desire port interface! Usually frame will put in queue!
			pm.mux.getPort(i).SendAsync(frame)
		}
	} else {
		// send the frame to the specific port as UnicastFrame!
		var portID = GetNextHop(frame)
		pm.mux.getPort(portID).SendAsync(frame)
	}
}
