/* For license and copyright information please see LEGAL file in repository */

package object

import (
	er "../error"
	"../protocol"
)

const errorEnglishDomain = "Object"
const errorPersianDomain = "ابجکت"

// Errors
var (
	ErrDeviceProblem = er.New("urn:giti:object.protocol:error:device-problem").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Device Problem",
		"Requested object exist on the device that had problem to access it!",
		"",
		"").Save()

	ErrExist = er.New("urn:giti:object.protocol:error:exist").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Object Exist",
		"Request to make new object can't proccessed due to UUID already exist",
		"",
		"").Save()

	ErrNotExist = er.New("urn:giti:object.protocol:error:not-exist").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Not Exist",
		"Requested object not exist",
		"",
		"").Save()

	ErrNotAuthorize = er.New("urn:giti:object.protocol:error:not-authorize").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Not Authorized",
		"Requested object belong to other app owner and this app can't access to it",
		"",
		"").Save()

	ErrMisMatchedStructureID = er.New("urn:giti:object.protocol:error:mis-matched-structure-id").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Mis Matched StructureID",
		"Can't read object data due to it StructureID not match with desire one",
		"",
		"").Save()

	ErrObjectNotFound = er.New("urn:giti:object.protocol:error:object-not-found").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Object Not Found",
		"Can't find a object with desire data",
		"",
		"").SetDetail(protocol.LanguagePersian, errorPersianDomain, "رکود یافت نشد",
		"رکورد با مشخصات مورد نظر یافت نشد",
		"",
		"").Save()

	ErrSourceNotChangeable = er.New("urn:giti:object.protocol:error:source-not-changeable").SetDetail(protocol.LanguageEnglish, errorEnglishDomain, "Source not Changeable",
		"Can't read from other source than object source",
		"You can't set/write data in this way",
		"Use Save() method in the object directory").Save()
)
