/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// SetRecordReq is request structure of SetRecord()
type SetRecordReq struct {
	RecordID [16]byte
	Record   []byte
}

// SetRecord use to respect all data in record and don't change something like RecordID or WriteTime!
// If data like OwnerAppID is wrong you can't get record anymore!
func SetRecord(c *ganjine.Cluster, req *SetRecordReq) (err error) {
	// Generate hash and copy it to Record
	var checksum = c.Server.Cryptography.ChecksumGenerator.Generate(req.Record[32:])
	copy(req.Record[:], checksum[:])

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

	// Set SetRecord ServiceID
	reqStream.ServiceID = 10488062
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return err
	}
	
	return resStream.Err
}

func (req *SetRecordReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, len(req.Record)+4) // +4 for sRPC ID instead get offset argument

	// Don't need to include req.RecordID! we just get it from upper due to Go is strongly type
	// and we don't want to use unsafe here in SDK!
	copy(buf[4:], req.Record)

	return
}
