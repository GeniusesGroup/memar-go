/* For license and copyright information please see LEGAL file in repository */

package services

import chaparkhane "../ChaparKhane"

var finishTransactionService = chaparkhane.Service{
	ID:              3962420401,
	Name:            "FinishTransaction",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         FinishTransaction,
	Description: []string{
		`use to approve transaction!
		Transaction Manager will set record and index! no further action need after this call!
		`,
	},
	TAGS: []string{""},
}

// FinishTransaction use to approve transaction!
// Transaction Manager will set record and index! no further action need after this call!
func FinishTransaction(s *chaparkhane.Server, st *chaparkhane.Stream) {}

type finishTransactionReq struct {
	IndexHash [32]byte
	Record    []byte
}

func finishTransaction(st *chaparkhane.Stream, req *finishTransactionReq) (err error) {
	return nil
}

func (req *finishTransactionReq) syllabDecoder(buf []byte) error {
	return nil
}
