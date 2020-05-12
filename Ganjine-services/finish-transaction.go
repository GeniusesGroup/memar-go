/* For license and copyright information please see LEGAL file in repository */

package services

import "../achaemenid"

var finishTransactionService = achaemenid.Service{
	ID:              3962420401,
	Name:            "FinishTransaction",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
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
func FinishTransaction(s *achaemenid.Server, st *achaemenid.Stream) {}

type finishTransactionReq struct {
	IndexHash [32]byte
	Record    []byte
}

func finishTransaction(st *achaemenid.Stream, req *finishTransactionReq) (err error) {
	return nil
}

func (req *finishTransactionReq) syllabDecoder(buf []byte) error {
	return nil
}
