/* For license and copyright information please see LEGAL file in repository */

package chapar

type port interface {
	// transmitting that must be non blocking and queue frames for congestion situations!
	// A situation might be occur that a port available when a frame queued but when the time to send is come, the port broken and sender don't know about this!
	Send(frame []byte) (err Error)

	SendAsync(frame []byte) (err Error)

	Receive(frame []byte)
}