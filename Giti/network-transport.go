/* For license and copyright information please see LEGAL file in repository */

package giti

/*
**********************************************************************************
Transport (OSI Layer 3-6: Network to Presentation)
**********************************************************************************
*/

// NetworkTransportMultiplexer indicate a transport packet multiplexer object methods must implemented by any transport connection!
type NetworkTransportMultiplexer interface {
	GetConnectionByID(personConnID [32]byte) NetworkTransportConnection

	HeaderID() (id byte)
	Receive(conn NetworkLinkConnection, packet []byte)	
	
}

// NetworkTransportConnection or App2AppConnection
type NetworkTransportConnection interface {
	MTU() int

	Streams

	MetricsConnection
}
