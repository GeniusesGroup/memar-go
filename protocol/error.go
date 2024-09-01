/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Error is base behaviors that any Error capsule must implement.
//
// It is similar to opaque error model describe here: https://dave.cheney.net/paste/gocon-spring-2016.pdf
// or this RFC: https://tools.ietf.org/html/rfc7807
type Error interface {
	DataType
	DataType_Equal[Error]

	// ADT
	MediaType

	// Below methods comment in favor of log_p.Event_Message interface.
	// Error() string
	// Stringer[String]
}
