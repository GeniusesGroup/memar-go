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

	// shutdown the listener when the application closes or force to closes by not recovered panic!
	// first closing open listener for income packet and refuse all new packet,
	// then closing all idle connections,
	// and then waiting indefinitely for connections to return to idle
	// and then shut down
	Shutdown()
}

// NetworkTransportConnection or App2AppConnection
type NetworkTransportConnection interface {
	/* Connection data */
	MTU() int

	/* Peer data */
	ThingID() [32]byte
	UserID() [32]byte
	DomainName() string

	Send(st Stream) (err Error)
	SendAsync(st Stream) (err Error)

	Streams

	ConnectionMetrics
}
