/* For license and copyright information please see LEGAL file in repository */

package gp

// Addr present GP address with needed methods
type Addr [14]byte

// Some global address
var (
	AddrNil = Addr{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
)

// GetSocietyID returns society ID.
func (addr Addr) GetSocietyID() uint32 {
	return uint32(addr[0]) | uint32(addr[1])<<8 | uint32(addr[2])<<16 | uint32(addr[3])<<24
}

// GetRouterID returns router ID in the society scope.
func (addr Addr) GetRouterID() uint32 {
	return uint32(addr[4]) | uint32(addr[5])<<8 | uint32(addr[6])<<16 | uint32(addr[7])<<24
}

// GetUserID returns user ID in the router scope.
func (addr Addr) GetUserID() uint32 {
	return uint32(addr[8]) | uint32(addr[9])<<8 | uint32(addr[10])<<16 | uint32(addr[11])<<24
}

// GetAppID returns app ID in the user scope.
func (addr Addr) GetAppID() uint16 {
	return uint16(addr[12]) | uint16(addr[13])<<8
}

// SetSocietyID set society ID in given Addr.
func (addr Addr) SetSocietyID(id uint32) {
	addr[0] = byte(id)
	addr[1] = byte(id >> 8)
	addr[2] = byte(id >> 16)
	addr[3] = byte(id >> 24)
}

// SetRouterID set router ID in given Addr.
func (addr Addr) SetRouterID(id uint32) {
	addr[4] = byte(id)
	addr[5] = byte(id >> 8)
	addr[6] = byte(id >> 16)
	addr[7] = byte(id >> 24)
}

// SetUserID set user ID in given Addr.
func (addr Addr) SetUserID(id uint32) {
	addr[8] = byte(id)
	addr[9] = byte(id >> 8)
	addr[10] = byte(id >> 16)
	addr[11] = byte(id >> 24)
}

// SetAppID set app ID in given Addr.
func (addr Addr) SetAppID(id uint32) {
	addr[12] = byte(id)
	addr[13] = byte(id >> 8)
}
