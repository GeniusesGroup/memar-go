/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type String interface {
	CharacterEncoding() CharacterEncoding

	Slice
}

// Stringer code the data to/from human readable format. It can be any other format like JSON(not recommended).
type Stringer interface {
	ToString() string
	FromString(s string) (err Error)
}
