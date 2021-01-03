/* For license and copyright information please see LEGAL file in repository */

package psdk

import (
	"../achaemenid"
	er "../error"
	"../ganjine"
	"../pehrest"
	"../srpc"
)

// HashTransactionRegister register new transaction on queue and get last record when transaction ready for this one!
func HashTransactionRegister(req *pehrest.HashTransactionRegisterReq) (res *pehrest.HashTransactionRegisterRes, err *er.Error) {
	var node *ganjine.Node = ganjine.Cluster.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return nil, ganjine.ErrNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		res, err = pehrest.HashTransactionRegister(req)
		return
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return nil, err
	}

	st.Service = &pehrest.HashTransactionRegisterService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(st)
	if err != nil {
		return nil, err
	}

	res = &pehrest.HashTransactionRegisterRes{}
	res.SyllabDecoder(srpc.GetPayload(st.IncomePayload))
	return res, st.Err
}
