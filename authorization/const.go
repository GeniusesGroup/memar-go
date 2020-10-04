/* For license and copyright information please see LEGAL file in repository */

package authorization

// CRUD indicate all CRUD.
type CRUD uint8

// CRUD
const (
	CRUDNone   CRUD = 0b00000000 // 0
	CRUDCreate CRUD = 0b00000001 // 1
	CRUDRead   CRUD = 0b00000010 // 2
	// CRUDCreateRead         CRUD = 0b00000011 // 3
	CRUDUpdate CRUD = 0b00000100 // 4
	// CRUDCreateUpdate       CRUD = 0b00000101 // 5
	// CRUDReadUpdate         CRUD = 0b00000110 // 6
	// CRUDCreateReadUpdate   CRUD = 0b00000111 // 7
	CRUDDelete CRUD = 0b00001000 // 8
	// CRUDCreateDelete       CRUD = 0b00001001 // 9
	// CRUDReadDelete         CRUD = 0b00001010 // 10
	// CRUDCreateReadDelete   CRUD = 0b00001011 // 11
	// CRUDUpdateDelete       CRUD = 0b00001100 // 12
	// CRUDCreateUpdateDelete CRUD = 0b00001101 // 13
	// CRUDReadUpdateDelete   CRUD = 0b00001110 // 14
	CRUDAll CRUD = 0b00001111 // 15
)

// Day indicate all Days.
type Day uint8

// Days
const (
	DayNone      Day = 0b00000000
	DaySaturDay  Day = 0b00000001
	DaySunDay    Day = 0b00000010
	DayMonDay    Day = 0b00000100
	DayTuesDay   Day = 0b00001000
	DayWednesDay Day = 0b00010000
	DayThursDay  Day = 0b00100000
	DayFriDay    Day = 0b01000000
	DayAll       Day = 0b01111111
)
