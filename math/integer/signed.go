/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"strconv"

	error_p "memar/process/error/protocol"
	"memar/math/boolean"
)

// https://en.wikipedia.org/wiki/Integer_(computer_science)
type Signed struct {
	number  int64
	bitSize int
}

//memar:impl memar/math/logic/protocol.Equivalence
func (s *Signed) Equal(with Signed) boolean.Boolean {
	return *s == with
}

//memar:impl memar/codec/string/protocol.Stringer
func (s *Signed) ToString() (str string, err error_p.Error) {
	str = strconv.FormatInt(int64(s.number), 10)
	return
}
func (s *Signed) FromString(str string) (err error_p.Error) {
	strconv.ParseInt(str, 10, s.bitSize)
	return
}
