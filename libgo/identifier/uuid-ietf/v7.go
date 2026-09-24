/* For license and copyright information please see the LEGAL file in the code repository */

package uuid_rfc4122

import (
	error_p "memar/process/error/protocol"
)

// V7 generate version 7 of Internet Engineering Task Force (IETF)
type V7 UUID

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self V5) Init(nameSpace [16]byte, value []byte) (err error_p.Error) {
	return
}
