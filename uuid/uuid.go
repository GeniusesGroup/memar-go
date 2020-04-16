/* For license and copyright information please see LEGAL file in repository */

package uuid

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// UUID representation compliant with specification
// described in https://tools.ietf.org/html/rfc4122.
type UUID [16]byte

// Nil is special form of UUID that is specified to have all
// 128 bits set to zero.
var Nil = UUID{}

// NewV4 returns random generated UUID.
func NewV4() (UUID, error) {
	id := UUID{}
	if _, err := io.ReadFull(rand.Reader, id[:]); err != nil {
		return Nil, err
	}
	// Set version to 4
	id[6] = (id[6] & 0x0f) | (0x04 << 4)
	// Set variant to RFC4122
	id[8] = (id[8]&(0xff>>2) | (0x02 << 6))

	return id, nil
}

// Equal returns true if u1 and u2 equals
func Equal(id1, id2 UUID) bool {
	return bytes.Equal(id1[:], id2[:])
}

// String implements Stringer interface and returns canonical string representation of UUID:
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func String(u UUID) string {
	buf := make([]byte, 36)
	hex.Encode(buf[0:8], u[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], u[10:])

	return string(buf)
}

const (
	personUserID = 1 + iota
	orgUserID
)

// NewPersonID use to make New UUID for person. first 4 bit set to 1!
func NewPersonID() (UUID, error) {
	id := UUID{}
	if _, err := io.ReadFull(rand.Reader, id[:]); err != nil {
		return id, err
	}

	// Set version to 4
	id[6] = (id[6] & 0x0f) | (0x04 << 4)
	// Set variant to RFC4122
	id[8] = (id[8]&(0xff>>2) | (0x02 << 6))

	// Set UUID to PersonType
	id[0] = (id[0]&(0x1>>1) | (0x02 << 6))

	return id, nil
}
