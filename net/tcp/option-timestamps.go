/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/binary"
	error_p "memar/error/protocol"
)

type optionTimestamps []byte

func (o optionTimestamps) Length() byte       { return o[0] }
func (o optionTimestamps) Timestamps() uint16 { return binary.BigEndian(o[1:]).Uint16() }
func (o optionTimestamps) NextOption() []byte { return o[8:] }

func (o optionTimestamps) Process(s *Stream) (err error_p.Error) {
	return
}
