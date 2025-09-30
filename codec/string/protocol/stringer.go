/* For license and copyright information please see the LEGAL file in the code repository */

package string_p

import (
	error_p "memar/error/protocol"
)

// Stringer code the data to/from human readable format. It can be any other format like JSON(not recommended).
type Stringer[STR String] interface {
	Stringer_To[STR]
	Stringer_From[STR]
}

type Stringer_To[STR String] interface {
	ToString() (str STR, err error_p.Error)
}
type Stringer_From[STR String] interface {
	FromString(str STR) (err error_p.Error)
}

// Stringer_GO is same as `fmt.Stringer`
// It just exist to prevent protocol package not couple to a implementation package(fmt).
type Stringer_GO interface {
	// It will clone the underlying buffer and return standalone string.
	String() string
}
