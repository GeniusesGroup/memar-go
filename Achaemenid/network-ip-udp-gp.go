/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"net"

	gp "../GP"
	"../log"
)

/*
-------------------------------NOTICE:-------------------------------
We just implement and support IP for transition period and not our goal!
Please have plan to transform your network to GP protocol!
*/

const (
	gpOverUDPPortNumber = 252
)

// MakeGPOverUDPNetwork start a UDP listener and response request in given stream handler
func MakeGPOverUDPNetwork(s *Server) (err error) {
	s.Networks.GPOverUDP, err = net.ListenUDP("udp", &net.UDPAddr{IP: s.Networks.localIP, Port: gpOverUDPPortNumber})
	if err != nil {
		log.Warn("UDP listen on port ", gpOverUDPPortNumber, " failed due to: ", err)
		return
	}

	log.Info("Begin listen UDP to serve GP over IP/UDP on port: ", gpOverUDPPortNumber)

	go handleGPEncapsulateInUDP(s, s.Networks.GPOverUDP)

	return
}

func handleGPEncapsulateInUDP(s *Server, udpConn *net.UDPConn) {
	var rwSize int
	var udpAddr *net.UDPAddr
	var err error
	var conn *Connection
	var st *Stream
	for {
		// Make a buffer to hold incoming data -- no packet can be bigger.
		// 1472 = 1500(Ethernet MTU) - 20(IP header) - 8(UDP header)
		var buf = make([]byte, 1472)

		rwSize, udpAddr, err = udpConn.ReadFromUDP(buf)
		if err != nil || rwSize < gp.MinPacketLen {
			// TODO::: attack??
			continue
		}

		conn, st = handleGP(s, buf[:rwSize])
		if conn != nil {
			// Just ignore packet and continue
			continue
		}

		rwSize, err = udpConn.WriteToUDP(st.OutcomePayload, udpAddr)
		if err != nil {
			if log.DebugMode {
				log.Debug("TCP - Writing error:", err.Error())
			}
			// Just ignore packet and continue
			continue
		}
	}
}
