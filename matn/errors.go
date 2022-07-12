/* For license and copyright information please see LEGAL file in repository */

package matn

import (
	er "../error"
	"../protocol"
)

const domainEnglish = "Text"
const domainPersian = "متن"

// Errors
var (
	ErrRecordNil         er.Error
	ErrRecordNotValid    er.Error
	ErrRecordNotExist    er.Error
	ErrRecordManipulated er.Error
)

func init() {
	ErrRecordNil.Init("domain/matn.protocol.error; name=record-nil")
	ErrRecordNil.SetDetail(protocol.LanguageEnglish, domainEnglish, "Record Nil",
		"Given record can't be nil",
		"",
		"",
		nil)

	ErrRecordNotValid.Init("domain/matn.protocol.error; name=record-not-valid")
	ErrRecordNotValid.SetDetail(protocol.LanguageEnglish, domainEnglish, "Record Not Valid",
		"Given recordID exist in storage devices but has different StructureID",
		"",
		"",
		nil)

	ErrRecordNotExist.Init("domain/matn.protocol.error; name=record-not-exist")
	ErrRecordNotExist.SetDetail(protocol.LanguageEnglish, domainEnglish, "Record Not Exist",
		"Given recordID not exist in any storage devices",
		"",
		"",
		nil)

	ErrRecordManipulated.Init("domain/matn.protocol.error; name=record-manipulated")
	ErrRecordManipulated.SetDetail(protocol.LanguageEnglish, domainEnglish, "Record Manipulated",
		"Index record has problem when engine try to read it from storage devices",
		"",
		"",
		nil)
}
