/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"net"

	"../log"
)

/*
-------------------------------NOTICE:-------------------------------
We just implement and support IP/UDP for transition period and not our goal!
Please have plan to transform your network to GP protocol!
*/

// udpNetwork store related data.
type udpNetwork struct {
	s    *Server
	port uint16
	conn *net.UDPConn
}

// MakeUDPNetwork start a UDP listener and response request in given stream handler
func MakeUDPNetwork(s *Server, port uint16) (err error) {
	// Can't make a network on a port that doesn't has a handler!
	if s.StreamProtocols.GetProtocolHandler(port) == nil {
		return ErrProtocolHandler
	}

	var udp = udpNetwork{
		s:    s,
		port: port,
	}

	udp.conn, err = net.ListenUDP("udp", &net.UDPAddr{IP: s.Networks.localIP[:], Port: int(port)})
	if err != nil {
		log.Warn("UDP listen on port ", port, " failed due to: ", err)
		return
	}

	s.Networks.RegisterUDPNetwork(&udp)
	log.Info("Begin listen UDP on port: ", port)

	go handleUDPListener(s, &udp)

	return
}

// handleUDPListener use to handle UDP networks packets with any application protocol.
func handleUDPListener(s *Server, udp *udpNetwork) {
	var readSize int
	var UDPAddr *net.UDPAddr
	var err error
	for {
		// Make a buffer to hold incoming data -- no packet can be bigger.
		// 1472 = 1500(Ethernet MTU) - 20(IP header) - 8(UDP header)
		var buf = make([]byte, 1472)

		readSize, UDPAddr, err = udp.conn.ReadFromUDP(buf)
		if err != nil {
			// TODO::: attack??
			continue
		}

		go handleUDPPacket(s, udp, buf[:readSize], UDPAddr)

	}
}

func handleUDPPacket(s *Server, udp *udpNetwork, packet []byte, udpAddr *net.UDPAddr) {
	var err error
	var st *Stream

	st, err = MakeNewStream()
	// Server can't make new stream or connection almost due to not enough resources!
	if err != nil {
		// log.Warn("Error writing on UDP due to ", err.Error())
		return
	}

	st.IncomePayload = packet
	s.StreamProtocols.GetProtocolHandler(udp.port)(s, st)

	/* Metrics data */
	st.Connection.BytesReceived += uint64(len(st.IncomePayload))

	_, err = udp.conn.WriteToUDP(st.OutcomePayload, udpAddr)
	if err != nil {
		// log.Warn("Error writing on UDP due to ", err.Error())
		return
	}
}
