/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// DeleteIndexHash use to delete exiting index hash with all related records IDs!
// It wouldn't delete related records! Use DeleteIndexHistory() instead if you want delete all records too!
func DeleteIndexHash(c *ganjine.Cluster, req *gs.DeleteIndexHashReq) (err error) {
	var node *ganjine.Node = c.GetNodeByIndexHash(req.IndexHash)
	if node == nil {
		return ErrNoNodeAvailable
	}

	// Check if desire node is local node!
	if node.Conn == nil {
		err = gs.DeleteIndexHash(req)
		return
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = node.Conn.MakeBidirectionalStream(0)
	if err != nil {
		return err
	}

	// Set DeleteIndexHash ServiceID
	reqStream.ServiceID = 3411747355
	reqStream.Payload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, reqStream)
	if err != nil {
		return err
	}

	return resStream.Err
}
