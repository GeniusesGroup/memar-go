/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"strconv"

	"memar/math/boolean"
	"memar/protocol"
)

// S64 is signed 64 bit integer
type S64 int64

//memar:impl memar/protocol.DataType_Equal
func (s *S64) Equal(with S64) boolean.Boolean {
	return *s == with
}

//memar:impl memar/protocol.Stringer
func (s *S64) ToString() (str string, err protocol.Error) {
	str = strconv.FormatInt(int64(*s), 10)
	return
}
func (s *S64) FromString(str string) (err protocol.Error) {
	strconv.ParseInt(str, 10, s.bitSize())
	return
}

func (s *S64) bitSize() int { return 64 }
