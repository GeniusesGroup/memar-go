/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

import (
	operation_p "memar/process/operation/protocol"
)

// Socket term in many textbooks refers to an entity that is uniquely identified by the socket number.
// In the original definition of socket given in RFC 147,[2] as it was related to the ARPA network in 1971,
// "the socket is specified as a 32-bit number with even sockets identifying receiving sockets and odd sockets identifying sending sockets."
// Today, however, socket communications are bidirectional.
// Any stateful network need to be implement a socket.
// https://en.wikipedia.org/wiki/Network_socket
type Socket interface {
	// PeerInitiated() bool // false means server-initiated

	// https://en.wikipedia.org/wiki/Berkeley_sockets
	// Internet socket APIs are usually based on the Berkeley sockets standard.
	// In the Berkeley sockets standard, sockets are a form of file descriptor,
	// due to the Unix philosophy that "everything is a file", and the analogies between sockets and files.
	// Both have functions to read, write, open, and close

	// string form of full address of socket to dial any time later.
	Field_NetworkAddresses

	Field_SocketStatus

	Socket_LowLevelAPIs
}

// Socket_LowLevelAPIs is low level APIs, don't use them in the services layer, if you don't know how it can be effect the application.
// It will use in chunks managing packages e.g. sRPC, QUIC, TCP, UDP, ... or Application layer protocols e.g. HTTP, ...
type Socket_LowLevelAPIs interface {
	Field_OSI_Physical
	Field_OSI_DataLink
	Field_OSI_Network
	Field_OSI_Transport
	Field_OSI_Session
	Field_OSI_Presentation
	Field_OSI_Application

	Method_SocketStatus

	Field_Socket_Buffers

	operation_p.Timeout
}
