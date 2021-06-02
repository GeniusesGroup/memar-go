/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"bytes"
	"sync"

	er "../error"
	"../giti"
)

// PortOS implement giti.LinkPort as endpoint port that handle frame by upper layer data!
type PortEndPoint struct {
	physicalConnection giti.PhysicalConnection
	portNumber         byte

	queue [][]byte

	statusLock    sync.RWMutex
	sendStatus    uint8
	receiveStatus uint8

	mux giti.LinkMultiplexer
}

// RegisterMultiplexer register given multiplexer to the port for further usage!
func (pe *PortEndPoint) RegisterMultiplexer(lm giti.LinkMultiplexer) {
	pe.mux = lm
}

// SetPortNumber set port number for given port
func (pe *PortEndPoint) SetPortNumber(num byte) { pe.portNumber = num }

// PortNumber return port number of exiting port
func (pe *PortEndPoint) PortNumber() (num byte) { return }

// Send send frame sync that block sender until frame send and sure received successfully!
func (pe *PortEndPoint) Send(frame []byte) (err *er.Error) {
	err = pe.physicalConnection.Send(frame)
	return
}

// SendAsync send frame async that will not block sender but frame might not send successfully
// if port occur problem after port queued frame and caller can't notify about this situation!
func (pm *PortEndPoint) SendAsync(frame []byte) (err *er.Error) {
	err = pm.physicalConnection.SendAsync(frame)
	return
}

// Receive use for
func (pe *PortEndPoint) Receive(frame []byte) {
	var nexHeaderID = GetNextHeader(frame)
	var path = GetPath(frame)
	var payload = GetPayload(frame)

	var conn = pe.mux.GetConnectionByPath(path)
	if conn == nil {
		conn = pe.mux.NewConnection(pe.physicalConnection, path)
	} else if !bytes.Equal(conn.path, path) {
		// TODO:::
	}

	pe.mux.GetUpperHandler(nextHeaderID).Receive(conn, payload)
	return

}
