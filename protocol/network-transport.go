/* For license and copyright information please see LEGAL file in repository */

package protocol

/*
**********************************************************************************
Transport (OSI Layer 3-6: Network to Presentation)
**********************************************************************************
*/

// NetworkTransportOSMultiplexer indicate a transport packet multiplexer methods must implemented by any transport connection in OS layer!
type NetworkTransportOSMultiplexer interface {
	HeaderID() (headerID NetworkLinkNextHeaderID)
	Receive(conn NetworkLinkConnection, packet []byte)

	RegisterAppMultiplexer(appMux NetworkTransportAppMultiplexer)
	UnRegisterAppMultiplexer(appMux NetworkTransportAppMultiplexer)

	Shutdown()
}

// NetworkTransportAppMultiplexer indicate a transport packet multiplexer methods must implemented by any transport multiplexer in app layer!
type NetworkTransportAppMultiplexer interface {
	HeaderID() (id NetworkLinkNextHeaderID)
	Receive(conn NetworkLinkConnection, packet []byte)
	// Send method exist in stream of each connection
}
