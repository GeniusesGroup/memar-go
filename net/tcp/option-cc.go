/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"libgo/binary"
	"libgo/protocol"
)

type optionCC []byte

func (o optionCC) Length() byte       { return o[0] }
func (o optionCC) CC() uint16         { return binary.BigEndian.Uint16(o[1:]) }
func (o optionCC) NextOption() []byte { return o[5:] }

func (o optionCC) Process(s *Stream) (err protocol.Error) {
	return
}
