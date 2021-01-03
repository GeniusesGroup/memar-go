/* For license and copyright information please see LEGAL file in repository */

package psdk

import (
	"../achaemenid"
	er "../error"
	"../ganjine"
	"../pehrest"
	"../srpc"
)

// HashGetValuesNumber get number of recordsID register for specific IndexHash
func HashGetValuesNumber(req *pehrest.HashGetValuesNumberReq) (res *pehrest.HashGetValuesNumberRes, err *er.Error) {
	var node *ganjine.Node = ganjine.Cluster.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return nil, ganjine.ErrNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		return pehrest.HashGetValuesNumber(req)
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return nil, err
	}

	st.Service = &pehrest.HashGetValuesNumberService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(st)
	if err != nil {
		return nil, err
	}

	res = &pehrest.HashGetValuesNumberRes{}
	res.SyllabDecoder(srpc.GetPayload(st.IncomePayload))
	return res, st.Err
}
