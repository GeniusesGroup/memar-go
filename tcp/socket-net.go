/* For license and copyright information please see LEGAL file in repository */

package tcp

import (
	"net"
	"time"

	"../protocol"
)

/*
********** net.Conn interface **********
// TODO::: concurrency??
*/

// Read is not concurrent safe. Use just by one goroutine.
func (s *Socket) Read(b []byte) (n int, err error) {
	err = s.checkSocket()
	if err != nil {
		return
	}

	if !s.recv.buf.Full() {
		err = s.blockInSelect()
	}
	n, err = s.recv.buf.Read(b)
	return
}
func (s *Socket) Write(b []byte) (n int, err error) {
	err = s.checkSocket()
	if err != nil {
		return
	}

	for len(b) > 0 {
		select {
		case <-s.writeTimer.Signal():
			// err =
			return
		default:
			var sendNumber int
			sendNumber, err = s.sendPayload(b)
			if err != nil {
				return
			}
			n += sendNumber
			b = b[sendNumber:]
		}
	}
	return
}
func (s *Socket) Close() (err error) {
	err = s.checkSocket()
	if err != nil {
		return
	}

	err = s.close()
	return
}
func (s *Socket) LocalAddr() net.Addr  { return nil }
func (s *Socket) RemoteAddr() net.Addr { return nil }
func (s *Socket) SetDeadline(t time.Time) (err error) {
	var d protocol.Duration
	if !t.IsZero() {
		d = protocol.Duration(time.Until(t))
		if d == 0 {
			d = -1 // don't confuse deadline right now with no deadline
		}
	}
	s.SetTimeout(d)
	return
}
func (s *Socket) SetReadDeadline(t time.Time) (err error) {
	var d protocol.Duration
	if !t.IsZero() {
		d = protocol.Duration(time.Until(t))
		if d == 0 {
			d = -1 // don't confuse deadline right now with no deadline
		}
	}
	err = s.SetReadTimeout(d)
	return
}
func (s *Socket) SetWriteDeadline(t time.Time) (err error) {
	var d protocol.Duration
	if !t.IsZero() {
		d = protocol.Duration(time.Until(t))
		if d == 0 {
			d = -1 // don't confuse deadline right now with no deadline
		}
	}
	err = s.SetWriteTimeout(d)
	return
}
