/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/binary"
	"memar/protocol"
)

type optionCCEcho []byte

func (o optionCCEcho) Length() byte       { return o[0] }
func (o optionCCEcho) CCEcho() uint16     { return binary.BigEndian(o[1:]).Uint16() }
func (o optionCCEcho) NextOption() []byte { return o[5:] }

func (o optionCCEcho) Process(s *Stream) (err protocol.Error) {
	return
}
