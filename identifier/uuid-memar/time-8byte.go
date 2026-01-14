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

type Time8 [8]byte

func (self Time8) UUID() [8]byte      { return self }
func (self Time8) ID() [3]byte        { return self.id() }
func (self Time8) IDasString() string { return base64.RawURLEncoding.EncodeToString(self[:8]) }
func (self Time8) ExistenceTime() time_p.Time {
	var time unix.Time
	time.ChangeTo(self.secondElapsed(), 0)
	return &time
}

//memar:impl memar/codec/string/protocol.Stringer
func (self Time8) ToString() (s string, err error_p.Error) {
	s = base64.RawURLEncoding.EncodeToString(self[:])
	return
}
func (self *Time8) FromString(s string) (err error_p.Error) {
	// TODO:::
	return
}

// New will generate 8 byte time based UUID.
// **CAUTION**: Use for ObjectID in a clustered software without any hash cause all writes always go to one node.
// 99.999999% collision free on distribution generation.
func (self *Time8) New() {
	var err error
	_, err = io.ReadFull(rand.Reader, self[5:])
	if err != nil {
		// TODO::: make random by other ways
	}

	// Set time to UUID
	var now = unix.Now()
	self.setSecondElapsed(now.SecondElapsed())
}

// NewHash generate 8 byte incremental by time + hash of data UUID
// CAUTION::: Use for ObjectID in a clustered software cause all writes always go to one node!
// 99.999% collision free on distribution generation.
func (self *Time8) NewHash(data []byte) {
	var uuid32 = sha3.Sum256(data)
	copy(self[5:], uuid32[:])

	// Set time to UUID
	var now = unix.Now()
	self.setSecondElapsed(now.SecondElapsed())
}

// NewRandom generate 8 byte random UUID.
func (self *Time8) NewRandom() {
	var err error
	_, err = io.ReadFull(rand.Reader, self[:])
	if err != nil {
		// TODO::: make random by other ways
	}
}

func (self Time8) id() (rid [3]byte) { copy(rid[:], self[5:]); return }
func (self Time8) secondElapsed() duration.Second {
	var sec [8]byte
	copy(sec[:], self[:])
	return duration.Second(binary.LittleEndian.Uint64(self[0:])) >> (64 - 40)
}
func (self *Time8) setSecondElapsed(sec duration.Second) {
	binary.LittleEndian.PutUint64(self[0:], (uint64(sec) << (64 - 40)))
}
