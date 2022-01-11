/* For license and copyright information please see LEGAL file in repository */

package user

import (
	"crypto/rand"
	"time"

	"../protocol"
)

func NewID(userType protocol.UserType) (id UUID) {
	setTime(id, time.Now().UnixMilli())
	id[8] = byte(userType)
	rand.Read(id[9:])
	return
}

type UUID [16]byte

func (id UUID) UUID() [16]byte                        { return id }
func (id UUID) ExistenceTime() protocol.TimeUnixMilli { return protocol.TimeUnixMilli(getTime(id)) }
func (id UUID) Type() protocol.UserType               { return protocol.UserType(id[8]) }
func (id UUID) ID() uint64                            { return getID(id) }

func getTime(uuid [16]byte) int64 {
	return int64(uuid[0]) | int64(uuid[1])<<8 | int64(uuid[2])<<16 | int64(uuid[3])<<24 | int64(uuid[4])<<32 | int64(uuid[5])<<40 | int64(uuid[6])<<48 | int64(uuid[7])<<56
}
func getID(uuid [16]byte) uint64 {
	return uint64(uuid[9]) | uint64(uuid[10])<<8 | uint64(uuid[11])<<16 | uint64(uuid[12])<<24 | uint64(uuid[13])<<32 | uint64(uuid[14])<<40 | uint64(uuid[15])<<48
}
func setTime(uuid [16]byte, time int64) {
	uuid[0] = byte(time)
	uuid[1] = byte(time >> 8)
	uuid[2] = byte(time >> 16)
	uuid[3] = byte(time >> 24)
	uuid[4] = byte(time >> 32)
	uuid[5] = byte(time >> 40)
	uuid[6] = byte(time >> 48)
	uuid[7] = byte(time >> 56)
}
