/* For license and copyright information please see LEGAL file in repository */

package services

import "../achaemenid"

var getRecordService = achaemenid.Service{
	ID:              4052491139,
	Name:            "GetRecord",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Handler:         GetRecord,
	Description: []string{
		`use to get a record by given ID! It must send to proper node otherwise get not found error!`,
	},
	TAGS: []string{""},
}

// GetRecord use to get a record by given ID! It must send to proper node otherwise get not found error!
func GetRecord(s *achaemenid.Server, st *achaemenid.Stream) {}

type getRecordReq struct {
	RecordID [16]byte
}

type getRecordRes struct {
	Record []byte
}

func getRecord(st *achaemenid.Stream, req *getRecordReq) (res *getRecordRes) {
	return nil
}

func (req *getRecordReq) validator() error {
	return nil
}

func (req *getRecordReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *getRecordRes) syllabEncoder(buf []byte) error {
	return nil
}
