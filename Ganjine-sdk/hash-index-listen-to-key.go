/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
	"../srpc"
)

// HashIndexListenToKey get the recordID by index hash when new record set!
// Must send this request to specific node that handle that range!!
func HashIndexListenToKey(c *ganjine.Cluster, req *gs.HashIndexListenToKeyReq) (err error) {
	var node *ganjine.Node = c.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return ganjine.ErrGanjineNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		// return gs.HashIndexListenToKey(req)
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &gs.HashIndexListenToKeyService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, st)
	if err != nil {
		return
	}

	// Sender can reuse exiting stream to send new record

	var res = &gs.HashIndexListenToKeyRes{}
	res.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	return st.Err
}
