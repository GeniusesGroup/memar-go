/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	er "libgo/error"
)

// Declare package errors
var (
	ErrEncodedIncludeNotDefinedKey er.Error
	ErrEncodedCorrupted            er.Error
	ErrEncodedIntegerCorrupted     er.Error
	ErrEncodedStringCorrupted      er.Error
	ErrEncodedArrayCorrupted       er.Error
	ErrEncodedSliceCorrupted       er.Error

// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Errors/JSON_bad_parse

// Decoding
// SyntaxError: JSON.parse: unterminated string literal
// SyntaxError: JSON.parse: bad control character in string literal
// SyntaxError: JSON.parse: bad character in string literal
// SyntaxError: JSON.parse: bad Unicode escape
// SyntaxError: JSON.parse: bad escape character
// SyntaxError: JSON.parse: unterminated string
// SyntaxError: JSON.parse: no number after minus sign
// SyntaxError: JSON.parse: unexpected non-digit
// SyntaxError: JSON.parse: missing digits after decimal point
// SyntaxError: JSON.parse: unterminated fractional number
// SyntaxError: JSON.parse: missing digits after exponent indicator
// SyntaxError: JSON.parse: missing digits after exponent sign
// SyntaxError: JSON.parse: exponent part is missing a number
// SyntaxError: JSON.parse: unexpected end of data
// SyntaxError: JSON.parse: unexpected keyword
// SyntaxError: JSON.parse: unexpected character
// SyntaxError: JSON.parse: end of data while reading object contents
// SyntaxError: JSON.parse: expected property name or '}'
// SyntaxError: JSON.parse: end of data when ',' or ']' was expected
// SyntaxError: JSON.parse: expected ',' or ']' after array element
// SyntaxError: JSON.parse: end of data when property name was expected
// SyntaxError: JSON.parse: expected double-quoted property name
// SyntaxError: JSON.parse: end of data after property name when ':' was expected
// SyntaxError: JSON.parse: expected ':' after property name in object
// SyntaxError: JSON.parse: end of data after property value in object
// SyntaxError: JSON.parse: expected ',' or '}' after property value in object
// SyntaxError: JSON.parse: expected ',' or '}' after property-value pair in object literal
// SyntaxError: JSON.parse: property names must be double-quoted strings
// SyntaxError: JSON.parse: expected property name or '}'
// SyntaxError: JSON.parse: unexpected character
// SyntaxError: JSON.parse: unexpected non-whitespace character after JSON data
)

// TODO::: use json.ietf.org or ??
func init() {
	ErrEncodedIncludeNotDefinedKey.Init("domain/json.ecma-international.org; type=error; name=encoded-include-not-defined-key")
	ErrEncodedCorrupted.Init("domain/json.ecma-international.org; type=error; name=encoded-corrupted")
	ErrEncodedIntegerCorrupted.Init("domain/json.ecma-international.org; type=error; name=encoded-integer-corrupted")
	ErrEncodedStringCorrupted.Init("domain/json.ecma-international.org; type=error; name=encoded-string-corrupted")
	ErrEncodedArrayCorrupted.Init("domain/json.ecma-international.org; type=error; name=encoded-array-corrupted")
	ErrEncodedSliceCorrupted.Init("domain/json.ecma-international.org; type=error; name=encoded-slice-corrupted")
}
