/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

const (
	Network_FrameID_Length = 1 // byte get on byte space
)

// Network_FrameID is Frame type ID like service ID but fixed ID with 8bit length. Just some few services get one byte length service ID
// Common services must register by 64bit unsigned integer.
type Network_FrameID byte

type Network_RawFrame []byte

type Network_Frame interface {
	// TODO::: due to method need custom args in each frame type, we can't uncomment bellow easily!
	// StaticFrameLen(args) int

	// FrameLen or FrameLength
	FrameLen() int
	NextFrame() []byte

	Process(soc Socket) (err Error)
	Do(soc Socket) (err Error)
}

type Network_Framer interface {
	FrameID() (fID Network_FrameID)
}
type Network_FrameWriter interface {
	WriteFrame(packet []byte) (n int, err Error)
}

// https://github.com/GeniusesGroup/RFCs/blob/master/networking.md#frames-number
const (
	Network_FrameID_Unset Network_FrameID = iota
	Network_FrameID_Asb
	Network_FrameID_Chapar
	Network_FrameID_GP
	// Network_FrameID_Ethernet
	// Network_FrameID_ATM

	Network_FrameID_Padding // Network_FrameID = 128 + iota
	Network_FrameID_Ping
	Network_FrameID_CallService
	Network_FrameID_OpenStream
	Network_FrameID_CloseStream
	Network_FrameID_Data
	Network_FrameID_Error
	Network_FrameID_Signature
)
