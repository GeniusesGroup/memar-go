/* For license and copyright information please see LEGAL file in repository */

package chapar

import "errors"

/*
-------------------------------NOTICE:-------------------------------
This protocol usually implement in hardware level not software!
We just want to show simpilicity of protocol and its functions here!
*/

const (
	// MinFrameLen is minimum Chapar frame length
	// 4 Byte Chapar header + 4 Byte min payload
	MinFrameLen = 7 // Due to len & index mismatch, so exactly MinFrameLen = 8
	// MaxFrameLen is maximum Chapar frame length
	MaxFrameLen = 8191 // Due to len & index mismatch, so exactly MinFrameLen = 8191
)

// Declare Errors Details
var (
	ErrIllegalFrameSize = errors.New("Chapar frame is too short(<8) or too long(>8192) than standard")
)

// MakeNewFrame makes new unicast||broadcast frame!
func MakeNewFrame(hopNum, nexHeaderID byte, path []byte, payloadSize int) (frame []byte) {
	frame = make([]byte, 3+len(path)+payloadSize)
	SetHopsNum(frame, hopNum)
	SetNextHeader(frame, nexHeaderID)
	SetPath(frame, path)
	return
}

// CheckFrame checks frame for any bad situation!
// Always check frame before use any other frame methods otherwise panic may occur!
func CheckFrame(frame []byte) (err error) {
	var len = len(frame)
	if len < MinFrameLen || len > MaxFrameLen {
		return ErrIllegalFrameSize
	}
	return
}

// GetNextHop returns next hop number.
func GetNextHop(frame []byte) byte {
	return frame[0]
}

// IncrementNextHop sets received port number and increment NextHop number in frame!
func IncrementNextHop(frame []byte, receivedPortNumber byte) {
	// The reasons of use SetPortNum method here:
	// - BroadcastFrame : To improve performance, previous switch just send frame without declare port, we must declare it now!
	// - UnicastFrame : To be sure receive port is same with declaration one in frame, we replace it always!
	// - Rule&Security : To be sure physical network port is same on sender and receiver switch, we must set it again here!
	SetPortNum(frame, frame[0], receivedPortNumber)

	frame[0]++
}

// GetHopsNum returns number of hops of the frame.
func GetHopsNum(frame []byte) byte {
	return frame[1]
}

// SetHopsNum writes given hops number to the frame.
func SetHopsNum(frame []byte, hopsNum byte) {
	frame[1] = hopsNum
}

// IsBroadcastFrame checks frame is broadcast frame or not!
func IsBroadcastFrame(frame []byte) bool {
	// Due to frame must have at least 1 hop so we use unused HopsNum==0 for multicast frames to all ports!
	// So both HopsNum==0x00 & HopsNum==0xff have 256 SwitchPortNum space in frame header!
	// Frame must have all Switch0PortNum to Switch254PortNum with 0 byte data in header
	// otherwise frame payload rewrite by switches!
	return GetHopsNum(frame) == 0x00
}

// GetNextHeader returns next header ID of the frame.
func GetNextHeader(frame []byte) byte {
	return frame[2]
}

// SetNextHeader sets next header id in the frame.
func SetNextHeader(frame []byte, nextHeaderID byte) {
	frame[2] = nextHeaderID
}

// GetPortNum returns port number of given hop number.
func GetPortNum(frame []byte, hopNum byte) byte {
	return frame[3+hopNum]
}

// SetPortNum set given port number in given hop number!
func SetPortNum(frame []byte, hopNum byte, portNum byte) {
	frame[3+hopNum] = portNum
}

// GetPath gets frame path in all hops.
func GetPath(frame []byte) []byte {
	return frame[3 : 3+GetHopsNum(frame)]
}

// SetPath sets given path in given the frame.
func SetPath(frame []byte, path []byte) {
	copy(frame[3:], path[:])
}

// GetReversePath gets frame path in all hops just in reverse.
func GetReversePath(frame []byte) []byte {
	return ReversePath(GetPath(frame))
}

// ReversePath returns a copy of given path in reverse.
func ReversePath(path []byte) (ReversePath []byte) {
	var i = len(path)
	ReversePath = make([]byte, i)
	i-- // Due to len & index mismatch
	var j int
	for i >= j {
		ReversePath[i], ReversePath[j] = path[j], path[i]
		i--
		j++
	}
	return
}

// GetPayload returns payload.
func GetPayload(frame []byte) []byte {
	return frame[GetHopsNum(frame)+3:]
}
