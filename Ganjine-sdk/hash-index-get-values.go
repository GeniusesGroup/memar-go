/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
	"../srpc"
)

// HashIndexGetValues get related RecordsID that set to given IndexKey before.
func HashIndexGetValues(c *ganjine.Cluster, req *gs.HashIndexGetValuesReq) (res *gs.HashIndexGetValuesRes, err error) {
	var node *ganjine.Node = c.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return nil, ganjine.ErrGanjineNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		res, err = gs.HashIndexGetValues(req)
		return
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return nil, err
	}

	st.Service = &gs.HashIndexGetValuesService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, st)
	if err != nil {
		return nil, err
	}

	res = &gs.HashIndexGetValuesRes{}
	res.SyllabDecoder(srpc.GetPayload(st.IncomePayload))
	return res, st.Err
}
