/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Errors is the interface that must implement by any Application.
type Errors interface {
	RegisterError(err Error)
	GetErrorByID(id MediaTypeID) (err Error)
	GetErrorByMediaType(mt string) (err Error)
}

// Error is similar to opaque error model describe here: https://dave.cheney.net/paste/gocon-spring-2016.pdf
// or this RFC: https://tools.ietf.org/html/rfc7807
type Error interface {
	// Init() must call protocol.App.RegisterError() to register the error in application
	// Init(mediatype string)

	// Check both flat or chain situation.
	Equal(Error) bool

	// Who cause the error?
	// Internal	: means calling process logic has runtime bugs like HTTP server error status codes ( 500 – 599 ).
	// Caller	: Opposite of internal that indicate caller give some data that cause the error like HTTP client error status codes ( 400 – 499 )
	Internal() bool
	Temporary() bool // opposite is permanent situation
	Timeout() bool

	// Notify error to user by graphic, sound and vibration (Haptic Feedback)
	Notify()

	// Add below method is not force by this interface but we must implement it to respect golang error interface as inner syntax
	// **ATTENTION** assign Error in error typed variables cause serious problems.
	// **ATTENTION** nil interface assign into other interface don't make nil value.
	Error() string

	Details
	MediaType
	Stringer
}
