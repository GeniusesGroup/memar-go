/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// GetRecordReq is request structure of GetRecord()
type GetRecordReq struct {
	RecordID [16]byte
}

// GetRecordRes is response structure of GetRecord()
type GetRecordRes struct {
	Record []byte
}

// GetRecord use get the specific record by its ID!
func GetRecord(c *ganjine.Cluster, req *GetRecordReq) (res *GetRecordRes, err error) {
	var node *ganjine.Node = c.GetNodeByRecordID(req.RecordID)
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

	// Set GetRecord ServiceID
	reqStream.ServiceID = 4052491139
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return nil, err
	}

	res = &GetRecordRes{}
	err = res.syllabDecoder(resStream.Payload[4:])
	if err != nil {
		return nil, err
	}
	
	return res, resStream.Err
}

func (req *GetRecordReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 16+4) // +4 for sRPC ID instead get offset argument

	copy(buf[4:], req.RecordID[:])

	return
}

func (res *GetRecordRes) syllabDecoder(buf []byte) (err error) {
	return
}
