/* For license and copyright information please see the LEGAL file in the code repository */

package ascii_p

import (
	error_p "memar/error/protocol"
	string_p "memar/string/protocol"
)

// ASCII abbreviated from American Standard Code for Information Interchange,
// is a character encoding standard for electronic communication.
// https://en.wikipedia.org/wiki/ASCII
type ASCII interface {
	string_p.String

	// TODO:::
}

// Stringer code the data to/from human readable format.
type Stringer interface {
	Stringer_To
	Stringer_From
}

// Stringer code the data to human readable format.
type Stringer_To interface {
	ToASCII() (str ASCII, err error_p.Error)
}

// Stringer_From code the data from human readable format.
type Stringer_From interface {
	FromASCII(str ASCII) (err error_p.Error)
}
