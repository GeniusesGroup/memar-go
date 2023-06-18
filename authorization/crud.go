/* For license and copyright information please see the LEGAL file in the code repository */

package authorization

import (
	"libgo/protocol"
)

// Check check given crud request exist in base CRUD!
func CheckCrud(base, request protocol.CRUD) (exist bool) {
	return request&base == request
}

// CRUD indicate all CRUD.
// CRUD == Create, Read, Update, Delete
type CRUD protocol.CRUD

// Set given crud to given CRUD!
// e.g. c.Set(CRUDCreate|CRUDDelete)
func (c *CRUD) Set(crud CRUD) {
	*c = crud
}

// Check check given crud exist in given CRUD!
func (c CRUD) Check(crud CRUD) (exist bool) {
	if crud&c == crud {
		return true
	}
	return false
}

// CRUD
// Can mix by binary OR e.g. CRUDCreate|CRUDRead
const (
	CRUD_None   CRUD = 0b00000000
	CRUD_Create CRUD = 0b00000001
	CRUD_Read   CRUD = 0b00000010
	CRUD_Update CRUD = 0b00000100
	CRUD_Delete CRUD = 0b00001000
	// Approve
	CRUD_All CRUD = 0b11111111
)

// String return name of crud in app language.
func (crud CRUD) String() string {
	switch protocol.AppLanguage {
	case protocol.LanguageEnglish:
		switch crud {
		case CRUD_None:
			return "None"
		case CRUD_Create:
			return "Create"
		case CRUD_Read:
			return "Read"
		case CRUD_Update:
			return "Update"
		case CRUD_Delete:
			return "Delete"
		case CRUD_All:
			return "All"
		}
	case protocol.LanguagePersian:
		switch crud {
		case CRUD_None:
			return "هیچ کدام"
		case CRUD_Create:
			return "ایجاد کردن"
		case CRUD_Read:
			return "خواندن"
		case CRUD_Update:
			return "بروزرسانی"
		case CRUD_Delete:
			return "حذف"
		case CRUD_All:
			return "همه"
		}
	}
	return ""
}
