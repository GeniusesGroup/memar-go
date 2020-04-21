/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var deleteIndexHashRecordService = chaparkhane.Service{
	ID:              3481200025,
	Name:            "DeleteIndexHashRecord",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         DeleteIndexHashRecord,
	Description: []string{
		"Delete a record ID from exiting index hash",
	},
	TAGS: []string{""},
}

// DeleteIndexHashRecord use to delete a record ID from exiting index hash!
func DeleteIndexHashRecord(s *chaparkhane.Server, st *chaparkhane.Stream) {}

type deleteIndexHashRecordReq struct {
	IndexHash [32]byte
	RecordID  [16]byte
}

func deleteIndexHashRecord(st *chaparkhane.Stream, req *deleteIndexHashRecordReq) (err error) {
	return nil
}

func (req *deleteIndexHashRecordReq) syllabDecoder(buf []byte) error {
	return nil
}
