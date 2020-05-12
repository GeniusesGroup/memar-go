/* For license and copyright information please see LEGAL file in repository */

package services

import "../achaemenid"

var writeRecordService = achaemenid.Service{
	ID:              3836795965,
	Name:            "WriteRecord",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Handler:         WriteRecord,
	Description: []string{
		"",
	},
	TAGS: []string{""},
}

// WriteRecord use to write some part of a record!
// Don't use this service until you force to use!
// Recalculate checksum do in database server that is not so efficient!
func WriteRecord(s *achaemenid.Server, st *achaemenid.Stream) {}

type writeRecordReq struct{}

type writeRecordRes struct{}

func writeRecord(st *achaemenid.Stream, req *writeRecordReq) (res *writeRecordRes, err error) {
	return res, nil
}

func (req *writeRecordReq) validator() error {
	return nil
}

func (req *writeRecordReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *writeRecordRes) syllabDecoder(buf []byte) error {
	return nil
}
