// For license and copyright information please see LEGAL file in repository

package udp

import (
	"../binary"
	"../protocol"
)

// Packet implement all methods to Get||Set data to a packet as a byte slice with 0-alloc
type Packet []byte

// CheckPacket will check packet for any bad situation.
// Always check packet before use any other packet methods otherwise panic occur.
func (p Packet) CheckPacket() protocol.Error {
	var packetLen = len(p)
	if packetLen < MinPacketLen {
		return ErrPacketTooShort
	}
	if packetLen < int(p.Length()) {
		return ErrPacketWrongLength
	}
	return nil
}

/*
********** Get Methods **********
 */
func (p Packet) SourcePort() uint16      { return binary.BigEndian.Uint16(p[0:]) }
func (p Packet) DestinationPort() uint16 { return binary.BigEndian.Uint16(p[2:]) }
func (p Packet) Length() uint16          { return binary.BigEndian.Uint16(p[4:]) }
func (p Packet) Checksum() uint16        { return binary.BigEndian.Uint16(p[6:]) }
func (p Packet) Payload() []byte         { return p[8:] }

/*
********** Set Methods **********
 */
func (p Packet) SetSourcePort(port uint16)      { binary.BigEndian.PutUint16(p[0:], port) }
func (p Packet) SetDestinationPort(port uint16) { binary.BigEndian.PutUint16(p[2:], port) }
func (p Packet) SetLength(v uint16)             { binary.BigEndian.PutUint16(p[4:], v) }
func (p Packet) SetChecksum(v uint16)           { binary.BigEndian.PutUint16(p[6:], v) }
func (p Packet) SetPayload(payload []byte)      { copy(p[8:], payload) }
