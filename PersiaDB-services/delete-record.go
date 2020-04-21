/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var deleteRecordService = chaparkhane.Service{
	ID:              1758631843,
	Name:            "DeleteRecord",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         DeleteRecord,
	Description: []string{
		`Delete specific record by given ID in all cluster!
		We don't suggest use this service, due to we strongly suggest think about data as immutable entity(stream and time)
		It won't delete record history or indexes associate to it!`,
	},
	TAGS: []string{""},
}

// DeleteRecord use to delete specific record by given ID in all cluster!
// We don't suggest use this service, due to we strongly suggest think about data as immutable entity(stream and time)
// It won't delete record history or indexes associate to it!
func DeleteRecord(s *chaparkhane.Server, st *chaparkhane.Stream) {}

type deleteRecordReq struct {
	RecordID [16]byte
}

func deleteRecord(st *chaparkhane.Stream, req *deleteRecordReq) (err error) {
	return nil
}

func (req *deleteRecordReq) syllabDecoder(buf []byte) error {
	return nil
}
