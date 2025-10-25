/* For license and copyright information please see the LEGAL file in the code repository */

package codec_p

import (
	container_p "memar/adt/container/protocol"
)

type Field_Length interface {
	// SerializationLength return value ln, that is the max number of bytes that will written as encode data by Encode()||Marshal()
	// 0 means no data and -1 means can't tell until full write.
	// Due to prevent performance penalty, Implementors can return max number instead of actual number of length.
	SerializationLength() Length
}

type Length = container_p.NumberOfElement
