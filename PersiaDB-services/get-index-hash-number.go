/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var getIndexHashNumberService = chaparkhane.Service{
	Name:            "GetIndexHashNumber",
	IssueDate:       0,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         GetIndexHashNumber,
	Description: []string{
		"",
	},
	TAGS: []string{""},
}

type getIndexHashNumberReq struct{}

type getIndexHashNumberRes struct{}

func getIndexHashNumber(st *chaparkhane.Stream, req *getIndexHashNumberReq) (res *getIndexHashNumberRes, err error) {
	return res, nil
}

// GetIndexHashNumber use to get number of recordsID register for specific IndexHash
func GetIndexHashNumber(s *chaparkhane.Server, st *chaparkhane.Stream) {}

func (req *getIndexHashNumberReq) validator() error {
	return nil
}

func (req *getIndexHashNumberReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *getIndexHashNumberRes) syllabEncoder(buf []byte) error {
	return nil
}
