/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Addr represents a network end point address.
// They can be any layer 2 or 3 or even 4.
type NetworkAddress interface {
	LocalAddr() Stringer[String]
	RemoteAddr() Stringer[String]
}
