/* For license and copyright information please see the LEGAL file in the code repository */

package uuid

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"io"

	"memar/binary"
	"memar/convert"
	"memar/protocol"
)

// RFC4122 representation compliant with specification described in https://tools.ietf.org/html/rfc4122.
// https://github.com/google/uuid/blob/master/uuid.go
type UUID [16]byte

// V1 generate version 1 RFC4122 include date-time and MAC address
// Use V1 for massive data that don't need to read much specially very close to write time!
// These type of records write sequently in cluster and don't need very much move in cluster expand process!
func V1() (uuid UUID) {
	return
}

// V4 generate version 4 RFC4122 include randomly numbers.
func V4() (uuid UUID) {
	var err error
	_, err = io.ReadFull(rand.Reader, uuid[:])
	if err != nil {
		panic(err)
	}

	// Set version to 4
	uuid[6] = (uuid[6] & 0x0f) | (0x04 << 4)
	// Set variant to RFC4122
	uuid[8] = (uuid[8]&(0xff>>2) | (0x02 << 6))
	return
}

// V5 generate version 5 RFC4122 include hash namespace and value
func V5(nameSpace [16]byte, value []byte) (uuid UUID) {
	return
}

// Equal returns true if uuid1 and uuid2 equals
func (uuid UUID) Equal(uuid2 UUID) bool {
	return bytes.Equal(uuid[:], uuid2[:])
}

// encode/parse by RFC4122
//
//memar:impl memar/protocol.Stringer
func (uuid UUID) ToString() (s string, err protocol.Error) {
	s = uuid.String()
	return
}
func (uuid UUID) FromString(s string) (err protocol.Error) {
	var text = convert.UnsafeStringToByteSlice(s)
	hex.Decode(uuid[0:4], text[:8])
	hex.Decode(uuid[4:6], text[9:13])
	hex.Decode(uuid[6:8], text[14:18])
	hex.Decode(uuid[8:10], text[19:23])
	hex.Decode(uuid[10:], text[24:])
	return
}

// String returns canonical string representation of RFC4122:
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func (uuid UUID) String() string {
	var buf [36]byte
	encodeHex(buf[:], uuid)
	return string(buf[:])
}

// URI returns the RFC 2141 URN form of uuid,
// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx,  or "" if uuid is invalid.
func (uuid UUID) URI() string {
	var buf [36 + 9]byte
	copy(buf[:], "urn:uuid:")
	encodeHex(buf[9:], uuid)
	return string(buf[:])
}

func encodeHex(dst []byte, uuid UUID) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}

func (uuid UUID) FirstUint64() (id uint64) { return binary.LittleEndian.Uint64(uuid[0:]) }
func (uuid UUID) LastUint64() (id uint64)  { return binary.LittleEndian.Uint64(uuid[8:]) }
func (uuid UUID) FirstUint32() (id uint32) { return binary.LittleEndian.Uint32(uuid[0:]) }
