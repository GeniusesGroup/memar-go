/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"bytes"

	"../protocol"
)

type port struct {
	portNumber         byte
	mux                *Multiplexer
	physicalConnection protocol.NetworkPhysicalConnection
}

func (p *port) PortNumber() byte { return p.portNumber }

// Send send frame sync that block sender until frame send and sure received successfully!
// transmitting that must be non blocking and queue frames for congestion situations!
// A situation might be occur that a port available when a frame queued but when the time to send is come, the port broken and sender don't know about this!
func (p *port) Send(frame []byte) (err protocol.Error) {
	err = p.physicalConnection.Send(frame)
	return
}

// SendAsync send frame async that will not block sender but frame might not send successfully
// if port occur problem after port queued frame and caller can't notify about this situation!
func (pm *port) SendAsync(frame []byte) (err protocol.Error) {
	err = pm.physicalConnection.SendAsync(frame)
	return
}

// Receive read the frame and call upper layer handler!
func (p *port) Receive(frame []byte) {
	var nexHeaderID = GetNextHeader(frame)
	var path = GetPath(frame)
	var payload = GetPayload(frame)

	var conn = p.mux.connections.getConnectionByPath(path)
	if conn == nil {
		conn = p.mux.connections.newConnection(p, frame)
	} else if !bytes.Equal(conn.pathFromPeer.Get(), path) {
		// TODO::: receive frame on alternative path, Any action needed??
	}

	p.mux.getTransportHandler(nextHeaderID).Receive(conn, payload)
	return
}
