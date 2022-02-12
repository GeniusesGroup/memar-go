/* For license and copyright information please see LEGAL file in repository */

package ipv6

import (
	"../binary"
	"../protocol"
)

// Packet implement all methods to Get||Set data to a packet as a byte slice with 0-alloc
type Packet []byte

// CheckPacket checks packet for any bad situation
// Always check packet before use any other packet methods otherwise panic may occur
func (p Packet) CheckPacket() protocol.Error {
	if len(p) < HeaderLen {
		return ErrPacketTooShort
	}
	return nil
}

/*
********** Get Methods **********
 */
func (p Packet) Version() uint8                  { return p[0] >> 4 }
func (p Packet) TrafficClass() uint8             { return p[0]<<4 | p[1]>>4 }
func (p Packet) FlowLabel() (fl [3]byte)         { copy(fl[:], p[1:]); fl[0] &= 0x0f; return }
func (p Packet) PayloadLength() uint16           { return binary.BigEndian.Uint16(p[4:]) }
func (p Packet) NextHeader() uint8               { return p[6] }
func (p Packet) HopLimit() uint8                 { return p[7] }
func (p Packet) SourceAddr() (srcAddr Addr)      { copy(srcAddr[:], p[8:]); return }
func (p Packet) DestinationAddr() (desAddr Addr) { copy(desAddr[:], p[24:]); return }
func (p Packet) Payload() []byte                 { return p[40:] }

/*
********** Set Methods **********
 */
func (p Packet) SetVersion(v uint8)              { p[0] = (v << 4) }
func (p Packet) SetTrafficClass(tc uint8)        { p[0] |= (tc >> 4); p[1] = (tc << 4) }
func (p Packet) SetFlowLabel(fl [3]byte)         { p[1] |= fl[0]; p[2] = fl[1]; p[3] = fl[2] }
func (p Packet) SetPayloadLength(ln uint16)      { binary.BigEndian.SetUint16(p[4:], ln) }
func (p Packet) SetNextHeader(nh uint8)          { p[6] = nh }
func (p Packet) SetHopLimit(hl uint8)            { p[7] = hl }
func (p Packet) SetSourceAddr(srcAddr Addr)      { copy(p[8:], srcAddr[:]) }
func (p Packet) SetDestinationAddr(desAddr Addr) { copy(p[24:], desAddr[:]) }
func (p Packet) SetPayload(payload []byte)       { copy(p[40:], payload) }
