/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/binary"
	error_p "memar/error/protocol"
)

type optionEchoReply []byte

func (o optionEchoReply) Length() byte       { return o[0] }
func (o optionEchoReply) EchoReply() uint16  { return binary.BigEndian(o[1:]).Uint16() }
func (o optionEchoReply) NextOption() []byte { return o[5:] }

func (o optionEchoReply) Process(s *Stream) (err error_p.Error) {
	return
}
