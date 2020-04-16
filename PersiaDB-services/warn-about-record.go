/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var warnAboutRecordService = chaparkhane.Service{
	Name:            "WarnAboutRecord",
	IssueDate:       0,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         WarnAboutRecord,
	Description: []string{
		"",
	},
	TAGS: []string{""},
}

type warnAboutRecordReq struct{}

type warnAboutRecordRes struct{}

func warnAboutRecord(st *chaparkhane.Stream, req *warnAboutRecordReq) (res *warnAboutRecordRes, err error) {
	return res, nil
}

// WarnAboutRecord will warn a node that a recordID with its range exist in other replication!
// It can be deleted or corrupted due to the record didn't match expected RecordStructureID, ...
func WarnAboutRecord(s *chaparkhane.Server, st *chaparkhane.Stream) {}

func (req *warnAboutRecordReq) validator() error {
	return nil
}

func (req *warnAboutRecordReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *warnAboutRecordRes) syllabEncoder(buf []byte) error {
	return nil
}

