/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

import (
	object_p "memar/computer/language/object/protocol"
)

/*
**********************************************************************************
Network (OSI Layer 2: Network)

https://en.wikipedia.org/wiki/Network_layer
**********************************************************************************
*/

// OSI_Network is the network layer that provides the means of transferring variable-length network packets
// from a source to a destination host via one or more networks.
type OSI_Network interface {
	object_p.LifeCycle
	Framer
	NetworkAddress // string form of address (for example, "ipv4://192.0.2.1", "ipv6://[2001:db8::1]")

	OSI_Network_LowLevelAPIs
}

// ConnectionLevelAPIs is low level APIs, don't use them in the services layer, if you don't know how it can be effect the application.
// It will use in chunks managing packages e.g. sRPC, QUIC, TCP, UDP, ... or Application layer protocols e.g. HTTP, ...
type OSI_Network_LowLevelAPIs interface {
	FrameWriter
}
