/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	persiaos "../PersiaOS-sdk"
	"../achaemenid"
)

var writeRecordService = achaemenid.Service{
	ID:              3836795965,
	Name:            "WriteRecord",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		`write some part of a record! Don't use this service until you force to use!
		Recalculate checksum do in database server that is not so efficient!`,
	},
	TAGS:        []string{""},
	SRPCHandler: WriteRecordSRPC,
}

// WriteRecordSRPC is sRPC handler of WriteRecord service.
func WriteRecordSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &WriteRecordReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Err = WriteRecord(req)
}

// WriteRecordReq is request structure of WriteRecord()
type WriteRecordReq struct {
	Type     requestType
	RecordID [32]byte
	Offset   uint64 // start location of write data
	Data     []byte
}

// WriteRecord write some part of a record! Don't use this service until you force to use!
func WriteRecord(req *WriteRecordReq) (err error) {
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

			// Set WriteRecord ServiceID
			reqStream.ServiceID = 3836795965
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
	err = persiaos.WriteStorageRecord(req.RecordID, req.Offset, req.Data)
	return
}

// SyllabDecoder decode from buf to req
func (req *WriteRecordReq) SyllabDecoder(buf []byte) (err error) {
	req.Type = requestType(buf[0])
	copy(req.RecordID[:], buf[1:])
	req.Offset = uint64(buf[17]) | uint64(buf[18])<<8 | uint64(buf[19])<<16 | uint64(buf[20])<<24 |
		uint64(buf[21])<<32 | uint64(buf[22])<<40 | uint64(buf[23])<<48 | uint64(buf[24])<<56
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	req.Data = buf[25:]
	return
}

// SyllabEncoder encode req to buf
func (req *WriteRecordReq) SyllabEncoder() (buf []byte) {
	var ln = len(req.Data)
	buf = make([]byte, ln+29) // 37=4+1+16+8+(4+4) >> first 4+ for sRPC ID instead get offset argument

	buf[4] = byte(req.Type)

	// RecordID
	copy(buf[5:], req.RecordID[:])
	// Offset
	buf[21] = byte(req.Offset)
	buf[22] = byte(req.Offset >> 8)
	buf[23] = byte(req.Offset >> 16)
	buf[24] = byte(req.Offset >> 24)
	buf[25] = byte(req.Offset >> 32)
	buf[26] = byte(req.Offset >> 40)
	buf[27] = byte(req.Offset >> 48)
	buf[28] = byte(req.Offset >> 56)

	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Data in buf
	// buf[29] = byte(41)
	// buf[30] = byte(41 >> 8)
	// buf[31] = byte(41 >> 16)
	// buf[32] = byte(41 >> 24)
	// encode slice length to the payload buffer.
	// buf[33] = byte(ln)
	// buf[34] = byte(ln >> 8)
	// buf[35] = byte(ln >> 16)
	// buf[36] = byte(ln >> 24)

	// Data
	copy(buf[29:], req.Data[:])

	return
}
