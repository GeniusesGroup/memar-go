/* For license and copyright information please see LEGAL file in repository */

package giti

type NetworkLinkHeaderID byte

// https://github.com/SabzCity/RFCs/blob/master/Chapar.md#next-header-standard-supported-protocols
const (
	NextHeaderSRPC NetworkLinkHeaderID = iota
	NextHeaderGP
	NextHeaderIPv4
	NextHeaderIPv6
	NextHeaderICMP
	NextHeaderNTP

	NextHeaderExperimental1 NetworkLinkHeaderID = 251
	NextHeaderExperimental2 NetworkLinkHeaderID = 252
	NextHeaderExperimental3 NetworkLinkHeaderID = 253
	NextHeaderExperimental4 NetworkLinkHeaderID = 254
	NextHeaderExperimental5 NetworkLinkHeaderID = 255
)

/*
**********************************************************************************
Link - (OSI Layer 2: Data Link)
**********************************************************************************
*/

// NetworkLinkMultiplexer indicate a link frame multiplexer object methods must implemented by any os!
type NetworkLinkMultiplexer interface {
	Receive(frame []byte)

	GetPort(id byte) NetworkLinkPort
	RegisterPort(port NetworkLinkPort)
	UnRegisterPort(port NetworkLinkPort)
	// RemovePort delete the port by given port number on given ports pool!
	RemovePort(portNumber byte)

	// GetUpperHandler return upper layer handler
	GetUpperHandler(id byte) NetworkTransportMultiplexer
	RegisterHandler(tm NetworkTransportMultiplexer)
	UnRegisterHandler(tm NetworkTransportMultiplexer)
	RemoveHandler(handlerID byte)

	// GetConnectionByPath get a connection by its path from Connections pool!!
	GetConnectionByPath(path []byte) (conn NetworkLinkConnection, err Error)
	NewConnection(port NetworkLinkPort, path []byte) (conn NetworkLinkConnection)
}

// NetworkLinkPort indicate a link port object methods must implemented by any driver!
type NetworkLinkPort interface {
	// RegisterMultiplexer register given multiplexer to the port for further usage!
	RegisterMultiplexer(lm NetworkLinkMultiplexer)

	PortNumber() (num byte)

	// transmitting that must be non blocking and queue frames for congestion situations!
	// A situation might be occur that a port available when a frame queued but when the time to send is come, the port broken and sender don't know about this!
	Send(frame []byte) (err Error)

	SendAsync(frame []byte) (err Error)

	Receive(frame []byte)
}

// NetworkLinkConnection or Device2DeviceConnection
type NetworkLinkConnection interface {
	MTU() int
	Send(nexHeaderID NetworkLinkHeaderID, payload WriterTo) (err Error)
}
