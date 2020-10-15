/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	persiaos "../PersiaOS-sdk"
	"../achaemenid"
	"../ganjine"
	lang "../language"
	"../srpc"
)

// GetRecordService store details about GetRecord service
var GetRecordService = achaemenid.Service{
	ID:                4052491139,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "",
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "GetRecord",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: `use to get a record by given ID! It must send to proper node otherwise get not found error!`,
	},
	TAGS: []string{""},

	SRPCHandler: GetRecordSRPC,
}

// GetRecordSRPC is sRPC handler of GetRecord service.
func GetRecordSRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var req = &GetRecordReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	var res *GetRecordRes
	res, st.Err = GetRecord(req)
	if st.Err != nil {
		return
	}

	st.OutcomePayload = res.SyllabEncoder()
}

// GetRecordReq is request structure of GetRecord()
type GetRecordReq struct {
	RecordID [32]byte
}

// GetRecordRes is response structure of GetRecord()
type GetRecordRes struct {
	Record []byte
}

// GetRecord get the specific record by its ID!
func GetRecord(req *GetRecordReq) (res *GetRecordRes, err error) {
	res = &GetRecordRes{}
	res.Record, err = persiaos.GetStorageRecord(req.RecordID)
	return
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *GetRecordReq) SyllabDecoder(buf []byte) {
	copy(req.RecordID[:], buf[:])
	return
}

// SyllabEncoder encode req to buf
func (req *GetRecordReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	copy(buf[4:], req.RecordID[:])
	return
}

func (req *GetRecordReq) syllabStackLen() (ln uint32) {
	return 32
}

func (req *GetRecordReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *GetRecordReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *GetRecordRes) SyllabDecoder(buf []byte) {
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	res.Record = buf
	return
}

// SyllabEncoder encode req to buf
func (res *GetRecordRes) SyllabEncoder() (buf []byte) {
	buf = make([]byte, res.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Record in buf
	// syllab.SetUInt32(buf, 4, res.syllabStackLen())
	// syllab.SetUInt32(buf, 8, uint32(len(res.Record)))
	copy(buf[4:], res.Record)
	return
}

func (res *GetRecordRes) syllabStackLen() (ln uint32) {
	return 0
}

func (res *GetRecordRes) syllabHeapLen() (ln uint32) {
	ln = uint32(len(res.Record))
	return
}

func (res *GetRecordRes) syllabLen() (ln uint64) {
	return uint64(res.syllabStackLen() + res.syllabHeapLen())
}
