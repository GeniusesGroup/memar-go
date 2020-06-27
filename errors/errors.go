/* For license and copyright information please see LEGAL file in repository */

package errors

import (
	"fmt"
	"hash/crc32"
)

var errPool = map[uint32]*ExtendedError{}

// ExtendedError : This is a extended implementation of error.
type ExtendedError struct {
	Name        string
	Text        string
	Code        uint32
	Information interface{}
}

// New returns an error that formats as the given name, text & code.
// Never change name due to it is very complicated to troubleshooting errors on SDK!
func New(name, text string) error {
	var err = ExtendedError{
		Name: name,
		Text: text,
		Code: crc32.ChecksumIEEE([]byte(name)),
	}
	if errPool[err.Code] != nil {
		fmt.Print("Exiting error >> ", errPool[err.Code].Code, " : ", errPool[err.Code].Text , "\n")
		fmt.Print("New error >> ", err.Code, " : ", err.Text , "\n")
		panic("Duplicate Error code exist ^^")
	}
	errPool[err.Code] = &err
	return &err
}

// GetErrByCode returns desire error if exist or nil!
func GetErrByCode(code uint32) error {
	return errPool[code]
}

// Error : Return text of error.
func (e *ExtendedError) Error() string {
	if e != nil {
		return fmt.Sprintf("Error Code %v: %v \n Additional information: %v \n", e.Code, e.Text, e.Information)
	}

	return "Error is empty"
}

// AddInformationToError : Add information to existing error and return it as new error(pointer)!
func (e *ExtendedError) AddInformationToError(information interface{}) error {
	return &ExtendedError{
		Text:        e.Text,
		Information: information,
		Code:        e.Code,
	}
}

// IsEqualError : Compare two error.
func (e *ExtendedError) IsEqualError(err error) bool {
	if e.Code == err.(*ExtendedError).Code {
		return true
	}

	return false
}
