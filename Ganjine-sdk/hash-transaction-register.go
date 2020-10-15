/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
	"../srpc"
)

// HashTransactionRegister register new transaction on queue and get last record when transaction ready for this one!
func HashTransactionRegister(c *ganjine.Cluster, req *gs.HashTransactionRegisterReq) (res *gs.HashTransactionRegisterRes, err error) {
	var node *ganjine.Node = c.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return nil, ganjine.ErrGanjineNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		res, err = gs.HashTransactionRegister(req)
		return
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return nil, err
	}

	st.Service = &gs.HashTransactionRegisterService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, st)
	if err != nil {
		return nil, err
	}

	res = &gs.HashTransactionRegisterRes{}
	res.SyllabDecoder(srpc.GetPayload(st.IncomePayload))
	return res, st.Err
}
