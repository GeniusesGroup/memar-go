/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// UTF-8 is a variable-length character encoding standard used for electronic communication.
// Defined by the Unicode Standard, the name is derived from Unicode Transformation Format – 8-bit.
// UTF-8 is capable of encoding all 1,112,064 valid character code points in Unicode using one to four one-byte code units.
type UTF8 interface {
	String

	// TODO:::
}

// Stringer code the data to/from human readable format.
type UTF8_Stringer interface {
	ToUTF8() UTF8
	FromUTF8(s UTF8) (err Error)
}
