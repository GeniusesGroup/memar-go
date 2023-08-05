/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"memar/binary"
	"memar/protocol"
)

// Some global address
var (
	AddrNil = Addr{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
)

// Addr present GP address with needed methods
type Addr [AddrLen]byte

func (addr Addr) PlanetID() uint16  { return binary.BigEndian.Uint16(addr[0:]) }
func (addr Addr) SocietyID() uint32 { return binary.BigEndian.Uint32(addr[2:]) }
func (addr Addr) RouterID() uint32  { return binary.BigEndian.Uint32(addr[6:]) }
func (addr Addr) UserID() uint32    { return binary.BigEndian.Uint32(addr[10:]) }
func (addr Addr) AppID() uint16     { return binary.BigEndian.Uint16(addr[14:]) }

func (addr *Addr) SetPlanetID(id uint16)  { binary.BigEndian.PutUint16(addr[0:], id) }
func (addr *Addr) SetSocietyID(id uint32) { binary.BigEndian.PutUint32(addr[2:], id) }
func (addr *Addr) SetRouterID(id uint32)  { binary.BigEndian.PutUint32(addr[6:], id) }
func (addr *Addr) SetUserID(id uint32)    { binary.BigEndian.PutUint32(addr[10:], id) }
func (addr *Addr) SetAppID(id uint16)     { binary.BigEndian.PutUint16(addr[14:], id) }

//memar:impl memar/protocol.Stringer
func (addr *Addr) ToString() string {
	// TODO::: ??
	return string(addr[:])
}
func (addr *Addr) FromString(s string) (err protocol.Error) {
	// TODO:::
	return
}
