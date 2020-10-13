/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"crypto/tls"
	"net"
	"time"

	"../log"
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

// MakeTCPNetwork start a TCP listener and response request by given stream handler
func MakeTCPNetwork(s *Server, port uint16) (err error) {
	// Can't make a network on a port that doesn't has a handler!
	if s.StreamProtocols.GetProtocolHandler(port) == nil {
		return ErrAchaemenidProtocolHandler
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
	defer s.PanicHandler()
	// TODO::: defer a function to remake tcp listener
	for {
		var err error
		var tcpConn *net.TCPConn
		tcpConn, err = tcpListener.AcceptTCP()
		if err != nil {
			if log.DebugMode {
				log.Debug("TCP - Accepting new connection occur error:", err)
			}
			continue
		}

		if log.DebugMode {
			log.Debug("TCP - Begin listen on:", tcpConn.RemoteAddr())
		}

		go handleTCPConn(s, tcp, tcpConn)
	}
}

// TODO::: Check some other idea here:
// https://github.com/xtaci/gaio
func handleTCPConn(s *Server, tcp *tcpNetwork, tcpConn net.Conn) {
	defer s.PanicHandler()
	var conn *Connection
	var rwSize int
	var err error
	var st *Stream
	for {
		// close the connection by Deadline and keep alive the connection.
		// set or reset 2 minutes timeout for the connection
		tcpConn.SetDeadline(time.Now().Add(120 * time.Second))
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
		// if err == io.EOF || rwSize == 0 {
		// 	// log.Warn("Closing error reading: ", err)
		// 	tcpConn.Close()
		// 	return
		// } else
		if err != nil {
			// Peer already closed the connection, So we close it too!
			if log.DebugMode {
				log.Debug("TCP - Read error:", err.Error())
			}
			tcpConn.Close()
			return
		}

		if conn == nil {
			// TODO::: add limit make connection per IP
			st, err = MakeNewStream()
		} else {
			st, err = conn.MakeIncomeStream(0)
		}
		// Server can't make new stream or connection almost due to not enough resources!
		if err != nil {
			if log.DebugMode {
				log.Debug("TCP - Make new connection error:", err.Error())
			}
			// TODO::: need to send message??
			tcpConn.Close()
			return
		}

		st.IncomePayload = buf[:rwSize]
		s.StreamProtocols.GetProtocolHandler(tcp.port)(s, st)
		// Can't continue listen on a tcp connection that don't have active Achaemenid connection!
		if st.Connection == nil {
			tcpConn.Close()
			return
		}
		// TODO::: is it worth to check conn==nil or just overwrite it in every loop!!??
		// if conn == nil {
		conn = st.Connection
		// }

		rwSize, err = tcpConn.Write(st.OutcomePayload)
		if err != nil {
			if log.DebugMode {
				log.Debug("TCP - Writing error:", err.Error())
			}
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
