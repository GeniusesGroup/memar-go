/* For license and copyright information please see LEGAL file in repository */

package giti

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
	AppManifest() ApplicationManifest
	Receive(conn NetworkLinkConnection, packet []byte)

	Shutdown()
}

// NetworkTransportConnection or App2AppConnection
type NetworkTransportConnection interface {
	MTU() int

	Send(st Stream) (err Error)
	SendAsync(st Stream) (err Error)

	Streams

	MetricsConnection
}
