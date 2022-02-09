/* For license and copyright information please see LEGAL file in repository */

package protocol

// Errors is the interface that must implement by any Application!
type Errors interface {
	RegisterError(err Error)
	GetErrorByID(id uint64) (err Error)
	GetErrorByMediaType(mt string) (err Error)
}

// New() function in any package must call Application.RegisterError() to save the error in application
type Error interface {
	ID() uint64 // copy of MediaType().ID() to improve performance
	MediaType() MediaType
	// Check both flat or chain situation.
	Equal(Error) bool

	// Notify by graphic, sound and vibration or just log it
	// Notify()

	// Add below method is not force by this interface but you must impelement it to respect golang error interface as inner syntax
	Error() string
	// Rarely use, But can use in logging, so It must always return very simple string as `"Error ID: " + err.URN().IDasString()` and GUI app can provide more human friendly details
	Stringer
}
