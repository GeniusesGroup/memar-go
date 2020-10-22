/* For license and copyright information please see LEGAL file in repository */

package json

import (
	errorr "../error"
	lang "../language"
)

// Declare Errors Details
var (
	ErrJSONNeededTypeNotExist = errorr.New().SetDetail(lang.EnglishLanguage, "JSON - Needed Type Not Exist",
		"Custom struct type exist in upper struct type that generator can't access it to know its fields").Save()

	ErrJSONFieldType = errorr.New().SetDetail(lang.EnglishLanguage, "JSON - FieldType",
		"Requested type may include function, interface, int, uint, ... type that can't encode||decode").Save()

	ErrJSONEncodedIncludeNotDeffiendKey = errorr.New().SetDetail(lang.EnglishLanguage, "JSON - Encoded Include Not Deffiend Key",
		"Given encoded json string include a key that must not be in the encoded string").Save()

	ErrJSONEncodedStringCorrupted = errorr.New().SetDetail(lang.EnglishLanguage, "JSON - Encoded String Corrupted",
		"Given encoded json string corruputed and not encode in the way that can decode").Save()
)
