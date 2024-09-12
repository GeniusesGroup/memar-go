/* For license and copyright information please see the LEGAL file in the code repository */

package error_p

import (
	// primitive_p "memar/computer/language/primitive/protocol"
	datatype_p "memar/datatype/protocol"
	mediatype_p "memar/mediatype/protocol"
)

// Error is base behaviors that any Error capsule must implement.
//
// Other suggestions:
// - opaque error model: https://dave.cheney.net/paste/gocon-spring-2016.pdf
// - RFC7807: https://tools.ietf.org/html/rfc7807
// - https://doc.rust-lang.org/stable/std/error/trait.Error.html
type Error interface {
	datatype_p.DataType
	mediatype_p.MediaType

	// Can't un-comment below due to Golang import cycle problem, So add it manually.
	// primitive_p.Equivalence[Error]
	Equal(with Error) bool

	// ADT

	// Below methods comment in favor of log_p.Event_Message interface.
	// Error() string
	// Stringer[String]
}
