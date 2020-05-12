/* For license and copyright information please see LEGAL file in repository */

package services

import "../achaemenid"

var setIndexService = achaemenid.Service{
	ID:              1881585857,
	Name:            "SetIndex",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Handler:         SetIndex,
	Description: []string{
		"",
	},
	TAGS: []string{""},
}

type setIndexReq struct{}

type setIndexRes struct{}

func setIndex(st *achaemenid.Stream, req *setIndexReq) (res *setIndexRes, err error) {
	return res, nil
}

// SetIndex will
func SetIndex(s *achaemenid.Server, st *achaemenid.Stream) {}

func (req *setIndexReq) validator() error {
	return nil
}

func (req *setIndexReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *setIndexRes) syllabEncoder(buf []byte) error {
	return nil
}
