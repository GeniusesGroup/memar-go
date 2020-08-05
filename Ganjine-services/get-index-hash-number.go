/* For license and copyright information please see LEGAL file in repository */

package gs

import "../achaemenid"

var getIndexHashNumberService = achaemenid.Service{
	ID:              222077451,
	Name:            "GetIndexHashNumber",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		"Get number of recordsID register for specific IndexHash",
	},
	TAGS:        []string{""},
	SRPCHandler: GetIndexHashNumberSRPC,
}

// GetIndexHashNumberSRPC is sRPC handler of GetIndexHashNumber service.
func GetIndexHashNumberSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &GetIndexHashNumberReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	var res *GetIndexHashNumberRes
	res, st.ReqRes.Err = GetIndexHashNumber(req)
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Payload = res.SyllabEncoder()
}

// GetIndexHashNumberReq is request structure of GetIndexHashNumber()
type GetIndexHashNumberReq struct {
	IndexHash [32]byte
}

// GetIndexHashNumberRes is response structure of GetIndexHashNumber()
type GetIndexHashNumberRes struct {
	RecordNumber uint64
}

// GetIndexHashNumber get number of recordsID register for specific IndexHash
func GetIndexHashNumber(req *GetIndexHashNumberReq) (res *GetIndexHashNumberRes, err error) {
	res = &GetIndexHashNumberRes{}
	_, res.RecordNumber = cluster.Node.HashIndex.GetRecordsIDBucketInfo(req.IndexHash)
	return
}

// SyllabDecoder decode from buf to req
func (req *GetIndexHashNumberReq) SyllabDecoder(buf []byte) (err error) {
	copy(req.IndexHash[:], buf[:])
	return
}

// SyllabEncoder encode req to buf
func (req *GetIndexHashNumberReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, 36) // 36=4+32 >> first 4+ for sRPC ID instead get offset argument

	// Index Hash
	copy(buf[4:], req.IndexHash[:])

	return
}

// SyllabDecoder decode from buf to req
func (res *GetIndexHashNumberRes) SyllabDecoder(buf []byte) (err error) {
	res.RecordNumber = uint64(buf[0]) | uint64(buf[1])<<8 | uint64(buf[2])<<16 | uint64(buf[3])<<24 |
		uint64(buf[4])<<32 | uint64(buf[5])<<40 | uint64(buf[6])<<48 | uint64(buf[7])<<56
	return
}

// SyllabEncoder encode req to buf
func (res *GetIndexHashNumberRes) SyllabEncoder() (buf []byte) {
	buf = make([]byte, 12) // 12=4+8 >> first 4+ for sRPC ID instead get offset argument

	buf[4] = byte(res.RecordNumber)
	buf[5] = byte(res.RecordNumber >> 8)
	buf[6] = byte(res.RecordNumber >> 16)
	buf[7] = byte(res.RecordNumber >> 24)
	buf[8] = byte(res.RecordNumber >> 32)
	buf[9] = byte(res.RecordNumber >> 40)
	buf[10] = byte(res.RecordNumber >> 48)
	buf[11] = byte(res.RecordNumber >> 56)
	return
}
