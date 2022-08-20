/* For license and copyright information please see LEGAL file in repository */

package protocol

/*
**********************************************************************************
Physical - (OSI Layer 1: Physical)
**********************************************************************************
*/

// NetworkPhysicalConnection or Hardware2HardwareConnection
type NetworkPhysicalConnection interface {
	RegisterLinkMultiplexer(linkMux NetworkLinkMultiplexer)
	UnRegisterLinkMultiplexer(linkMux NetworkLinkMultiplexer)

	Send(frame []byte) (err Error)
	SendAsync(frame []byte) (err Error)

	Shutdown()
}
