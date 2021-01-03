/* For license and copyright information please see LEGAL file in repository */

package psdk

import (
	"../achaemenid"
	er "../error"
	"../ganjine"
	"../pehrest"
)

// HashDeleteKey use to delete exiting index hash with all related records IDs!
// It wouldn't delete related records! Use DeleteIndexHistory() instead if you want delete all records too!
func HashDeleteKey(req *pehrest.HashDeleteKeyReq) (err *er.Error) {
	var node *ganjine.Node = ganjine.Cluster.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		return ganjine.ErrNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		return pehrest.HashDeleteKey(req)
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &pehrest.HashDeleteKeyService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(st)
	if err != nil {
		return
	}

	return st.Err
}
