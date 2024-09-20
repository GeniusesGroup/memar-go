/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"memar/binary"
	error_p "memar/error/protocol"
)

/*
TCP POC-service-profile Option:

	Kind: 10  Length: 3 bytes

	                              1 bit        1 bit    6 bits
	    +----------+----------+------------+----------+--------+
	    |  Kind=10 | Length=3 | Start_flag | End_flag | Filler |
	    +----------+----------+------------+----------+--------+
*/
type optionPartialOrderServiceProfile []byte

func (o optionPartialOrderServiceProfile) Length() byte { return o[0] }
func (o optionPartialOrderServiceProfile) PartialOrderServiceProfile() uint16 {
	return binary.BigEndian(o[1:]).Uint16()
	//
}
func (o optionPartialOrderServiceProfile) NextOption() []byte { return o[3:] }

func (o optionPartialOrderServiceProfile) Process(s *Stream) (err error_p.Error) {
	return
}
