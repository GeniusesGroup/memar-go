/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"crypto/tls"
	"io"
	"net"
	"time"

	etime "../earth-time"
	er "../error"
	"../log"
)

/*
-------------------------------NOTICE:-------------------------------
We just implement and support TCP over IP for transition period and not our goal!
Please have plan to transform your network to GP protocol!
*/

const (
	tcpKeepAliveDuration       = 120
	tcpKeepAliveDurationString = "120"
	// TODO::: make decision for 8192 byte below!! 8192 is max Chapar protocol payload size.
	tcpBufferSize = 8192
)

// tcpNetwork store related data.
type tcpNetwork struct {
	port        uint16
	listener    *net.TCPListener
	tlsListener net.Listener
	certificate tls.Certificate
}

// MakeTCPNetwork start a TCP listener and response request by given stream handler
func MakeTCPNetwork(port uint16) (err error) {
	// Can't make a network on a port that doesn't has a handler!
	if Server.StreamProtocols.GetProtocolHandler(port) == nil {
		return ErrProtocolHandler
	}

	var tcp = tcpNetwork{
		port: port,
	}

	tcp.listener, err = net.ListenTCP("tcp", &net.TCPAddr{IP: Server.Networks.localIP[:], Port: int(port)})
	if err != nil {
		log.Warn("TCP -  listen on port ", tcp.listener.Addr(), " failed due to: ", err)
		return
	}

	Server.Networks.RegisterTCPNetwork(&tcp)
	log.Info("TCP - Begin listen on ", tcp.listener.Addr())

	go handleTCPListener(&tcp, tcp.listener)

	return
}

// handleTCPListener use to handle TCP networks connections with any application protocol.
func handleTCPListener(tcp *tcpNetwork, tcpListener *net.TCPListener) {
	defer Server.PanicHandler()
	// TODO::: defer a function to remake tcp listener
	for {
		var err error
		var tcpConn *net.TCPConn
		tcpConn, err = tcpListener.AcceptTCP()
		if err != nil {
			if log.DebugMode {
				log.Debug("TCP - Accepting new connection occur error:", tcp.listener.Addr(), err)
			}
			continue
		}

		if log.DebugMode {
			log.Debug("TCP - New connection:", tcpConn.RemoteAddr())
		}

		go handleTCPConn(tcp, tcpConn)
	}
}

// TODO::: Check some other idea here:
// https://github.com/xtaci/gaio
func handleTCPConn(tcp *tcpNetwork, tcpConn net.Conn) {
	// TODO::: improve handle panic and log more data in log.DebugMode
	defer Server.PanicHandler()

	var conn *Connection
	var rwSize int
	var goErr error
	var err *er.Error
	var st *Stream
	for {
		// close the connection by Deadline and keep alive the connection.
		// set or reset 2 minutes timeout for the connection
		tcpConn.SetDeadline(time.Now().Add(tcpKeepAliveDuration * time.Second))
		// TODO::: TCP keep-alive function means send packet to peer and keep connection alive until close by some way! Why need this to waste resources!!??
		// tcpConn.(*net.TCPConn).SetKeepAlive(true)

		// Make a buffer to hold incoming data.
		var buf = make([]byte, tcpBufferSize)
		// Read the incoming connection into the buffer.
		rwSize, goErr = tcpConn.Read(buf)
		if goErr == io.EOF || rwSize == 0 {
			if log.DebugMode {
				log.Debug("TCP - Closed by peer:", goErr.Error())
			}
			// Peer already closed the connection, So we close it too!
			tcpConn.Close()
			return
		} else if goErr != nil {
			if log.DebugMode {
				log.Debug("TCP - Read error:", tcpConn.RemoteAddr(), goErr.Error())
			}
			// Peer already closed the connection, So we close it too!
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
				log.Debug("TCP - Make new Achaemenid stream error:", tcpConn.RemoteAddr(), err.Error())
			}
			// TODO::: need to send message??
			tcpConn.Close()
			return
		}

		st.tcpConn = tcpConn // TODO::: due to some TCP restrection, some client send body seprately after send header, so need pass tcpConn to HTTP layer!
		st.IncomePayload = buf[:rwSize]
		Server.StreamProtocols.GetProtocolHandler(tcp.port)(st)
		// Can't continue listen on a tcp connection that don't have active Achaemenid connection!
		if st.Connection == nil {
			if log.DebugMode {
				log.Debug("TCP - Make new Achaemenid connection error on this conn:", tcpConn.RemoteAddr())
			}
			tcpConn.Close()
			return
		}

		/* Metrics data */
		st.Connection.BytesReceived += uint64(rwSize)

		if conn == nil {
			conn = st.Connection
			copy(conn.IPAddr[:], tcpConn.RemoteAddr().(*net.TCPAddr).IP)
			conn.LastUsage = etime.Now()
		}

		rwSize, goErr = tcpConn.Write(st.OutcomePayload)
		if goErr != nil {
			if log.DebugMode {
				log.Debug("TCP - Writing error:", tcpConn.RemoteAddr(), goErr.Error())
			}
			tcpConn.Close()
			return
		}
		if rwSize != len(st.OutcomePayload) {
			// TODO:::
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
