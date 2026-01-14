/* For license and copyright information please see the LEGAL file in the code repository */

package uuid_rfc4122

import (
	"bytes"
	"encoding/hex"

	"memar/buffer/byteslice/convert"
	"memar/codec/binary"
	error_p "memar/process/error/protocol"
)

// RFC4122 representation compliant with specification described in https://tools.ietf.org/html/rfc4122.
// https://github.com/google/uuid/blob/master/uuid.go
type UUID [16]byte

// Equal returns true if uuid1 and uuid2 equals
func (self UUID) Equal(to UUID) bool {
	return bytes.Equal(self[:], to[:])
}

// encode/parse by RFC4122
//
//memar:impl memar/codec/string/protocol.Stringer
func (self UUID) ToString() (s string, err error_p.Error) {
	s = self.String()
	return
}
func (self UUID) FromString(s string) (err error_p.Error) {
	var text = convert.UnsafeStringToByteSlice(s)
	hex.Decode(self[0:4], text[:8])
	hex.Decode(self[4:6], text[9:13])
	hex.Decode(self[6:8], text[14:18])
	hex.Decode(self[8:10], text[19:23])
	hex.Decode(self[10:], text[24:])
	return
}

// String returns canonical string representation of RFC4122:
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func (self UUID) String() string {
	var buf [36]byte
	encodeHex(buf[:], self)
	return string(buf[:])
}

// URI returns the RFC 2141 URN form of uuid,
// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx,  or "" if uuid is invalid.
func (self UUID) URI() string {
	var buf [36 + 9]byte
	copy(buf[:], "urn:uuid:")
	encodeHex(buf[9:], self)
	return string(buf[:])
}

func (self UUID) FirstUint64() (id uint64) { return binary.LittleEndian.Uint64(self[0:]) }
func (self UUID) LastUint64() (id uint64)  { return binary.LittleEndian.Uint64(self[8:]) }
func (self UUID) FirstUint32() (id uint32) { return binary.LittleEndian.Uint32(self[0:]) }

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
