/* For license and copyright information please see LEGAL file in repository */

package authorization

// CRUD indicate all CRUD.
type CRUD uint8

// CRUD
const (
	CRUDNone   CRUD = 0b00000000 // 0
	CRUDCreate CRUD = 0b00000001 // 1
	CRUDRead   CRUD = 0b00000010 // 2
	CRUDUpdate CRUD = 0b00000100 // 4
	CRUDDelete CRUD = 0b00001000 // 8
	CRUDAll    CRUD = 0b11111111 // 255

	// CRUDCreateRead         CRUD = 0b00000011 // 3
	// CRUDCreateUpdate       CRUD = 0b00000101 // 5
	// CRUDReadUpdate         CRUD = 0b00000110 // 6
	// CRUDCreateReadUpdate   CRUD = 0b00000111 // 7
	// CRUDCreateDelete       CRUD = 0b00001001 // 9
	// CRUDReadDelete         CRUD = 0b00001010 // 10
	// CRUDCreateReadDelete   CRUD = 0b00001011 // 11
	// CRUDUpdateDelete       CRUD = 0b00001100 // 12
	// CRUDCreateUpdateDelete CRUD = 0b00001101 // 13
	// CRUDReadUpdateDelete   CRUD = 0b00001110 // 14
)

// Set given crud to given CRUD!
// e.g. c.Set(CRUDCreate|CRUDDelete)
func (c CRUD) Set(crud CRUD) {
	c = crud
}

// Check check given crud exist in given CRUD!
func (c CRUD) Check(crud CRUD) (exist bool) {
	if crud&c == crud {
		return true
	}
	return false
}
