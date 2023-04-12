/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"libgo/protocol"
)

type Frame []byte

// Init initialize new unicast||broadcast frame.
func (f *Frame) Init(nexHeaderID NextHeaderID, path []byte, payloadLen int) {
	var pathLen byte = byte(len(path))
	if pathLen == 0 {
		pathLen = maxHopCount // broadcast frame
	}
	var payloadLoc int = int(fixedHeaderLength + pathLen)
	var frameLength int = payloadLoc + payloadLen
	// TODO::: get from pool??
	*f = make([]byte, frameLength)

	f.SetHopCount(pathLen)
	f.SetNextHeader(byte(nexHeaderID))
	// Set path for unicast. it will not copy if path is 0 for broadcast frame as we want!
	f.SetPath(path)
	return
}
func (f *Frame) Reinit() {}
func (f *Frame) Deinit() {
	// TODO::: back it to pool??
}

// Getter methods to get frame fields.
func (f Frame) NextHop() byte { return f[0] }
func (f Frame) HopCount() byte {
	if f.IsBroadcastFrame() {
		return maxHopCount // broadcast frame
	}
	return f[1]
}
func (f Frame) NextHeader() protocol.NetworkTransport_HeaderID {
	return protocol.NetworkTransport_HeaderID(f[2])
}
func (f Frame) NextPortNum() byte        { return f[fixedHeaderLength+f.NextHop()] }
func (f Frame) PortNum(hopNum byte) byte { return f[fixedHeaderLength+hopNum] }
func (f Frame) Path() []byte             { return f[fixedHeaderLength : fixedHeaderLength+f.HopCount()] }
func (f Frame) Payload() []byte          { return f[fixedHeaderLength+f.HopCount():] }

// Setter methods to set frame fields.
func (f Frame) SetHopCount(hopCount byte)            { f[1] = hopCount }
func (f Frame) SetNextHeader(linkHeaderID byte)      { f[2] = linkHeaderID }
func (f Frame) SetPortNum(hopNum byte, portNum byte) { f[fixedHeaderLength+hopNum] = portNum }
func (f Frame) SetPath(path []byte)                  { copy(f[fixedHeaderLength:], path) }

// CheckFrame checks frame for any bad situation.
// Always check frame before use any other frame methods otherwise panic may occur.
func (f Frame) CheckFrame() (err protocol.Error) {
	var ln = len(f)
	if ln < MinFrameLen {
		return &ErrShortFrameLength
	}
	if ln > MaxFrameLen {
		return &ErrLongFrameLength
	}
	return
}

// IsBroadcastFrame checks frame is broadcast frame or not.
// spec: https://github.com/GeniusesGroup/RFCs/blob/master/Chapar.md#frame-types
func (f Frame) IsBroadcastFrame() bool { return f[1] == broadcastHopCount }

// IncrementNextHop sets received port number and increment NextHop number in frame.
func (f Frame) IncrementNextHop(receivedPortNumber byte) (lastHop bool) {
	var nextHop = f.NextHop()
	var hopCount = f.HopCount()
	// spec: https://github.com/GeniusesGroup/RFCs/blob/master/Chapar.md#rules
	f.SetPortNum(nextHop, receivedPortNumber)
	if nextHop == hopCount-1 { // -1 due to hop count start from 1 not 0
		return true
	}
	f[0]++
	return false
}

// NewFrame makes new unicast||broadcast frame.
func NewFrame(nexHeaderID NextHeaderID, path []byte, payloadLen int) (frame Frame) {
	frame.Init(nexHeaderID, path, payloadLen)
	return
}
