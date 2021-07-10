/* For license and copyright information please see LEGAL file in repository */

package giti

import (
	"crypto/sha512"
	// "../convert"
)

// IDCalculator calculate service||error||data-structure ID by given urn
func IDCalculator(urn string) (id uint64) {
	// var hash = sha512.Sum512(convert.UnsafeStringToByteSlice(urn))
	var hash = sha512.Sum512([]byte(urn))
	id = uint64(hash[0]) | uint64(hash[1])<<8 | uint64(hash[2])<<16 | uint64(hash[3])<<24 | uint64(hash[4])<<32 | uint64(hash[5])<<40 | uint64(hash[6])<<48 | uint64(hash[7])<<56
	return
}
