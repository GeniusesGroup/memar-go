/* For license and copyright information please see the LEGAL file in the code repository */

package uuid

import (
	"bytes"
	"encoding/hex"
)

// Equal returns true if uuid1 and uuid2 equals
func Equal(uuid1, uuid2 [16]byte) bool {
	return bytes.Equal(uuid1[:], uuid2[:])
}

// ToString returns canonical string representation of UUID:
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func ToString(uuid [16]byte) string {
	buf := make([]byte, 36)
	hex.Encode(buf[0:8], uuid[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], uuid[10:])

	return string(buf)
}

// FromString will parsing UUID from string input
func FromString(s string) (uuid [16]byte) {
	return
}
