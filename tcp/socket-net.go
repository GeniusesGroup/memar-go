/* For license and copyright information please see LEGAL file in repository */

package tcp

import (
	"net"
	"syscall"
	"time"
)

/*
********** net.Conn interface **********
 */

func (s *Socket) ok() bool { return s != nil }

// Read is not concurrent safe. Use just by one goroutine.
func (s *Socket) Read(b []byte) (n int, err error) {
	if !s.recv.buf.Full() {
		select {
		case <-s.readDeadline.C:
			// break
		case <-s.recv.pushFlag:
			// break
		}
		// TODO::: receive push flag and fin flag??
	}
	n, err = s.recv.buf.Read(b)
	return
}
func (s *Socket) Write(b []byte) (n int, err error) {
	var sendNumber int
	for len(b) != 0 {
		sendNumber, err = s.sendPayload(b)
		if err != nil {
			return
		}
		n += sendNumber
		b = b[sendNumber:]
	}
	return
}
func (s *Socket) Close() (err error) {
	if !s.ok() {
		return syscall.EINVAL
	}
	return
}
func (s *Socket) LocalAddr() net.Addr  { return nil }
func (s *Socket) RemoteAddr() net.Addr { return nil }
func (s *Socket) SetDeadline(t time.Time) (err error) {
	if !s.ok() {
		return syscall.EINVAL
	}
	err = s.SetReadDeadline(t)
	if err != nil {
		return
	}
	err = s.SetWriteDeadline(t)
	return
}
func (s *Socket) SetReadDeadline(t time.Time) (err error) {
	if !s.ok() {
		return syscall.EINVAL
	}
	s.readDeadline = t
	// TODO:::
	return
}
func (s *Socket) SetWriteDeadline(t time.Time) (err error) {
	if !s.ok() {
		return syscall.EINVAL
	}
	s.writeDeadline = t
	// TODO:::
	return
}
