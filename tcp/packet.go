// For license and copyright information please see LEGAL file in repository

package tcp

import (
	"../binary"
)

// Packet implement all methods to Get||Set data to a packet as a byte slice with 0-alloc
type Packet []byte

/*
********** Get Methods **********
 */
func (p Packet) SourcePort() uint16      { return binary.BigEndian.Uint16(p[0:]) }
func (p Packet) DestinationPort() uint16 { return binary.BigEndian.Uint16(p[2:]) }
func (p Packet) SequenceNumber() uint32  { return binary.BigEndian.Uint32(p[4:8]) }
func (p Packet) AckNumber() uint32       { return binary.BigEndian.Uint32(p[8:]) }
func (p Packet) DataOffset() uint8       { return (p[12] >> 4) * 4 }
func (p Packet) Window() uint16          { return binary.BigEndian.Uint16(p[14:]) }
func (p Packet) GetChecksum() uint16     { return binary.BigEndian.Uint16(p[16:18]) }
func (p Packet) UrgentPointer() uint16   { return binary.BigEndian.Uint16(p[18:20]) }
func (p Packet) Options() []byte         { return p[21:p.DataOffset()] }
func (p Packet) Payload() []byte         { return p[p.DataOffset():] }

/*
********** Set Methods **********
 */
func (p Packet) SetSourcePort(port uint16)      { binary.BigEndian.PutUint16(p[0:], port) }
func (p Packet) SetDestinationPort(port uint16) { binary.BigEndian.PutUint16(p[2:], port) }
func (p Packet) SetSequenceNumber(v uint32)     { binary.BigEndian.PutUint32(p[4:], v) }
func (p Packet) SetAckNumber(v uint32)          { binary.BigEndian.PutUint32(p[8:], v) }
func (p Packet) SetDataOffset(v uint8)          { p[12] = byte((v/4)<<4) | byte(p[12]>>4) }
func (p Packet) SetFlagPartOne(flags byte)      { p[12] = p[12] | flags }
func (p Packet) SetFlagPartTwo(flags byte)      { p[13] = flags }
func (p Packet) SetWindow(v uint16)             { binary.BigEndian.PutUint16(p[14:], v) }
func (p Packet) SetChecksum(v uint16)           { binary.BigEndian.PutUint16(p[16:], v) }
func (p Packet) SetUrgentPointer(v uint16)      { binary.BigEndian.PutUint16(p[18:], v) }
func (p Packet) SetOptions(o []byte)            { copy(p[21:], o) }
func (p Packet) SetPayload(payload []byte)      { copy(p[p.DataOffset():], payload) }

/*
********** Flags **********
 */
func (p Packet) FlagReserved1() bool { return p[12]&Flag_Reserved1 != 0 }
func (p Packet) FlagReserved2() bool { return p[12]&Flag_Reserved2 != 0 }
func (p Packet) FlagReserved3() bool { return p[12]&Flag_Reserved3 != 0 }
func (p Packet) FlagNS() bool        { return p[12]&Flag_NS != 0 }
func (p Packet) FlagCWR() bool       { return p[13]&Flag_CWR != 0 }
func (p Packet) FlagECE() bool       { return p[13]&Flag_ECE != 0 }
func (p Packet) FlagURG() bool       { return p[13]&Flag_URG != 0 }
func (p Packet) FlagACK() bool       { return p[13]&Flag_ACK != 0 }
func (p Packet) FlagPSH() bool       { return p[13]&Flag_PSH != 0 }
func (p Packet) FlagRST() bool       { return p[13]&Flag_RST != 0 }
func (p Packet) FlagSYN() bool       { return p[13]&Flag_SYN != 0 }
func (p Packet) FlagFIN() bool       { return p[13]&Flag_FIN != 0 }

func (p Packet) SetFlagReserved1() { p[12] |= Flag_Reserved1 }
func (p Packet) SetFlagReserved2() { p[12] |= Flag_Reserved2 }
func (p Packet) SetFlagReserved3() { p[12] |= Flag_Reserved3 }
func (p Packet) SetFlagNS()        { p[12] |= Flag_NS }
func (p Packet) SetFlagCWR()       { p[13] |= Flag_CWR }
func (p Packet) SetFlagECE()       { p[13] |= Flag_ECE }
func (p Packet) SetFlagURG()       { p[13] |= Flag_URG }
func (p Packet) SetFlagACK()       { p[13] |= Flag_ACK }
func (p Packet) SetFlagPSH()       { p[13] |= Flag_PSH }
func (p Packet) SetFlagRST()       { p[13] |= Flag_RST }
func (p Packet) SetFlagSYN()       { p[13] |= Flag_SYN }
func (p Packet) SetFlagFIN()       { p[13] |= Flag_FIN }
