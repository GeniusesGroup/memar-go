/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var deleteIndexRecordService = chaparkhane.Service{
	Name:            "DeleteIndexRecord",
	IssueDate:       0,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         DeleteIndexRecord,
	Description: []string{
		"",
	},
	TAGS: []string{""},
}

type deleteIndexRecordReq struct {
	IndexHash [32]byte
	RecordID  [16]byte
}

type deleteIndexRecordRes struct{}

func deleteIndexRecord(st *chaparkhane.Stream, req *deleteIndexRecordReq) (res *deleteIndexRecordRes, err error) {
	return res, nil
}

// DeleteIndexRecord will
func DeleteIndexRecord(s *chaparkhane.Server, st *chaparkhane.Stream) {}

func (req *deleteIndexRecordReq) validator() error {
	return nil
}

func (req *deleteIndexRecordReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *deleteIndexRecordRes) syllabEncoder(buf []byte) error {
	return nil
}
