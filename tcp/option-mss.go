/* For license and copyright information please see LEGAL file in repository */

package tcp

import (
	"../binary"
	"../protocol"
)

/*
type optionMSS struct {
	Length byte
	MSS    uint16 // Max Segment Length
}
*/
type optionMSS []byte

func (o optionMSS) Length() byte       { return o[0] }
func (o optionMSS) MSS() uint16        { return binary.BigEndian.Uint16(o[1:]) }
func (o optionMSS) NextOption() []byte { return o[3:] }

func (o optionMSS) Process(s *Socket) (err protocol.Error) {
	return
}
