/* For license and copyright information please see LEGAL file in repository */

package uip

import "errors"

// Declare Errors Details
var (
	ErrUIPPacketTooShort = errors.New("UIP packet is empty or too short than standard header. It must include 44Byte header plus 16Byte min Payload")
)
