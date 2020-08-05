/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// FindRecords get related RecordsID that set to given indexHash before!
// get 64 related index to given IndexHash even if just one of them use!
func FindRecords(c *ganjine.Cluster, req *gs.FindRecordsReq) (res *gs.FindRecordsRes, err error) {
	var node *ganjine.Node = c.GetNodeByIndexHash(req.IndexHash)
	if node == nil {
		return nil, ErrNoNodeAvailable
	}

	// Check if desire node is local node!
	if node.Conn == nil {
		res, err = gs.FindRecords(req)
		return
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = node.Conn.MakeBidirectionalStream(0)
	if err != nil {
		return nil, err
	}

	// Set FindRecords ServiceID
	reqStream.ServiceID = 1992558377
	reqStream.Payload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, reqStream)
	if err != nil {
		return nil, err
	}

	res = &gs.FindRecordsRes{}
	err = res.SyllabDecoder(resStream.Payload[4:])
	if err != nil {
		return nil, err
	}

	return res, resStream.Err
}
