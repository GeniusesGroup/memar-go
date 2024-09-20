/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	error_p "memar/error/protocol"
)

type optionSACKPermitted []byte

func (o optionSACKPermitted) Length() byte { return o[0] }

// func (o optionSACKPermitted) SACKPermitted() uint16 { return binary.BigEndian(o[1:]).Uint16() }
func (o optionSACKPermitted) NextOption() []byte { return o[1:] }

func (o optionSACKPermitted) Process(s *Stream) (err error_p.Error) {
	return
}
