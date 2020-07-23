/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"crypto/tls"
	"io"
	"net"
	"time"

	"../errors"
	"../log"
	"../uuid"
)

/*
-------------------------------NOTICE:-------------------------------
We just implement and support TCP over IP for transition period and not our goal!
Please have plan to transform your network to GP protocol!
*/

// tcpNetwork store related data.
type tcpNetwork struct {
	s           *Server
	port        uint16
	listener    *net.TCPListener
	tlsListener net.Listener
	certificate tls.Certificate
}

// Errors
var (
	ErrNoStreamProtocolHandler = errors.New("NoStreamProtocolHandler", "Can't make a network on a port that doesn't has a stream handler!")
)

// MakeTCPNetwork start a TCP listener and response request by given stream handler
func MakeTCPNetwork(s *Server, port uint16) (err error) {
	// Can't make a network on a port that doesn't has a handler!
	if s.StreamProtocols.GetProtocolHandler(port) == nil {
		return ErrNoStreamProtocolHandler
	}

	var tcp = tcpNetwork{
		s:    s,
		port: port,
	}

	tcp.listener, err = net.ListenTCP("tcp", &net.TCPAddr{IP: s.Networks.localIP, Port: int(port)})
	if err != nil {
		log.Warn("TCP listen on port ", tcp.listener.Addr(), " failed due to: ", err)
		return
	}

	s.Networks.RegisterTCPNetwork(&tcp)
	log.Info("Begin listen TCP on ", tcp.listener.Addr())

	go handleTCPListener(s, &tcp, tcp.listener)

	return
}

// handleTCPListener use to handle TCP networks connections with any application protocol.
func handleTCPListener(s *Server, tcp *tcpNetwork, tcpListener *net.TCPListener) {
	for {
		var err error
		var tcpConn *net.TCPConn
		tcpConn, err = tcpListener.AcceptTCP()
		if err != nil {
			// log.Warn("TCP accepting occur error: ", err)
			continue
		}
		// log.Info("Begin listen TCP conn on: ", tcpConn.RemoteAddr())

		go handleTCPConn(s, tcp, tcpConn)
	}
}

// TODO::: Check some other idea here:
// https://github.com/xtaci/gaio
func handleTCPConn(s *Server, tcp *tcpNetwork, tcpConn net.Conn) {
	var conn *Connection
	var rwSize int
	var err error
	var reqStream, resStream *Stream
	for {
		// close the connection by Deadline and keep alive the connection.
		// set or reset 1 minutes timeout for the connection
		tcpConn.SetDeadline(time.Now().Add(60 * time.Second))
		// TODO::: TCP keep-alive function means send packet to peer and keep connection alive until close by some way! Why need this to waste resources!!??
		// tcpConn.(*net.TCPConn).SetKeepAlive(true)

		// Make a buffer to hold incoming data.
		// TODO::: make decision for 8192 byte below!! 8192 is max Chapar protocol payload size.
		var buf = make([]byte, 4096)

		// TODO::: check below performance!
		// var buf bytes.Buffer
		// io.Copy(&buf, conn)
		// log.Warn("total size:", buf.Len())

		// Read the incoming connection into the buffer.
		rwSize, err = tcpConn.Read(buf)
		if err == io.EOF || rwSize == 0 {
			// Peer already closed the connection, So we close it too!
			// log.Warn("Closing error reading: ", err)
			tcpConn.Close()
			return
		} else if err != nil {
			// log.Warn("Other error reading: ", err.Error())
			tcpConn.Close()
			return
		}

		if conn == nil {
			// TODO::: add limit make connection per IP
			reqStream, resStream, err = MakeBidirectionalStream()
		} else {
			reqStream, resStream, err = conn.MakeBidirectionalStream(0)
		}
		// Server can't make new stream or connection almost due to not enough resources!
		if err != nil {
			// log.Warn("Error writing: ", err.Error())
			tcpConn.Close()
			return
		}

		reqStream.Payload = buf[:rwSize]
		s.StreamProtocols.GetProtocolHandler(tcp.port)(s, reqStream)
		// Can't continue listen on a tcp connection that don't have active Achaemenid connection!
		if reqStream.Connection == nil {
			tcpConn.Close()
			return
		}
		// TODO::: is it worth to check conn==nil or just overwrite it in every loop!!??
		// if conn == nil {
		conn = reqStream.Connection
		// }

		rwSize, err = tcpConn.Write(resStream.Payload)
		if err != nil {
			// log.Warn("Error writing: ", err.Error())
			tcpConn.Close()
			return
		}
	}
}

// shutdown the listener when the application closes or force to closes by not recovered panic!
func (tcp *tcpNetwork) shutdown() {
	if tcp == nil {
		return
	}
	if tcp.listener != nil {
		tcp.listener.Close()
	}
	if tcp.tlsListener != nil {
		tcp.tlsListener.Close()
	}
}

// makeNewGuestConnection make new connection and register on given stream due to it is first attempt connect to server!
func makeNewGuestConnection(s *Server, st *Stream) (err error) {
	if s.Manifest.TechnicalInfo.GuestMaxConnections == 0 {
		return ErrGuestUsersConnectionNotAllow
	} else if s.Manifest.TechnicalInfo.GuestMaxConnections > 0 && s.Connections.GuestConnectionCount > s.Manifest.TechnicalInfo.GuestMaxConnections {
		return ErrMaxOpenedGuestConnectionReached
	}

	st.Connection = &Connection{
		ID:         uuid.NewV4(),
		StreamPool: map[uint32]*Stream{},
		UserType:   0, // guest
	}
	st.ReqRes.Connection = st.Connection

	s.Connections.RegisterConnection(st.Connection)
	s.Connections.GuestConnectionCount++

	return
}

func getIPPort(c net.Conn) (addr [16]byte) {
	var ipAddr = c.RemoteAddr().(*net.TCPAddr)
	if ipAddr == nil {
		return
	}
	addr[0] = ipAddr.IP[0]
	addr[1] = ipAddr.IP[1]
	addr[2] = ipAddr.IP[2]
	addr[3] = ipAddr.IP[3]
	addr[4] = byte(ipAddr.Port)
	addr[5] = byte(ipAddr.Port >> 8)
	return
}
