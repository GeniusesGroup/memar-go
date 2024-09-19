//go:build lang_per

/* For license and copyright information please see the LEGAL file in the code repository */

package operation

// String return name of crud in app language.
func (at *ActionType) String() string {
	switch at {
	case ActionType_None:
		return "هیچ کدام"
	case ActionType_Create:
		return "ایجاد کردن"
	case ActionType_Read:
		return "خواندن"
	case ActionType_Update:
		return "بروزرسانی"
	case ActionType_Delete:
		return "حذف"
	case ActionType_All:
		return "همه"
	}
	return ""
}
