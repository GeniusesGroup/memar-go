/* For license and copyright information please see LEGAL file in repository */

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

// GetFirstUint64 use to get first 64bit of UUID as uint64
func GetFirstUint64(uuid [16]byte) (id uint64) {
	id = uint64(uuid[0]) | uint64(uuid[1])<<8 | uint64(uuid[2])<<16 | uint64(uuid[3])<<24 |
		uint64(uuid[4])<<32 | uint64(uuid[5])<<40 | uint64(uuid[6])<<48 | uint64(uuid[7])<<56
	return
}

// GetLastUint64 use to get last 64bit of UUID as uint64
func GetLastUint64(uuid [16]byte) (id uint64) {
	id = uint64(uuid[8]) | uint64(uuid[9])<<8 | uint64(uuid[10])<<16 | uint64(uuid[11])<<24 |
		uint64(uuid[12])<<32 | uint64(uuid[13])<<40 | uint64(uuid[14])<<48 | uint64(uuid[15])<<56
	return
}

// GetFirstUint32 use to get first 32bit of UUID as uint32
func GetFirstUint32(uuid [16]byte) (id uint32) {
	id = uint32(uuid[0]) | uint32(uuid[1])<<8 | uint32(uuid[2])<<16 | uint32(uuid[3])<<24
	return
}
