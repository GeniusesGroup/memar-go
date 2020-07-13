/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// DeleteIndexHashReq is request structure of DeleteIndexHash()
type DeleteIndexHashReq struct {
	IndexHash [32]byte
}

// DeleteIndexHash use to delete exiting index hash with all related records IDs!
// It wouldn't delete related records! Use DeleteIndexHistory() instead if you want delete all records too!
func DeleteIndexHash(c *ganjine.Cluster, req *DeleteIndexHashReq) (err error) {
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

	// Set DeleteIndexHash ServiceID
	reqStream.ServiceID = 3411747355
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return err
	}
	
	return resStream.Err
}

func (req *DeleteIndexHashReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 32+4) // +4 for sRPC ID instead get offset argument

	copy(buf[4:], req.IndexHash[:])

	return
}
