/* For license and copyright information please see LEGAL file in repository */

package gs

import "../achaemenid"

var deleteIndexHashService = achaemenid.Service{
	ID:              3411747355,
	Name:            "DeleteIndexHash",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		`Delete just exiting index hash without any related record!
		It wouldn't delete related records! Use DeleteIndexHistory() instead if you want delete all records too!`,
	},
	TAGS:        []string{""},
	SRPCHandler: DeleteIndexHashSRPC,
}

// DeleteIndexHashSRPC is sRPC handler of DeleteIndexHash service.
func DeleteIndexHashSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &DeleteIndexHashReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Err = DeleteIndexHash(req)
}

// DeleteIndexHashReq is request structure of DeleteIndexHash()
type DeleteIndexHashReq struct {
	Type      requestType
	IndexHash [32]byte
}

// DeleteIndexHash delete just exiting index hash without any related record.
func DeleteIndexHash(req *DeleteIndexHashReq) (err error) {
	if req.Type == RequestTypeBroadcast {
		// tell other node that this node handle request and don't send this request to other nodes!
		req.Type = RequestTypeStandalone
		var reqEncoded = req.SyllabEncoder()

		// send request to other related nodes
		var i uint8
		for i = 1; i < cluster.TotalReplications; i++ {
			// Make new request-response streams
			var reqStream, resStream *achaemenid.Stream
			reqStream, resStream, err = cluster.Replications[i].Nodes[cluster.Node.ID].Conn.MakeBidirectionalStream(0)
			if err != nil {
				// TODO::: Can we easily return error if two nodes did their job and not have enough resource to send request to final node??
				return
			}

			// Set DeleteIndexHash ServiceID
			reqStream.ServiceID = 3411747355
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
	cluster.Node.HashIndex.DeleteIndexHash(req.IndexHash)
	return
}

// SyllabDecoder decode from buf to req
func (req *DeleteIndexHashReq) SyllabDecoder(buf []byte) (err error) {
	req.Type = requestType(buf[0])
	copy(req.IndexHash[:], buf[1:])
	return
}

// SyllabEncoder encode req to buf
func (req *DeleteIndexHashReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, 37) // 37=4+1+32 >> first 4+ for sRPC ID instead get offset argument

	buf[4] = byte(req.Type)
	copy(buf[5:], req.IndexHash[:])

	return
}
