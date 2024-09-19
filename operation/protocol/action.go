/* For license and copyright information please see the LEGAL file in the code repository */

package operation_p

type Field_ActionType interface {
	ActionType() ActionType
}

// ActionType indicate all ActionType.
// Famous action types are CRUD(Create, Read, Update, Delete)
// https://en.wikipedia.org/wiki/Create,_read,_update_and_delete
type ActionType uint64

// Set given ActionType to the AT.
// e.g. c.Set(ActionType_Create|ActionType_Delete)
func (at *ActionType) Set(with ActionType) {
	*at = with
}

// Check check given ActionType exist in the AT!
func (at ActionType) Check(with ActionType) (exist bool) {
	return with&at == with
}

const ActionType_None ActionType = 0

// Action types
// Can mix by binary OR e.g. ActionType_Create|ActionType_Read
const (
	ActionType_Create ActionType = (1 << iota)
	ActionType_Read
	ActionType_Update
	ActionType_Delete
	ActionType_List // Find, Filter
	ActionType_Notify
	ActionType_Approve

	ActionType_All ActionType = 0b1111111111111111111111111111111111111111111111111111111111111111
)
