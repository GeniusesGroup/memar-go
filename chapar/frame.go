/* For license and copyright information please see LEGAL file in repository */

package chapar

import (
	"../protocol"
)

const (
	// MinFrameLen is minimum Chapar frame length
	// 4 Byte Chapar header + 8 Byte min payload
	MinFrameLen = 12
	// MaxFrameLen is maximum Chapar frame length
	MaxFrameLen = 8192

	fixedHeaderLength      byte = 3 // without path part
	maxHopCount            byte = 255
	broadcastHopCount      byte = 0
	maxBroadcastPayloadLen      = MaxFrameLen - (fixedHeaderLength + maxHopCount)
)

// MakeNewFrame makes new unicast||broadcast frame!
func MakeNewFrame(nexHeaderID protocol.NetworkLinkNextHeaderID, path []byte, payload []byte) (frame []byte) {
	var pathLen byte = byte(len(path))
	if pathLen == 0 {
		pathLen = maxHopCount // broadcast frame
	}
	var payloadLoc int = int(fixedHeaderLength + pathLen)
	var frameLength int = payloadLoc + len(payload)
	frame = make([]byte, frameLength)

	SetHopCount(frame, pathLen)
	SetNextHeader(frame, byte(nexHeaderID))
	// Set path for unicast. it will not copy if path is 0 for broadcast frame as we want!
	SetPath(frame, path)
	// copy payload to frame
	copy(frame[payloadLoc:], payload)
	return
}

// CheckFrame checks frame for any bad situation!
// Always check frame before use any other frame methods otherwise panic may occur!
func CheckFrame(frame []byte) (err protocol.Error) {
	var len = len(frame)
	if len < MinFrameLen {
		return ErrShortFrameLength
	}
	if len > MaxFrameLen {
		return ErrLongFrameLength
	}
	return
}

// GetNextHop returns next hop number.
func GetNextHop(frame []byte) byte {
	return frame[0]
}

// IncrementNextHop sets received port number and increment NextHop number in frame!
func IncrementNextHop(frame []byte, receivedPortNumber byte) {
	// spec: https://github.com/GeniusesGroup/RFCs/blob/master/Chapar.md#rules
	SetPortNum(frame, frame[0], receivedPortNumber)
	frame[0]++
}

// GetHopCount returns the number of intermediate network devices indicate in frame.
func GetHopCount(frame []byte) byte {
	if IsBroadcastFrame(frame) {
		return maxHopCount
	}
	return frame[1]
}

// SetHopCount writes given hop number to the frame.
func SetHopCount(frame []byte, hopCount byte) {
	frame[1] = hopCount
}

// IsBroadcastFrame checks frame is broadcast frame or not!
// spec: https://github.com/GeniusesGroup/RFCs/blob/master/Chapar.md#frame-types
func IsBroadcastFrame(frame []byte) bool {
	return frame[1] == broadcastHopCount
}

// GetNextHeader returns next header ID of the frame.
func GetNextHeader(frame []byte) byte {
	return frame[2]
}

// SetNextHeader sets next header id in the frame.
func SetNextHeader(frame []byte, linkHeaderID byte) {
	frame[2] = linkHeaderID
}

// GetNextPortNum returns next port number set in the frame.
func GetNextPortNum(frame []byte) byte {
	return frame[fixedHeaderLength+GetNextHop(frame)]
}

// GetPortNum returns port number of given hop number.
// First hopNum is hopNum==1 not hopNum==0. Don't read hopNum==0 due to it is use for broadcast frame.
func GetPortNum(frame []byte, hopNum byte) byte {
	return frame[fixedHeaderLength+hopNum]
}

// SetPortNum set given port number in given hop number!
// First hopNum is hopNum==1 not hopNum==0. Don't set hopNum==0 due to it is use for broadcast frame.
func SetPortNum(frame []byte, hopNum byte, portNum byte) {
	frame[fixedHeaderLength+hopNum] = portNum
}

// GetPath gets frame path in all hops.
func GetPath(frame []byte) []byte {
	return frame[fixedHeaderLength : fixedHeaderLength+GetHopCount(frame)]
}

// SetPath sets given path in given the frame.
func SetPath(frame []byte, path []byte) {
	copy(frame[fixedHeaderLength:], path[:])
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
	return frame[fixedHeaderLength+GetHopCount(frame):]
}
