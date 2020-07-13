/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// RegisterTransactionReq is request structure of RegisterTransaction()
type RegisterTransactionReq struct {
	IndexHash [32]byte
	RecordID  [16]byte
}

// RegisterTransactionRes is response structure of RegisterTransaction()
type RegisterTransactionRes struct {
	Record []byte
}

// RegisterTransaction use read some part of the specific record by its ID!
func RegisterTransaction(c *ganjine.Cluster, req *RegisterTransactionReq) (res *RegisterTransactionRes, err error) {
	var node *ganjine.Node = c.GetNodeByIndexHash(req.IndexHash)
	if node == nil {
		return nil, ErrNoNodeAvailableToHandleRequests
	}

	// Make new request-response streams
	var conn *achaemenid.Connection = node.GetConnection()
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = conn.MakeBidirectionalStream(0)
	if err != nil {
		return nil, err
	}

	// Set RegisterTransaction ServiceID
	reqStream.ServiceID = 3840530512
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return nil, err
	}

	res = &RegisterTransactionRes{}
	err = res.syllabDecoder(resStream.Payload[4:])
	if err != nil {
		return nil , err
	}

	return res, resStream.Err
}

func (req *RegisterTransactionReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 48+4) // +4 for sRPC ID instead get offset argument

	copy(buf[4:], req.IndexHash[:])
	copy(buf[36:], req.RecordID[:])

	return
}

func (res *RegisterTransactionRes) syllabDecoder(buf []byte) error {
	return nil
}
