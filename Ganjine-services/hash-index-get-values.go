/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	"../achaemenid"
	"../authorization"
	"../convert"
	"../ganjine"
	lang "../language"
	"../srpc"
	"../syllab"
)

// HashIndexGetValuesService store details about HashIndexGetValues service
var HashIndexGetValuesService = achaemenid.Service{
	ID:                121153927,
	CRUD:              authorization.CRUDRead,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // English name of favor service just to show off!
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "HashIndexGetValues",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: `Get related RecordsID that set to given indexHash before.
Request 32 RecordsID to given IndexKey even if just one of them use!
Suggest not get more than 65536 related RecordID in single request!`,
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashIndexGetValuesSRPC,
}

// HashIndexGetValuesSRPC is sRPC handler of HashIndexGetValues service.
func HashIndexGetValuesSRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var req = &HashIndexGetValuesReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	var res *HashIndexGetValuesRes
	res, st.Err = HashIndexGetValues(req)
	// Check if any error occur in bussiness logic
	if st.Err != nil {
		return
	}

	st.OutcomePayload = res.SyllabEncoder()
}

// HashIndexGetValuesReq is request structure of HashIndexGetValues()
type HashIndexGetValuesReq struct {
	IndexKey [32]byte
	Offset   uint64
	Limit    uint64 // It is better to be modulus of 32||128 if storage devices use 4K clusters!
}

// HashIndexGetValuesRes is response structure of HashIndexGetValues()
type HashIndexGetValuesRes struct {
	IndexValues [][32]byte
}

// HashIndexGetValues returns related IndexValues that set to given indexKey before in local storages.
func HashIndexGetValues(req *HashIndexGetValuesReq) (res *HashIndexGetValuesRes, err error) {
	var hashIndex = ganjine.HashIndex{
		RecordID: req.IndexKey,
	}

	res = &HashIndexGetValuesRes{}
	res.IndexValues, err = hashIndex.Get(req.Offset, req.Limit)
	return
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashIndexGetValuesReq) SyllabDecoder(buf []byte) {
	copy(req.IndexKey[:], buf[:])
	req.Offset = syllab.GetUInt64(buf, 32)
	req.Limit = syllab.GetUInt64(buf, 40)
}

// SyllabEncoder encode req to buf
func (req *HashIndexGetValuesReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	copy(buf[4:], req.IndexKey[:])
	syllab.SetUInt64(buf, 36, req.Offset)
	syllab.SetUInt64(buf, 44, req.Limit)
	return
}

func (req *HashIndexGetValuesReq) syllabStackLen() (ln uint32) {
	return 48
}

func (req *HashIndexGetValuesReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashIndexGetValuesReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *HashIndexGetValuesRes) SyllabDecoder(buf []byte) {
	buf = buf[8:]
	res.IndexValues = convert.UnsafeByteSliceTo32ByteArraySlice(buf)
}

// SyllabEncoder encode res to buf
func (res *HashIndexGetValuesRes) SyllabEncoder() (buf []byte) {
	buf = make([]byte, res.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	// Due to just have one field in res structure we skip set address of res.IndexValues in buf
	// syllab.SetUInt32(buf, 4, res.syllabStackLen())
	syllab.SetUInt32(buf, 8, uint32(len(res.IndexValues)))
	copy(buf[res.syllabStackLen():], convert.Unsafe32ByteArraySliceToByteSlice(res.IndexValues))
	return
}

func (res *HashIndexGetValuesRes) syllabStackLen() (ln uint32) {
	return 8
}

func (res *HashIndexGetValuesRes) syllabHeapLen() (ln uint32) {
	ln = uint32(len(res.IndexValues) * 32)
	return
}

func (res *HashIndexGetValuesRes) syllabLen() (ln int) {
	return int(res.syllabStackLen() + res.syllabHeapLen())
}
