/* For license and copyright information please see LEGAL file in repository */

package protocol

// CRUD indicate all CRUD.
type CRUD uint8

// CRUD
// Can mix by binary OR e.g. CRUDCreate|CRUDRead
const (
	CRUDNone   CRUD = 0b00000000
	CRUDCreate CRUD = 0b00000001
	CRUDRead   CRUD = 0b00000010
	CRUDUpdate CRUD = 0b00000100
	CRUDDelete CRUD = 0b00001000
	// Approve
	CRUDAll    CRUD = 0b11111111
)

// String return name of crud in app language.
func (crud CRUD) String() string {
	switch AppLanguage {
	case LanguageEnglish:
		switch crud {
		case CRUDNone:
			return "None"
		case CRUDCreate:
			return "Create"
		case CRUDRead:
			return "Read"
		case CRUDUpdate:
			return "Update"
		case CRUDDelete:
			return "Delete"
		case CRUDAll:
			return "All"
		}
	case LanguagePersian:
		switch crud {
		case CRUDNone:
			return "هیچ کدام"
		case CRUDCreate:
			return "ایجاد کردن"
		case CRUDRead:
			return "خواندن"
		case CRUDUpdate:
			return "بروزرسانی"
		case CRUDDelete:
			return "حذف"
		case CRUDAll:
			return "همه"
		}
	}
	return ""
}
