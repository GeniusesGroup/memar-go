/* For license and copyright information please see LEGAL file in repository */

package syllab

import (
	errorr "../error"
	lang "../language"
)

// Declare Errors Details
var (
	ErrSyllabNeededTypeNotExist = errorr.New().SetDetail(lang.EnglishLanguage, "Syllab - Needed Type Not Exist",
		"Custom struct type exist in upper struct type that generator can't access it to know its fields").Save()

	ErrSyllabFieldType = errorr.New().SetDetail(lang.EnglishLanguage, "Syllab - Field Type",
		"Requested type may include function, interface, int, uint, ... field type that can't encode||decode").Save()

	ErrSyllabArrayLen = errorr.New().SetDetail(lang.EnglishLanguage, "Syllab - Array Len",
		"Length of array larger than 32 bit space that syllab can encode||decode").Save()

	ErrSyllabDecodeSmallSlice = errorr.New().SetDetail(lang.EnglishLanguage, "Syllab - Decode Small Slice",
		"Given slice smaller than expected to decode data from it").Save()

	ErrSyllabDecodeHeapOverFlow = errorr.New().SetDetail(lang.EnglishLanguage, "Syllab - Decode Heap OverFlow",
		"Encoded syllab want to access to out of slice.").Save()
)
