/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"memar/protocol"
)

type Frame []byte

// Init initialize new unicast||broadcast frame.
//
//memar:impl memar/protocol.ObjectLifeCycle
func (f *Frame) Init(path []byte) (err protocol.Error) {
	var pathLen byte = byte(len(path))
	if pathLen == 0 {
		pathLen = maxHopCount // broadcast frame
	}

	f.SetFrameType(protocol.Network_FrameType_Chapar)
	f.SetHopCount(pathLen)
	// Set path for unicast. it will not copy if path is 0 for broadcast frame as we want.
	f.SetPath(path)
	return
}
func (f *Frame) Reinit() (err protocol.Error) {
	// TODO::: ??
	return
}
func (f *Frame) Deinit() (err protocol.Error) {
	// TODO::: ??
	return
}

//memar:impl memar/protocol.Network_Framer
func (f Frame) FrameType() protocol.Network_FrameType { return protocol.Network_FrameType(f[0]) }

// Getter methods to get frame fields.
func (f Frame) HopCount() byte {
	if f.IsBroadcastFrame() {
		return maxHopCount // broadcast frame
	}
	return f[1]
}
func (f Frame) NextHop() byte            { return f[2] }
func (f Frame) NextPortNum() byte        { return f[frameFixedLength+f.NextHop()] }
func (f Frame) PortNum(hopNum byte) byte { return f[frameFixedLength+hopNum] }
func (f Frame) Path() []byte             { return f[frameFixedLength : frameFixedLength+f.HopCount()] }

// Setter methods to set frame fields.
func (f Frame) SetFrameType(fID protocol.Network_FrameType) { f[0] = byte(fID) }
func (f Frame) SetHopCount(hopCount byte)                   { f[1] = hopCount }
func (f Frame) SetPortNum(hopNum byte, portNum byte)        { f[frameFixedLength+hopNum] = portNum }
func (f Frame) SetPath(path []byte)                         { copy(f[frameFixedLength:], path) }

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
	if f.FrameType() != protocol.Network_FrameType_Chapar {
		return &ErrBadFrameType
	}
	return
}

// IsBroadcastFrame checks frame is broadcast frame or not.
// spec: https://github.com/GeniusesGroup/memar/blob/main/networking-osi_2-Chapar.md#frame-types
func (f Frame) IsBroadcastFrame() bool { return f[1] == broadcastHopCount }

// IncrementNextHop sets received port number and increment NextHop number in frame.
func (f Frame) IncrementNextHop(receivedPortNumber byte) (lastHop bool) {
	var nextHop = f.NextHop()
	var hopCount = f.HopCount()
	// spec: https://github.com/GeniusesGroup/memar/blob/main/networking-osi_2-Chapar.md#rules
	f.SetPortNum(nextHop, receivedPortNumber)
	if nextHop == hopCount-1 { // -1 due to hop count start from 1 not 0
		return true
	}
	f[2]++
	return false
}

//memar:impl memar/protocol.Network_Frame
func (f Frame) StaticFrameLen(pathLen byte) (frameLength int) { return int(frameFixedLength + pathLen) }
func (f Frame) FrameLen() (frameLength int)                   { return int(frameFixedLength + f.HopCount()) }
func (f Frame) NextFrame() []byte                             { return f[frameFixedLength+f.HopCount():] }
func (f Frame) Process(soc protocol.Socket) (err protocol.Error) {
	return
}
func (f Frame) Do(soc protocol.Socket) (err protocol.Error) {
	return
}
