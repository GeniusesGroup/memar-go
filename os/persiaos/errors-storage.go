/* For license and copyright information please see LEGAL file in repository */

package persiaos

import (
	er "../../error"
	"../../protocol"
)

const errorEnglishStorageDomain = "PersiaOS - Storage"
const errorPersianStorageDomain = "سیستم عامل پارس - ذخیره سازی"

// Errors - Storage
var (
	ErrStorageDeviceProblem = er.New("urn:giti:storage.persiaos:error:device-problem").SetDetail(protocol.LanguageEnglish, errorEnglishStorageDomain, "Device Problem",
		"Requested record exist on the device that had problem to access it!",
		"",
		"").Save()

	ErrStorageNotExist = er.New("urn:giti:storage.persiaos:error:not-exist").SetDetail(protocol.LanguageEnglish, errorEnglishStorageDomain, "Not Exist",
		"Requested record not exist",
		"",
		"").Save()

	ErrStorageNotAuthorize = er.New("urn:giti:storage.persiaos:error:not-authorize").SetDetail(protocol.LanguageEnglish, errorEnglishStorageDomain, "Not Authorized",
		"Requested record belong to other app owner and this app can't access to it",
		"",
		"").Save()
)
