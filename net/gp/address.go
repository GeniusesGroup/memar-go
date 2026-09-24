/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"memar/codec/binary"
	error_p "memar/process/error/protocol"
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
func (addr Addr) ThingID() uint32    { return binary.BigEndian.Uint32(addr[10:]) }
func (addr Addr) AppID() uint16     { return binary.BigEndian.Uint16(addr[14:]) }

func (addr *Addr) SetPlanetID(id uint16)  { binary.BigEndian.PutUint16(addr[0:], id) }
func (addr *Addr) SetSocietyID(id uint32) { binary.BigEndian.PutUint32(addr[2:], id) }
func (addr *Addr) SetRouterID(id uint32)  { binary.BigEndian.PutUint32(addr[6:], id) }
func (addr *Addr) SetThingID(id uint32)    { binary.BigEndian.PutUint32(addr[10:], id) }
func (addr *Addr) SetAppID(id uint16)     { binary.BigEndian.PutUint16(addr[14:], id) }

//memar:impl memar/codec/string/protocol.Stringer
func (addr *Addr) ToString() (str string, err error_p.Error) {
	// TODO::: ??
	str = string(addr[:])
	return
}
func (addr *Addr) FromString(str string) (err error_p.Error) {
	// TODO:::
	return
}
