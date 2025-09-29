//go:build lang_eng

/* For license and copyright information please see the LEGAL file in the code repository */

package operation

// String return name of crud in app language.
func (at *ActionType) String() string {
	switch at {
	case ActionType_None:
		return "None"
	case ActionType_Create:
		return "Create"
	case ActionType_Read:
		return "Read"
	case ActionType_Update:
		return "Update"
	case ActionType_Delete:
		return "Delete"
	case ActionType_All:
		return "All"
	}
	return ""
}
