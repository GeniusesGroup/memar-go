/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/binary"
	"memar/protocol"
)

type optionWindowScale []byte

func (o optionWindowScale) Length() byte        { return o[0] }
func (o optionWindowScale) WindowScale() uint16 { return binary.BigEndian(o[1:]).Uint16() }
func (o optionWindowScale) NextOption() []byte  { return o[2:] }

// handler options -> stream
func (o optionWindowScale) Process(s *Stream) (err protocol.Error) {
	return
}
