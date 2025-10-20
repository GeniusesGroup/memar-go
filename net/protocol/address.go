/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

import (
	string_p "memar/codec/string/protocol"
)

// Addr represents a network end point address.
// They can be any layer 2 or 3 or 4 or even socket!
type Field_NetworkAddresses interface {
	// LocalAddr or Source address
	LocalAddr() NetworkAddress
	
	// RemoteAddr or Destination address
	RemoteAddr() NetworkAddress
}

type NetworkAddress interface {
	string_p.Stringer
}
