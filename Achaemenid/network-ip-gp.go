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
	gpOverIPProtocolNumber = "ip:252"
)

// MakeGPOverIPNetwork start a IP packet listener and response request in given stream handler
func MakeGPOverIPNetwork(s *Server) (err error) {
	s.Networks.GPOverIP, err = net.ListenIP(gpOverIPProtocolNumber, &net.IPAddr{IP: s.Networks.localIP})
	if err != nil {
		log.Warn("IP listen on protocol number "+gpOverIPProtocolNumber+" failed due to: ", err)
		return
	}

	log.Info("Begin listen IP on protocol number: ", gpOverIPProtocolNumber)

	go handleGPEncapsulateInIP(s, s.Networks.GPOverIP)

	return
}

func handleGPEncapsulateInIP(s *Server, IPConn *net.IPConn) {
	var rwSize int
	var IPAddr *net.IPAddr
	var err error
	var conn *Connection
	for {
		// Make a buffer to hold incoming data -- no packet can be bigger.
		// 1480 = 1500(Ethernet MTU) - 20(IP header)
		var buf = make([]byte, 1480)

		rwSize, IPAddr, err = IPConn.ReadFromIP(buf)
		if err != nil || rwSize < gp.MinPacketLen {
			// TODO::: attack??
			continue
		}

		conn, _ = handleGP(s, buf[:rwSize])
		if conn != nil {
			conn.IPAddr = IPAddr
		}
	}
}
