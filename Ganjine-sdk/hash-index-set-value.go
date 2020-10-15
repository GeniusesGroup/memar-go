/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// HashIndexSetValue set a record ID to new||exiting index hash!
func HashIndexSetValue(c *ganjine.Cluster, req *gs.HashIndexSetValueReq) (err error) {
	var node *ganjine.Node = c.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return ganjine.ErrGanjineNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		return gs.HashIndexSetValue(req)
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return err
	}

	st.Service = &gs.HashIndexSetValueService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, st)
	if err != nil {
		return err
	}
	return st.Err
}
