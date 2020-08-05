/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// ListenToIndex get the recordID by index hash when new record set!
// Must send this request to specific node that handle that range!!
func ListenToIndex(c *ganjine.Cluster, req *gs.ListenToIndexReq) (res *gs.ListenToIndexRes, err error) {
	var node *ganjine.Node = c.GetNodeByIndexHash(req.IndexHash)
	if node == nil {
		return nil, ErrNoNodeAvailable
	}

	// Check if desire node is local node!
	if node.Conn == nil {
		res, err = gs.ListenToIndex(req)
		return
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = node.Conn.MakeBidirectionalStream(0)
	if err != nil {
		return nil, err
	}

	// Set listenToIndex ServiceID
	reqStream.ServiceID = 2145882122
	reqStream.Payload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, reqStream)
	if err != nil {
		return nil, err
	}

	// Sender can reuse exiting stream to send new record

	res = &gs.ListenToIndexRes{}
	err = res.SyllabDecoder(resStream.Payload[4:])
	if err != nil {
		return nil, err
	}

	return res, resStream.Err
}
