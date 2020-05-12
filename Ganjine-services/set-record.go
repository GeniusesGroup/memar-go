/* For license and copyright information please see LEGAL file in repository */

package services

import "../achaemenid"

var setRecordService = achaemenid.Service{
	ID:              10488062,
	Name:            "SetRecord",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Handler:         SetRecord,
	Description: []string{
		"",
	},
	TAGS: []string{""},
}

type setRecordReq struct {
	RecordID [16]byte
	Data     []byte
	Indexes  [][32]byte
}

type setRecordRes struct{}

func setRecord(st *achaemenid.Stream, req *setRecordReq) (res *setRecordRes, err error) {
	return res, nil
}

// SetRecord use to write a whole record and will replace old record if it is exist!
func SetRecord(s *achaemenid.Server, st *achaemenid.Stream) {}

func (req *setRecordReq) validator() error {
	return nil
}

func (req *setRecordReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *setRecordRes) syllabEncoder(buf []byte) error {
	return nil
}
