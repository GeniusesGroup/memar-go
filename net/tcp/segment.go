// For license and copyright information please see the LEGAL file in the code repository

package tcp

import (
	"libgo/binary"
	"libgo/protocol"
)

// Segment implement all methods to Get||Set data to a segment as a byte slice with 0-alloc
// https://datatracker.ietf.org/doc/html/rfc793#section-3.1
type Segment []byte

// CheckSegment will check segment for any bad situation.
// Always check segment before use any other segment methods otherwise panic occur.
func (s Segment) CheckSegment() protocol.Error {
	var segmentLen = len(s)
	if segmentLen < CNF_Segment_MinSize {
		return &ErrSegmentTooShort
	}
	if segmentLen < int(s.DataOffset()) {
		return &ErrSegmentWrongLength
	}
	return nil
}

/*
********** Get Methods **********
 */
func (s Segment) SourcePort() uint16      { return binary.BigEndian.Uint16(s[0:]) }
func (s Segment) DestinationPort() uint16 { return binary.BigEndian.Uint16(s[2:]) }
func (s Segment) SequenceNumber() uint32  { return binary.BigEndian.Uint32(s[4:]) }
func (s Segment) AckNumber() uint32       { return binary.BigEndian.Uint32(s[8:]) }
func (s Segment) DataOffset() uint8       { return (s[12] >> 4) * 4 }
func (s Segment) Window() uint16          { return binary.BigEndian.Uint16(s[14:]) }
func (s Segment) Checksum() uint16        { return binary.BigEndian.Uint16(s[16:]) }
func (s Segment) UrgentPointer() uint16   { return binary.BigEndian.Uint16(s[18:]) }
func (s Segment) Options() []byte         { return s[20:s.DataOffset()] }
func (s Segment) Payload() []byte         { return s[s.DataOffset():] }

/*
********** Set Methods **********
 */
func (s Segment) SetSourcePort(port uint16)      { binary.BigEndian.PutUint16(s[0:], port) }
func (s Segment) SetDestinationPort(port uint16) { binary.BigEndian.PutUint16(s[2:], port) }
func (s Segment) SetSequenceNumber(v uint32)     { binary.BigEndian.PutUint32(s[4:], v) }
func (s Segment) SetAckNumber(v uint32)          { binary.BigEndian.PutUint32(s[8:], v) }
func (s Segment) SetDataOffset(ln uint8)         { s[12] = byte((ln/4)<<4) | s[12] }
func (s Segment) SetFlagPartOne(flags byte)      { s[12] = s[12] | flags }
func (s Segment) SetFlagPartTwo(flags byte)      { s[13] = flags }
func (s Segment) SetWindow(v uint16)             { binary.BigEndian.PutUint16(s[14:], v) }
func (s Segment) SetChecksum(v uint16)           { binary.BigEndian.PutUint16(s[16:], v) }
func (s Segment) SetUrgentPointer(v uint16)      { binary.BigEndian.PutUint16(s[18:], v) }
func (s Segment) SetOptions(o []byte)            { copy(s[20:], o) }
func (s Segment) SetPayload(payload []byte)      { copy(s[s.DataOffset():], payload) }

/*
********** Flags **********
 */
func (s Segment) FlagReserved1() bool { return s[12]&byte(flag_Reserved1) != 0 }
func (s Segment) FlagReserved2() bool { return s[12]&byte(flag_Reserved2) != 0 }
func (s Segment) FlagReserved3() bool { return s[12]&byte(flag_Reserved3) != 0 }
func (s Segment) FlagNS() bool        { return s[12]&byte(flag_NS) != 0 }
func (s Segment) FlagCWR() bool       { return s[13]&byte(flag_CWR) != 0 }
func (s Segment) FlagECE() bool       { return s[13]&byte(flag_ECE) != 0 }
func (s Segment) FlagURG() bool       { return s[13]&byte(flag_URG) != 0 }
func (s Segment) FlagACK() bool       { return s[13]&byte(flag_ACK) != 0 }
func (s Segment) FlagPSH() bool       { return s[13]&byte(flag_PSH) != 0 }
func (s Segment) FlagRST() bool       { return s[13]&byte(flag_RST) != 0 }
func (s Segment) FlagSYN() bool       { return s[13]&byte(flag_SYN) != 0 }
func (s Segment) FlagFIN() bool       { return s[13]&byte(flag_FIN) != 0 }

func (s Segment) SetFlagReserved1() { s[12] |= byte(flag_Reserved1) }
func (s Segment) SetFlagReserved2() { s[12] |= byte(flag_Reserved2) }
func (s Segment) SetFlagReserved3() { s[12] |= byte(flag_Reserved3) }
func (s Segment) SetFlagNS()        { s[12] |= byte(flag_NS) }
func (s Segment) SetFlagCWR()       { s[13] |= byte(flag_CWR) }
func (s Segment) SetFlagECE()       { s[13] |= byte(flag_ECE) }
func (s Segment) SetFlagURG()       { s[13] |= byte(flag_URG) }
func (s Segment) SetFlagACK()       { s[13] |= byte(flag_ACK) }
func (s Segment) SetFlagPSH()       { s[13] |= byte(flag_PSH) }
func (s Segment) SetFlagRST()       { s[13] |= byte(flag_RST) }
func (s Segment) SetFlagSYN()       { s[13] |= byte(flag_SYN) }
func (s Segment) SetFlagFIN()       { s[13] |= byte(flag_FIN) }

func (s Segment) UnsetFlagReserved1() { s[12] &= ^byte(flag_Reserved1) }
func (s Segment) UnsetFlagReserved2() { s[12] &= ^byte(flag_Reserved2) }
func (s Segment) UnsetFlagReserved3() { s[12] &= ^byte(flag_Reserved3) }
func (s Segment) UnsetFlagNS()        { s[12] &= ^byte(flag_NS) }
func (s Segment) UnsetFlagCWR()       { s[13] &= ^byte(flag_CWR) }
func (s Segment) UnsetFlagECE()       { s[13] &= ^byte(flag_ECE) }
func (s Segment) UnsetFlagURG()       { s[13] &= ^byte(flag_URG) }
func (s Segment) UnsetFlagACK()       { s[13] &= ^byte(flag_ACK) }
func (s Segment) UnsetFlagPSH()       { s[13] &= ^byte(flag_PSH) }
func (s Segment) UnsetFlagRST()       { s[13] &= ^byte(flag_RST) }
func (s Segment) UnsetFlagSYN()       { s[13] &= ^byte(flag_SYN) }
func (s Segment) UnsetFlagFIN()       { s[13] &= ^byte(flag_FIN) }
