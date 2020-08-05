/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// GetRecord get the specific record by its ID!
func GetRecord(c *ganjine.Cluster, req *gs.GetRecordReq) (res *gs.GetRecordRes, err error) {
	// TODO::: First read from local OS (related lib) as cache
	// TODO::: Write to local OS as cache if not enough storage exist do GC(Garbage Collector)
	
	var node *ganjine.Node = c.GetNodeByRecordID(req.RecordID)
	if node == nil {
		return nil, ErrNoNodeAvailable
	}

	// Check if desire node is local node!
	if node.Conn == nil {
		res, err = gs.GetRecord(req)
		return
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = node.Conn.MakeBidirectionalStream(0)
	if err != nil {
		return nil, err
	}

	// Set GetRecord ServiceID
	reqStream.ServiceID = 4052491139
	reqStream.Payload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, reqStream)
	if err != nil {
		return nil, err
	}

	res = &gs.GetRecordRes{}
	err = res.SyllabDecoder(resStream.Payload[4:])
	if err != nil {
		return nil, err
	}

	return res, resStream.Err
}
