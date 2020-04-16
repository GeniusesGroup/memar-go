/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var deleteIndexHashService = chaparkhane.Service{
	Name:            "DeleteIndexHash",
	IssueDate:       0,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         DeleteIndexHash,
	Description: []string{
		"",
	},
	TAGS: []string{""},
}

type deleteIndexHashReq struct {
	IndexHash [32]byte
}

func deleteIndexHash(st *chaparkhane.Stream, req *deleteIndexHashReq) (err error) {
	return nil
}

// DeleteIndexHash will
func DeleteIndexHash(s *chaparkhane.Server, st *chaparkhane.Stream) {}

func (req *deleteIndexHashReq) validator() error {
	return nil
}

func (req *deleteIndexHashReq) syllabDecoder(buf []byte) error {
	return nil
}
