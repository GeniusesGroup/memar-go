/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// HashIndexDeleteKeyHistory delete all record associate to given index and delete index itself!
func HashIndexDeleteKeyHistory(c *ganjine.Cluster, req *gs.HashIndexDeleteKeyHistoryReq) (err error) {
	var node *ganjine.Node = c.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return ganjine.ErrGanjineNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		return gs.HashIndexDeleteKeyHistory(req)
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return err
	}

	st.Service = &gs.HashIndexDeleteKeyHistoryService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, st)
	if err != nil {
		return err
	}
	return st.Err
}
