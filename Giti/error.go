/* For license and copyright information please see LEGAL file in repository */

package giti

type Error interface {
	URN() string
	ID() uint64
	IDasString() string
	Detail(LanguageID) ErrorDetail
	Equal(Error) bool

	// Save finalize needed logic on the error and save to a Errors global variable pools to retrieve an error by ID or URN.
	Save() Error

	// Add below method is not force by this interface but you must impelement it to respect golang error interface as inner syntax!!
	Error() string
}

type ErrorDetail interface {
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
