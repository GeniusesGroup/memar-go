/* For license and copyright information please see LEGAL file in repository */

package chapar

/*
-------------------------------NOTICE:-------------------------------
This protocol must implement in hardware level not software!
We just want to show simpilicity of protocol and its functions here!
*/

// Frame use for frame methods!
type Frame []byte

const (
	// FrameLen is minimum frame length
	// 3 Byte header + 72 Byte min payload
	FrameLen = 75
)

// CheckFrame will check frame for any bad situation!
// Always check frame before use any other frame methods otherwise panic may occur!
func (f Frame) CheckFrame() (err error) {
	if len(f) < FrameLen {
		// return FrameTooShort
	}

	return nil
}

// GetNextHop will return NextHop in memory safe way!
func (f Frame) GetNextHop() (NextHop byte) {
	return f[0]
}

// IncrementNextHop will increment NextHop number in frame!
func (f Frame) IncrementNextHop() {
	f[0]++
}

// GetTotalHop will return TotalHop in memory safe way!
func (f Frame) GetTotalHop() (TotalHop byte) {
	return f[1]
}

// IsBroadcastFrame will check frame is broadcast frame or not!
func (f Frame) IsBroadcastFrame() (broadcast bool) {
	// Due to frame must have at least 1 hop so we use unused TotalHop==0 for multicast farmes to all ports!
	// So both TotalHop==0x00 & TotalHop==0xff have 256 SwitchPortNum space in frame header!
	// Frame must have all Switch0PortNum to Switch254PortNum with 0 byte data in header
	// otherwise frame payload rewrite by switches!
	if f.GetTotalHop() == 0x00 {
		return true
	}
	return false
}

// GetNextHeader will return NextHeader in memory safe way!
func (f Frame) GetNextHeader() (NextHeader byte) {
	return f[2]
}

// GetSwitchPortNum will return SwitchPortNum of i in memory safe way!
func (f Frame) GetSwitchPortNum(i byte) (SwitchPortNum byte) {
	return f[i+3]
}

// SetSwitchPortNum will set SwitchPortNum of i!
func (f Frame) SetSwitchPortNum(i byte, SwitchPortNum byte) {
	f[i+3] = SwitchPortNum
}

// GetPayload will return Payload in memory safe way!
func (f Frame) GetPayload() (Payload []byte) {
	return f[f.GetTotalHop()+3:]
}
