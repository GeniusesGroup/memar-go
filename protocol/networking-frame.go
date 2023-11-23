/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

const (
	Network_FrameType_Length = 1 // byte get on byte space
)

// Network_FrameType is Frame type ID like service ID but fixed ID with 8bit length. Just some few services get one byte length service ID
// Common services must register by 64bit unsigned integer.
type Network_FrameType byte

type Network_Frame interface {
	// TODO::: due to method need custom args in each frame type, we can't uncomment bellow easily!
	// StaticFrameLen(args) int

	// FrameLen or FrameLength
	FrameLen() int
	NextFrame() []byte // Network_Frame

	Process(sk Socket) (err Error)
	Do(sk Socket) (err Error)
}

type Network_Framer interface {
	FrameType() Network_FrameType
}
type Network_FrameWriter interface {
	WriteFrame(packet []byte) (n int, err Error)
}

// https://github.com/GeniusesGroup/memar/blob/main/networking.md#frames-number
const (
	Network_FrameType_Unset Network_FrameType = iota
	Network_FrameType_Asb
	Network_FrameType_Chapar
	Network_FrameType_GP
	// Network_FrameType_Ethernet
	// Network_FrameType_ATM

	Network_FrameType_Padding // Network_FrameType = 128 + iota
	Network_FrameType_CallService
	Network_FrameType_OpenStream
	Network_FrameType_CloseStream
	Network_FrameType_Data
	Network_FrameType_Error
	Network_FrameType_Signature
)
