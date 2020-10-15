/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	persiaos "../PersiaOS-sdk"
	"../achaemenid"
	"../ganjine"
	lang "../language"
	"../srpc"
	"../syllab"
)

// ReadRecordService store details about ReadRecord service
var ReadRecordService = achaemenid.Service{
	ID:                108857663,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "",
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "ReadRecord",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: `use to read some part of a record! It must send to proper node otherwise get not found error!
Mostly use to get metadata first to know about record size before get it to split to some shorter part!`,
	},
	TAGS: []string{""},

	SRPCHandler: ReadRecordSRPC,
}

// ReadRecordSRPC is sRPC handler of ReadRecord service.
func ReadRecordSRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var req = &ReadRecordReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	var res *ReadRecordRes
	res, st.Err = ReadRecord(req)
	if st.Err != nil {
		return
	}

	st.OutcomePayload = res.SyllabEncoder()
}

// ReadRecordReq is request structure of ReadRecord()
type ReadRecordReq struct {
	RecordID [32]byte
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
	res.Record, err = persiaos.ReadStorageRecord(req.RecordID, req.Offset, req.Limit)
	return
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *ReadRecordReq) SyllabDecoder(buf []byte) {
	copy(req.RecordID[:], buf[:])
	req.Offset = syllab.GetUInt64(buf, 32)
	req.Limit = syllab.GetUInt64(buf, 40)
	return
}

// SyllabEncoder encode req to buf
func (req *ReadRecordReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	copy(buf[4:], req.RecordID[:])
	syllab.SetUInt64(buf, 36, req.Offset)
	syllab.SetUInt64(buf, 44, req.Limit)
	return
}

func (req *ReadRecordReq) syllabStackLen() (ln uint32) {
	return 48
}

func (req *ReadRecordReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *ReadRecordReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *ReadRecordRes) SyllabDecoder(buf []byte) {
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	res.Record = buf
	return
}

// SyllabEncoder encode req to buf
func (res *ReadRecordRes) SyllabEncoder() (buf []byte) {
	buf = make([]byte, res.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Record in buf
	// syllab.SetUInt32(buf, 4, res.syllabStackLen())
	// syllab.SetUInt32(buf, 8, uint32(len(res.Record)))
	copy(buf[4:], res.Record)
	return
}

func (res *ReadRecordRes) syllabStackLen() (ln uint32) {
	return 0
}

func (res *ReadRecordRes) syllabHeapLen() (ln uint32) {
	ln = uint32(len(res.Record))
	return
}

func (res *ReadRecordRes) syllabLen() (ln uint64) {
	return uint64(res.syllabStackLen() + res.syllabHeapLen())
}
