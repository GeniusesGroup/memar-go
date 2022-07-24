/* For license and copyright information please see LEGAL file in repository */

package tcp

import "github.com/GeniusesGroup/libgo/binary"

type optionEchoReply []byte

func (o optionEchoReply) Length() byte       { return o[0] }
func (o optionEchoReply) EchoReply() uint16  { return binary.BigEndian.Uint16(o[1:]) }
func (o optionEchoReply) NextOption() []byte { return o[5:] }

func (o optionEchoReply) Process(s *Socket) error {
	return nil
}
