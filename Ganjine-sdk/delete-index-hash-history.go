/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// DeleteIndexHashHistoryReq is request structure of DeleteIndexHashHistory()
type DeleteIndexHashHistoryReq struct {
	IndexHash [32]byte
}

// DeleteIndexHashHistory use to delete all record associate to given index and delete index itself!
func DeleteIndexHashHistory(c *ganjine.Cluster, req *DeleteIndexHashHistoryReq) (err error) {
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

	// Set DeleteIndexHashHistory ServiceID
	reqStream.ServiceID = 691384835
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return err
	}
	return resStream.Err
}

func (req *DeleteIndexHashHistoryReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 32+4) // +4 for sRPC ID instead get offset argument

	copy(buf[4:], req.IndexHash[:])

	return
}
