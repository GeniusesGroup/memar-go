/* For license and copyright information please see LEGAL file in repository */

package ganjine

// DataStructure store related data about each unique data structure
type DataStructure struct {
	ID              uint32
	Name            string
	IssueDate       int64
	ExpiryDate      int64
	ExpireInFavorOf uint32 // Other DataStructure.ID
	Status          uint8
	KeysName        []string // By order
	KeysType        []string // By order
	KeysComment     []string // By order
	Description     []string // By app language order
	TAGS            []string
}

// CompleteStructure use to complete Keys--- DataStructure by given data!
func (ds *DataStructure) CompleteStructure(structure interface{}) {}
