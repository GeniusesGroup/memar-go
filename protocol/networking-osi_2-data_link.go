/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

/*
**********************************************************************************
Link - (OSI Layer 2: Data Link)

It can use to network hardware devices in a computers or connect two or more computers.
**********************************************************************************
*/

type NetworkLink_NextHeaderID uint64

// NetworkLink_Multiplexer indicate a link frame multiplexer object methods must implemented by any os.
type NetworkLink_Multiplexer interface {
	HeaderID() (headerID NetworkPhysical_NextHeaderID)

	// Send usually use to send a broadcast frame
	Send(frame []byte) (err Error)
	Receive(conn NetworkPhysical_Connection, frame []byte) (err Error)

	RegisterNetworkMux(transMux NetworkNetwork_Multiplexer)
	UnRegisterNetworkMux(transMux NetworkNetwork_Multiplexer)

	ObjectLifeCycle
}

// NetworkLink_Connection or Device2DeviceConnection
type NetworkLink_Connection interface {
	NewFrame(nexHeaderID NetworkLink_NextHeaderID, payloadLen int) (frame []byte, payload []byte, err Error)

	// Send transmitting in blocking mode and block caller until physical layer copy frame to its hardware.
	// Return error not related to frame situation, just about any hardware error.
	// A situation might be occur that the port available when a frame queued but when the time to send is come,
	// the port broken and sender don't know about this.
	// Due to speed matters in link layer, and it is very rare situation, it is better to ignore suddenly port unavailability.
	Send(frame []byte) (err Error)

	NetworkMTU
	NetworkAddress // string form of address (for example, "MAC://aa:bb:cc:dd:ee:ff", "Chapar://[1:242:20]")
}
