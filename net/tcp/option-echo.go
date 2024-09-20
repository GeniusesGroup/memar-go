/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/binary"
	error_p "memar/error/protocol"
)

type optionEcho []byte

func (o optionEcho) Length() byte       { return o[0] }
func (o optionEcho) Echo() uint16       { return binary.BigEndian(o[1:]).Uint16() }
func (o optionEcho) NextOption() []byte { return o[5:] }

func (o optionEcho) Process(s *Stream) (err error_p.Error) {
	return
}
