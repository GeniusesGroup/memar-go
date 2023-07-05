/* For license and copyright information please see the LEGAL file in the code repository */

package ipv4

import (
	"libgo/binary"
	"libgo/protocol"
)

// Packet implement all methods to Get||Set data to a packet as a byte slice with 0-alloc
type Packet []byte

// CheckPacket checks packet for any bad situation
// Always check packet before use any other packet methods otherwise panic may occur
func (p Packet) CheckPacket() protocol.Error {
	var packetLen = len(p)
	if packetLen < MinHeaderLen {
		return &ErrPacketTooShort
	}
	if packetLen < int(p.IHL()) {
		return &ErrPacketWrongLength
	}
	return nil
}

/*
********** Get Methods **********
 */
func (p Packet) Version() uint8                  { return p[0] >> 4 }
func (p Packet) IHL() uint8                      { return (p[0] & 0x0f) * 4 }
func (p Packet) DSCP() uint8                     { return p[1] >> 2 }
func (p Packet) TotalLength() uint16             { return binary.BigEndian(p[2:]).Uint16() }
func (p Packet) Identification() (id [2]byte)    { copy(id[:], p[4:]); return }
func (p Packet) FragmentOffset() uint16          { return binary.BigEndian(p[6:]).Uint16() & 0b1110000000000000 }
func (p Packet) TimeToLive() uint8               { return p[8] }
func (p Packet) Protocol() uint8                 { return p[9] }
func (p Packet) HeaderChecksum() (check [2]byte) { copy(check[:], p[10:]); return }
func (p Packet) SourceAddr() (srcAddr Addr)      { copy(srcAddr[:], p[12:]); return }
func (p Packet) DestinationAddr() (desAddr Addr) { copy(desAddr[:], p[16:]); return }
func (p Packet) Options() []byte                 { return p[20:] }
func (p Packet) Payload() []byte                 { return p[p.IHL():] }

/*
********** Set Methods **********
 */
func (p Packet) SetVersion(v uint8)              { p[0] = (v << 4) }
func (p Packet) SetIHL(ln uint8)                 { p[0] |= (ln / 4) }
func (p Packet) SetDSCP(dscp uint8)              { p[1] |= (dscp >> 2) }
func (p Packet) SetTotalLength(tl uint16)        { binary.BigEndian(p[2:]).PutUint16(tl) }
func (p Packet) SetIdentification(id [2]byte)    { copy(p[4:], id[:]) }
func (p Packet) SetFragmentOffset(fo uint16)     { binary.BigEndian(p[6:]).PutUint16(fo) }
func (p Packet) SetTimeToLive(ttl uint8)         { p[8] = ttl }
func (p Packet) SetProtocol(proto uint8)         { p[9] = proto }
func (p Packet) SetHeaderChecksum(check [2]byte) { copy(p[10:], check[:]); return }
func (p Packet) SetSourceAddr(srcAddr Addr)      { copy(p[12:], srcAddr[:]); return }
func (p Packet) SetDestinationAddr(desAddr Addr) { copy(p[16:], desAddr[:]); return }
func (p Packet) SetOptions(opts []byte)          { copy(p[20:], opts) }
func (p Packet) SetPayload(payload []byte)       { copy(p[p.IHL():], payload) }

/*
********** Flags **********
 */
func (p Packet) FlagECT() bool      { return p.FlagECT0() || p.FlagECT1() }
func (p Packet) FlagECT0() bool     { return p[1]&flag_ECT0 != 0 }
func (p Packet) FlagECT1() bool     { return p[1]&flag_ECT1 != 0 }
func (p Packet) FlagCE() bool       { return p[1]&flag_CE != 0 }
func (p Packet) FlagReserved() bool { return p[6]&flag_Reserved != 0 }
func (p Packet) FlagDF() bool       { return p[6]&flag_DF != 0 }
func (p Packet) FlagMF() bool       { return p[6]&flag_MF != 0 }

func (p Packet) SetFlagECT()      { p.SetFlagECT0() }
func (p Packet) SetFlagECT0()     { p[1] |= flag_ECT0 }
func (p Packet) SetFlagECT1()     { p[1] |= flag_ECT1 }
func (p Packet) SetFlagCE()       { p[1] |= flag_CE }
func (p Packet) SetFlagReserved() { p[6] |= flag_Reserved }
func (p Packet) SetFlagDF()       { p[6] |= flag_DF }
func (p Packet) SetFlagMF()       { p[6] |= flag_MF }

func (p Packet) UnsetFlagECT()      { p[1] &= ^flag_CE }
func (p Packet) UnsetFlagECT0()     { p[1] &= ^flag_ECT0 }
func (p Packet) UnsetFlagECT1()     { p[1] &= ^flag_ECT1 }
func (p Packet) UnsetFlagCE()       { p[1] &= ^flag_CE }
func (p Packet) UnsetFlagReserved() { p[6] &= ^flag_Reserved }
func (p Packet) UnsetFlagDF()       { p[6] &= ^flag_DF }
func (p Packet) UnsetFlagMF()       { p[6] &= ^flag_MF }
