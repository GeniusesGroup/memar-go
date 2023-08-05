/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/binary"
	"memar/protocol"
)

type optionCCNew []byte

func (o optionCCNew) Length() byte       { return o[0] }
func (o optionCCNew) CCNew() uint16      { return binary.BigEndian(o[1:]).Uint16() }
func (o optionCCNew) NextOption() []byte { return o[5:] }

func (o optionCCNew) Process(s *Stream) (err protocol.Error) {
	return
}
