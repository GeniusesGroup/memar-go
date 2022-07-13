/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../authorization"
	"../convert"
	"../protocol"
	"../srpc"
	"../syllab"
)

// HashGetValuesService store details about HashGetValues service
var HashGetValuesService = service.Service{
	URN:                "urn:giti:index.protocol:service:hash-get-values",
	Domain:             DomainName,
	ID:                 6441982709696635966,
	IssueDate:          1587282740,
	ExpiryDate:         0,
	ExpireInFavorOfURN: "",
	ExpireInFavorOfID:  0,
	Status:             protocol.Software_PreAlpha,

	Authorization: authorization.Service{
		CRUD:     authorization.CRUDRead,
		UserType: protocol.UserType_App,
	},

	Detail: map[protocol.LanguageID]service.ServiceDetail{
		protocol.LanguageEnglish: {
			Name: "Index Hash - Get Values",
			Description: `Get related RecordsID that set to given indexHash before.
Request 32 RecordsID to given IndexKey even if just one of them use!
Suggest not get more than 65536 related RecordID in single request!`,
			TAGS: []string{},
		},
	},

	SRPCHandler: HashGetValuesSRPC,
}

// HashGetValues get related RecordsID that set to given IndexKey before.
func HashGetValues(req HashGetValuesReq) (res HashGetValuesRes, err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByStorage(req.MediaTypeID, req.IndexKey)
	if err != nil {
		return
	}

	if node.Node.State == protocol.ApplicationState_LocalNode {
		res, err = HashGetValues(req)
		return
	}

	var st protocol.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &HashGetValuesService
	st.OutcomePayload = req.ToSyllab()

	err = node.Conn.Send(st)
	if err != nil {
		return
	}

	res = &HashGetValuesRes{}
	res.FromSyllab(srpc.GetPayload(st.IncomePayload))
	return
}

// HashGetValuesSRPC is sRPC handler of HashGetValues service.
func HashGetValuesSRPC(st protocol.Stream) {
	if st.Connection.UserID != protocol.OS.AppManifest().AppUUID() {
		// TODO::: Attack??
		err = authorization.ErrUserNotAllow
		return
	}

	var req = &HashGetValuesReq{}
	req.FromSyllab(srpc.GetPayload(st.IncomePayload))

	var res *HashGetValuesRes
	res, err = HashGetValues(req)
	// Check if any error occur in bussiness logic
	if err != nil {
		return
	}

	st.OutcomePayload = res.ToSyllab()
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
func HashGetValues(req *HashGetValuesReq) (res *HashGetValuesRes, err protocol.Error) {
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

// FromSyllab decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashGetValuesReq) FromSyllab(payload []byte, stackIndex uint32) {
	copy(req.IndexKey[:], buf[:])
	req.Offset = syllab.GetUInt64(buf, 32)
	req.Limit = syllab.GetUInt64(buf, 40)
}

// ToSyllab encode req to buf
func (req *HashGetValuesReq) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	buf = make([]byte, req.LenAsSyllab()+4) // +4 for sRPC ID instead get offset argument
	copy(buf[4:], req.IndexKey[:])
	syllab.SetUInt64(buf, 36, req.Offset)
	syllab.SetUInt64(buf, 44, req.Limit)
	return
}

func (req *HashGetValuesReq) LenOfSyllabStack() uint32 {
	return 48
}

func (req *HashGetValuesReq) LenOfSyllabHeap() (ln uint32) {
	return
}

func (req *HashGetValuesReq) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}

// FromSyllab decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *HashGetValuesRes) FromSyllab(payload []byte, stackIndex uint32) {
	buf = buf[8:]
	res.IndexValues = convert.UnsafeByteSliceTo32ByteArraySlice(buf)
}

// ToSyllab encode res to buf
func (res *HashGetValuesRes) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	// Due to just have one field in res structure we skip set address of res.IndexValues in buf
	// syllab.SetUInt32(buf, 4, res.LenOfSyllabStack())
	syllab.SetUInt32(buf, 8, uint32(len(res.IndexValues)))
	copy(buf[res.LenOfSyllabStack():], convert.Unsafe32ByteArraySliceToByteSlice(res.IndexValues))
	return
}

func (res *HashGetValuesRes) LenOfSyllabStack() uint32 {
	return 8
}

func (res *HashGetValuesRes) LenOfSyllabHeap() (ln uint32) {
	ln = uint32(len(res.IndexValues) * 32)
	return
}

func (res *HashGetValuesRes) LenAsSyllab() uint64 {
	return uint64(res.LenOfSyllabStack() + res.LenOfSyllabHeap())
}
