/* For license and copyright information please see LEGAL file in repository */

package error

import "fmt"

// Error interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
	AddInformationToError(information interface{}) error
	IsEqualError(err error) bool
}

// ExtendedError : This is a extended implementation of error.
type ExtendedError struct {
	Text        string
	Information interface{}
	Code        uint32
	HTTPStatus  uint
}

// NewError : Returns an error that formats as the given text, Code and httpStatus code.
func NewError(text string, errorCode uint32, httpStatus uint) error {
	return &ExtendedError{
		Text:       text,
		Code:       errorCode,
		HTTPStatus: httpStatus}
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
		HTTPStatus:  e.HTTPStatus}
}

// IsEqualError : Compare two error.
func (e *ExtendedError) IsEqualError(err error) bool {
	if e.Code == err.(*ExtendedError).Code {
		return true
	}

	return false
}
