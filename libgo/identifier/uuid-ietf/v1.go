/* For license and copyright information please see the LEGAL file in the code repository */

package uuid_rfc4122

import (
	error_p "memar/process/error/protocol"
)

// V1 generate version 1 RFC4122 include date-time and MAC address.
// Use V1 for massive data that don't need to read much specially very close to write time.
// These type of records write sequently in cluster and don't need very much move in cluster expand process.
type V1 UUID

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self V1) Init() (err error_p.Error) {
	return
}
