/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// ReadRecordReq is request structure of ReadRecord()
type ReadRecordReq struct {
	RecordID [16]byte
	Offset   uint64 // Do something like block storage API
	Limit    uint64 // Do something like block storage API
}

// ReadRecordRes is response structure of ReadRecord()
type ReadRecordRes struct {
	Data []byte
}

// ReadRecord use read some part of the specific record by its ID!
func ReadRecord(c *ganjine.Cluster, req *ReadRecordReq) (res *ReadRecordRes, err error) {
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

	// Set ReadRecord ServiceID
	reqStream.ServiceID = 108857663
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return nil, err
	}

	res = &ReadRecordRes{}
	err = res.syllabDecoder(resStream.Payload[4:])
	if err != nil {
		return nil , err
	}

	return res, resStream.Err
}

func (req *ReadRecordReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 32+4) // +4 for sRPC ID instead get offset argument

	copy(buf[4:], req.RecordID[:])

	return
}

func (res *ReadRecordRes) syllabDecoder(buf []byte) error {
	return nil
}
