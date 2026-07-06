/* For license and copyright information please see the LEGAL file in the code repository */

package json_p

import (
	container_p "memar/adt/container/protocol"
)

// Field_Length is same as CodecLength
type Field_Length interface {
	// JSON_Length return whole calculated length of JSON encoded of the struct
	// It is NOT include of first and last curly braces as `{}`
	JSON_Length() Length
}

type Length = container_p.NumberOfElement
