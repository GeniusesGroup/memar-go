/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

import (
	container_p "memar/adt/container/protocol"
	error_p "memar/process/error/protocol"
)

type Frame interface {
	// TODO::: due to method need custom args in each frame type, we can't uncomment bellow easily.
	// StaticFrameLen(args) int

	// FrameLen or FrameLength
	FrameLength() container_p.NumberOfElement
	NextFrame() []byte // Frame
	// Buffer

	Process(sk Socket) (err error_p.Error)
	Do(sk Socket) (err error_p.Error)
}

type FrameWriter interface {
	WriteFrame(packet Packet) (n container_p.NumberOfElement, err error_p.Error)
}
