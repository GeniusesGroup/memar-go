/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	"../achaemenid"
	"../authorization"
	"../ganjine"
	lang "../language"
	"../srpc"
)

// HashIndexListenToKeyService store details about HashIndexListenToKey service
var HashIndexListenToKeyService = achaemenid.Service{
	ID:                1301212853,
	CRUD:              authorization.CRUDRead,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // English name of favor service just to show off!
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "HashIndexListenToKey",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: `get records to given index hash when new record set!
Request Must send to specific node that handle that hash index range!!
This service has a lot of use cases like:
- any geospatial usage e.g. tracking device or user, ...
- following content author like telegram channels or instagram live video!`,
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashIndexListenToKeySRPC,
}

// HashIndexListenToKeySRPC is sRPC handler of HashIndexListenToKey service.
func HashIndexListenToKeySRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var req = &HashIndexListenToKeyReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	var res *HashIndexListenToKeyRes
	res, st.Err = HashIndexListenToKey(req)
	if st.Err != nil {
		return
	}

	st.OutcomePayload = res.SyllabEncoder()
}

// HashIndexListenToKeyReq is request structure of HashIndexListenToKey()
type HashIndexListenToKeyReq struct {
	IndexKey       [32]byte
	ReceiveChannel chan []byte `syllab:"-"`
}

// HashIndexListenToKeyRes is response structure of HashIndexListenToKey()
type HashIndexListenToKeyRes struct {
	Record []byte
}

// HashIndexListenToKey get the recordID by index hash when new record set!
func HashIndexListenToKey(req *HashIndexListenToKeyReq) (res *HashIndexListenToKeyRes, err error) {
	// TODO::: it can't be simple byte, maybe channel
	return
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashIndexListenToKeyReq) SyllabDecoder(buf []byte) {
	copy(req.IndexKey[:], buf[:])
	return
}

// SyllabEncoder encode req to buf
func (req *HashIndexListenToKeyReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	copy(buf[4:], req.IndexKey[:])
	return
}

func (req *HashIndexListenToKeyReq) syllabStackLen() (ln uint32) {
	return 32
}

func (req *HashIndexListenToKeyReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashIndexListenToKeyReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}

// SyllabDecoder decode from buf to res
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *HashIndexListenToKeyRes) SyllabDecoder(buf []byte) {
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	res.Record = buf
}

// SyllabEncoder encode req to buf
func (res *HashIndexListenToKeyRes) SyllabEncoder() (buf []byte) {
	buf = make([]byte, res.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Record in buf
	// syllab.SetUInt32(buf, 4, res.syllabStackLen())
	// syllab.SetUInt32(buf, 8, uint32(len(res.Record)))
	copy(buf[4:], res.Record)
	return
}

func (res *HashIndexListenToKeyRes) syllabStackLen() (ln uint32) {
	return 0
}

func (res *HashIndexListenToKeyRes) syllabHeapLen() (ln uint32) {
	ln = uint32(len(res.Record))
	return
}

func (res *HashIndexListenToKeyRes) syllabLen() (ln uint64) {
	return uint64(res.syllabStackLen() + res.syllabHeapLen())
}
