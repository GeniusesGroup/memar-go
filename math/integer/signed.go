/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"strconv"

	"memar/math/boolean"
	"memar/protocol"
)

// https://en.wikipedia.org/wiki/Integer_(computer_science)
type Signed struct {
	number  int64
	bitSize int
}

//memar:impl memar/protocol.DataType_Equal
func (s *Signed) Equal(with Signed) boolean.Boolean {
	return *s == with
}

//memar:impl memar/protocol.Stringer
func (s *Signed) ToString() (str string, err protocol.Error) {
	str = strconv.FormatInt(int64(s.number), 10)
	return
}
func (s *Signed) FromString(str string) (err protocol.Error) {
	strconv.ParseInt(str, 10, s.bitSize)
	return
}
