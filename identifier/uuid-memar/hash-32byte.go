/* For license and copyright information please see the LEGAL file in the code repository */

package memar_uuid

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"reflect"
	"unsafe"

	"golang.org/x/crypto/sha3"

	"memar/codec/binary"
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
	"memar/time/duration"
	time_p "memar/time/protocol"
	"memar/time/unix"
)

type Hash32 [32]byte

//memar:impl memar/protocol.UUID_Hash
func (self Hash32) UUID() [32]byte     { return self }
func (self Hash32) ID() datatype_p.ID  { return self.id() }
func (self Hash32) IDasString() string { return base64.RawURLEncoding.EncodeToString(self[:8]) }

//memar:impl memar/codec/string/protocol.Stringer
func (self Hash32) ToString() (s string, err error_p.Error) {
	s = base64.RawURLEncoding.EncodeToString(self[:])
	return
}
func (self *Hash32) FromString(s string) (err error_p.Error) {
	// TODO:::
	return
}

//memar:impl memar/protocol.UUID
func (self Hash32) ExistenceTime() time_p.Time {
	var time unix.Time
	time.ChangeTo(self.secondElapsed(), self.nanoSecondElapsed())
	return &time
}

// New will generate 32 byte time based UUID.
// **CAUTION**: Use for ObjectID in a clustered software without any hash cause all writes always go to one node.
// 99.999999% collision free on distribution generation!
func (self *Hash32) New() {
	var err error
	_, err = io.ReadFull(rand.Reader, self[12:])
	if err != nil {
		// TODO::: make random by other ways
	}

	// Set time to UUID
	var now = unix.Now()
	self.setSecondElapsed(now.SecondElapsed())
	self.setNanoInSecondElapsed(now.NanoInSecondElapsed())
}

// NewHash generate 32 byte incremental by time + hash of data UUID
// CAUTION::: Use for ObjectID in a clustered software cause all writes always go to one node!
// 99.999% collision free on distribution generation.
func (self *Hash32) NewHash(data []byte) {
	*self = sha3.Sum256(data)
}

func (self *Hash32) NewHashString(data string) {
	self.NewHash((unsafeStringToByteSlice(data)))
}

// NewRandom generate 32 byte random UUID.
// CAUTION::: Not use in distribution platforms!
func (self *Hash32) NewRandom() {
	var err error
	_, err = io.ReadFull(rand.Reader, self[:])
	if err != nil {
		// TODO::: make random by other ways
	}
}

func (self Hash32) id() datatype_p.ID {
	return datatype_p.ID(binary.LittleEndian.Uint64(self[0:]))
}
func (self Hash32) secondElapsed() duration.Second {
	return duration.Second(binary.LittleEndian.Uint64(self[0:]))
}
func (self Hash32) nanoSecondElapsed() duration.NanoInSecond {
	return duration.NanoInSecond(binary.LittleEndian.Uint32(self[8:]))
}
func (self *Hash32) setSecondElapsed(sec duration.Second) {
	binary.LittleEndian.PutUint64(self[0:], uint64(sec))
}
func (self *Hash32) setNanoInSecondElapsed(nsec duration.NanoInSecond) {
	binary.LittleEndian.PutUint32(self[8:], uint32(nsec))
}

func IDfromString(IDasString string) (id uint64, err error_p.Error) {
	var IDasSlice = unsafeStringToByteSlice(IDasString)
	var ID [8]byte
	var _, goErr = base64.RawURLEncoding.Decode(ID[:], IDasSlice)
	if goErr != nil {
		// err =
		return
	}
	id = binary.LittleEndian.Uint64(ID[0:])
	return
}

func unsafeStringToByteSlice(req string) (res []byte) {
	var reqStruct = (*reflect.StringHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len
	resStruct.Cap = reqStruct.Len
	return
}
