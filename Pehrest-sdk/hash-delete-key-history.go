/* For license and copyright information please see LEGAL file in repository */

package psdk

import (
	"../achaemenid"
	er "../error"
	"../ganjine"
	"../pehrest"
)

// HashDeleteKeyHistory delete all record associate to given index and delete index itself!
func HashDeleteKeyHistory(req *pehrest.HashDeleteKeyHistoryReq) (err *er.Error) {
	var node *ganjine.Node = ganjine.Cluster.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return ganjine.ErrNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		return pehrest.HashDeleteKeyHistory(req)
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &pehrest.HashDeleteKeyHistoryService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(st)
	if err != nil {
		return
	}
	return st.Err
}
