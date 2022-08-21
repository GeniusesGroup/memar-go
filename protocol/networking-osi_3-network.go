/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

/*
**********************************************************************************
Network (OSI Layer 2: Network)
**********************************************************************************
*/

// NetworkNetwork_Multiplexer indicate a transport packet multiplexer.
// OS and APP part must implement in dedicate structures.
type NetworkNetwork_Multiplexer interface {
	HeaderID() (headerID NetworkLink_NextHeaderID)

	// Receiver must release packet slice and don't use it after return. So almost in most cases dev must copy packet payload to the stream.
	Receive(conn NetworkLink_Connection, packet []byte)

	Shutdown()
}
