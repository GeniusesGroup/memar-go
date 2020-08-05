/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// SetRecord respect all data in record and don't change something like RecordID or WriteTime!
// If data like OwnerAppID is wrong you can't get record anymore!
func SetRecord(c *ganjine.Cluster, req *gs.SetRecordReq) (err error) {
	// Generate hash and copy it to Record
	var checksum = c.Server.Cryptography.ChecksumGenerator.Generate(req.Record[32:])
	copy(req.Record[:], checksum[:])

	var recordID [16]byte
	copy(recordID[:], req.Record[32:])
	var node *ganjine.Node = c.GetNodeByRecordID(recordID)
	if node == nil {
		return ErrNoNodeAvailable
	}

	// Check if desire node is local node!
	if node.Conn == nil {
		err = gs.SetRecord(req)
		return
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = node.Conn.MakeBidirectionalStream(0)
	if err != nil {
		return err
	}

	// Set SetRecord ServiceID
	reqStream.ServiceID = 10488062
	reqStream.Payload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, reqStream)
	if err != nil {
		return err
	}

	return resStream.Err
}
