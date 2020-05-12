/* For license and copyright information please see LEGAL file in repository */

package persiaos

// GPIncomeBufferPool store
var GPIncomeBufferPool GPBufferPool

// GPBufferPool use to
type GPBufferPool struct {
	PoolCapacity       uint16
	ActiveFreeLocation uintptr
	LastPacketGet      uintptr
	Pool               uintptr
}

func init() {
	// make array by available memory and set it in pool
	// available space in pool must respect MTU too! e.g. if MTU = 8192 make like [256][8192]byte
}
