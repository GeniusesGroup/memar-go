/* For license and copyright information please see the LEGAL file in the code repository */

package operation_p

type Field_CRUD interface {
	// TODO::: CRUDType or CRUD??
	CRUDType() CRUD
}

// CRUD indicate all CRUD.
// CRUD == Create, Read, Update, Delete
type CRUD uint8
