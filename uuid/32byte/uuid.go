/* For license and copyright information please see LEGAL file in repository */

package uuid

import (
	"crypto/rand"
	"encoding/base64"
	"io"

	"golang.org/x/crypto/sha3"

	"github.com/GeniusesGroup/libgo/binary"
	"github.com/GeniusesGroup/libgo/protocol"
	"github.com/GeniusesGroup/libgo/time/unix"
)

type UUID [32]byte

func (id UUID) UUID() [32]byte     { return id }
func (id UUID) ID() protocol.ID    { return id.id() }
func (id UUID) IDasString() string { return base64.RawURLEncoding.EncodeToString(id[:8]) }
func (id UUID) ToString() string   { return "TODO:::" }
func (id UUID) ExistenceTime() protocol.Time {
	var time unix.Time
	time.ChangeTo(unix.SecElapsed(id.secondElapsed()), id.nanoSecondElapsed())
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
	id.setNanoSecondElapsed(now.NanoSecondElapsed())
}

// NewHash generate 32 byte incremental by time + hash of data UUID
// CAUTION::: Use for ObjectID in a clustered software cause all writes always go to one node!
// 99.999% collision free on distribution generation.
func (id *UUID) NewHash(data []byte) {
	*id = sha3.Sum256(data)
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

func (id UUID) id() protocol.ID             { return protocol.ID(binary.LittleEndian.Uint64(id[0:])) }
func (id UUID) secondElapsed() int64        { return int64(binary.LittleEndian.Uint64(id[0:])) }
func (id UUID) nanoSecondElapsed() int32    { return int32(binary.LittleEndian.Uint32(id[8:])) }
func (id *UUID) setSecondElapsed(sec int64) { binary.LittleEndian.PutUint64(id[0:], uint64(sec)) }
func (id *UUID) setNanoSecondElapsed(nsec int32) {
	binary.LittleEndian.PutUint32(id[8:], uint32(nsec))
}
