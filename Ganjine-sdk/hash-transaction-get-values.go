/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
	"../srpc"
)

// HashTransactionGetValues find records by indexes that store before in consistently!
// It will get index from transaction managers not indexes nodes!
func HashTransactionGetValues(c *ganjine.Cluster, req *gs.HashTransactionGetValuesReq) (res *gs.HashTransactionGetValuesRes, err error) {
	var node *ganjine.Node = c.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return nil, ganjine.ErrGanjineNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		return gs.HashTransactionGetValues(req)
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return nil, err
	}

	st.Service = &gs.HashTransactionGetValuesService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, st)
	if err != nil {
		return nil, err
	}

	res = &gs.HashTransactionGetValuesRes{}
	res.SyllabDecoder(srpc.GetPayload(st.IncomePayload))
	return res, st.Err
}
