/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

import "errors"

// Declare Errors Details
var (
	ErrUIPPacketTooShort = errors.New("UIP packet is empty or too short than standard header. It must include 44Byte header plus 16Byte min Payload")

	ErrSRPCServiceNotFound = errors.New("Requested sRPC Service is out range of services in this version of service")
	ErrSRPCPayloadEmpty    = errors.New("Stream data payload can't be empty")

	ErrStreamPayloadEmpty = errors.New("Stream data payload can't be empty")
)
