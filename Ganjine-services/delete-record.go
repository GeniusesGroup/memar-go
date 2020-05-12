/* For license and copyright information please see LEGAL file in repository */

package services

import "../achaemenid"

var deleteRecordService = achaemenid.Service{
	ID:              1758631843,
	Name:            "DeleteRecord",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
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
func DeleteRecord(s *achaemenid.Server, st *achaemenid.Stream) {}

type deleteRecordReq struct {
	RecordID [16]byte
}

func deleteRecord(st *achaemenid.Stream, req *deleteRecordReq) (err error) {
	return nil
}

func (req *deleteRecordReq) syllabDecoder(buf []byte) error {
	return nil
}
