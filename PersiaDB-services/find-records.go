/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var findRecordsService = chaparkhane.Service{
	Name:            "FindRecords",
	IssueDate:       0,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         FindRecords,
	Description: []string{
		`Use to find records by indexes that store before!
		Suggest not get more than 65535 related RecordID in single request!
		`,
	},
	TAGS: []string{""},
}

type findRecordsReq struct {
	IndexHash [32]byte
	Offset    uint64
	Limit     uint64 // It is better to be modulus of 64 or even 256 if storage devices use 4K clusters!
}

type findRecordsRes struct {
	RecordIDs [][32]byte
}

func findRecords(st *chaparkhane.Stream, req *findRecordsReq) (res *findRecordsRes, err error) {
	return nil, nil
}

// FindRecords use to find records by indexes that store before!
// Suggest not get more than 65535 related RecordID in single request!
func FindRecords(s *chaparkhane.Server, st *chaparkhane.Stream) {}

func (req *findRecordsReq) validator() error {
	return nil
}

func (req *findRecordsReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *findRecordsRes) syllabEncoder(buf []byte) error {
	return nil
}
