/* For license and copyright information please see LEGAL file in repository */

package gs

import "../achaemenid"

var registerTransactionService = achaemenid.Service{
	ID:              3840530512,
	Name:            "RegisterTransaction",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		`Register new transaction on queue and get last record when transaction ready for this one!
Requester must send FinishTransaction() immediately, otherwise Transaction manager will drop this request from queue and chain!
transaction write can be on secondary indexes not primary indexes, due to primary index must always unique!
transaction manager on any node in a replication must sync with master replication corresponding node manager!
Get a record by ID when record ready to submit! Usually use in transaction queue to act when record ready to read!
Must send this request to specific node that handle that range!!`,
	},
	TAGS:        []string{"transactional authority", "index lock ticket"},
	SRPCHandler: RegisterTransactionSRPC,
}

// RegisterTransactionSRPC is sRPC handler of RegisterTransaction service.
func RegisterTransactionSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &RegisterTransactionReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	var res *RegisterTransactionRes
	res, st.ReqRes.Err = RegisterTransaction(req)
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Payload = res.SyllabEncoder()
}

// RegisterTransactionReq is request structure of RegisterTransaction()
type RegisterTransactionReq struct {
	Type      requestType
	IndexHash [32]byte
	RecordID  [32]byte
}

// RegisterTransactionRes is response structure of RegisterTransaction()
type RegisterTransactionRes struct {
	Record []byte
}

// RegisterTransaction register new transaction on queue and get last record when transaction ready for this one!
func RegisterTransaction(req *RegisterTransactionReq) (res *RegisterTransactionRes, err error) {
	res = &RegisterTransactionRes{}

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

		// Do for i=0 as local node
		res.Record, err = cluster.TransactionManager.RegisterTransaction(req.IndexHash, req.RecordID)
	} else {
		// Don't send last record due to Master node will give it to requester!
		_, err = cluster.TransactionManager.RegisterTransaction(req.IndexHash, req.RecordID)
	}

	return
}

// SyllabDecoder decode from buf to req
func (req *RegisterTransactionReq) SyllabDecoder(buf []byte) (err error) {
	req.Type = requestType(buf[0])
	copy(req.IndexHash[:], buf[1:])
	copy(req.RecordID[:], buf[33:])
	return
}

// SyllabEncoder encode req to buf
func (req *RegisterTransactionReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, 53) // 52=4+1+32+16 >> first 4+ for sRPC ID instead get offset argument

	buf[4] = byte(req.Type)
	copy(buf[5:], req.IndexHash[:])
	copy(buf[37:], req.RecordID[:])

	return
}

// SyllabDecoder decode from buf to req
func (res *RegisterTransactionRes) SyllabDecoder(buf []byte) (err error) {
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	res.Record = buf
	return
}

// SyllabEncoder encode req to buf
func (res *RegisterTransactionRes) SyllabEncoder() (buf []byte) {
	var ln = len(res.Record)
	buf = make([]byte, ln+4) // 12=4+(4+4) >> first 4+ for sRPC ID instead get offset argument

	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Record in buf
	// buf[4] = byte(8)
	// buf[5] = byte(8 >> 8)
	// buf[6] = byte(8 >> 16)
	// buf[7] = byte(8 >> 24)
	// encode slice length to the payload buffer.
	// buf[8] = byte(ln)
	// buf[9] = byte(ln >> 8)
	// buf[10] = byte(ln >> 16)
	// buf[11] = byte(ln >> 24)

	copy(buf[4:], res.Record)

	return
}
