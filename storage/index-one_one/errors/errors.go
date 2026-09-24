/* For license and copyright information please see the LEGAL file in the code repository */

package pehrest

import (
	er "memar/error"
	"memar/protocol"
)

const domainEnglish = "Index"
const domainPersian = "پهرست"

// Errors
var (
	ErrRecordNil              er.Error
	ErrRecordNotValid         er.Error
	ErrRecordNotExist         er.Error
	ErrRecordManipulated      er.Error
	ErrIndexValueAlreadyExist er.Error
)

func init() {
	ErrRecordNil.Init("domain/index.protocol; type=error; name=record-nil")
	ErrRecordNil.SetDetail(protocol.LanguageEnglish, domainEnglish, "Record Nil",
		"Given record can't be nil",
		"",
		"",
		nil)

	ErrRecordNotValid.Init("domain/index.protocol; type=error; name=record-not-valid")
	ErrRecordNotValid.SetDetail(protocol.LanguageEnglish, domainEnglish, "Record Not Valid",
		"Given recordID exist in storage devices but has diffrent StructureID",
		"",
		"",
		nil)

	ErrRecordNotExist.Init("domain/index.protocol; type=error; name=record-not-exist")
	ErrRecordNotExist.SetDetail(protocol.LanguageEnglish, domainEnglish, "Record Not Exist",
		"Given recordID not exist in any storage devices",
		"",
		"",
		nil)

	ErrRecordManipulated.Init("domain/index.protocol; type=error; name=record-manipulated")
	ErrRecordManipulated.SetDetail(protocol.LanguageEnglish, domainEnglish, "Record Manipulated",
		"Index record has problem when engine try to read it from storage devices",
		"",
		"",
		nil)

	ErrIndexValueAlreadyExist.Init("domain/index.protocol; type=error; name=index-value-already-exist")
	ErrIndexValueAlreadyExist.SetDetail(protocol.LanguageEnglish, domainEnglish, "Index Value Already Exist",
		"",
		"",
		"",
		nil)
}
