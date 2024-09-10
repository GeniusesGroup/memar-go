/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

import (
	string_p "memar/string/protocol"
)

// Addr represents a network end point address.
// They can be any layer 2 or 3 or even 4.
type NetworkAddress interface {
	LocalAddr() string_p.Stringer[string_p.String]
	RemoteAddr() string_p.Stringer[string_p.String]
}
