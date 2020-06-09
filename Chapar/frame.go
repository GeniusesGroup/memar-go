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
	MaxFrameLen = 8191 // Due to len & index mismatch, so exactly MinFrameLen = 8192
)

// Declare Errors Details
var (
	ErrIllegalFrameSize = errors.New("Chapar frame is too short(<8) or too long(>8192) than standard")
)

// MakeNewFrame makes new unicast||broadcast frame!
func MakeNewFrame(nexHeaderID byte, path []byte, payloadSize int) (frame []byte, payloadLoc int) {
	var pathLen = len(path)
	if pathLen == 0 {
		pathLen = 255 // broadcast frame
	}
	payloadLoc = 3 + pathLen
	frame = make([]byte, payloadLoc+payloadSize)
	SetHopCount(frame, byte(pathLen))
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
	// spec: https://github.com/SabzCity/RFCs/blob/master/Chapar.md#rules
	SetPortNum(frame, frame[0], receivedPortNumber)
	frame[0]++
}

// GetHopCount returns the number of intermediate network devices indicate in frame.
func GetHopCount(frame []byte) byte {
	if frame[1] == 0 {
		return 255 // broadcast frame
	}
	return frame[1]
}

// SetHopCount writes given hop number to the frame.
func SetHopCount(frame []byte, hopCount byte) {
	frame[1] = hopCount
}

// IsBroadcastFrame checks frame is broadcast frame or not!
func IsBroadcastFrame(frame []byte) bool {
	// spec: https://github.com/SabzCity/RFCs/blob/master/Chapar.md#frame-types
	return frame[1] == 0x00
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
// First hopNum is hopNum==1 not hopNum==0. Don't read hopNum==0 due to it is use for broadcast frame.
func GetPortNum(frame []byte, hopNum byte) byte {
	return frame[3+hopNum]
}

// SetPortNum set given port number in given hop number!
// First hopNum is hopNum==1 not hopNum==0. Don't set hopNum==0 due to it is use for broadcast frame.
func SetPortNum(frame []byte, hopNum byte, portNum byte) {
	frame[3+hopNum] = portNum
}

// GetPath gets frame path in all hops.
func GetPath(frame []byte) []byte {
	return frame[3 : 3+GetHopCount(frame)]
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
	return frame[GetHopCount(frame)+3:]
}

// SetPayload copy payload to frame. make frame by MakeNewFrame() to have needed data!
func SetPayload(frame []byte, payloadLoc int, p []byte) {
	copy(frame[payloadLoc:], p)
}
