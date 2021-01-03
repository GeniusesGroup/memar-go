/* For license and copyright information please see LEGAL file in repository */

package psdk

import (
	"../achaemenid"
	er "../error"
	"../ganjine"
	"../pehrest"
)

// HashSetValue set a record ID to new||exiting index hash!
func HashSetValue(req *pehrest.HashSetValueReq) (err *er.Error) {
	var node *ganjine.Node = ganjine.Cluster.GetNodeByRecordID(req.IndexKey)
	if node == nil {
		err = ganjine.ErrNoNodeAvailable
		return
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		err = pehrest.HashSetValue(req)
		return
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &pehrest.HashSetValueService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler( st)
	if err != nil {
		return
	}
	return st.Err
}
