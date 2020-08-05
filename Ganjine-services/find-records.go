/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	"unsafe"

	"../achaemenid"
)

var findRecordsService = achaemenid.Service{
	ID:              1992558377,
	Name:            "FindRecords",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		`Use to find records by indexes that store before!
		Suggest not get more than 65536 related RecordID in single request!
		`,
	},
	TAGS:        []string{""},
	SRPCHandler: FindRecordsSRPC,
}

// FindRecordsSRPC is sRPC handler of FindRecords service.
func FindRecordsSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &FindRecordsReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	var res *FindRecordsRes
	res, st.ReqRes.Err = FindRecords(req)
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Payload = res.SyllabEncoder()
}

// FindRecordsReq is request structure of FindRecords()
type FindRecordsReq struct {
	IndexHash [32]byte
	Offset    uint64
	Limit     uint64 // It is better to be modulus of 64 or 256 if storage devices use 4K clusters!
}

// FindRecordsRes is response structure of FindRecords()
type FindRecordsRes struct {
	RecordIDs [][16]byte
}

// FindRecords get related RecordsID that set to given indexHash before!
func FindRecords(req *FindRecordsReq) (res *FindRecordsRes, err error) {
	res = &FindRecordsRes{
		RecordIDs: cluster.Node.HashIndex.GetIndexRecords(req.IndexHash, req.Offset, req.Limit),
	}
	return
}

// SyllabDecoder decode from buf to req
func (req *FindRecordsReq) SyllabDecoder(buf []byte) (err error) {
	copy(req.IndexHash[:], buf[:])
	req.Offset = uint64(buf[32]) | uint64(buf[33])<<8 | uint64(buf[34])<<16 | uint64(buf[35])<<24 |
		uint64(buf[36])<<32 | uint64(buf[37])<<40 | uint64(buf[38])<<48 | uint64(buf[39])<<56
	req.Limit = uint64(buf[40]) | uint64(buf[41])<<8 | uint64(buf[42])<<16 | uint64(buf[43])<<24 |
		uint64(buf[44])<<32 | uint64(buf[45])<<40 | uint64(buf[46])<<48 | uint64(buf[47])<<56

	return
}

// SyllabEncoder encode req to buf
func (req *FindRecordsReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, 52) // 52=4+32+8+8 >> first 4+ for sRPC ID instead get offset argument

	// Index Hash
	copy(buf[4:], req.IndexHash[:])
	// Offset
	buf[36] = byte(req.Offset)
	buf[37] = byte(req.Offset >> 8)
	buf[38] = byte(req.Offset >> 16)
	buf[39] = byte(req.Offset >> 24)
	buf[40] = byte(req.Offset >> 32)
	buf[41] = byte(req.Offset >> 40)
	buf[42] = byte(req.Offset >> 48)
	buf[43] = byte(req.Offset >> 56)
	// Limit
	buf[44] = byte(req.Limit)
	buf[45] = byte(req.Limit >> 8)
	buf[46] = byte(req.Limit >> 16)
	buf[47] = byte(req.Limit >> 24)
	buf[48] = byte(req.Limit >> 32)
	buf[49] = byte(req.Limit >> 40)
	buf[50] = byte(req.Limit >> 48)
	buf[51] = byte(req.Limit >> 56)

	return
}

// SyllabDecoder decode from buf to req
func (res *FindRecordsRes) SyllabDecoder(buf []byte) (err error) {
	var sliceLen = uint32(buf[4]) | uint32(buf[5])<<8 | uint32(buf[6])<<16 | uint32(buf[7])<<24
	buf = buf[8:]
	res.RecordIDs = *(*[][16]byte)(unsafe.Pointer(&buf))
	res.RecordIDs = res.RecordIDs[:sliceLen]
	return
}

// SyllabEncoder encode req to buf
func (res *FindRecordsRes) SyllabEncoder() (buf []byte) {
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
