/* For license and copyright information please see LEGAL file in repository */

package services

import "../achaemenid"

var findRecordsService = achaemenid.Service{
	ID:              1992558377,
	Name:            "FindRecords",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Handler:         FindRecords,
	Description: []string{
		`Use to find records by indexes that store before!
		Suggest not get more than 65535 related RecordID in single request!
		`,
	},
	TAGS: []string{""},
}

// FindRecords use to find records by indexes that store before!
// Suggest not get more than 65535 related RecordID in single request!
func FindRecords(s *achaemenid.Server, st *achaemenid.Stream) {}

type findRecordsReq struct {
	IndexHash [32]byte
	Offset    uint64
	Limit     uint64 // It is better to be modulus of 64 or even 256 if storage devices use 4K clusters!
}

type findRecordsRes struct {
	RecordIDs [][16]byte
}

func findRecords(st *achaemenid.Stream, req *findRecordsReq) (res *findRecordsRes, err error) {
	return nil, nil
}

func (req *findRecordsReq) validator() error {
	return nil
}

func (req *findRecordsReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *findRecordsRes) syllabEncoder(buf []byte) error {
	return nil
}
