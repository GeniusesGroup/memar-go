/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// listenToIndexReq is request structure of listenToIndex()
type listenToIndexReq struct {
	IndexHash [32]byte
}

// listenToIndexRes is response structure of listenToIndex()
type listenToIndexRes struct {
	// Record []byte TODO::: it can't be simple byte, maybe channel
}

// listenToIndex use to get the recordID by index hash when new record set!
// Must send this request to specific node that handle that range!!
func listenToIndex(c *ganjine.Cluster, req *listenToIndexReq) (res *listenToIndexRes, err error) {
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

	// Set listenToIndex ServiceID
	reqStream.ServiceID = 2145882122
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return nil, err
	}

	// Sender can reuse exiting stream to send new record

	res = &listenToIndexRes{}
	err = res.syllabDecoder(resStream.Payload[4:])
	if err != nil {
		return nil, err
	}

	return res, resStream.Err
}

func (req *listenToIndexReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 32+4) // +4 for sRPC ID instead get offset argument

	copy(buf[4:], req.IndexHash[:])

	return
}

func (res *listenToIndexRes) syllabDecoder(buf []byte) error {
	return nil
}
