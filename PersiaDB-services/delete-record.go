/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var deleteRecordService = chaparkhane.Service{
	Name:            "DeleteRecord",
	IssueDate:       0,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         DeleteRecord,
	Description: []string{
		"",
	},
	TAGS: []string{""},
}

type deleteRecordReq struct{
	RecordID [32]byte
}

type deleteRecordRes struct{}

func deleteRecord(st *chaparkhane.Stream, req *deleteRecordReq) (res *deleteRecordRes, err error) {
	return res, nil
}

// DeleteRecord use to send delete request to all cluster to delete a record!
// We don't suggest use this service, due to we strongly suggest think about data as immutable entity(stream and time)
// It won't delete record history just given specific ID.
func DeleteRecord(s *chaparkhane.Server, st *chaparkhane.Stream) {}

func (req *deleteRecordReq) validator() error {
	return nil
}

func (req *deleteRecordReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *deleteRecordRes) syllabEncoder(buf []byte) error {
	return nil
}
