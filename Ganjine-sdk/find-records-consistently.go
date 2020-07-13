/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// FindRecordsConsistentlyReq is request structure of FindRecordsConsistently()
type FindRecordsConsistentlyReq struct {
	IndexHash [32]byte
}

// FindRecordsConsistentlyRes is response structure of FindRecordsConsistently()
type FindRecordsConsistentlyRes struct {
	RecordID [][16]byte
}

// FindRecordsConsistently use to find records by indexes that store before in consistently!
// It will get index from transaction managers not indexes nodes!
func FindRecordsConsistently(c *ganjine.Cluster, req *FindRecordsConsistentlyReq) (res *FindRecordsConsistentlyRes, err error) {
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

	// Set FindRecordsConsistently ServiceID
	reqStream.ServiceID = 480215407
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return nil, err
	}

	res = &FindRecordsConsistentlyRes{}
	err = res.syllabDecoder(resStream.Payload[4:])
	if err != nil {
		return nil, err
	}
	
	return res, resStream.Err
}

func (req *FindRecordsConsistentlyReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 32+4) // +4 for sRPC ID instead get offset argument

	copy(buf[4:], req.IndexHash[:])

	return
}

func (res *FindRecordsConsistentlyRes) syllabDecoder(buf []byte) error {
	return nil
}
