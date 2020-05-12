/* For license and copyright information please see LEGAL file in repository */

package services

import "../achaemenid"

var registerTransactionService = achaemenid.Service{
	ID:              3840530512,
	Name:            "RegisterTransaction",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Handler:         RegisterTransaction,
	Description: []string{
		`Register new transaction on queue and get last record when transaction ready for this one!
		Requester must send FinishTransaction() immediately, otherwise Transaction manager will drop this request from queue and chain!
		`,
	},
	TAGS: []string{"transactional authority", "index lock ticket"},
}

// transaction write can be on secondary indexes not primary indexes, due to primary index must always unique!
// transaction manager on any node in a replication must sync with master replication corresponding node manager!
// Get a record by ID when record ready to submit! Usually use in transaction queue to act when record ready to read!
// Must send this request to specific node that handle that range!!

// RegisterTransaction use to register new transaction on queue and get last record when transaction ready for this one!
// Requester must send FinishTransaction() immediately, otherwise Transaction manager will drop this request from queue and chain!
func RegisterTransaction(s *achaemenid.Server, st *achaemenid.Stream) {}

type registerTransactionReq struct {
	IndexHash [32]byte
}

type registerTransactionRes struct {
	Data []byte
}

func registerTransaction(st *achaemenid.Stream, req *registerTransactionReq) (res *registerTransactionRes, err error) {
	return res, nil
}

func (req *registerTransactionReq) syllabDecoder(buf []byte) error {
	return nil
}
