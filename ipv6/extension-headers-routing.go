/* For license and copyright information please see LEGAL file in repository */

package ipv6

// Routing is IPv6 extension header with NextHeader==43
type Routing struct {
	NextHeader       uint8
	HdrExtLen        uint8
	RoutingType      uint8
	SegmentsLeft     uint8
	TypeSpecificData []byte
	Options          []byte
}
