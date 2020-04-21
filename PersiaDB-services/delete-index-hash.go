/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var deleteIndexHashService = chaparkhane.Service{
	ID:              3411747355,
	Name:            "DeleteIndexHash",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         DeleteIndexHash,
	Description: []string{
		`delete exiting index hash with all related records IDs!
		It wouldn't delete related records! Use DeleteIndexHistory() instead if you want delete all records too!`,
	},
	TAGS: []string{""},
}

// DeleteIndexHash use to delete exiting index hash with all related records IDs!
// It wouldn't delete related records! Use DeleteIndexHistory() instead if you want delete all records too!
func DeleteIndexHash(s *chaparkhane.Server, st *chaparkhane.Stream) {}

type deleteIndexHashReq struct {
	IndexHash [32]byte
}

func deleteIndexHash(st *chaparkhane.Stream, req *deleteIndexHashReq) (err error) {
	return nil
}

func (req *deleteIndexHashReq) syllabDecoder(buf []byte) error {
	return nil
}
