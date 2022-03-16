/* For license and copyright information please see LEGAL file in repository */

package protocol

// Errors is the interface that must implement by any Application.
type Errors interface {
	RegisterError(err Error)
	GetErrorByID(id uint64) (err Error)
	GetErrorByMediaType(mt string) (err Error)
}

// New() function in any package must call Application.RegisterError() to save the error in application
// It is similar to opaque error model describe here: https://dave.cheney.net/paste/gocon-spring-2016.pdf
type Error interface {
	ID() uint64         // copy of MediaType().ID() to improve performance
	IDasString() string // copy of MediaType().IDasString() to improve performance
	MediaType() MediaType

	// Check both flat or chain situation.
	Equal(Error) bool

	// Who cause the error?
	// Internal	: means calling process logic has runtime bugs like HTTP server error status codes ( 500 – 599 ).
	// Caller	: Opposite of internal that indicate caller give some data that cause the error like HTTP client error status codes ( 400 – 499 )
	Internal() bool
	Temporary() bool // opposite is permanent situation

	// Notify by graphic, sound and vibration or just log it
	// Notify()

	// Add below method is not force by this interface but you must impelement it to respect golang error interface as inner syntax
	Error() string
	// Rarely use, But can use in logging, so It must always return very simple string as "err.ID()" and GUI app can provide more human friendly details
	Stringer
}
