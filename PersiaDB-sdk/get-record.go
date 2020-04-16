/* For license and copyright information please see LEGAL file in repository */

package persiadb


// GetRecordReq is request structure of GetRecord()
type GetRecordReq struct {
	RecordID [32]byte
}

// GetRecordRes is response structure of GetRecord()
type GetRecordRes struct {
	Record []byte
}

// GetRecord use get the specific record by its ID!
func GetRecord(req *GetRecordReq) (res *GetRecordRes) {
	return nil
}