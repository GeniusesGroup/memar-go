/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"libgo/binary"
	"libgo/protocol"
)

type optionEchoReply []byte

func (o optionEchoReply) Length() byte       { return o[0] }
func (o optionEchoReply) EchoReply() uint16  { return binary.BigEndian.Uint16(o[1:]) }
func (o optionEchoReply) NextOption() []byte { return o[5:] }

func (o optionEchoReply) Process(s *Stream) (err protocol.Error) {
	return
}
