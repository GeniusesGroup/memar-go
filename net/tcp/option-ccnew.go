/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"libgo/binary"
	"libgo/protocol"
)

type optionCCNew []byte

func (o optionCCNew) Length() byte       { return o[0] }
func (o optionCCNew) CCNew() uint16      { return binary.BigEndian.Uint16(o[1:]) }
func (o optionCCNew) NextOption() []byte { return o[5:] }

func (o optionCCNew) Process(s *Stream) (err protocol.Error) {
	return
}
