/* For license and copyright information please see LEGAL file in repository */

package ipv6

// HopByHop is IPv6 extension header with NextHeader==0
type HopByHop struct {
	NextHeader uint8
	HdrExtLen  uint8
	Options    []byte
}
