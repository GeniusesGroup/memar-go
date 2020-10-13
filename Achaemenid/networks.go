/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"net"

	"../log"
)

type networks struct {
	server    *Server
	gp        []*gpNetwork
	localIP   net.IP
	GPOverIP  *net.IPConn
	GPOverUDP *net.UDPConn
	udp       []*udpNetwork
	tcp       []*tcpNetwork
}

// Init use to register all implemented networks!
// Usually Dev must register needed network by hand, not use this method to register all networks
func (n *networks) Init(s *Server) (err error) {
	n.server = s
	n.localIP = getLocalIP()

	// GP network just need register once for full port range!
	log.Info("try to register GP network...")
	err = MakeGPNetwork(s)
	if err != nil {
		return
	}

	// GP over IP network.
	log.Info("try to register GP over IP network...")
	err = MakeGPOverIPNetwork(s)
	if err != nil {
		return
	}

	// GP over IP/UDP network.
	log.Info("try to register GP over IP/UDP network...")
	err = MakeGPOverUDPNetwork(s)
	if err != nil {
		return
	}

	return
}

// RegisterUDPNetwork use to register a established udp network!
func (n *networks) RegisterUDPNetwork(udp *udpNetwork) {
	n.udp = append(n.udp, udp)
}

// RegisterTCPNetwork use to register a established tcp network!
func (n *networks) RegisterTCPNetwork(tcp *tcpNetwork) {
	n.tcp = append(n.tcp, tcp)
}

func (n *networks) shutdown() {
	var err error
	// TODO:::

	// first closing open listener for income packet and refuse all new packet,
	// then closing all idle connections,
	// and then waiting indefinitely for connections to return to idle
	// and then shut down

	if n.GPOverIP != nil {
		err = n.GPOverIP.Close()
		if err != nil {
			log.Warn("Closing IP/GP network face this error: ", err)
		}
	}

	if n.GPOverUDP != nil {
		err = n.GPOverUDP.Close()
		if err != nil {
			log.Warn("Closing IP/UDP/GP network face this error: ", err)
		}
	}
}

func getLocalIP() (ipAddr net.IP) {
	var addrs, err = net.InterfaceAddrs()
	if err != nil {
		return
	}

	var addr net.Addr
	for _, addr = range addrs {
		switch addr := addr.(type) {
		case *net.IPNet:
			if addr.IP.IsGlobalUnicast() {
				return addr.IP
			}
		}
	}

	return
}
