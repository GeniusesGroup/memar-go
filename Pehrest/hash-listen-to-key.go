/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../achaemenid"
	"../authorization"
	er "../error"
	"../ganjine"
	lang "../language"
	"../srpc"
)

// HashListenToKeyService store details about HashListenToKey service
var HashListenToKeyService = achaemenid.Service{
	ID:                115550110,
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
		lang.LanguageEnglish: "Index Hash - Listen To Key",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: `get records to given index hash when new record set!
Request Must send to specific node that handle that hash index range!!
This service has a lot of use cases like:
- any geospatial usage e.g. tracking device or user, ...
- following content author like telegram channels or instagram live video!`,
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashListenToKeySRPC,
}

// HashListenToKeySRPC is sRPC handler of HashListenToKey service.
func HashListenToKeySRPC(st *achaemenid.Stream) {
	if st.Connection.UserID != achaemenid.Server.AppID {
		// TODO::: Attack??
		st.Err = ganjine.ErrNotAuthorizeRequest
		return
	}

	var req = &HashListenToKeyReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	var res *HashListenToKeyRes
	res, st.Err = HashListenToKey(req)
	if st.Err != nil {
		return
	}

	st.OutcomePayload = res.SyllabEncoder()
}

// HashListenToKeyReq is request structure of HashListenToKey()
type HashListenToKeyReq struct {
	IndexKey       [32]byte
	ReceiveChannel chan []byte `syllab:"-"`
}

// HashListenToKeyRes is response structure of HashListenToKey()
type HashListenToKeyRes struct {
	Record []byte
}

// HashListenToKey get the recordID by index hash when new record set!
func HashListenToKey(req *HashListenToKeyReq) (res *HashListenToKeyRes, err *er.Error) {
	// TODO::: it can't be simple byte, maybe channel
	return
}

/*
	-- Syllab Encoder & Decoder --
*/

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashListenToKeyReq) SyllabDecoder(buf []byte) {
	copy(req.IndexKey[:], buf[:])
	return
}

// SyllabEncoder encode req to buf
func (req *HashListenToKeyReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	copy(buf[4:], req.IndexKey[:])
	return
}

func (req *HashListenToKeyReq) syllabStackLen() (ln uint32) {
	return 32
}

func (req *HashListenToKeyReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashListenToKeyReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}

// SyllabDecoder decode from buf to res
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *HashListenToKeyRes) SyllabDecoder(buf []byte) {
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	res.Record = buf
}

// SyllabEncoder encode req to buf
func (res *HashListenToKeyRes) SyllabEncoder() (buf []byte) {
	buf = make([]byte, res.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Record in buf
	// syllab.SetUInt32(buf, 4, res.syllabStackLen())
	// syllab.SetUInt32(buf, 8, uint32(len(res.Record)))
	copy(buf[4:], res.Record)
	return
}

func (res *HashListenToKeyRes) syllabStackLen() (ln uint32) {
	return 0
}

func (res *HashListenToKeyRes) syllabHeapLen() (ln uint32) {
	ln = uint32(len(res.Record))
	return
}

func (res *HashListenToKeyRes) syllabLen() (ln uint64) {
	return uint64(res.syllabStackLen() + res.syllabHeapLen())
}
