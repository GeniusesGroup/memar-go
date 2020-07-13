/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// DeleteIndexHashRecordReq is request structure of DeleteIndexHashRecord()
type DeleteIndexHashRecordReq struct {
	IndexHash [32]byte
	RecordID  [16]byte
}

// DeleteIndexHashRecord use to delete a record ID from exiting index hash!
func DeleteIndexHashRecord(c *ganjine.Cluster, req *DeleteIndexHashRecordReq) (err error) {
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

	// Set DeleteIndexHashRecord ServiceID
	reqStream.ServiceID = 3481200025
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return err
	}
	return resStream.Err
}

func (req *DeleteIndexHashRecordReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 48+4) // +4 for sRPC ID instead get offset argument

	copy(buf[4:], req.IndexHash[:])
	copy(buf[36:], req.RecordID[:])

	return
}
