/* For license and copyright information please see the LEGAL file in the code repository */

package memar_uuid

import (
	"crypto/rand"
	"encoding/base64"
	"io"

	"golang.org/x/crypto/sha3"

	"memar/codec/binary"
	error_p "memar/process/error/protocol"
	"memar/time/duration"
	time_p "memar/time/protocol"
	"memar/time/unix"
)

type Time16 [16]byte

func (self Time16) UUID() [16]byte { return self }
func (self Time16) ID() [4]byte    { return self.id() }
func (self Time16) ExistenceTime() time_p.Time {
	var time = self.ExistenceUnixTime()
	return &time
}

func (self Time16) ExistenceUnixTime() (time unix.Time) {
	time.ChangeTo(self.secondElapsed(), self.nanoSecondElapsed())
	return
}

//memar:impl memar/codec/string/protocol.Stringer
func (self Time16) ToString() (s string, err error_p.Error) {
	s = base64.RawURLEncoding.EncodeToString(self[:])
	return
}
func (self *Time16) FromString(s string) (err error_p.Error) {
	// TODO:::
	return
}

// New will generate 16 byte time based UUID.
// **CAUTION**: Use for ObjectID in a clustered software without any hash cause all writes always go to one node.
// 99.999999% collision free on distribution generation.
func (self *Time16) New() {
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

// NewHash generate 16 byte incremental by time + hash of data UUID
// CAUTION::: Use for ObjectID in a clustered software cause all writes always go to one node!
// 99.999% collision free on distribution generation.
func (self *Time16) NewHash(data []byte) {
	var uuid32 = sha3.Sum256(data)
	copy(self[12:], uuid32[:])

	// Set time to UUID
	var now = unix.Now()
	self.setSecondElapsed(now.SecondElapsed())
	self.setNanoInSecondElapsed(now.NanoInSecondElapsed())
}

// NewRandom generate 16 byte random UUID.
func (self *Time16) NewRandom() {
	var err error
	_, err = io.ReadFull(rand.Reader, self[:])
	if err != nil {
		// TODO::: make random by other ways
	}
}

func (self Time16) id() (rid [4]byte) { copy(rid[:], self[12:]); return }
func (self Time16) secondElapsed() duration.Second {
	return duration.Second(binary.LittleEndian.Uint64(self[0:]))
}
func (self Time16) nanoSecondElapsed() duration.NanoInSecond {
	return duration.NanoInSecond(binary.LittleEndian.Uint32(self[8:]))
}
func (self *Time16) setSecondElapsed(sec duration.Second) {
	binary.LittleEndian.PutUint64(self[0:], uint64(sec))
}
func (self *Time16) setNanoInSecondElapsed(nsec duration.NanoInSecond) {
	binary.LittleEndian.PutUint32(self[8:], uint32(nsec))
}
