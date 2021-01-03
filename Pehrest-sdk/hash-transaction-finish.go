/* For license and copyright information please see LEGAL file in repository */

package psdk

import (
	"../achaemenid"
	er "../error"
	"../ganjine"
	"../pehrest"
)

// HashTransactionFinish approve transaction!
// Transaction Manager will set record and index! no further action need after this call!
func HashTransactionFinish(req *pehrest.HashTransactionFinishReq) (err *er.Error) {
	var node *ganjine.Node = ganjine.Cluster.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return ganjine.ErrNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		return pehrest.HashTransactionFinish(req)
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &pehrest.HashTransactionFinishService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler( st)
	if err != nil {
		return
	}
	return st.Err
}
