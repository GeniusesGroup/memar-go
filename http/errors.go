/* For license and copyright information please see LEGAL file in repository */

package http

import "errors"

// Declare Errors Details
var (
	ErrHTTPPacketTooShort = errors.New("Received HTTP Packet size is smaller than expected and can't use")
	ErrHTTPPacketTooLong = errors.New("Received HTTP Packet size is larger than expected and can't use")
)
