/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"sync"

	er "../error"
	"../giti"
)

// PortMultiplexer implement giti.LinkPort as multiplexer port
type PortMultiplexer struct {
	physicalConnection giti.PhysicalConnection
	portNumber         byte

	queue [][]byte

	statusLock    sync.RWMutex
	sendStatus    uint8
	receiveStatus uint8

	mux giti.LinkMultiplexer
}

// RegisterMultiplexer register given multiplexer to the port for further usage!
func (pm *PortMultiplexer) RegisterMultiplexer(lm giti.LinkMultiplexer) {
	pm.mux = lm
}

// SetPortNumber set port number for given port
func (pm *PortMultiplexer) SetPortNumber(num byte) { pm.portNumber = num }

// PortNumber return port number of exiting port
func (pm *PortMultiplexer) PortNumber() (num byte) { return pm.portNumber }

// Send send frame sync that block sender until frame send and sure received successfully!
func (pm *PortMultiplexer) Send(frame []byte) (err *er.Error) {
	err = pm.physicalConnection.Send(frame)
	return
}

// SendAsync send frame async that will not block sender but frame might not send successfully
// if port occur problem after port queued frame and caller can't notify about this situation!
func (pm *PortMultiplexer) SendAsync(frame []byte) (err *er.Error) {
	err = pm.physicalConnection.SendAsync(frame)
	return
}

// Receive
func (pm *PortMultiplexer) Receive(frame []byte) {
	IncrementNextHop(frame, pm.PortNumber())

	if IsBroadcastFrame(frame) {
		// send the frame to all ports as BroadcastFrame!
		var i byte
		for i = 0; i <= 255; i++ {
			// Send frame to desire port interface! Usually frame will put in queue!
			pm.mux.GetPort(i).SendAsync(frame)
		}
	} else {
		// send the frame to the specific port as UnicastFrame!
		var portID = GetNextHop(frame)
		pm.mux.GetPort(portID).SendAsync(frame)
	}
}
