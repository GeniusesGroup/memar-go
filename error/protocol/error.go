/* For license and copyright information please see the LEGAL file in the code repository */

package error_p

import (
	// primitive_p "memar/computer/language/primitive/protocol"
	datatype_p "memar/datatype/protocol"
)

// Error is base behaviors that any Error capsule must implement.
//
// It is similar to opaque error model describe here: https://dave.cheney.net/paste/gocon-spring-2016.pdf
// or this RFC: https://tools.ietf.org/html/rfc7807
type Error interface {
	datatype_p.DataType

	// Can't un-comment below due to Golang import cycle problem, So add it manually!
	// primitive_p.Equivalence[Error]
	Equal(with Error) bool

	// ADT
	// MediaType

	// Below methods comment in favor of log_p.Event_Message interface.
	// Error() string
	// Stringer[String]
}
