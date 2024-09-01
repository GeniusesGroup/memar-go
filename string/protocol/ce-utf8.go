/* For license and copyright information please see the LEGAL file in the code repository */

package string_p

import (
	"memar/protocol"
)

// UTF-8 is a variable-length character encoding standard used for electronic communication.
// Defined by the Unicode Standard, the name is derived from Unicode Transformation Format – 8-bit.
// UTF-8 is capable of encoding all 1,112,064 valid character code points in Unicode using one to four one-byte code units.
// https://en.wikipedia.org/wiki/UTF-8
type UTF8 interface {
	String

	// TODO:::
}

// Stringer_UTF8 code the data to/from human readable format.
type Stringer_UTF8 interface {
	ToUTF8() (str UTF8, err protocol.Error)
	FromUTF8(str UTF8) (err protocol.Error)
}
