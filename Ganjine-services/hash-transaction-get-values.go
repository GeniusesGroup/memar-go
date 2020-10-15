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

// HashTransactionGetValuesService store details about HashTransactionGetValues service
var HashTransactionGetValuesService = achaemenid.Service{
	ID:                2051910519,
	CRUD:              authorization.CRUDRead,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // English name of favor service just to show off!
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "HashTransactionGetValues",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: `Find records by indexes that store before in consistently!
It will get index from transaction managers not indexes nodes!`,
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashTransactionGetValuesSRPC,
}

// HashTransactionGetValuesSRPC is sRPC handler of HashTransactionGetValues service.
func HashTransactionGetValuesSRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var req = &HashTransactionGetValuesReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	var res *HashTransactionGetValuesRes
	res, st.Err = HashTransactionGetValues(req)
	if st.Err != nil {
		return
	}

	st.OutcomePayload = res.SyllabEncoder()
}

// HashTransactionGetValuesReq is request structure of HashTransactionGetValues()
type HashTransactionGetValuesReq struct {
	IndexKey [32]byte
}

// HashTransactionGetValuesRes is response structure of HashTransactionGetValues()
type HashTransactionGetValuesRes struct {
	IndexValues [][32]byte // Max 128 record return
}

// HashTransactionGetValues find records by indexes that store before in consistently!
func HashTransactionGetValues(req *HashTransactionGetValuesReq) (res *HashTransactionGetValuesRes, err error) {
	res = &HashTransactionGetValuesRes{
		// get index from transaction managers not indexes nodes
		IndexValues: cluster.TransactionManager.GetIndexRecords(req.IndexKey),
	}
	return
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashTransactionGetValuesReq) SyllabDecoder(buf []byte) {
	copy(req.IndexKey[:], buf[:])
	return
}

// SyllabEncoder encode req to buf
func (req *HashTransactionGetValuesReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	copy(buf[4:], req.IndexKey[:])
	return
}

func (req *HashTransactionGetValuesReq) syllabStackLen() (ln uint32) {
	return 32
}

func (req *HashTransactionGetValuesReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashTransactionGetValuesReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *HashTransactionGetValuesRes) SyllabDecoder(buf []byte) {
	buf = buf[8:]
	res.IndexValues = convert.UnsafeByteSliceTo32ByteArraySlice(buf)
	return
}

// SyllabEncoder encode req to buf
func (res *HashTransactionGetValuesRes) SyllabEncoder() (buf []byte) {
	buf = make([]byte, res.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	// Due to just have one field in res structure we skip set address of res.IndexValues in buf
	// syllab.SetUInt32(buf, 4, res.syllabStackLen())
	syllab.SetUInt32(buf, 8, uint32(len(res.IndexValues)))
	copy(buf[res.syllabStackLen():], convert.Unsafe32ByteArraySliceToByteSlice(res.IndexValues))
	return
}

func (res *HashTransactionGetValuesRes) syllabStackLen() (ln uint32) {
	return 8
}

func (res *HashTransactionGetValuesRes) syllabHeapLen() (ln uint32) {
	return
}

func (res *HashTransactionGetValuesRes) syllabLen() (ln uint64) {
	return uint64(res.syllabStackLen() + res.syllabHeapLen())
}
