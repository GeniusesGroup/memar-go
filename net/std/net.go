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
<<<<<<< HEAD
func (s *Socket) Write(b []byte) (n int, err error) {
	n, err = s.Unmarshal(b)
	return
}
func (s *Socket) Close() (err error) {
	err = s.Check()
=======
func (sk *Socket) Write(b []byte) (n int, err error) {
	n, err = sk.Unmarshal(b)
	return
}
func (sk *Socket) Close() (err error) {
	err = sk.Check()
>>>>>>> 31bf680 ([net/std] some minor fixes)
	if err != nil {
		return
	}

<<<<<<< HEAD
	err = s.Close()
	return
}
func (s *Socket) LocalAddr() net.Addr {
	var err = s.Check()
=======
	err = sk.Close()
	return
}
func (sk *Socket) LocalAddr() net.Addr {
	var err = sk.Check()
>>>>>>> 31bf680 ([net/std] some minor fixes)
	if err != nil {
		return nil
	}
	return &net.TCPAddr{
<<<<<<< HEAD
		// IP:   net.IP(s.connection.LocalAddr()),
		// Port: int(s.sourcePort),
	}
}
func (s *Socket) RemoteAddr() net.Addr {
	var err = s.Check()
=======
		// IP:   net.IP(sk.connection.LocalAddr()),
		// Port: int(sk.sourcePort),
	}
}
func (sk *Socket) RemoteAddr() net.Addr {
	var err = sk.Check()
>>>>>>> 31bf680 ([net/std] some minor fixes)
	if err != nil {
		return nil
	}
	return &net.TCPAddr{
<<<<<<< HEAD
		// IP:   net.IP(s.connection.RemoteAddr()),
		// Port: int(s.destinationPort),
	}
}
func (s *Socket) SetDeadline(t time.Time) (err error) {
=======
		// IP:   net.IP(sk.connection.RemoteAddr()),
		// Port: int(sk.destinationPort),
	}
}
func (sk *Socket) SetDeadline(t time.Time) (err error) {
>>>>>>> 31bf680 ([net/std] some minor fixes)
	var d = untilTo(t)
	sk.SetTimeout(d)
	return
}
<<<<<<< HEAD
func (s *Socket) SetReadDeadline(t time.Time) (err error) {
=======
func (sk *Socket) SetReadDeadline(t time.Time) (err error) {
>>>>>>> 31bf680 ([net/std] some minor fixes)
	var d = untilTo(t)
	err = sk.SetReadTimeout(d)
	return
}
<<<<<<< HEAD
func (s *Socket) SetWriteDeadline(t time.Time) (err error) {
=======
func (sk *Socket) SetWriteDeadline(t time.Time) (err error) {
>>>>>>> 31bf680 ([net/std] some minor fixes)
	var d = untilTo(t)
	err = sk.SetWriteTimeout(d)
	return
}

//memar:impl std/net.TCPConn
<<<<<<< HEAD
func (s *Socket) CloseRead() (err error)                         { return }
func (s *Socket) CloseWrite() (err error)                        { return }
func (s *Socket) SetLinger(sec int) (err error)                  { return }
func (s *Socket) SetKeepAlive(keepalive bool) (err error)        { return }
func (s *Socket) SetKeepAlivePeriod(d time.Duration) (err error) { return }
func (s *Socket) SetNoDelay(noDelay bool) (err error)            { return }
=======
func (sk *Socket) CloseRead() (err error)                         { return }
func (sk *Socket) CloseWrite() (err error)                        { return }
func (sk *Socket) SetLinger(sec int) (err error)                  { return }
func (sk *Socket) SetKeepAlive(keepalive bool) (err error)        { return }
func (sk *Socket) SetKeepAlivePeriod(d time.Duration) (err error) { return }
func (sk *Socket) SetNoDelay(noDelay bool) (err error)            { return }
>>>>>>> 31bf680 ([net/std] some minor fixes)

func untilTo(t time.Time) (d protocol.Duration) {
	if !t.IsZero() {
		d = protocol.Duration(time.Until(t))
		if d == 0 {
			d = -1 // don't confuse deadline right now with no deadline
		}
	}
	return
}
