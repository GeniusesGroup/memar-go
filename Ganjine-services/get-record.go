/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	psdk "../PersiaOS-sdk"
	"../achaemenid"
)

var getRecordService = achaemenid.Service{
	ID:              4052491139,
	Name:            "GetRecord",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		`use to get a record by given ID! It must send to proper node otherwise get not found error!`,
	},
	TAGS:        []string{""},
	SRPCHandler: GetRecordSRPC,
}

// GetRecordSRPC is sRPC handler of GetRecord service.
func GetRecordSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &GetRecordReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	var res *GetRecordRes
	res, st.ReqRes.Err = GetRecord(req)
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Payload = res.SyllabEncoder()
}

// GetRecordReq is request structure of GetRecord()
type GetRecordReq struct {
	RecordID [16]byte
}

// GetRecordRes is response structure of GetRecord()
type GetRecordRes struct {
	Record []byte
}

// GetRecord get the specific record by its ID!
func GetRecord(req *GetRecordReq) (res *GetRecordRes, err error) {
	res = &GetRecordRes{}
	res.Record, err = psdk.GetStorageRecord(req.RecordID)
	return
}

// SyllabDecoder decode from buf to req
func (req *GetRecordReq) SyllabDecoder(buf []byte) (err error) {
	copy(req.RecordID[:], buf[:])
	return
}

// SyllabEncoder encode req to buf
func (req *GetRecordReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, 20) // 20=4+16 >> first 4+ for sRPC ID instead get offset argument

	copy(buf[4:], req.RecordID[:])

	return
}

// SyllabDecoder decode from buf to req
func (res *GetRecordRes) SyllabDecoder(buf []byte) (err error) {
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	res.Record = buf
	return
}

// SyllabEncoder encode req to buf
func (res *GetRecordRes) SyllabEncoder() (buf []byte) {
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
