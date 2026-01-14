/* For license and copyright information please see the LEGAL file in the code repository */

package uuid

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"reflect"
	"unsafe"

	"golang.org/x/crypto/sha3"

	"memar/binary"
	"memar/protocol"
	"memar/time/duration"
	time_p "memar/time/protocol"
	"memar/time/unix"
)

// UID is the same as the UUID.
// Use this type when embed in other struct to solve field & method same name problem(UUID struct and UUID() method) to satisfy interfaces.
type UID = UUID

type UUID [32]byte

//memar:impl memar/protocol.UUID_Hash
func (id UUID) UUID() [32]byte          { return id }
func (id UUID) ID() protocol.DataTypeID { return id.id() }
func (id UUID) IDasString() string      { return base64.RawURLEncoding.EncodeToString(id[:8]) }

//memar:impl memar/string/protocol.Stringer
func (id UUID) ToString() (s string, err protocol.Error) {
	s = base64.RawURLEncoding.EncodeToString(id[:])
	return
}
func (id *UUID) FromString(s string) (err protocol.Error) {
	// TODO:::
	return
}

//memar:impl memar/protocol.UUID
func (id UUID) ExistenceTime() time_p.Time {
	var time unix.Time
	time.ChangeTo(id.secondElapsed(), id.nanoSecondElapsed())
	return &time
}

// New will generate 32 byte time based UUID.
// **CAUTION**: Use for ObjectID in a clustered software without any hash cause all writes always go to one node.
// 99.999999% collision free on distribution generation!
func (id *UUID) New() {
	var err error
	_, err = io.ReadFull(rand.Reader, id[12:])
	if err != nil {
		// TODO::: make random by other ways
	}

	// Set time to UUID
	var now = unix.Now()
	id.setSecondElapsed(now.SecondElapsed())
	id.setNanoInSecondElapsed(now.NanoInSecondElapsed())
}

// NewHash generate 32 byte incremental by time + hash of data UUID
// CAUTION::: Use for ObjectID in a clustered software cause all writes always go to one node!
// 99.999% collision free on distribution generation.
func (id *UUID) NewHash(data []byte) {
	*id = sha3.Sum256(data)
}

func (id *UUID) NewHashString(data string) {
	id.NewHash((unsafeStringToByteSlice(data)))
}

// NewRandom generate 32 byte random UUID.
// CAUTION::: Not use in distribution platforms!
func (id *UUID) NewRandom() {
	var err error
	_, err = io.ReadFull(rand.Reader, id[:])
	if err != nil {
		// TODO::: make random by other ways
	}
}

func (id UUID) id() protocol.DataTypeID {
	return protocol.DataTypeID(binary.LittleEndian.Uint64(id[0:]))
}
func (id UUID) secondElapsed() duration.Second {
	return duration.Second(binary.LittleEndian.Uint64(id[0:]))
}
func (id UUID) nanoSecondElapsed() duration.NanoInSecond {
	return duration.NanoInSecond(binary.LittleEndian.Uint32(id[8:]))
}
func (id *UUID) setSecondElapsed(sec duration.Second) {
	binary.LittleEndian.PutUint64(id[0:], uint64(sec))
}
func (id *UUID) setNanoInSecondElapsed(nsec duration.NanoInSecond) {
	binary.LittleEndian.PutUint32(id[8:], uint32(nsec))
}

func IDfromString(IDasString string) (id uint64, err protocol.Error) {
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
