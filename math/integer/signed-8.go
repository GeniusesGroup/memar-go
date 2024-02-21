/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"strconv"

	"memar/math/boolean"
	"memar/protocol"
)

// S8 is signed 8 bit integer
type S8 int8

//memar:impl memar/protocol.DataType_Equal
func (s *S8) Equal(with S8) boolean.Boolean {
	return *s == with
}

// TODO::: not efficient enough code
//
//memar:impl memar/protocol.Stringer
func (s *S8) ToString() (str string, err protocol.Error) {
	str = strconv.FormatInt(int64(*s), 10)
	return
}
func (s *S8) FromString(str string) (err protocol.Error) {
	strconv.ParseInt(str, 10, s.bitSize())
	return
}

func (s *S8) bitSize() int { return 8 }
