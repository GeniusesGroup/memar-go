/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var getIndexHashNumberService = chaparkhane.Service{
	ID:              222077451,
	Name:            "GetIndexHashNumber",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         GetIndexHashNumber,
	Description: []string{
		"Get number of recordsID register for specific IndexHash",
	},
	TAGS: []string{""},
}

// GetIndexHashNumber use to get number of recordsID register for specific IndexHash
func GetIndexHashNumber(s *chaparkhane.Server, st *chaparkhane.Stream) {}

type getIndexHashNumberReq struct {
	IndexHash [32]byte
}

type getIndexHashNumberRes struct {
	RecordNumber uint64
}

func getIndexHashNumber(st *chaparkhane.Stream, req *getIndexHashNumberReq) (res *getIndexHashNumberRes, err error) {
	return res, nil
}

func (req *getIndexHashNumberReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *getIndexHashNumberRes) syllabEncoder(buf []byte) error {
	return nil
}
