/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	persiaos "../PersiaOS-sdk"
	"../achaemenid"
)

var deleteRecordService = achaemenid.Service{
	ID:              1758631843,
	Name:            "DeleteRecord",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		`Delete specific record by given ID in all cluster!
		We don't suggest use this service, due to we strongly suggest think about data as immutable entity(stream and time)
		It won't delete record history or indexes associate to it!`,
	},
	TAGS:        []string{""},
	SRPCHandler: DeleteRecordSRPC,
}

// DeleteRecordSRPC is sRPC handler of DeleteRecord service.
func DeleteRecordSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &DeleteRecordReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Err = DeleteRecord(req)
}

// DeleteRecordReq is request structure of DeleteRecord()
type DeleteRecordReq struct {
	Type     requestType
	RecordID [32]byte
}

// DeleteRecord delete specific record by given ID in all cluster!
func DeleteRecord(req *DeleteRecordReq) (err error) {
	if req.Type == RequestTypeBroadcast {
		// tell other node that this node handle request and don't send this request to other nodes!
		req.Type = RequestTypeStandalone
		var reqEncoded = req.SyllabEncoder()

		// send request to other related nodes
		var i uint8
		for i = 1; i < cluster.Replications.TotalZones; i++ {
			// Make new request-response streams
			var reqStream, resStream *achaemenid.Stream
			reqStream, resStream, err = cluster.Replications.Zones[i].Nodes[cluster.Node.ID].Conn.MakeBidirectionalStream(0)
			if err != nil {
				// TODO::: Can we easily return error if two nodes did their job and not have enough resource to send request to final node??
				return
			}

			// Set DeleteRecord ServiceID
			reqStream.ServiceID = 1758631843
			reqStream.Payload = reqEncoded

			err = achaemenid.SrpcOutcomeRequestHandler(server, reqStream)
			if err != nil {
				// TODO::: Can we easily return error if two nodes do their job and just one node connection lost??
				return
			}

			// TODO::: Can we easily return response error without handle some known situations??
			err = resStream.Err
		}
	}

	// Do for i=0 as local node
	err = persiaos.DeleteStorageRecord(req.RecordID)
	return
}

// SyllabDecoder decode from buf to req
func (req *DeleteRecordReq) SyllabDecoder(buf []byte) (err error) {
	req.Type = requestType(buf[0])
	copy(req.RecordID[:], buf[1:])
	return
}

// SyllabEncoder encode req to buf
func (req *DeleteRecordReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, 21) // 21=4+1+16 >> first 4+ for sRPC ID instead get offset argument

	buf[4] = byte(req.Type)
	copy(buf[5:], req.RecordID[:])

	return
}
