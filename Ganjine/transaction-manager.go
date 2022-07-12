/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../protocol"
)

// TransactionManager or transactional authority store
type TransactionManager struct {
}

func (tm *TransactionManager) init() {}

// GetIndexRecords return related records ID to given index.
func (tm *TransactionManager) GetIndexRecords(indexHash [32]byte) (recordsID [][32]byte) {
	return
}

// RegisterTransaction register new transaction on queue and get last record when transaction ready for this one!
func (tm *TransactionManager) RegisterTransaction(indexHash [32]byte, recordID [32]byte) (Record []byte, err protocol.Error) {
	return
}

// FinishTransaction approve transaction!
func (tm *TransactionManager) FinishTransaction(indexHash [32]byte, record []byte) (err protocol.Error) {
	return
}
