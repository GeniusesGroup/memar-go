/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/binary"
	"memar/protocol"
)

/*
	type optionMSS struct {
		Length byte
		MSS    uint16 // Max Segment Length
	}
*/
type optionMSS []byte

func (o optionMSS) Length() byte       { return o[0] }
func (o optionMSS) MSS() uint16        { return binary.BigEndian(o[1:]).Uint16() }
func (o optionMSS) NextOption() []byte { return o[3:] }

func (o optionMSS) Process(s *Stream) (err protocol.Error) {
	return
}
