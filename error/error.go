/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"crypto/sha512"
	"fmt"
	"strconv"

	lang "../language"
	"../log"
)

// Error is a extended implementation of error.
// Never change urn due to it adds unnecessary complicated troubleshooting errors on SDK!
type Error struct {
	urn         string // "urn:giti:{{domain-name}}:error:{{error-name}}"
	id          uint64
	idAsString  string
	detail      map[lang.Language]Detail
	Chain       *Error
	information interface{}
	JSON        []byte
	Syllab      []byte
}

// Detail store detail about an error
type Detail struct {
	Domain string
	Short  string
	Long   string
}

// New returns a new error!
func New(urn string) *Error {
	var err = Error{
		urn:    urn,
		detail: make(map[lang.Language]Detail),
	}
	return &err
}

// GetCode return id of error if err id exist.
func GetCode(err error) uint64 {
	if err == nil {
		return 0
	}
	var exErr *Error
	exErr = err.(*Error)
	if exErr != nil {
		return exErr.id
	}
	// if error not nil but not Error, pass biggest number!
	return 18446744073709551615
}

// Save finalize needed logic on given error and save to Errors global variable.
func (e *Error) Save() *Error {
	if e.urn == "" {
		log.Fatal("Error must have URN to save it in platform errors pools! >> ", *e)
	}

	e.IDCalculator()

	e.Syllab = e.syllabEncoder()
	e.JSON = e.jsonEncoder()

	Errors.AddError(e)
	return e
}

// SetDetail add short and long text detail to existing error and return it.
func (e *Error) SetDetail(lang lang.Language, domain, short, long string) *Error {
	e.detail[lang] = Detail{
		Domain: domain,
		Short:  short,
		Long:   long,
	}

	return e
}

// IDCalculator set error ID by error urn
func (e *Error) IDCalculator() {
	var hash = sha512.Sum512([]byte(e.urn))
	e.id = uint64(hash[0]) | uint64(hash[1])<<8 | uint64(hash[2])<<16 | uint64(hash[3])<<24 | uint64(hash[4])<<32 | uint64(hash[5])<<40 | uint64(hash[6])<<48 | uint64(hash[7])<<56
	e.idAsString = strconv.FormatUint(e.id, 10)
	return
}

// Detail return detail of the error in desire language!
func (e *Error) Detail(lang lang.Language) Detail {
	return e.detail[lang]
}

// URN return URN of error.
func (e *Error) URN() string {
	return e.urn
}

// ID return id of error.
func (e *Error) ID() uint64 {
	return e.id
}

// IDasString return id of error as string.
func (e *Error) IDasString() string {
	return e.idAsString
}

// Error return full details of error in text.
func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("Error Code: %v\n Short detail: %v\n Long detail: %v\n Error Additional information: %v\n", e.id, e.detail[log.Language].Short, e.detail[log.Language].Long, e.information)
}

// AddInformation add to existing error and return it as new error(pointer) with chain errors!
func (e *Error) AddInformation(information interface{}) *Error {
	return &Error{
		Chain:       e,
		information: information,
	}
}

// Equal compare two Error.
func (e *Error) Equal(err *Error) bool {
	if e == nil && err == nil {
		return true
	}
	if e != nil && err != nil && e.id == err.id {
		return true
	}
	return false
}

// IsEqual compare two error.
func (e *Error) IsEqual(err error) bool {
	var exErr = err.(*Error)
	if exErr != nil && e.id == exErr.id {
		return true
	}
	return false
}
