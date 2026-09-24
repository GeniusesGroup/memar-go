/* For license and copyright information please see the LEGAL file in the code repository */

package buffer_p

import (
	container_p "memar/computer/adt/container/protocol"
)

type Field_Indexes interface {
	ReadIndex() Index
	WriteIndex() Index
}

type Method_Indexes interface {
	SetReadIndex(ri Index)
	SetWriteIndex(wi Index)
}

type Index = container_p.ElementIndex
