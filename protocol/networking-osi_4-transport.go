/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

/*
**********************************************************************************
Transport (OSI Layer 4: Transport)
**********************************************************************************
*/

type NetworkTransport_HeaderID = ID

// NetworkTransport_Multiplexer indicate a transport segment multiplexer.
// OS and APP part must implement in dedicate structures.
type NetworkTransport_Multiplexer interface {
	HeaderID() NetworkTransport_HeaderID

	// Receiver must release segment slice and don't use it after return. So almost in most cases dev must copy segment payload to the stream.
	Receive(conn Connection, segment []byte)

	Shutdown()
}
