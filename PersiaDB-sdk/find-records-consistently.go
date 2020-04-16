/* For license and copyright information please see LEGAL file in repository */

package persiadb

type findRecordsConsistentlyReq struct {
	IndexHash [32]byte
	Offset    uint64 // Do something like block storage API
	Limit     uint16 // can't get more than 65535 related record in single request
}

type findRecordsConsistentlyRes struct {
	RecordIDs [][32]byte
}

func findRecordsConsistently(req *findRecordsConsistentlyReq) {}

// FindRecordsConsistently use to find records by indexes that store before in consistently!
// It will get index by master replication not any close to business logic app!
func FindRecordsConsistently() {}
