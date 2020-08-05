/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	psdk "../PersiaOS-sdk"
	"../achaemenid"
)

var readRecordService = achaemenid.Service{
	ID:              108857663,
	Name:            "ReadRecord",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		`use to read some part of a record! It must send to proper node otherwise get not found error!
		Mostly use to get metadata first to know about record size before get it to split to some shorter part!
		`,
	},
	TAGS:        []string{""},
	SRPCHandler: ReadRecordSRPC,
}

// ReadRecordSRPC is sRPC handler of ReadRecord service.
func ReadRecordSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &ReadRecordReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	var res *ReadRecordRes
	res, st.ReqRes.Err = ReadRecord(req)
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Payload = res.SyllabEncoder()
}

// ReadRecordReq is request structure of ReadRecord()
type ReadRecordReq struct {
	RecordID [16]byte
	Offset   uint64 // Do something like block storage API
	Limit    uint64 // Do something like block storage API
}

// ReadRecordRes is response structure of ReadRecord()
type ReadRecordRes struct {
	Record []byte
}

// ReadRecord read some part of the specific record by its ID!
func ReadRecord(req *ReadRecordReq) (res *ReadRecordRes, err error) {
	res = &ReadRecordRes{}
	res.Record, err = psdk.ReadStorageRecord(req.RecordID, req.Offset, req.Limit)
	return
}

// SyllabDecoder decode from buf to req
func (req *ReadRecordReq) SyllabDecoder(buf []byte) (err error) {
	copy(req.RecordID[:], buf[:])
	req.Offset = uint64(buf[16]) | uint64(buf[17])<<8 | uint64(buf[18])<<16 | uint64(buf[19])<<24 |
		uint64(buf[20])<<32 | uint64(buf[21])<<40 | uint64(buf[22])<<48 | uint64(buf[23])<<56
	req.Limit = uint64(buf[24]) | uint64(buf[25])<<8 | uint64(buf[26])<<16 | uint64(buf[27])<<24 |
		uint64(buf[28])<<32 | uint64(buf[29])<<40 | uint64(buf[30])<<48 | uint64(buf[31])<<56

	return
}

// SyllabEncoder encode req to buf
func (req *ReadRecordReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, 36) // 36=4+16+8+8 >> first 4+ for sRPC ID instead get offset argument

	copy(buf[4:], req.RecordID[:])
	// Offset
	buf[20] = byte(req.Offset)
	buf[21] = byte(req.Offset >> 8)
	buf[22] = byte(req.Offset >> 16)
	buf[23] = byte(req.Offset >> 24)
	buf[24] = byte(req.Offset >> 32)
	buf[25] = byte(req.Offset >> 40)
	buf[26] = byte(req.Offset >> 48)
	buf[27] = byte(req.Offset >> 56)
	// Limit
	buf[28] = byte(req.Limit)
	buf[29] = byte(req.Limit >> 8)
	buf[30] = byte(req.Limit >> 16)
	buf[31] = byte(req.Limit >> 24)
	buf[32] = byte(req.Limit >> 32)
	buf[33] = byte(req.Limit >> 40)
	buf[34] = byte(req.Limit >> 48)
	buf[35] = byte(req.Limit >> 56)

	return
}

// SyllabDecoder decode from buf to req
func (res *ReadRecordRes) SyllabDecoder(buf []byte) (err error) {
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	res.Record = buf
	return
}

// SyllabEncoder encode req to buf
func (res *ReadRecordRes) SyllabEncoder() (buf []byte) {
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
