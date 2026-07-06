/* For license and copyright information please see the LEGAL file in the code repository */

package std

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
func (sk *Socket) Read(b []byte) (n int, err error) {
	err = sk.Check()
	if err != nil {
		return
	}

	if !s.Full() {
		err = s.blockInSelect()
	}
	// TODO::: check and wrap above error?
	n, err = sk.recv.buf.Read(b)
	return
}
func (sk *Socket) Write(b []byte) (n int, err error) {
	n, err = sk.Unmarshal(b)
	return
}
func (sk *Socket) Close() (err error) {
	err = sk.Check()
	if err != nil {
		return
	}

	err = sk.Close()
	return
}
func (sk *Socket) LocalAddr() net.Addr {
	var err = sk.Check()
	if err != nil {
		return nil
	}
	return &net.TCPAddr{
		// IP:   net.IP(sk.connection.LocalAddr()),
		// Port: int(sk.sourcePort),
	}
}
func (sk *Socket) RemoteAddr() net.Addr {
	var err = sk.Check()
	if err != nil {
		return nil
	}
	return &net.TCPAddr{
		// IP:   net.IP(sk.connection.RemoteAddr()),
		// Port: int(sk.destinationPort),
	}
}
func (sk *Socket) SetDeadline(t time.Time) (err error) {
	var d = untilTo(t)
	sk.SetTimeout(d)
	return
}
func (sk *Socket) SetReadDeadline(t time.Time) (err error) {
	var d = untilTo(t)
	err = sk.SetReadTimeout(d)
	return
}
func (sk *Socket) SetWriteDeadline(t time.Time) (err error) {
	var d = untilTo(t)
	err = sk.SetWriteTimeout(d)
	return
}

//memar:impl std/net.TCPConn
func (sk *Socket) CloseRead() (err error)                         { return }
func (sk *Socket) CloseWrite() (err error)                        { return }
func (sk *Socket) SetLinger(sec int) (err error)                  { return }
func (sk *Socket) SetKeepAlive(keepalive bool) (err error)        { return }
func (sk *Socket) SetKeepAlivePeriod(d time.Duration) (err error) { return }
func (sk *Socket) SetNoDelay(noDelay bool) (err error)            { return }

func untilTo(t time.Time) (d protocol.Duration) {
	if !t.IsZero() {
		d = protocol.Duration(time.Until(t))
		if d == 0 {
			d = -1 // don't confuse deadline right now with no deadline
		}
	}
	return
}
