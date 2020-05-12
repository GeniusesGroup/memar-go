/* For license and copyright information please see LEGAL file in repository */

package services

import "../achaemenid"

var findRecordsConsistentlyService = achaemenid.Service{
	ID:              480215407,
	Name:            "FindRecordsConsistently",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Handler:         FindRecordsConsistently,
	Description: []string{
		`Find records by indexes that store before in consistently!
		It will get index from transaction managers not indexes nodes!
		`,
	},
	TAGS: []string{""},
}

// FindRecordsConsistently use to find records by indexes that store before in consistently!
// It will get index from transaction managers not indexes nodes!
func FindRecordsConsistently(s *achaemenid.Server, st *achaemenid.Stream) {}

type findRecordsConsistentlyReq struct {
	IndexHash [32]byte
}

type findRecordsConsistentlyRes struct {
	RecordIDs [][16]byte
}

func findRecordsConsistently(st *achaemenid.Stream, req *findRecordsConsistentlyReq) (res *findRecordsConsistentlyRes, err error) {
	return nil, nil
}

func (req *findRecordsConsistentlyReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *findRecordsConsistentlyRes) syllabEncoder(buf []byte) error {
	return nil
}
