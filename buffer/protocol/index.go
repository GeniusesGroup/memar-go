/* For license and copyright information please see the LEGAL file in the code repository */

package buffer_p

import (
	container_p "memar/adt/container/protocol"
)

type Index = container_p.ElementIndex

type Field_Indexes interface {
	ReadIndex() Index
	WriteIndex() Index
}

type Method_Indexes interface {
	SetReadIndex(ri Index)
	SetWriteIndex(wi Index)
}
