/* For license and copyright information please see LEGAL file in repository */

package uuid

import (
	"crypto/rand"
	"io"
)

/*
UUID representation compliant with specification described in https://tools.ietf.org/html/rfc4122.

Use V1 for massive data that don't need to read much specially very close to write time!
These type of records write sequently in cluster and don't need very much move in cluster expand proccess!
*/

// NewV1 generate version 1 UUID include date-time and MAC address
func NewV1() (uuid [16]byte) {
	return
}

// NewV4 generate version 4 UUID include randomly numbers.
func NewV4() (uuid [16]byte) {
	var err error
	_, err = io.ReadFull(rand.Reader, uuid[:])
	if err != nil {
		// TODO::: make random by other ways
	}

	// Set version to 4
	uuid[6] = (uuid[6] & 0x0f) | (0x04 << 4)
	// Set variant to RFC4122
	uuid[8] = (uuid[8]&(0xff>>2) | (0x02 << 6))

	return
}

const (
	personUserID = 1 + iota
	orgUserID
)

// SetUserType use to set user type in any UUID generated!
func SetUserType(uuid [16]byte, userType uint8) [16]byte {
	//  first 4 bit use!
	return uuid
}
