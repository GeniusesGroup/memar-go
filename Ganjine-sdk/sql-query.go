/* For license and copyright information please see LEGAL file in repository */

package ganjine

// SQLParser do Parse "Select" SQL query and make ReqQueryData
// We don't support other SQL type like "Insert", "Replace", "Update", "Delete", ...
func SQLParser(selectQuery string) *ReqQueryData {
	reqQueryData := ReqQueryData{}

	return &reqQueryData
}

// ReqQueryData : The request structure of "QueryData()"
type ReqQueryData struct {
	DataVersion []uint16 // It can be empty that means global searching.
	MediaType   []uint16 // It can be empty that means all MediaTypes.
	Conditions  []Condition
}

// Condition can be any indexed data e.g. UUID, TAGS, ...
type Condition struct {
	FieldName         string // It can be empty that means full text queries!
	FieldValue        string // It can be empty that means null value and FieldOperator must be =
	FieldOperator     uint8  // 0:= , 1:> , 2:< , 3:!= , 4:>= , 5:<= , 6:LIKE(regexp) ,
	ConditionOperator uint8  // 0:AND(Must have) , 1:OR(One of 2 condition) , 2:NOT(Must not have)
}

// ResQueryData : The response structure of "QueryData()".
type ResQueryData struct {
	ObjectsUUID [][16]byte
}

// QueryData is func to find Objects UUID with related condition in data or metadata.
func QueryData(logicRequest *ReqQueryData) (*ResQueryData, error) {
	logicResponse := ResQueryData{}

	// We don't return error if no data found, instead we return with empty array

	return &logicResponse, nil
}
