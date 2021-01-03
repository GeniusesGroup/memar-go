/* For license and copyright information please see LEGAL file in repository */

package psdk

import (
	"../achaemenid"
	er "../error"
	"../ganjine"
	"../pehrest"
	"../srpc"
)

// HashListenToKey get the recordID by index hash when new record set!
// Must send this request to specific node that handle that range!!
func HashListenToKey(req *pehrest.HashListenToKeyReq) (err *er.Error) {
	var node *ganjine.Node = ganjine.Cluster.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return ganjine.ErrNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		// return pehrest.HashListenToKey(req)
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &pehrest.HashListenToKeyService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler( st)
	if err != nil {
		return
	}

	// Sender can reuse exiting stream to send new record

	var res = &pehrest.HashListenToKeyRes{}
	res.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	return st.Err
}
