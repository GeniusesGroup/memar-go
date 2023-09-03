/* For license and copyright information please see the LEGAL file in the code repository */

package net

// TODO::: Is it ok to import "net" & "time" package? move this file to internal package? or build tag?
import (
	"net"
	"time"

	"memar/protocol"
)

// Non of below methods are concurrent safe. Use just by one goroutine.
// TODO::: concurrency safe??
//
//memar:impl std/net.Conn
func (s *Socket) Read(b []byte) (n int, err error) {
	err = s.Check()
	if err != nil {
		return
	}

	if !s.Full() {
		err = s.blockInSelect()
	}
	// TODO::: check and wrap above error?
	n, err = s.recv.buf.Read(b)
	return
}
func (s *Socket) Write(b []byte) (n int, err error) {
	n, err = s.Unmarshal(b)
	return
}
func (s *Socket) Close() (err error) {
	err = s.Check()
	if err != nil {
		return
	}

	err = s.Close()
	return
}
func (s *Socket) LocalAddr() net.Addr {
	var err = s.Check()
	if err != nil {
		return nil
	}
	return &net.TCPAddr{
		// IP:   net.IP(s.connection.LocalAddr()),
		// Port: int(s.sourcePort),
	}
}
func (s *Socket) RemoteAddr() net.Addr {
	var err = s.Check()
	if err != nil {
		return nil
	}
	return &net.TCPAddr{
		// IP:   net.IP(s.connection.RemoteAddr()),
		// Port: int(s.destinationPort),
	}
}
func (s *Socket) SetDeadline(t time.Time) (err error) {
	var d = untilTo(t)
	s.SetTimeout(d)
	return
}
func (s *Socket) SetReadDeadline(t time.Time) (err error) {
	var d = untilTo(t)
	err = s.SetReadTimeout(d)
	return
}
func (s *Socket) SetWriteDeadline(t time.Time) (err error) {
	var d = untilTo(t)
	err = s.SetWriteTimeout(d)
	return
}

//memar:impl std/net.TCPConn
func (s *Socket) CloseRead() (err error)                         { return }
func (s *Socket) CloseWrite() (err error)                        { return }
func (s *Socket) SetLinger(sec int) (err error)                  { return }
func (s *Socket) SetKeepAlive(keepalive bool) (err error)        { return }
func (s *Socket) SetKeepAlivePeriod(d time.Duration) (err error) { return }
func (s *Socket) SetNoDelay(noDelay bool) (err error)            { return }

func untilTo(t time.Time) (d protocol.Duration) {
	if !t.IsZero() {
		d = protocol.Duration(time.Until(t))
		if d == 0 {
			d = -1 // don't confuse deadline right now with no deadline
		}
	}
	return
}
