/* For license and copyright information please see LEGAL file in repository */

package user

import (
	"crypto/rand"

	"../binary"
	"../protocol"
	"../time/unix"
)

func NewID(userType protocol.UserType) (id UUID) {
	var now = unix.Now()
	id.setSecondElapsed(now.SecondElapsed())
	id.setNanoSecondElapsed(now.NanoSecondElapsed())
	id[12] = byte(userType)
	// TODO::: moved to heap: id, Is it any way to prevent heap alloc??
	rand.Read(id[13:])
	return
}

// 0...7 >> second elapsed
// 8...11 >> nano second elapsed
// 12...12 >> type
// 13...15 >> random id
type UUID [16]byte

func (id UUID) UUID() [16]byte          { return id }
func (id UUID) Type() protocol.UserType { return protocol.UserType(id[8]) }
func (id UUID) ID() [3]byte             { return id.id() }
func (id UUID) ExistenceTime() protocol.Time {
	var time unix.Time
	time.ChangeTo(unix.SecElapsed(id.secondElapsed()), id.nanoSecondElapsed())
	return &time
}

func (id UUID) id() (rid [3]byte)               { copy(rid[:], id[13:]); return }
func (id UUID) secondElapsed() int64            { return int64(binary.LittleEndian.Uint64(id[0:])) }
func (id UUID) nanoSecondElapsed() int32        { return int32(binary.LittleEndian.Uint32(id[8:])) }
func (id UUID) setSecondElapsed(sec int64)      { binary.LittleEndian.PutUint64(id[0:], uint64(sec)) }
func (id UUID) setNanoSecondElapsed(nsec int32) { binary.LittleEndian.PutUint32(id[8:], uint32(nsec)) }
