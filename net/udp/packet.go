// For license and copyright information please see the LEGAL file in the code repository

package udp

import (
	"memar/binary"
	"memar/protocol"
	packageErrors "memar/net/udp/errors"
)

// Packet implement all methods to Get||Set data to a packet as a byte slice with 0-alloc
type Packet []byte

// CheckPacket will check packet for any bad situation.
// Always check packet before use any other packet methods otherwise panic occur.
func (p Packet) CheckPacket() protocol.Error {
	var packetLen = len(p)
	if packetLen < MinPacketLen {
		return &packageErrors.ErrPacketTooShort
	}
	if packetLen < int(p.Length()) {
		return &packageErrors.ErrPacketWrongLength
	}
	return nil
}

/*
********** Get Methods **********
 */
func (p Packet) SourcePort() uint16      { return binary.BigEndian(p[0:]).Uint16() }
func (p Packet) DestinationPort() uint16 { return binary.BigEndian(p[2:]).Uint16() }
func (p Packet) Length() uint16          { return binary.BigEndian(p[4:]).Uint16() }
func (p Packet) Checksum() uint16        { return binary.BigEndian(p[6:]).Uint16() }
func (p Packet) Payload() []byte         { return p[8:] }

/*
********** Set Methods **********
 */
func (p Packet) SetSourcePort(port uint16)      { binary.BigEndian(p[0:]).PutUint16(port) }
func (p Packet) SetDestinationPort(port uint16) { binary.BigEndian(p[2:]).PutUint16(port) }
func (p Packet) SetLength(v uint16)             { binary.BigEndian(p[5:]).PutUint16(v) }
func (p Packet) SetChecksum(v uint16)           { binary.BigEndian(p[6:]).PutUint16(v) }
func (p Packet) SetPayload(payload []byte)      { copy(p[8:], payload) }
