/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// DeleteRecordReq is request structure of DeleteRecord()
type DeleteRecordReq struct {
	RecordID [16]byte
}

// DeleteRecord use to delete specific record by given ID in all cluster!
// We don't suggest use this service, due to we strongly suggest think about data as immutable entity(stream and time)
// It won't delete record history or indexes associate to it!
func DeleteRecord(c *ganjine.Cluster, req *DeleteRecordReq) (err error) {
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

	// Set DeleteRecord ServiceID
	reqStream.ServiceID = 1758631843
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return err
	}
	return resStream.Err
}

func (req *DeleteRecordReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 16+4) // +4 for sRPC ID instead get offset argument

	copy(buf[4:], req.RecordID[:])

	return
}
