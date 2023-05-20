/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

/*
**********************************************************************************
Physical - (OSI Layer 1: Physical)
**********************************************************************************
*/

type NetworkPhysical_NextHeaderID byte

const (
	NetworkPhysical_Unset NetworkPhysical_NextHeaderID = iota
	NetworkPhysical_sRPC
	NetworkPhysical_Chapar
	NetworkPhysical_Ethernet
	NetworkPhysical_ATM
)

// NetworkPhysical_Connection or Hardware2Hardware_Connection
type NetworkPhysical_Connection interface {
	RegisterLinkMultiplexer(linkMux NetworkLink_Multiplexer)
	UnRegisterLinkMultiplexer(linkMux NetworkLink_Multiplexer)

	// Send transmitting in blocking mode and block caller until physical layer copy frame to its hardware.
	// Return error not related to frame situation, just about any hardware error.
	// A situation might be occur that the port available when a frame queued but when the time to send is come,
	// the port broken and sender don't know about this.
	// Due to speed matters in link layer, and it is very rare situation, it is better to ignore suddenly port unavailability.
	Send(frame []byte) (err Error)

	ObjectLifeCycle
}
