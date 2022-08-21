/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

/*
**********************************************************************************
Link - (OSI Layer 2: Data Link)

It can use to network hardware devices in a computers or connect two or more computers.
**********************************************************************************
*/

type NetworkLink_NextHeaderID byte

// https://github.com/GeniusesGroup/RFCs/blob/master/Chapar.md#next-header-standard-supported-protocols
const (
	NetworkLink_Unset NetworkLink_NextHeaderID = iota
	NetworkLink_SRPC
	NetworkLink_GP
	NetworkLink_IPv6
	NetworkLink_NTP

	NetworkLink_Experimental1 NetworkLink_NextHeaderID = 251 // Use for non supported protocols like IPv4, ...
	NetworkLink_Experimental2 NetworkLink_NextHeaderID = 252
	NetworkLink_Experimental3 NetworkLink_NextHeaderID = 253
	NetworkLink_Experimental4 NetworkLink_NextHeaderID = 254
	NetworkLink_Experimental5 NetworkLink_NextHeaderID = 255
)

// NetworkLink_Multiplexer indicate a link frame multiplexer object methods must implemented by any os.
type NetworkLink_Multiplexer interface {
	// Send usually use to send a broadcast frame
	Send(frame []byte) (err Error)
	Receive(conn NetworkPhysical_Connection, frame []byte)

	RegisterNetworkMux(transMux NetworkNetwork_Multiplexer)
	UnRegisterNetworkMux(transMux NetworkNetwork_Multiplexer)

	Shutdown()
}

// NetworkLink_Connection or Device2DeviceConnection
type NetworkLink_Connection interface {
	MTU() int

	NewFrame(nexHeaderID NetworkLink_NextHeaderID, payloadLen int) (frame []byte, payload []byte, err Error)

	// Send transmitting in non blocking mode and just queue frames for congestion situations.
	// Return error not related to frame situation, just about any hardware error.
	// A situation might be occur that the port available when a frame queued but when the time to send is come, the port broken and sender don't know about this.
	// Due to speed matters in link layer, and it is very rare situation, it is better to ignore suddenly port unavailability.
	// After return, caller can't reuse payload array anymore.
	Send(frame []byte) (err Error)
}
