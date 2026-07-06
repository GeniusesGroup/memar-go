/* For license and copyright information please see the LEGAL file in the code repository */

package uuid_rfc4122

import (
	"crypto/rand"
	"io"

	error_p "memar/process/error/protocol"
)

// V4 generate version 4 RFC4122 include randomly numbers.
type V4 UUID

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self V4) Init() (err error_p.Error) {
	var goErr error
	_, goErr = io.ReadFull(rand.Reader, self[:])
	if goErr != nil {
		// err = &
		return
	}

	// Set version to 4
	self[6] = (self[6] & 0x0f) | (0x04 << 4)
	// Set variant to RFC4122
	self[8] = (self[8]&(0xff>>2) | (0x02 << 6))
	return
}
