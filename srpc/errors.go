/* For license and copyright information please see LEGAL file in repository */

package srpc

import "errors"

// Declare Errors Details
var (
	ErrSRPCPacketTooShort = errors.New("Received sRPC Packet size is smaller than expected and can't use")
)
