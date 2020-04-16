/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var setIndexService = chaparkhane.Service{
	Name:            "SetIndex",
	IssueDate:       0,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         SetIndex,
	Description: []string{
		"",
	},
	TAGS: []string{""},
}

type setIndexReq struct{}

type setIndexRes struct{}

func setIndex(st *chaparkhane.Stream, req *setIndexReq) (res *setIndexRes, err error) {
	return res, nil
}

// SetIndex will
func SetIndex(s *chaparkhane.Server, st *chaparkhane.Stream) {}

func (req *setIndexReq) validator() error {
	return nil
}

func (req *setIndexReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *setIndexRes) syllabEncoder(buf []byte) error {
	return nil
}
