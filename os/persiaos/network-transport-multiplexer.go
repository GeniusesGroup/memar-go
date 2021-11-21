/* For license and copyright information please see LEGAL file in repository */

package persiaos

// RegisterNetworkTransportMultiplexer will register multiplexer only if it is GP multiplexer.
func (os *os) RegisterNetworkTransportMultiplexer(appMux protocol.NetworkTransportMultiplexer) {
	switch appMux.HeaderID() {
	case protocol.NetworkLinkNextHeaderGP:
		os.gp.RegisterAppMultiplexer(appMux)
	case protocol.NetworkLinkNextHeaderIPv4:
		os.ipv4.RegisterAppMultiplexer(appMux)
	}
}
