/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// DeleteIndexHashRecord delete a record ID from exiting index hash!
func DeleteIndexHashRecord(c *ganjine.Cluster, req *gs.DeleteIndexHashRecordReq) (err error) {
	var node *ganjine.Node = c.GetNodeByIndexHash(req.IndexHash)
	if node == nil {
		return ErrNoNodeAvailable
	}

	// Check if desire node is local node!
	if node.Conn == nil {
		err = gs.DeleteIndexHashRecord(req)
		return
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = node.Conn.MakeBidirectionalStream(0)
	if err != nil {
		return err
	}

	// Set DeleteIndexHashRecord ServiceID
	reqStream.ServiceID = 3481200025
	reqStream.Payload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, reqStream)
	if err != nil {
		return err
	}
	return resStream.Err
}
