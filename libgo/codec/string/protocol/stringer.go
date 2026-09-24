/* For license and copyright information please see the LEGAL file in the code repository */

package go_string_p

import (
	error_p "memar/process/error/protocol"
)

// Stringer code the data to/from human readable format. It can be any other format like JSON(not recommended).
type Stringer interface {
	Stringer_To
	Stringer_From
}

type Stringer_To interface {
	ToString() (str string, err error_p.Error)
}
type Stringer_From interface {
	FromString(str string) (err error_p.Error)
}
