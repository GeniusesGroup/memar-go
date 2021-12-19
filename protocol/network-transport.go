/* For license and copyright information please see LEGAL file in repository */

package protocol

/*
**********************************************************************************
Transport (OSI Layer 3-6: Network to Presentation)
**********************************************************************************
*/

// NetworkTransportMultiplexer indicate a transport packet multiplexer.
// OS and APP part must impelement in dedicate structures.
type NetworkTransportMultiplexer interface {
	HeaderID() (headerID NetworkLinkNextHeaderID)
	Receive(conn NetworkLinkConnection, packet []byte)
	Send(packet []byte) (err Error) // to send data use Send() that exist in stream of each connection

	Shutdown()
}
