/* For license and copyright information please see the LEGAL file in the code repository */

package string_p

import (
	error_p "memar/process/error/protocol"
)

// Stringer code the data to/from human readable format. It can be any other format like JSON(not recommended).
type Stringer interface {
	Stringer_To
	Stringer_From
}

type Stringer_To interface {
	ToString() (str String, err error_p.Error)
}
type Stringer_From interface {
	FromString(str String) (err error_p.Error)
}

// Stringer_GO is same as `fmt.Stringer`
// It just exist to prevent protocol package not couple to a implementation package(fmt).
type Stringer_GO interface {
	// It will clone the underlying buffer and return standalone string.
	ToGoString() (str string, err error_p.Error)
	FromGoString(str string) (err error_p.Error)
}
