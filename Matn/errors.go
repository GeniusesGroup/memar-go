/* For license and copyright information please see LEGAL file in repository */

package matn

import (
	er "../error"
	lang "../language"
)

const errorEnglishDomain = "Text"
const errorPersianDomain = "متن"

// Errors
var (
	ErrRecordNil = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Record Nil",
		"Given record can't be nil").Save()

	ErrRecordNotValid = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Record Not Valid",
		"Given recordID exist in storage devices but has diffrent StructureID").Save()

	ErrRecordNotExist = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Record Not Exist",
		"Given recordID not exist in any storage devices").Save()

	ErrRecordManipulated = er.New().SetDetail(lang.LanguageEnglish, errorEnglishDomain, "Record Manipulated",
		"Index record has problem when engine try to read it from storage devices").Save()
)
