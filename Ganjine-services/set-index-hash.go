/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	"../achaemenid"
	"../ganjine"
)

var setIndexHashService = achaemenid.Service{
	ID:              1881585857,
	Name:            "SetIndex",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		"set a record ID to new||exiting index hash.",
	},
	TAGS:        []string{""},
	SRPCHandler: SetIndexHashSRPC,
}

// SetIndexHashSRPC is sRPC handler of SetIndex service.
func SetIndexHashSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &SetIndexHashReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Err = SetIndexHash(req)
}

// SetIndexHashReq is request structure of SetIndexHash()
type SetIndexHashReq struct {
	Type      requestType
	IndexHash [32]byte
	RecordID  [32]byte
}

// SetIndexHash set a record ID to new||exiting index hash.
func SetIndexHash(req *SetIndexHashReq) (err error) {
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

			// Set SetIndexHash ServiceID
			reqStream.ServiceID = 1881585857
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
	var hashIndex = ganjine.HashIndex{
		RecordID:   req.IndexHash,
		OwnerAppID: server.Manifest.AppID, // TODO::: just need copy if record is not exist and want to create new one
	}
	err = hashIndex.AppendRecordID(req.RecordID)
	return
}

// SyllabDecoder decode from buf to req
func (req *SetIndexHashReq) SyllabDecoder(buf []byte) (err error) {
	req.Type = requestType(buf[0])
	copy(req.IndexHash[:], buf[1:])
	copy(req.RecordID[:], buf[33:])
	return
}

// SyllabEncoder encode req to buf
func (req *SetIndexHashReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, 53) // 53=4+1+32+16 >> first 4+ for sRPC ID instead get offset argument

	buf[4] = byte(req.Type)
	copy(buf[5:], req.IndexHash[:])
	copy(buf[37:], req.RecordID[:])

	return
}
