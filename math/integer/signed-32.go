/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"strconv"

	error_p "memar/process/error/protocol"
	"memar/math/boolean"
)

// 32 is signed 32 bit integer
type S32 int64

//memar:impl memar/math/logic/protocol.Equivalence
func (s *S32) Equal(with S32) boolean.Boolean {
	return *s == with
}

//memar:impl memar/codec/string/protocol.Stringer
func (s *S32) ToString() (str string, err error_p.Error) {
	str = strconv.FormatInt(int64(*s), 10)
	return
}
func (s *S32) FromString(str string) (err error_p.Error) {
	strconv.ParseInt(str, 10, s.bitSize())
	return
}

func (s *S32) bitSize() int { return 32 }
