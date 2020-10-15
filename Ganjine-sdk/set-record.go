/* For license and copyright information please see LEGAL file in repository */

package gsdk

import (
	"../achaemenid"
	"../ganjine"
	gs "../ganjine-services"
)

// SetRecord respect all data in record and don't change something like RecordID or WriteTime!
// If data like OwnerAppID is wrong you can't get record anymore!
func SetRecord(c *ganjine.Cluster, req *gs.SetRecordReq) (err error) {
	var recordID [32]byte
	copy(recordID[:], req.Record[:])
	var node *ganjine.Node = c.GetNodeByRecordID(recordID)
	if node == nil {
		return ganjine.ErrGanjineNoNodeAvailable
	}

	if node.Node.State == achaemenid.NodeStateLocalNode {
		err = gs.SetRecord(req)
		return
	}

	var st *achaemenid.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return err
	}

	st.Service = &gs.SetRecordService
	st.OutcomePayload = req.SyllabEncoder()

	err = achaemenid.SrpcOutcomeRequestHandler(c.Server, st)
	if err != nil {
		return err
	}

	return st.Err
}
