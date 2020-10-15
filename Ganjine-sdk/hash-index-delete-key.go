/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// HashIndexDeleteKey use to delete exiting index hash with all related records IDs!
// It wouldn't delete related records! Use DeleteIndexHistory() instead if you want delete all records too!
func HashIndexDeleteKey(c *ganjine.Cluster, req *gs.HashIndexDeleteKeyReq) (err error) {
	var node *ganjine.Node = c.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return ganjine.ErrGanjineNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		return gs.HashIndexDeleteKey(req)
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return err
	}

	st.Service = &gs.HashIndexDeleteKeyService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, st)
	if err != nil {
		return err
	}

	return st.Err
}
