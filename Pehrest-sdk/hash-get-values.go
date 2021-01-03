/* For license and copyright information please see LEGAL file in repository */

package psdk

import (
	"../achaemenid"
	er "../error"
	"../ganjine"
	"../pehrest"
	"../srpc"
)

// HashGetValues get related RecordsID that set to given IndexKey before.
func HashGetValues(req *pehrest.HashGetValuesReq) (res *pehrest.HashGetValuesRes, err *er.Error) {
	var node *ganjine.Node = ganjine.Cluster.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return nil, ganjine.ErrNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		res, err = pehrest.HashGetValues(req)
		return
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return nil, err
	}

	st.Service = &pehrest.HashGetValuesService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler( st)
	if err != nil {
		return nil, err
	}

	res = &pehrest.HashGetValuesRes{}
	res.SyllabDecoder(srpc.GetPayload(st.IncomePayload))
	return res, st.Err
}
