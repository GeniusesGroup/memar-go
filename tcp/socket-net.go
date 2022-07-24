/* For license and copyright information please see LEGAL file in repository */

package tcp

import (
	"net"
	"time"
)

/*
********** net.Conn interface **********
// TODO::: concurrency safe??
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
	n, err = s.Unmarshal(b)
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
func (s *Socket) LocalAddr() net.Addr {
	var err = s.checkSocket()
	if err != nil {
		return nil
	}
	return &net.TCPAddr{
		IP:   net.IP(s.connection.LocalAddr()),
		Port: s.sourcePort,
	}
}
func (s *Socket) RemoteAddr() net.Addr {
	var err = s.checkSocket()
	if err != nil {
		return nil
	}
	return &net.TCPAddr{
		IP:   net.IP(s.connection.RemoteAddr()),
		Port: s.destinationPort,
	}
}
func (s *Socket) SetDeadline(t time.Time) (err error) {
	var d = getDuration(t)
	s.SetTimeout(d)
	return
}
func (s *Socket) SetReadDeadline(t time.Time) (err error) {
	var d = getDuration(t)
	err = s.SetReadTimeout(d)
	return
}
func (s *Socket) SetWriteDeadline(t time.Time) (err error) {
	var d = getDuration(t)
	err = s.SetWriteTimeout(d)
	return
}
