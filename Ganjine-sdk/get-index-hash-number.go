/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// GetIndexHashNumber get number of recordsID register for specific IndexHash
func GetIndexHashNumber(c *ganjine.Cluster, req *gs.GetIndexHashNumberReq) (res *gs.GetIndexHashNumberRes, err error) {
	var node *ganjine.Node = c.GetNodeByIndexHash(req.IndexHash)
	if node == nil {
		return nil, ErrNoNodeAvailable
	}

	// Check if desire node is local node!
	if node.Conn == nil {
		res, err = gs.GetIndexHashNumber(req)
		return
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = node.Conn.MakeBidirectionalStream(0)
	if err != nil {
		return nil, err
	}

	// Set GetIndexHashNumber ServiceID
	reqStream.ServiceID = 222077451
	reqStream.Payload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, reqStream)
	if err != nil {
		return nil, err
	}

	res = &gs.GetIndexHashNumberRes{}
	err = res.SyllabDecoder(resStream.Payload[4:])
	if err != nil {
		return nil, err
	}

	return res, resStream.Err
}
