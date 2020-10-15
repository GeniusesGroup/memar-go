/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
	"../srpc"
)

// ReadRecord read some part of the specific record by its ID!
func ReadRecord(c *ganjine.Cluster, req *gs.ReadRecordReq) (res *gs.ReadRecordRes, err error) {
	// TODO::: First read from local OS (related lib) as cache
	// TODO::: Write to local OS as cache if not enough storage exist do GC(Garbage Collector)

	var node *ganjine.Node = c.GetNodeByRecordID(req.RecordID)
	if node == nil {
		return nil, ganjine.ErrGanjineNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		return gs.ReadRecord(req)
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return nil, err
	}

	st.Service = &gs.ReadRecordService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, st)
	if err != nil {
		return nil, err
	}

	res = &gs.ReadRecordRes{}
	res.SyllabDecoder(srpc.GetPayload(st.IncomePayload))
	return res, st.Err
}
