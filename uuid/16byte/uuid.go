/* For license and copyright information please see LEGAL file in repository */

package uuid

import (
	"crypto/rand"
	"io"

	"golang.org/x/crypto/sha3"

	"github.com/GeniusesGroup/libgo/binary"
	"github.com/GeniusesGroup/libgo/protocol"
	"github.com/GeniusesGroup/libgo/time/unix"
)

type UUID [16]byte

func (id UUID) UUID() [16]byte { return id }
func (id UUID) ID() [4]byte    { return id.id() }
func (id UUID) ExistenceTime() protocol.Time {
	var time unix.Time
	time.ChangeTo(unix.SecElapsed(id.secondElapsed()), id.nanoSecondElapsed())
	return &time
}
func (id UUID) ToString() string { return "TODO:::" }

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
	id.setNanoSecondElapsed(now.NanoSecondElapsed())
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
	id.setNanoSecondElapsed(now.NanoSecondElapsed())
}

// NewRandom generate 16 byte random UUID.
func (id *UUID) NewRandom() {
	var err error
	_, err = io.ReadFull(rand.Reader, id[:])
	if err != nil {
		// TODO::: make random by other ways
	}
}

func (id UUID) id() (rid [4]byte)           { copy(rid[:], id[12:]); return }
func (id UUID) secondElapsed() int64        { return int64(binary.LittleEndian.Uint64(id[0:])) }
func (id UUID) nanoSecondElapsed() int32    { return int32(binary.LittleEndian.Uint32(id[8:])) }
func (id *UUID) setSecondElapsed(sec int64) { binary.LittleEndian.PutUint64(id[0:], uint64(sec)) }
func (id *UUID) setNanoSecondElapsed(nsec int32) {
	binary.LittleEndian.PutUint32(id[8:], uint32(nsec))
}
