/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// WriteRecordReq is request structure of WriteRecord()
type WriteRecordReq struct {
	RecordID [16]byte
	Offset   uint64 // Do something like block storage API
	Limit    uint64 // Do something like block storage API
	Data     []byte
}

// WriteRecord use to write some part of a record!
// Don't use this service until you force to use!
func WriteRecord(c *ganjine.Cluster, req *WriteRecordReq) (err error) {
	var node *ganjine.Node = c.GetNodeByRecordID(req.RecordID)
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

	// Set WriteRecord ServiceID
	reqStream.ServiceID = 3836795965
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return err
	}

	return resStream.Err
}

func (req *WriteRecordReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 32+len(req.Data)+4) // +4 for sRPC ID instead get offset argument

	// RecordID
	copy(buf[4:], req.RecordID[:])
	// Offset
	buf[20] = byte(req.Offset)
	buf[21] = byte(req.Offset >> 8)
	buf[22] = byte(req.Offset >> 16)
	buf[23] = byte(req.Offset >> 24)
	buf[24] = byte(req.Offset >> 32)
	buf[25] = byte(req.Offset >> 40)
	buf[26] = byte(req.Offset >> 48)
	buf[27] = byte(req.Offset >> 56)
	// Limit
	buf[28] = byte(req.Limit)
	buf[29] = byte(req.Limit >> 8)
	buf[30] = byte(req.Limit >> 16)
	buf[31] = byte(req.Limit >> 24)
	buf[32] = byte(req.Limit >> 32)
	buf[33] = byte(req.Limit >> 40)
	buf[34] = byte(req.Limit >> 48)
	buf[35] = byte(req.Limit >> 56)
	// Data
	copy(buf[32:], req.Data[:])

	return
}
