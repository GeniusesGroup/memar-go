/* For license and copyright information please see LEGAL file in repository */

package dos

import (
	"../../protocol"
)

type netTransMux struct {
	transportsMux []protocol.NetworkTransportAppMultiplexer
}

// RegisterUDPNetwork use to register a established udp network!
func (n *netTransMux) RegisterNetworkTransportMultiplexer(tMux protocol.NetworkTransportAppMultiplexer) {
	// TODO::: register in OS
	n.transportsMux = append(n.transportsMux, tMux)
}

func (n *netTransMux) shutdown() {
	for i := 0; i < len(n.transportsMux); i++ {
		n.transportsMux[i].Shutdown()
	}
}
