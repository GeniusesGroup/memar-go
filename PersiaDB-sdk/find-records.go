/* For license and copyright information please see LEGAL file in repository */

package persiadb

import (
	chaparkhane "../ChaparKhane"
	persiadb "../persiaDB"
)

// FindRecordsReq is request structure of FindRecords()
type FindRecordsReq struct {
	IndexHash [32]byte
	Offset    uint64
	Limit     uint64 // It is better to be modulus of 64 or even 256 if storage devices use 4K clusters!
}

// FindRecordsRes is response structure of FindRecords()
type FindRecordsRes struct {
	RecordIDs [][32]byte
}

// FindRecords will get related RecordsID that set to given indexHash before!
// get 64 related index to given IndexHash even if just one of them use!
func FindRecords(s *chaparkhane.Server, c *persiadb.Cluster, req *FindRecordsReq) (res *FindRecordsRes, err error) {
	var _ = c.FindIndexNodeID(req.IndexHash)

	return res, nil
}

func (req *FindRecordsReq) validator() error {
	return nil
}

func (req *FindRecordsReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *FindRecordsRes) syllabEecoder(buf []byte) error {
	return nil
}
