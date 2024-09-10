/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

import (
	adt_p "memar/adt/protocol"
	error_p "memar/error/protocol"
)

type Frame interface {
	// TODO::: due to method need custom args in each frame type, we can't uncomment bellow easily.
	// StaticFrameLen(args) int

	// FrameLen or FrameLength
	FrameLength() adt_p.NumberOfElement
	NextFrame() []byte // Frame
	// Buffer

	Process(sk Socket) (err error_p.Error)
	Do(sk Socket) (err error_p.Error)
}

type Framer interface {
	FrameType() FrameType
}
type FrameWriter interface {
	WriteFrame(packet Packet) (n adt_p.NumberOfElement, err error_p.Error)
}
