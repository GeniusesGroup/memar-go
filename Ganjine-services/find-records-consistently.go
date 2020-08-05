/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	"unsafe"

	"../achaemenid"
)

var findRecordsConsistentlyService = achaemenid.Service{
	ID:              480215407,
	Name:            "FindRecordsConsistently",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		`Find records by indexes that store before in consistently!
		It will get index from transaction managers not indexes nodes!
		`,
	},
	TAGS:        []string{""},
	SRPCHandler: FindRecordsConsistentlySRPC,
}

// FindRecordsConsistentlySRPC is sRPC handler of FindRecordsConsistently service.
func FindRecordsConsistentlySRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &FindRecordsConsistentlyReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	var res *FindRecordsConsistentlyRes
	res, st.ReqRes.Err = FindRecordsConsistently(req)
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Payload = res.SyllabEncoder()
}

// FindRecordsConsistentlyReq is request structure of FindRecordsConsistently()
type FindRecordsConsistentlyReq struct {
	IndexHash [32]byte
}

// FindRecordsConsistentlyRes is response structure of FindRecordsConsistently()
type FindRecordsConsistentlyRes struct {
	RecordIDs [][16]byte // Max 256 record return
}

// FindRecordsConsistently find records by indexes that store before in consistently!
func FindRecordsConsistently(req *FindRecordsConsistentlyReq) (res *FindRecordsConsistentlyRes, err error) {
	res = &FindRecordsConsistentlyRes{
		// get index from transaction managers not indexes nodes
		RecordIDs: cluster.TransactionManager.GetIndexRecords(req.IndexHash),
	}
	return
}

// SyllabDecoder decode from buf to req
func (req *FindRecordsConsistentlyReq) SyllabDecoder(buf []byte) (err error) {
	copy(req.IndexHash[:], buf[:])
	return
}

// SyllabEncoder encode req to buf
func (req *FindRecordsConsistentlyReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, 36) // 36=4+32 >> first 4+ for sRPC ID instead get offset argument

	copy(buf[4:], req.IndexHash[:])

	return
}

// SyllabDecoder decode from buf to req
func (res *FindRecordsConsistentlyRes) SyllabDecoder(buf []byte) (err error) {
	var sliceLen = uint32(buf[4]) | uint32(buf[5])<<8 | uint32(buf[6])<<16 | uint32(buf[7])<<24
	buf = buf[8:]
	res.RecordIDs = *(*[][16]byte)(unsafe.Pointer(&buf))
	res.RecordIDs = res.RecordIDs[:sliceLen]
	return
}

// SyllabEncoder encode req to buf
func (res *FindRecordsConsistentlyRes) SyllabEncoder() (buf []byte) {
	var ln = len(res.RecordIDs)
	buf = make([]byte, 0, (ln*16)+12) // 12=4+(4+4) >> first 4+ for sRPC ID instead get offset argument

	// encode slice address to the payload buffer.
	// Due to just have one field in res structure we skip set address of res.RecordIDs in buf
	// buf[5] = byte(8)
	// buf[6] = byte(8 >> 8)
	// buf[7] = byte(8 >> 16)
	// buf[8] = byte(8 >> 24)

	// encode slice length to the payload buffer.
	buf[9] = byte(ln)
	buf[10] = byte(ln >> 8)
	buf[11] = byte(ln >> 16)
	buf[12] = byte(ln >> 24)

	for i := 0; i < ln; i++ {
		buf = append(buf, res.RecordIDs[i][:]...)
	}

	return
}
