/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	psdk "../PersiaOS-sdk"
	"../achaemenid"
)

var setRecordService = achaemenid.Service{
	ID:              10488062,
	Name:            "SetRecord",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		"Write a whole record or replace old record if it is exist!",
	},
	TAGS:        []string{""},
	SRPCHandler: SetRecordSRPC,
}

// SetRecordSRPC is sRPC handler of SetRecord service.
func SetRecordSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &SetRecordReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Err = SetRecord(req)
}

// SetRecordReq is request structure of SetRecord()
type SetRecordReq struct {
	Type   requestType
	Record []byte
}

// SetRecord respect all data in record and don't change something like RecordID or WriteTime!
// If data like OwnerAppID is wrong you can't get record anymore!
func SetRecord(req *SetRecordReq) (err error) {
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

			// Set SetRecord ServiceID
			reqStream.ServiceID = 10488062
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
	err = psdk.SetStorageRecord(req.Record)
	return
}

// SyllabDecoder decode from buf to req
func (req *SetRecordReq) SyllabDecoder(buf []byte) (err error) {
	req.Type = requestType(buf[0])
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	req.Record = buf[1:]
	return
}

// SyllabEncoder encode req to buf
func (req *SetRecordReq) SyllabEncoder() (buf []byte) {
	var ln = len(req.Record)
	buf = make([]byte, ln+5) // 13=4+1+(4+4) >> first 4+ for sRPC ID instead get offset argument

	buf[4] = byte(req.Type)

	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Record in buf
	// buf[5] = byte(8)
	// buf[6] = byte(8 >> 8)
	// buf[7] = byte(8 >> 16)
	// buf[8] = byte(8 >> 24)
	// encode slice length to the payload buffer.
	// buf[9] = byte(ln)
	// buf[10] = byte(ln >> 8)
	// buf[11] = byte(ln >> 16)
	// buf[12] = byte(ln >> 24)

	copy(buf[5:], req.Record)

	return
}
