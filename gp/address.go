/* For license and copyright information please see LEGAL file in repository */

package gp

// Addr present GP address with needed methods
type Addr [16]byte

// Some global address
var (
	AddrNil = Addr{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
)

// PlanetID returns planet ID.
func (addr Addr) PlanetID() uint16 {
	return uint16(addr[0]) | uint16(addr[1])<<8
}

// SocietyID returns society ID.
func (addr Addr) SocietyID() uint32 {
	return uint32(addr[2]) | uint32(addr[3])<<8 | uint32(addr[4])<<16 | uint32(addr[5])<<24
}

// RouterID returns router ID in the society scope.
func (addr Addr) RouterID() uint32 {
	return uint32(addr[6]) | uint32(addr[7])<<8 | uint32(addr[8])<<16 | uint32(addr[9])<<24
}

// UserID returns user ID in the router scope.
func (addr Addr) UserID() uint32 {
	return uint32(addr[10]) | uint32(addr[11])<<8 | uint32(addr[12])<<16 | uint32(addr[13])<<24
}

// AppID returns app ID in the user scope.
func (addr Addr) AppID() uint16 {
	return uint16(addr[14]) | uint16(addr[15])<<8
}

// SetPlanetID set planet ID in given Addr.
func (addr Addr) SetPlanetID(id uint16) {
	addr[0] = byte(id)
	addr[1] = byte(id >> 8)
}

// SetSocietyID set society ID in given Addr.
func (addr Addr) SetSocietyID(id uint32) {
	addr[2] = byte(id)
	addr[3] = byte(id >> 8)
	addr[4] = byte(id >> 16)
	addr[5] = byte(id >> 24)
}

// SetRouterID set router ID in given Addr.
func (addr Addr) SetRouterID(id uint32) {
	addr[6] = byte(id)
	addr[7] = byte(id >> 8)
	addr[8] = byte(id >> 16)
	addr[9] = byte(id >> 24)
}

// SetUserID set user ID in given Addr.
func (addr Addr) SetUserID(id uint32) {
	addr[10] = byte(id)
	addr[11] = byte(id >> 8)
	addr[12] = byte(id >> 16)
	addr[13] = byte(id >> 24)
}

// SetAppID set app ID in given Addr.
func (addr Addr) SetAppID(id uint32) {
	addr[14] = byte(id)
	addr[15] = byte(id >> 8)
}
