/* For license and copyright information please see the LEGAL file in the code repository */

package operation_p

type Field_CRUD interface {
	// TODO::: CRUDType or CRUD??
	CRUDType() CRUD
}

// CRUD indicate all CRUD.
// CRUD == Create, Read, Update, Delete
type CRUD uint8

// CRUD
// Can mix by binary OR e.g. CRUDCreate|CRUDRead
const (
	CRUD_None   CRUD = 0b00000000
	CRUD_Create CRUD = 0b00000001
	CRUD_Read   CRUD = 0b00000010
	CRUD_Update CRUD = 0b00000100
	CRUD_Delete CRUD = 0b00001000
	CRUD_Notify CRUD = 0b00010000
	// Approve
	CRUD_All CRUD = 0b11111111
)
