/* For license and copyright information please see LEGAL file in repository */

package giti

type NetworkLinkNextHeaderID byte

// https://github.com/SabzCity/RFCs/blob/master/Chapar.md#next-header-standard-supported-protocols
const (
	NetworkLinkNextHeaderSRPC NetworkLinkNextHeaderID = iota
	NetworkLinkNextHeaderGP
	NetworkLinkNextHeaderIPv4
	NetworkLinkNextHeaderIPv6
	NetworkLinkNextHeaderICMP
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
	SendBroadcastAsync(nextHeaderID NetworkLinkNextHeaderID, payload Codec) (err Error)
	Receive(frame []byte)

	RegisterTransportHandler(osMux NetworkTransportOSMultiplexer)
	UnRegisterTransportHandler(osMux NetworkTransportOSMultiplexer)

	Shutdown()
}

// NetworkLinkConnection or Device2DeviceConnection
type NetworkLinkConnection interface {
	MTU() int

	Send(nextHeaderID NetworkLinkNextHeaderID, payload Codec) (err Error)

	// transmitting that must be non blocking and queue frames for congestion situations!
	// A situation might be occur that a port available when a frame queued but when the time to send is come, the port broken and sender don't know about this!
	SendAsync(nextHeaderID NetworkLinkNextHeaderID, payload Codec) (err Error)
}
