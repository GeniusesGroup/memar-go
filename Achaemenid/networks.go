/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import "../log"

type networks struct {
	gp  []*gpNetwork
	tcp []*tcpNetwork
}

// Init use to register all implemented networks!
// Usually Dev must register needed network by hand, not use this method to register all networks
func (n *networks) Init(s *Server) (err error) {
	// GP network just need register once for full port range!
	log.Info("try to register GP network")
	err = MakeGPNetwork(s)
	if err != nil {
		return
	}

	log.Info("try to register TCP on port 80 to listen for HTTP protocol")
	err = MakeTCPNetwork(s, 80)
	if err != nil {
		return
	}
	log.Info("try to register TCP on port 8080 to listen for HTTP protocol for dev phase")
	err = MakeTCPNetwork(s, 8080)
	if err != nil {
		return
	}
	log.Info("try to register TCP/TLS on port 443 to listen for HTTPs protocol")
	err = MakeTCPTLSNetwork(s, 443)
	if err != nil {
		return
	}

	return
}

// RegisterTCPNetwork use to register a established tcp network!
func (n *networks) RegisterTCPNetwork(tcp *tcpNetwork) {
	n.tcp = append(n.tcp, tcp)
}

func (n *networks) shutdown() {
	// TODO:::

	// first closing open listener for income packet and refuse all new packet,
	// then closing all idle connections,
	// and then waiting indefinitely for connections to return to idle
	// and then shut down
}
