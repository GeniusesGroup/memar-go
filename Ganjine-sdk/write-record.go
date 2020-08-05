/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// WriteRecord write some part of a record! Don't use this service until you force to use!
func WriteRecord(c *ganjine.Cluster, req *gs.WriteRecordReq) (err error) {
	var node *ganjine.Node = c.GetNodeByRecordID(req.RecordID)
	if node == nil {
		return ErrNoNodeAvailable
	}

	// Check if desire node is local node!
	if node.Conn == nil {
		err = gs.WriteRecord(req)
		return
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = node.Conn.MakeBidirectionalStream(0)
	if err != nil {
		return err
	}

	// Set WriteRecord ServiceID
	reqStream.ServiceID = 3836795965
	reqStream.Payload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, reqStream)
	if err != nil {
		return err
	}

	return resStream.Err
}
