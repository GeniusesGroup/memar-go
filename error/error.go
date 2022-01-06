/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"fmt"

	"../protocol"
	"../urn"
)

// New returns a new error
// "urn:giti:{{domain-name}}:error:{{error-name}}"
func New(urn string) *Error {
	if protocol.AppDevMode && urn == "" {
		// This condition will just be true in the dev phase.
		panic("Error must have URN to save it in platform errors pools")
	}

	var err = Error{
		detail: make(map[protocol.LanguageID]*Detail),
	}
	err.urn.Init(urn)
	return &err
}

// GetID return id of error if err id exist.
func GetID(err error) uint64 {
	if err == nil {
		return 0
	}
	var exErr *Error = err.(*Error)
	if exErr != nil {
		return exErr.urn.ID()
	}
	// if error not nil but not Error, pass biggest number!
	return 18446744073709551615
}

// Error is a extended implementation of error.
// Never change urn due to it adds unnecessary complicated troubleshooting errors on SDK!
// Read more : https://github.com/GeniusesGroup/RFCs/blob/master/Error.md
type Error struct {
	urn     urn.Giti
	detail  map[protocol.LanguageID]*Detail
	details []protocol.ErrorDetail

	stringMethod string
	errorMethod  string
}

// SetDetail add error text details to existing error and return it.
func (e *Error) SetDetail(lang protocol.LanguageID, domain, summary, overview, userAction, devAction string) *Error {
	var errDetail = Detail{
		languageID: lang,
		domain:     domain,
		summary:    summary,
		overview:   overview,
		userAction: userAction,
		devAction:  devAction,
	}
	e.detail[lang] = &errDetail
	e.details = append(e.details, &errDetail)
	return e
}

// Save finalize needed logic on given error and register in the application
func (e *Error) Save() *Error {
	e.updateStrings()
	// Force to check by runtime check, due to testing package not let us by any const!
	if protocol.App != nil {
		protocol.App.RegisterError(e)
	}
	return e
}

func (e *Error) URN() protocol.GitiURN                                { return &e.urn }
func (e *Error) Details() []protocol.ErrorDetail                      { return e.details }
func (e *Error) Detail(lang protocol.LanguageID) protocol.ErrorDetail { return e.detail[lang] }
func (e *Error) ToString() string                                     { return e.stringMethod }

// Equal compare two Error.
func (e *Error) Equal(err protocol.Error) bool {
	if e == nil && err == nil {
		return true
	}
	if e != nil && err != nil && e.urn.ID() == err.URN().ID() {
		return true
	}
	return false
}

// IsEqual compare two error.
func (e *Error) IsEqual(err protocol.Error) bool {
	var exErr = err.(*Error)
	if exErr != nil && e.urn.ID() == exErr.urn.ID() {
		return true
	}
	return false
}

// Go compatibility methods. Unwrap provides compatibility for Go 1.13 error chains.
func (e *Error) Error() string { return e.errorMethod }
func (e *Error) Cause() error  { return e }
func (e *Error) Unwrap() error { return e }

func (e *Error) updateStrings() {
	e.stringMethod = "Error ID: " + e.urn.IDasString()
	e.errorMethod = fmt.Sprintf("Error ID: %s\n	Summary: %s\n	Overview: %s\n", e.urn.IDasString(), e.detail[protocol.AppLanguage].summary, e.detail[protocol.AppLanguage].overview)
}
