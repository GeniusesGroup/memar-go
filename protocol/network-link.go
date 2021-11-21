/* For license and copyright information please see LEGAL file in repository */

package protocol

type NetworkLinkNextHeaderID byte

// https://github.com/GeniusesGroup/RFCs/blob/master/Chapar.md#next-header-standard-supported-protocols
const (
	NetworkLinkNextHeaderSRPC NetworkLinkNextHeaderID = iota
	NetworkLinkNextHeaderGP
	NetworkLinkNextHeaderIPv4
	NetworkLinkNextHeaderIPv6
	NetworkLinkNextHeaderNTP

	NetworkLinkNextHeaderExperimental1 NetworkLinkNextHeaderID = 251
	NetworkLinkNextHeaderExperimental2 NetworkLinkNextHeaderID = 252
	NetworkLinkNextHeaderExperimental3 NetworkLinkNextHeaderID = 253
	NetworkLinkNextHeaderExperimental4 NetworkLinkNextHeaderID = 254
	NetworkLinkNextHeaderExperimental5 NetworkLinkNextHeaderID = 255
)

/*
**********************************************************************************
Link - (OSI Layer 2: Data Link)
**********************************************************************************
*/

// NetworkLinkMultiplexer indicate a link frame multiplexer object methods must implemented by any os!
type NetworkLinkMultiplexer interface {
	// Send usually use to send a broadcast frame
	Send(frame []byte) (err Error)
	Receive(frame []byte)

	RegisterTransportHandler(osMux NetworkTransportOSMultiplexer)
	UnRegisterTransportHandler(osMux NetworkTransportOSMultiplexer)

	Shutdown()
}

// NetworkLinkConnection or Device2DeviceConnection
type NetworkLinkConnection interface {
	MTU() int
	// Send transmitting in non blocking mode and just queue frames for congestion situations.
	// Return error not realted to frame situation, just about any hardware error.
	// A situation might be occur that the port available when a frame queued but when the time to send is come, the port broken and sender don't know about this!
	// Due to speed matters in link layer, and it is very rare situation, it is better to ignore suddenly port unavailability.
	Send(nextHeaderID NetworkLinkNextHeaderID, payload Codec) (err Error)
}
