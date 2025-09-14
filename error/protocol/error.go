/* For license and copyright information please see the LEGAL file in the code repository */

package error_p

import (
	adt_p "memar/adt/protocol"
	datatype_p "memar/datatype/protocol"
	logic_p "memar/math/logic/protocol"
	mediatype_p "memar/mediatype/protocol"
)

type Field_Error interface {
	Error() Error 
}

// Error is base behaviors that any Error capsule must implement.
// Error MUST NOT mix with `Log Event`.
// Error has static data for any client, to tell about system fault situation when processing client request.
// Log event carry on static and dynamic data for developers, to indicate system fault situation, and help them troubleshoot potential bugs.
//
// Other frameworks:
// - RFC7807: https://tools.ietf.org/html/rfc7807
// - https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Error
// - https://doc.rust-lang.org/stable/std/error/trait.Error.html
// - https://learn.microsoft.com/en-us/dotnet/api/system.exception
// - opaque error model: https://dave.cheney.net/paste/gocon-spring-2016.pdf
type Error interface {
	datatype_p.DataType
	mediatype_p.MediaType

	logic_p.Equivalence[Error]

	adt_p.ADT

	// Below methods comment in favor of log_p.Event_Message interface.
	// Error() string
	// string_p.Stringer
}
