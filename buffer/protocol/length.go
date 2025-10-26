/* For license and copyright information please see the LEGAL file in the code repository */

package buffer_p

import (
	container_p "memar/adt/container/protocol"
)

type Length = container_p.NumberOfElement

type Field_Lengths interface {
	// UnreadLength returns how many bytes are not read(ReadIndex to WriteIndex) in the buffer.
	UnreadLength() Length
}
