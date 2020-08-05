/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// RegisterTransaction register new transaction on queue and get last record when transaction ready for this one!
func RegisterTransaction(c *ganjine.Cluster, req *gs.RegisterTransactionReq) (res *gs.RegisterTransactionRes, err error) {
	var node *ganjine.Node = c.GetNodeByIndexHash(req.IndexHash)
	if node == nil {
		return nil, ErrNoNodeAvailable
	}

	// Check if desire node is local node!
	if node.Conn == nil {
		res, err = gs.RegisterTransaction(req)
		return
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = node.Conn.MakeBidirectionalStream(0)
	if err != nil {
		return nil, err
	}

	// Set RegisterTransaction ServiceID
	reqStream.ServiceID = 3840530512
	reqStream.Payload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, reqStream)
	if err != nil {
		return nil, err
	}

	res = &gs.RegisterTransactionRes{}
	err = res.SyllabDecoder(resStream.Payload[4:])
	if err != nil {
		return nil, err
	}

	return res, resStream.Err
}
