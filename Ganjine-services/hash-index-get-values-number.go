/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	"../achaemenid"
	"../authorization"
	"../ganjine"
	lang "../language"
	"../srpc"
	"../syllab"
)

// HashIndexGetValuesNumberService store details about HashIndexGetValuesNumber service
var HashIndexGetValuesNumberService = achaemenid.Service{
	ID:                3748529454,
	CRUD:              authorization.CRUDRead,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // English name of favor service just to show off!
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "HashIndexGetValuesNumber",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: "Get number of recordsID register for specific IndexKey",
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashIndexGetValuesNumberSRPC,
}

// HashIndexGetValuesNumberSRPC is sRPC handler of HashIndexGetValuesNumber service.
func HashIndexGetValuesNumberSRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var req = &HashIndexGetValuesNumberReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	var res *HashIndexGetValuesNumberRes
	res, st.Err = HashIndexGetValuesNumber(req)
	if st.Err != nil {
		return
	}

	st.OutcomePayload = res.SyllabEncoder()
}

// HashIndexGetValuesNumberReq is request structure of HashIndexGetValuesNumber()
type HashIndexGetValuesNumberReq struct {
	IndexKey [32]byte
}

// HashIndexGetValuesNumberRes is response structure of HashIndexGetValuesNumber()
type HashIndexGetValuesNumberRes struct {
	IndexValuesNumber uint64
}

// HashIndexGetValuesNumber get number of IndexValues register for specific IndexKey
func HashIndexGetValuesNumber(req *HashIndexGetValuesNumberReq) (res *HashIndexGetValuesNumberRes, err error) {
	var hashIndex = ganjine.HashIndex{
		RecordID: req.IndexKey,
	}
	err = hashIndex.ReadHeader()
	res = &HashIndexGetValuesNumberRes{}
	res.IndexValuesNumber = hashIndex.IndexValuesNumber
	return
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashIndexGetValuesNumberReq) SyllabDecoder(buf []byte) {
	copy(req.IndexKey[:], buf[:])
	return
}

// SyllabEncoder encode req to buf
func (req *HashIndexGetValuesNumberReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	copy(buf[4:], req.IndexKey[:])
	return
}

func (req *HashIndexGetValuesNumberReq) syllabStackLen() (ln uint32) {
	return 32
}

func (req *HashIndexGetValuesNumberReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashIndexGetValuesNumberReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *HashIndexGetValuesNumberRes) SyllabDecoder(buf []byte) {
	res.IndexValuesNumber = syllab.GetUInt64(buf, 0)
	return
}

// SyllabEncoder encode req to buf
func (res *HashIndexGetValuesNumberRes) SyllabEncoder() (buf []byte) {
	buf = make([]byte, res.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt64(buf, 4, res.IndexValuesNumber)
	return
}

func (res *HashIndexGetValuesNumberRes) syllabStackLen() (ln uint32) {
	return 8
}

func (res *HashIndexGetValuesNumberRes) syllabHeapLen() (ln uint32) {
	return
}

func (res *HashIndexGetValuesNumberRes) syllabLen() (ln uint64) {
	return uint64(res.syllabStackLen() + res.syllabHeapLen())
}
