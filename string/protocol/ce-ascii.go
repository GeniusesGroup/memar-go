/* For license and copyright information please see the LEGAL file in the code repository */

package string_p

import (
	error_p "memar/error/protocol"
)

// ASCII abbreviated from American Standard Code for Information Interchange,
// is a character encoding standard for electronic communication.
// https://en.wikipedia.org/wiki/ASCII
type ASCII interface {
	String

	// TODO:::
}

// Stringer_ASCII code the data to/from human readable format.
type Stringer_ASCII interface {
	Stringer_To_ASCII
	Stringer_From_ASCII
}

// Stringer_ASCII code the data to human readable format.
type Stringer_To_ASCII interface {
	ToASCII() (str ASCII, err error_p.Error)
}

// Stringer_From_ASCII code the data from human readable format.
type Stringer_From_ASCII interface {
	FromASCII(str ASCII) (err error_p.Error)
}
