/* For license and copyright information please see LEGAL file in repository */

package protocol

// Errors is the interface that must implement by any Application!
type Errors interface {
	SaveError(err Error)
	GetErrorByID(id uint64) (err Error)
	GetErrorByURN(urn string) (err Error)
}

// New() function in any package must call Application.SaveError() to save the error in application
// The goals is to retrieve an error by ID or URN.
type Error interface {
	URN() GitiURN
	IDasString() string
	Details() []ErrorDetail
	Detail(LanguageID) ErrorDetail
	Equal(Error) bool

	// Add below method is not force by this interface but you must impelement it to respect golang error interface as inner syntax!!
	Error() string
}

type ErrorDetail interface {
	Language() LanguageID
	// Domain return locale domain name that error belongs to it!
	Domain() string
	// Short return locale general short error detail
	Short() string
	// Long return locale general long error detail
	Long() string
	// UserAction return locale user action that user do when face this error
	UserAction() string
	// DevAction return locale technical advice for developers
	DevAction() string
}
