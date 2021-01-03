/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../achaemenid"
	"../authorization"
	"../convert"
	er "../error"
	"../ganjine"
	lang "../language"
	"../srpc"
	"../syllab"
)

// HashGetValuesService store details about HashGetValues service
var HashGetValuesService = achaemenid.Service{
	ID:                183406116,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // English name of favor service just to show off!
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Authorization: authorization.Service{
		CRUD:     authorization.CRUDRead,
		UserType: authorization.UserTypeApp,
	},

	Name: map[lang.Language]string{
		lang.LanguageEnglish: "Index Hash - Get Values",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: `Get related RecordsID that set to given indexHash before.
Request 32 RecordsID to given IndexKey even if just one of them use!
Suggest not get more than 65536 related RecordID in single request!`,
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashGetValuesSRPC,
}

// HashGetValuesSRPC is sRPC handler of HashGetValues service.
func HashGetValuesSRPC(st *achaemenid.Stream) {
	if st.Connection.UserID != achaemenid.Server.AppID {
		// TODO::: Attack??
		st.Err = ganjine.ErrNotAuthorizeRequest
		return
	}

	var req = &HashGetValuesReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	var res *HashGetValuesRes
	res, st.Err = HashGetValues(req)
	// Check if any error occur in bussiness logic
	if st.Err != nil {
		return
	}

	st.OutcomePayload = res.SyllabEncoder()
}

// HashGetValuesReq is request structure of HashGetValues()
type HashGetValuesReq struct {
	IndexKey [32]byte
	Offset   uint64
	Limit    uint64 // It is better to be modulus of 32||128 if storage devices use 4K clusters!
}

// HashGetValuesRes is response structure of HashGetValues()
type HashGetValuesRes struct {
	IndexValues [][32]byte
}

// HashGetValues returns related IndexValues that set to given indexKey before in local storages.
func HashGetValues(req *HashGetValuesReq) (res *HashGetValuesRes, err *er.Error) {
	var hashIndex = IndexHash{
		RecordID: req.IndexKey,
	}

	res = &HashGetValuesRes{}
	res.IndexValues, err = hashIndex.Get(req.Offset, req.Limit)
	return
}

/*
	-- Syllab Encoder & Decoder --
*/

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashGetValuesReq) SyllabDecoder(buf []byte) {
	copy(req.IndexKey[:], buf[:])
	req.Offset = syllab.GetUInt64(buf, 32)
	req.Limit = syllab.GetUInt64(buf, 40)
}

// SyllabEncoder encode req to buf
func (req *HashGetValuesReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	copy(buf[4:], req.IndexKey[:])
	syllab.SetUInt64(buf, 36, req.Offset)
	syllab.SetUInt64(buf, 44, req.Limit)
	return
}

func (req *HashGetValuesReq) syllabStackLen() (ln uint32) {
	return 48
}

func (req *HashGetValuesReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashGetValuesReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *HashGetValuesRes) SyllabDecoder(buf []byte) {
	buf = buf[8:]
	res.IndexValues = convert.UnsafeByteSliceTo32ByteArraySlice(buf)
}

// SyllabEncoder encode res to buf
func (res *HashGetValuesRes) SyllabEncoder() (buf []byte) {
	buf = make([]byte, res.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	// Due to just have one field in res structure we skip set address of res.IndexValues in buf
	// syllab.SetUInt32(buf, 4, res.syllabStackLen())
	syllab.SetUInt32(buf, 8, uint32(len(res.IndexValues)))
	copy(buf[res.syllabStackLen():], convert.Unsafe32ByteArraySliceToByteSlice(res.IndexValues))
	return
}

func (res *HashGetValuesRes) syllabStackLen() (ln uint32) {
	return 8
}

func (res *HashGetValuesRes) syllabHeapLen() (ln uint32) {
	ln = uint32(len(res.IndexValues) * 32)
	return
}

func (res *HashGetValuesRes) syllabLen() (ln int) {
	return int(res.syllabStackLen() + res.syllabHeapLen())
}
