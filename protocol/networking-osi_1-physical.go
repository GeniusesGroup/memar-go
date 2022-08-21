/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

/*
**********************************************************************************
Physical - (OSI Layer 1: Physical)
**********************************************************************************
*/

type NetworkPhysical_NextHeaderID byte

// https://github.com/GeniusesGroup/RFCs/blob/master/Chapar.md#next-header-standard-supported-protocols
const (
	NetworkPhysical_Unset NetworkPhysical_NextHeaderID = iota
	NetworkPhysical_SRPC
	NetworkPhysical_Chapar
	NetworkPhysical_Ethernet
	NetworkPhysical_ATM
)

// NetworkPhysical_Connection or Hardware2Hardware_Connection
type NetworkPhysical_Connection interface {
	RegisterLinkMultiplexer(linkMux NetworkLink_Multiplexer)
	UnRegisterLinkMultiplexer(linkMux NetworkLink_Multiplexer)

	Send(frame []byte) (err Error)
	SendAsync(frame []byte) (err Error)

	Shutdown()
}
