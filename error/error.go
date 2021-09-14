/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"fmt"
	"strconv"

	"../log"
	"../protocol"
	"../urn"
)

// Error is a extended implementation of error.
// Never change urn due to it adds unnecessary complicated troubleshooting errors on SDK!
// Read more : https://github.com/SabzCity/RFCs/blob/master/Giti-Error.md
type Error struct {
	urn         urn.Giti
	idAsString  string
	detail      map[protocol.LanguageID]*Detail
	details     []protocol.ErrorDetail
	Chain       *Error
	Information interface{}
	JSON        []byte
	Syllab      []byte
}

// New returns a new error!
// "urn:giti:{{domain-name}}:error:{{error-name}}"
func New(urn string) *Error {
	if urn == "" {
		log.Fatal("Error must have URN to save it in platform errors pools!")
	}

	var err = Error{
		detail: make(map[protocol.LanguageID]*Detail),
	}
	err.urn.Init(urn)
	err.idAsString = strconv.FormatUint(err.urn.ID(), 10)
	return &err
}

// SetDetail add error text details to existing error and return it.
func (e *Error) SetDetail(lang protocol.LanguageID, domain, short, long, userAction, devAction string) *Error {
	var errDetail = Detail{
		languageID: lang,
		domain:     domain,
		short:      short,
		long:       long,
		userAction: userAction,
		devAction:  devAction,
	}
	e.detail[lang] = &errDetail
	e.details = append(e.details, &errDetail)
	return e
}

// Save finalize needed logic on given error and register in the application
func (e *Error) Save() *Error {
	e.Syllab = e.ToSyllab()
	e.JSON = e.ToJSON()

	// Force to check by runtime check, due to testing package not let us by any const!
	if protocol.App != nil {
		protocol.App.RegisterError(e)
	}
	return e
}

func (e *Error) URN() protocol.GitiURN                                { return &e.urn }
func (e *Error) IDasString() string                                   { return e.idAsString }
func (e *Error) Details() []protocol.ErrorDetail                      { return e.details }
func (e *Error) Detail(lang protocol.LanguageID) protocol.ErrorDetail { return e.detail[lang] }

// Error return full details of error in text.
func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("Error Code: %v\n Short detail: %v\n Long detail: %v\n Error Additional information: %v\n", e.urn.ID(), e.detail[protocol.AppLanguage].short, e.detail[protocol.AppLanguage].long, e.Information)
}

// AddInformation add to existing error and return it as new error(pointer) with chain errors!
func (e *Error) AddInformation(information interface{}) *Error {
	return &Error{
		Chain:       e,
		Information: information,
	}
}

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
func (e *Error) IsEqual(err error) bool {
	var exErr = err.(*Error)
	if exErr != nil && e.urn.ID() == exErr.urn.ID() {
		return true
	}
	return false
}

// GetCode return id of error if err id exist.
func GetCode(err error) uint64 {
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
