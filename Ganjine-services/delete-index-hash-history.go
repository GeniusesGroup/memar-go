/* For license and copyright information please see LEGAL file in repository */

package services

import "../achaemenid"

var deleteIndexHashHistoryService = achaemenid.Service{
	ID:              691384835,
	Name:            "DeleteIndexHashHistory",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Handler:         DeleteIndexHashHistory,
	Description: []string{
		"Delete all record associate to given index and delete index itself!",
	},
	TAGS: []string{""},
}

// DeleteIndexHashHistory use to delete all record associate to given index and delete index itself!
func DeleteIndexHashHistory(s *achaemenid.Server, st *achaemenid.Stream) {}

type deleteIndexHashHistoryReq struct {
	IndexHash [32]byte
}

func deleteIndexHashHistory(st *achaemenid.Stream, req *deleteIndexHashHistoryReq) (err error) {
	return nil
}

func (req *deleteIndexHashHistoryReq) syllabDecoder(buf []byte) error {
	return nil
}
