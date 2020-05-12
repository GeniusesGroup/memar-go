/* For license and copyright information please see LEGAL file in repository */

package services

import "../achaemenid"

var getIndexHashNumberService = achaemenid.Service{
	ID:              222077451,
	Name:            "GetIndexHashNumber",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Handler:         GetIndexHashNumber,
	Description: []string{
		"Get number of recordsID register for specific IndexHash",
	},
	TAGS: []string{""},
}

// GetIndexHashNumber use to get number of recordsID register for specific IndexHash
func GetIndexHashNumber(s *achaemenid.Server, st *achaemenid.Stream) {}

type getIndexHashNumberReq struct {
	IndexHash [32]byte
}

type getIndexHashNumberRes struct {
	RecordNumber uint64
}

func getIndexHashNumber(st *achaemenid.Stream, req *getIndexHashNumberReq) (res *getIndexHashNumberRes, err error) {
	return res, nil
}

func (req *getIndexHashNumberReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *getIndexHashNumberRes) syllabEncoder(buf []byte) error {
	return nil
}
