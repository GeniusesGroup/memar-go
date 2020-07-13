/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// GetIndexHashNumberReq is request structure of GetIndexHashNumber()
type GetIndexHashNumberReq struct {
	IndexHash [32]byte
}

// GetIndexHashNumberRes is response structure of GetIndexHashNumber()
type GetIndexHashNumberRes struct {
	RecordNumber uint64
}

// GetIndexHashNumber use to get number of recordsID register for specific IndexHash
func GetIndexHashNumber(c *ganjine.Cluster, req *GetIndexHashNumberReq) (res *GetIndexHashNumberRes, err error) {
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

	// Set GetIndexHashNumber ServiceID
	reqStream.ServiceID = 222077451
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return nil, err
	}

	res = &GetIndexHashNumberRes{}
	err = res.syllabDecoder(resStream.Payload[4:])
	if err != nil {
		return nil, err
	}

	return res, resStream.Err
}

func (req *GetIndexHashNumberReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 32+4) // +4 for sRPC ID instead get offset argument

	// Index Hash
	copy(buf[4:], req.IndexHash[:])

	return
}

func (res *GetIndexHashNumberRes) syllabDecoder(buf []byte) error {
	return nil
}
