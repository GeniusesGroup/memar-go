/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var readRecordService = chaparkhane.Service{
	ID:              108857663,
	Name:            "ReadRecord",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         ReadRecord,
	Description: []string{
		`use to read some part of a record! It must send to proper node otherwise get not found error!
		Mostly use to get metadata first to know about record size before get it to split to some shorter part!
		`,
	},
	TAGS: []string{""},
}

// ReadRecord use to read some part of a record! It must send to proper node otherwise get not found error!
// Mostly use to get metadata first to know about record size before get it to split to some shorter part!
func ReadRecord(s *chaparkhane.Server, st *chaparkhane.Stream) {}

type readRecordReq struct {
	RecordID [16]byte
	Offset   uint64 // Do something like block storage API
	Limit    uint64 // Do something like block storage API
}

type readRecordRes struct {
	Record []byte
}

func readRecord(st *chaparkhane.Stream, req *readRecordReq) (res *readRecordRes) {
	// Check Cache first by ID

	// Retrive Data from storage engine.

	// Cache object by ID

	return nil
}

func (req *readRecordReq) validator() error {
	return nil
}

func (req *readRecordReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *readRecordRes) syllabEncoder(buf []byte) error {
	return nil
}
