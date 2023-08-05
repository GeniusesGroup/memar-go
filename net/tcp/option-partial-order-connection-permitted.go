/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/protocol"
)

type optionPartialOrderConnectionPermitted []byte

func (o optionPartialOrderConnectionPermitted) Length() byte { return o[0] }

//	func (o optionPartialOrderConnectionPermitted) PartialOrderConnectionPermitted() uint16 {
//		return binary.BigEndian(o[1:]).Uint16()
//	}
func (o optionPartialOrderConnectionPermitted) NextOption() []byte { return o[3:] }

func (o optionPartialOrderConnectionPermitted) Process(s *Stream) (err protocol.Error) {
	return
}
