/* For license and copyright information please see LEGAL file in repository */

package gs

import "../achaemenid"

var finishTransactionService = achaemenid.Service{
	ID:              3962420401,
	Name:            "FinishTransaction",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		`use to approve transaction!
		Transaction Manager will set record and index! no further action need after this call!
		`,
	},
	TAGS:        []string{""},
	SRPCHandler: FinishTransactionSRPC,
}

// FinishTransactionSRPC is sRPC handler of FinishTransaction service.
func FinishTransactionSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &FinishTransactionReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Err = FinishTransaction(req)
}

// FinishTransactionReq is request structure of FinishTransaction()
type FinishTransactionReq struct {
	Type      requestType
	IndexHash [32]byte
	Record    []byte
}

// FinishTransaction approve transaction!
func FinishTransaction(req *FinishTransactionReq) (err error) {
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

			// Set FinishTransaction ServiceID
			reqStream.ServiceID = 3962420401
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
	err = cluster.TransactionManager.FinishTransaction(req.IndexHash, req.Record)
	return
}

// SyllabDecoder decode from buf to req
func (req *FinishTransactionReq) SyllabDecoder(buf []byte) (err error) {
	req.Type = requestType(buf[0])
	copy(req.IndexHash[:], buf[1:])
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	copy(req.Record[:], buf[33:])
	return
}

// SyllabEncoder encode req to buf
func (req *FinishTransactionReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, len(req.Record)+37) // 37=4+1+32+(4+4) >> first 4+ for sRPC ID instead get offset argument

	buf[4] = byte(req.Type)
	copy(buf[5:], req.IndexHash[:])

	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Record in buf
	// buf[6] = byte(8)
	// buf[7] = byte(8 >> 8)
	// buf[8] = byte(8 >> 16)
	// buf[9] = byte(8 >> 24)
	// encode slice length to the payload buffer.
	// buf[10] = byte(ln)
	// buf[11] = byte(ln >> 8)
	// buf[12] = byte(ln >> 16)
	// buf[13] = byte(ln >> 24)

	copy(buf[37:], req.Record[:])

	return
}
