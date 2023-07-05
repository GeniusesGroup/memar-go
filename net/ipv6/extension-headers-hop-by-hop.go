/* For license and copyright information please see the LEGAL file in the code repository */

package ipv6

// HopByHop is IPv6 extension header with NextHeader==0
type HopByHop struct {
	NextHeader uint8
	HdrExtLen  uint8
	Options    []byte
}
