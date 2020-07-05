/* For license and copyright information please see LEGAL file in repository */

package errors

import (
	"fmt"
	"hash/crc32"
	"strconv"

	"../log"
)

var errPool = map[uint32]*ExtendedError{}

// ExtendedError : This is a extended implementation of error.
type ExtendedError struct {
	name        string
	text        string
	code        uint32
	information interface{}
}

// New returns an error that formats as the given name, text & code.
// Never change name due to it is very complicated to troubleshooting errors on SDK!
func New(name, text string) error {
	var err = ExtendedError{
		name: name,
		text: text,
		code: crc32.ChecksumIEEE([]byte(name)),
	}
	if errPool[err.code] != nil {
		log.Warn("Duplicate Error code exist")
		log.Warn("Exiting error >> ", errPool[err.code].code, " : ", errPool[err.code].text, "\n")
		log.Fatal("New error >> ", err.code, " : ", err.text, "\n")
	}
	errPool[err.code] = &err
	return &err
}

// GetErrByCode returns desire error if exist or nil!
func GetErrByCode(code uint32) error {
	return errPool[code]
}

// GetCode return code of error if err code exist.
func GetCode(err error) uint32 {
	if err == nil {
		return 0
	}
	var exErr *ExtendedError
	exErr = err.(*ExtendedError)
	if exErr != nil {
		return exErr.code
	}
	// if error not nil but not ExtendedError, pass biggest number!
	return 4294967295
}

// Error returns code of error.
func (e *ExtendedError) Error() string {
	if e == nil {
		return "0"
	}
	return strconv.FormatUint(uint64(e.code), 10)
}

// Text return full details of error in text.
func (e *ExtendedError) Text() string {
	if e != nil {
		return fmt.Sprintf("Error Code %v: %v \n Additional information: %v \n", e.code, e.text, e.information)
	}

	return "Error is empty"
}

// AddInformation add to existing error and return it as new error(pointer)!
func (e *ExtendedError) AddInformation(information interface{}) error {
	return &ExtendedError{
		name:        e.name,
		text:        e.text,
		code:        e.code,
		information: information,
	}
}

// IsEqual compare two error.
func (e *ExtendedError) IsEqual(err error) bool {
	var exErr = err.(*ExtendedError)
	if exErr != nil && e.code == exErr.code {
		return true
	}

	return false
}
