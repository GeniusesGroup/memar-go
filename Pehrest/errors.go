/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	er "../error"
	"../protocol"
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
	ErrRecordNil.Init("urn:giti:index.protocol:error:record-nil")
	ErrRecordNil.SetDetail(protocol.LanguageEnglish, domainEnglish, "Record Nil",
		"Given record can't be nil",
		"",
		"",
		nil)

	ErrRecordNotValid.Init("urn:giti:index.protocol:error:record-not-valid")
	ErrRecordNotValid.SetDetail(protocol.LanguageEnglish, domainEnglish, "Record Not Valid",
		"Given recordID exist in storage devices but has diffrent StructureID",
		"",
		"",
		nil)

	ErrRecordNotExist.Init("urn:giti:index.protocol:error:record-not-exist")
	ErrRecordNotExist.SetDetail(protocol.LanguageEnglish, domainEnglish, "Record Not Exist",
		"Given recordID not exist in any storage devices",
		"",
		"",
		nil)

	ErrRecordManipulated.Init("urn:giti:index.protocol:error:record-manipulated")
	ErrRecordManipulated.SetDetail(protocol.LanguageEnglish, domainEnglish, "Record Manipulated",
		"Index record has problem when engine try to read it from storage devices",
		"",
		"",
		nil)

	ErrIndexValueAlreadyExist.Init("urn:giti:index.protocol:error:index-value-already-exist")
	ErrIndexValueAlreadyExist.SetDetail(protocol.LanguageEnglish, domainEnglish, "Index Value Already Exist",
		"",
		"",
		"",
		nil)
}
