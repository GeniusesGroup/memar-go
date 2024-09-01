/* For license and copyright information please see the LEGAL file in the code repository */

package ut16

import (
	"crypto/rand"
	"encoding/base64"
	"io"

	"golang.org/x/crypto/sha3"

	"memar/binary"
	"memar/protocol"
	"memar/time/duration"
	time_p "memar/time/protocol"
	"memar/time/unix"
)

type UUID [16]byte

func (id UUID) UUID() [16]byte { return id }
func (id UUID) ID() [4]byte    { return id.id() }
func (id UUID) ExistenceTime() time_p.Time {
	var time = id.ExistenceUnixTime()
	return &time
}

func (id UUID) ExistenceUnixTime() (time unix.Time) {
	time.ChangeTo(id.secondElapsed(), id.nanoSecondElapsed())
	return
}

//memar:impl memar/string/protocol.Stringer
func (id UUID) ToString() (s string, err protocol.Error) {
	s = base64.RawURLEncoding.EncodeToString(id[:])
	return
}
func (id *UUID) FromString(s string) (err protocol.Error) {
	// TODO:::
	return
}

// New will generate 16 byte time based UUID.
// **CAUTION**: Use for ObjectID in a clustered software without any hash cause all writes always go to one node.
// 99.999999% collision free on distribution generation.
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

// NewHash generate 16 byte incremental by time + hash of data UUID
// CAUTION::: Use for ObjectID in a clustered software cause all writes always go to one node!
// 99.999% collision free on distribution generation.
func (id *UUID) NewHash(data []byte) {
	var uuid32 = sha3.Sum256(data)
	copy(id[12:], uuid32[:])

	// Set time to UUID
	var now = unix.Now()
	id.setSecondElapsed(now.SecondElapsed())
	id.setNanoInSecondElapsed(now.NanoInSecondElapsed())
}

// NewRandom generate 16 byte random UUID.
func (id *UUID) NewRandom() {
	var err error
	_, err = io.ReadFull(rand.Reader, id[:])
	if err != nil {
		// TODO::: make random by other ways
	}
}

func (id UUID) id() (rid [4]byte) { copy(rid[:], id[12:]); return }
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
