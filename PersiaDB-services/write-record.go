/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var writeRecordService = chaparkhane.Service{
	Name:            "WriteRecord",
	IssueDate:       0,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         WriteRecord,
	Description: []string{
		"",
	},
	TAGS: []string{""},
}

type writeRecordReq struct{}

type writeRecordRes struct{}

func writeRecord(st *chaparkhane.Stream, req *writeRecordReq) (res *writeRecordRes, err error) {
	return res, nil
}

// WriteRecord use to write some part of a record!
// Don't use this service until you force to use!
// Recalculate checksum do in database server that is not efficient!
func WriteRecord(s *chaparkhane.Server, st *chaparkhane.Stream) {}

func (req *writeRecordReq) validator() error {
	return nil
}

func (req *writeRecordReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *writeRecordRes) syllabDecoder(buf []byte) error {
	return nil
}
