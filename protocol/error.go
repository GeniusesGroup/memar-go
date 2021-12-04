/* For license and copyright information please see LEGAL file in repository */

package protocol

// Errors is the interface that must implement by any Application!
type Errors interface {
	RegisterError(err Error)
	GetErrorByID(id uint64) (err Error)
	GetErrorByURN(urn string) (err Error)
	ErrorsSDK(humanLanguage LanguageID, machineLanguage MediaType) (sdk []byte, err Error)
}

// New() function in any package must call Application.RegisterError() to save the error in application
// https://github.com/GeniusesGroup/RFCs/blob/master/Error.md
type Error interface {
	URN() GitiURN
	Details() []ErrorDetail
	Detail(LanguageID) ErrorDetail
	Equal(Error) bool

	// Notify by graphic, sound and vibration or just log it
	// Notify()

	// Add below method is not force by this interface but you must impelement it to respect golang error interface as inner syntax!!
	Error() string
	// Use in logging, so It must always return very simple string as `"Error ID: " + err.URN().IDasString()` and GUI app can provide more human friendly details.
	Stringer
}

type ErrorDetail interface {
	Language() LanguageID
	// Domain return locale domain name that error belongs to it!
	Domain() string
	// Summary return locale general summary error text that gives the main points in a concise form
	Summary() string
	// Overview return locale general error text that gives the main ideas without explaining all the details.
	Overview() string
	// UserAction return locale user action that user do when face this error
	UserAction() string
	// DevAction return locale technical advice for developers
	DevAction() string
}
