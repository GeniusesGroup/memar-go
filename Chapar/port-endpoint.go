/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"bytes"
	"sync"

	"../giti"
)

type portEndPoint struct {
	portNumber byte
	mux        *Multiplexer
}

// Send send frame sync that block sender until frame send and sure received successfully!
func (pe *portEndPoint) Send(frame []byte) (err giti.Error) {
	err = pe.mux.physicalConnection.Send(frame)
	return
}

// SendAsync send frame async that will not block sender but frame might not send successfully
// if port occur problem after port queued frame and caller can't notify about this situation!
func (pm *portEndPoint) SendAsync(frame []byte) (err giti.Error) {
	err = pm.mux.physicalConnection.SendAsync(frame)
	return
}

// Receive read the frame and call upper layer handler!
func (pe *portEndPoint) Receive(frame []byte) {
	var nexHeaderID = GetNextHeader(frame)
	var path = GetPath(frame)
	var payload = GetPayload(frame)

	var conn = pe.mux.connections.getConnectionByPath(path)
	if conn == nil {
		conn = pe.mux.connections.newConnection(pe, frame)
	} else if !bytes.Equal(conn.pathFromPeer.Get(), path) {
		// TODO::: receive frame on alternative path, Any action needed??
	}

	pe.mux.getTransportHandler(nextHeaderID).Receive(conn, payload)
	return
}
