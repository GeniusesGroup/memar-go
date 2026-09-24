/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	container_p "memar/computer/adt/container/protocol"
	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
)

type Frame []byte

// Init initialize new unicast||broadcast frame.
//
//memar:impl memar/computer/capsule/protocol.LifeCycle
func (f *Frame) Init(path []byte) (err error_p.Error) {
	var pathLen byte = byte(len(path))
	if pathLen == 0 {
		pathLen = maxHopCount // broadcast frame
	}

	f.setFrameType(net_p.FrameType_Chapar)
	f.SetHopCount(pathLen)
	// Set path for unicast. it will not copy if path is 0 for broadcast frame as we want.
	f.SetPath(path)
	return
}
func (f *Frame) Reinit() (err error_p.Error) {
	// TODO::: ??
	return
}
func (f *Frame) Deinit() (err error_p.Error) {
	// TODO::: ??
	return
}

//memar:impl memar/net/protocol.Field_FrameType
func (f Frame) FrameType() net_p.FrameType { return net_p.FrameType(f[0]) }

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
func (f Frame) setFrameType(fID net_p.FrameType)     { f[0] = byte(fID) }
func (f Frame) SetHopCount(hopCount byte)            { f[1] = hopCount }
func (f Frame) SetPortNum(hopNum byte, portNum byte) { f[frameFixedLength+hopNum] = portNum }
func (f Frame) SetPath(path []byte)                  { copy(f[frameFixedLength:], path) }

// CheckFrame checks frame for any bad situation.
// Always check frame before use any other frame methods otherwise panic may occur.
func (f Frame) CheckFrame() (err error_p.Error) {
	var ln = len(f)
	if ln < MinFrameLen {
		// err = &chapar_errs.ShortFrameLength
		return
	}
	if ln > MaxFrameLen {
		// err =  &chapar_errs.LongFrameLength
		return
	}
	if f.FrameType() != net_p.FrameType_Chapar {
		// err = &chapar_errs.BadFrameType
		return
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
func (f Frame) FrameLength() (frameLength container_p.NumberOfElement) {
	return container_p.NumberOfElement(frameFixedLength + f.HopCount())
}
func (f Frame) NextFrame() []byte { return f[frameFixedLength+f.HopCount():] }
func (f Frame) Process(soc net_p.Socket) (err error_p.Error) {
	return
}
func (f Frame) Do(soc net_p.Socket) (err error_p.Error) {
	return
}
