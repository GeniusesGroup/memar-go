/* For license and copyright information please see LEGAL file in repository */

package psdk

import (
	"../achaemenid"
	er "../error"
	"../ganjine"
	"../pehrest"
	"../srpc"
)

// HashTransactionGetValues find records by indexes that store before in consistently!
// It will get index from transaction managers not indexes nodes!
func HashTransactionGetValues(req *pehrest.HashTransactionGetValuesReq) (res *pehrest.HashTransactionGetValuesRes, err *er.Error) {
	var node *ganjine.Node = ganjine.Cluster.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return nil, ganjine.ErrNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		return pehrest.HashTransactionGetValues(req)
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return nil, err
	}

	st.Service = &pehrest.HashTransactionGetValuesService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(st)
	if err != nil {
		return nil, err
	}

	res = &pehrest.HashTransactionGetValuesRes{}
	res.SyllabDecoder(srpc.GetPayload(st.IncomePayload))
	return res, st.Err
}
