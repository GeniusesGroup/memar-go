/* For license and copyright information please see LEGAL file in repository */

package json

import (
	er "../error"
	lang "../language"
)

const errorEnglishDomain = "JSON"
const errorPersianDomain = "JSON"

// Declare Errors Details
var (
	ErrJSONNeededTypeNotExist = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Needed Type Not Exist",
		"Custom struct type exist in upper struct type that generator can't access it to know its fields").Save()

	ErrJSONFieldType = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "FieldType",
		"Requested type may include function, interface, int, uint, ... type that can't encode||decode").Save()

	ErrJSONEncodedIncludeNotDeffiendKey = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Encoded Include Not Deffiend Key",
		"Given encoded json string include a key that must not be in the encoded string").Save()

	ErrJSONEncodedCorrupted = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Encoded Corrupted",
		"Given encoded json string corruputed and not encode in the way that can decode").Save()

	ErrJSONEncodedIntegerCorrupted = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Encoded Integer Corrupted",
		"Given encoded json in Integer part corruputed and not encode in the way that can decode").Save()

	ErrJSONEncodedStringCorrupted = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Encoded String Corrupted",
		"Given encoded json in string part corruputed and not encode in the way that can decode").Save()

	ErrJSONEncodedArrayCorrupted = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Encoded Array Corrupted",
		"Given encoded json in array part corruputed and not encode in the way that can decode").Save()

	ErrJSONEncodedSliceCorrupted = er.New().SetDetail(lang.EnglishLanguage, errorEnglishDomain, "Encoded Slice Corrupted",
		"Given encoded json in slice part corruputed and not encode in the way that can decode").Save()
)
