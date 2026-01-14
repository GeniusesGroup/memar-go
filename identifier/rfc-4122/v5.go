/* For license and copyright information please see the LEGAL file in the code repository */

package uuid_rfc4122

import (
	error_p "memar/process/error/protocol"
)

// V5 generate version 5 RFC4122 include hash namespace and value
type V5 UUID

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self V5) Init(nameSpace [16]byte, value []byte) (err error_p.Error) {
	return
}
