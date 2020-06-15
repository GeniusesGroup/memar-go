/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"crypto/tls"
	"io"
	"net"
	"strconv"
	"time"
)

/*
-------------------------------NOTICE:-------------------------------
We just implement and support TCP over IP for transition period and not our goal!
Please have plan to transform your network to GP protocol!
*/

// tcpNetwork use to store related data.
type tcpNetwork struct {
	s             *Server
	port          uint16
	streamHandler StreamHandler
	tcpListener   net.Listener
	tlsListener   net.Listener
}

// MakeTCPNetwork start a TCP listener and response request by given stream handler
func MakeTCPNetwork(s *Server, port uint16, streamHandler StreamHandler) (err error) {
	var p string = ":" + strconv.FormatUint(uint64(port), 10)

	var tcpListener net.Listener
	tcpListener, err = net.Listen("tcp", p)
	if err != nil {
		Log("TCP listen on port "+p+" failed due to: ", p, err)
		return
	}

	Log("Begin listen TCP on port: ", p)

	// register TCP network to s.Networks
	var tcp = tcpNetwork{
		s:             s,
		port:          port,
		streamHandler: streamHandler,
		tcpListener:   tcpListener,
	}
	s.Networks.RegisterTCPNetwork(&tcp)

	go handleTCPListener(s, &tcp, tcpListener)

	return
}

// MakeTCPTLSNetwork start a TCP/TLS listener and response request in given stream handler
func MakeTCPTLSNetwork(s *Server, port uint16, streamHandler StreamHandler) (err error) {
	var p string = ":" + strconv.FormatUint(uint64(port), 10)

	var tcpListener net.Listener
	tcpListener, err = net.Listen("tcp", p)
	if err != nil {
		Log("TCP/TLS listen on port "+p+" failed due to: ", p, err)
		return
	}

	Log("Begin listen TCP/TLS on port: ", p)

	// TODO::: convert tls package usage and implement tls by this package connections!
	var config = new(tls.Config)
	config.NextProtos = []string{"http/1.1", "sRPC"}
	config.PreferServerCipherSuites = true
	config.Certificates = make([]tls.Certificate, 1)
	// TODO::: get certificate from letsencrypt
	// var certPEMBlock, keyPEMBlock []byte
	config.Certificates[0], err = tls.X509KeyPair(certificate, privateKey)
	if err != nil {
		return
	}

	var tlsListener = tls.NewListener(tcpListener, config)

	// register TCP network to s.Networks
	var tcp = tcpNetwork{
		s:             s,
		port:          port,
		streamHandler: streamHandler,
		tcpListener:   tcpListener,
		tlsListener:   tlsListener,
	}
	s.Networks.RegisterTCPNetwork(&tcp)

	go handleTCPListener(s, &tcp, tlsListener)

	return
}

// handleTCPListener use to handle TCP networks packet with any application protocol.
// It is just support TCP+TLS not un-secure socket!
func handleTCPListener(s *Server, tcp *tcpNetwork, ln net.Listener) {
	for {
		var err error
		var tcpConn net.Conn
		tcpConn, err = ln.Accept()
		if err != nil {
			// Log("TCP accepting occur error: ", err)
			continue
		}

		// Log("Begin listen TCP conn on: ", conn.RemoteAddr())

		// set 1 minutes timeout for each connection
		tcpConn.SetDeadline(time.Now().Add(60 * time.Second))
		// conn.SetKeepAlive(true)

		// Due to nature of TCP and IPv4 NAT we have a lot problem and work to do!!
		var conn = Connection{
			StreamPool: map[uint32]*Stream{},
		}

		go handleTCPConn(s, tcp, tcpConn, &conn)

		// TODO::: handle conn status
	}
}

// TODO::: Check some other idea here:
// https://github.com/xtaci/gaio
func handleTCPConn(s *Server, tcp *tcpNetwork, tcpConn net.Conn, conn *Connection) {
	for {
		var err error
		var readSize int

		// Make a buffer to hold incoming data.
		// TODO::: make decision for 4096 byte below!!
		var buf = make([]byte, 4096)

		// TODO::: check below performance!
		// var buf bytes.Buffer
		// io.Copy(&buf, conn)
		// Log("total size:", buf.Len())

		// Read the incoming connection into the buffer.
		readSize, err = tcpConn.Read(buf)
		if err == io.EOF || readSize == 0 {
			// Peer already closed the connection, So we close it too!
			// Log("Error reading: ", err)
			tcpConn.Close()
			return
		} else if err != nil {
			// Log("Error reading: ", err.Error())
			tcpConn.Close()
			return
		}

		var reqStream, resStream *Stream
		reqStream, resStream, err = conn.MakeBidirectionalStream(0)

		reqStream.Payload = buf[:readSize]
		tcp.streamHandler(s, reqStream)

		readSize, err = tcpConn.Write(resStream.Payload)
		if err != nil {
			// Log("Error writing: ", err.Error())
			tcpConn.Close()
			return
		}

		// close the connection by Deadline and keep alive the connection.
		// tcpConn.Close()
	}
}

// shutdown the listener when the application closes or force to closes by not recovered panic!
func (tcp *tcpNetwork) shutdown() {
	tcp.tcpListener.Close()
	// tcp.tlsListener.Close()
}

func getIPPort(c net.Conn) (addr [16]byte) {
	ipAddr, ok := c.RemoteAddr().(*net.TCPAddr)
	if !ok {
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
