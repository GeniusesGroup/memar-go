/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../achaemenid"
	"../authorization"
	er "../error"
	"../ganjine"
	lang "../language"
	"../srpc"
	"../syllab"
)

// HashGetValuesNumberService store details about HashGetValuesNumber service
var HashGetValuesNumberService = achaemenid.Service{
	ID:                2503912670,
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
		lang.LanguageEnglish: "Index Hash - Get Values Number",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: "Get number of recordsID register for specific index hash key",
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashGetValuesNumberSRPC,
}

// HashGetValuesNumberSRPC is sRPC handler of HashGetValuesNumber service.
func HashGetValuesNumberSRPC(st *achaemenid.Stream) {
	if st.Connection.UserID != achaemenid.Server.AppID {
		// TODO::: Attack??
		st.Err = ganjine.ErrNotAuthorizeRequest
		return
	}

	var req = &HashGetValuesNumberReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	var res *HashGetValuesNumberRes
	res, st.Err = HashGetValuesNumber(req)
	if st.Err != nil {
		return
	}

	st.OutcomePayload = res.SyllabEncoder()
}

// HashGetValuesNumberReq is request structure of HashGetValuesNumber()
type HashGetValuesNumberReq struct {
	IndexKey [32]byte
}

// HashGetValuesNumberRes is response structure of HashGetValuesNumber()
type HashGetValuesNumberRes struct {
	IndexValuesNumber uint64
}

// HashGetValuesNumber get number of IndexValues register for specific IndexKey
func HashGetValuesNumber(req *HashGetValuesNumberReq) (res *HashGetValuesNumberRes, err *er.Error) {
	var hashIndex = IndexHash{
		RecordID: req.IndexKey,
	}
	err = hashIndex.ReadHeader()
	res = &HashGetValuesNumberRes{
		IndexValuesNumber: hashIndex.IndexValuesNumber,
	}
	return
}

/*
	-- Syllab Encoder & Decoder --
*/

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashGetValuesNumberReq) SyllabDecoder(buf []byte) {
	copy(req.IndexKey[:], buf[:])
	return
}

// SyllabEncoder encode req to buf
func (req *HashGetValuesNumberReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	copy(buf[4:], req.IndexKey[:])
	return
}

func (req *HashGetValuesNumberReq) syllabStackLen() (ln uint32) {
	return 32
}

func (req *HashGetValuesNumberReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashGetValuesNumberReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *HashGetValuesNumberRes) SyllabDecoder(buf []byte) {
	res.IndexValuesNumber = syllab.GetUInt64(buf, 0)
	return
}

// SyllabEncoder encode req to buf
func (res *HashGetValuesNumberRes) SyllabEncoder() (buf []byte) {
	buf = make([]byte, res.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt64(buf, 4, res.IndexValuesNumber)
	return
}

func (res *HashGetValuesNumberRes) syllabStackLen() (ln uint32) {
	return 8
}

func (res *HashGetValuesNumberRes) syllabHeapLen() (ln uint32) {
	return
}

func (res *HashGetValuesNumberRes) syllabLen() (ln uint64) {
	return uint64(res.syllabStackLen() + res.syllabHeapLen())
}
