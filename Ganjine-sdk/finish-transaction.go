/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// FinishTransactionReq is request structure of FinishTransaction()
type FinishTransactionReq struct {
	IndexHash [32]byte
	Record    []byte
}

// FinishTransaction use to approve transaction!
// Transaction Manager will set record and index! no further action need after this call!
func FinishTransaction(c *ganjine.Cluster, req *FinishTransactionReq) (err error) {
	var node *ganjine.Node = c.GetNodeByIndexHash(req.IndexHash)
	if node == nil {
		return ErrNoNodeAvailableToHandleRequests
	}

	// Make new request-response streams
	var conn *achaemenid.Connection = node.GetConnection()
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = conn.MakeBidirectionalStream(0)
	if err != nil {
		return err
	}

	// Set FinishTransaction ServiceID
	reqStream.ServiceID = 3962420401
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return err
	}
	return resStream.Err
}

func (req *FinishTransactionReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 32+len(req.Record)+4) // +4 for sRPC ID instead get offset argument

	copy(buf[4:], req.IndexHash[:])
	copy(buf[36:], req.Record[:])

	return
}
