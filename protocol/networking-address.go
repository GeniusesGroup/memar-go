/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type ProtocolID uint64

// Addr represents a network end point address.
// They can be any layer 2 or 3 or even 4.
type NetworkAddress interface {
	ProtocolID() ProtocolID
	LocalAddr() Stringer
	RemoteAddr() Stringer
}
