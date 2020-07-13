/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// SetIndexHashReq is request structure of SetIndexHash()
type SetIndexHashReq struct {
	IndexHash [32]byte
	RecordID  [16]byte
}

// SetIndexHash use to set a record ID to new||exiting index hash!
func SetIndexHash(c *ganjine.Cluster, req *SetIndexHashReq) (err error) {
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

	// Set SetIndexHash ServiceID
	reqStream.ServiceID = 1881585857
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return err
	}
	return resStream.Err
}

func (req *SetIndexHashReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 48+4) // +4 for sRPC ID instead get offset argument

	copy(buf[4:], req.IndexHash[:])
	copy(buf[36:], req.RecordID[:])

	return
}
