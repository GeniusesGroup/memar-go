/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../authorization"
	"../protocol"
	"../srpc"
	"../syllab"
)

// HashGetValuesNumberService store details about HashGetValuesNumber service
var HashGetValuesNumberService = service.Service{
	URN:                "urn:giti:index.protocol:service:hash-get-values-number",
	Domain:             DomainName,
	ID:                 8061689463948451244,
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
			Name:        "Index Hash - Get Values Number",
			Description: "Get number of recordsID register for specific index hash key",
			TAGS:        []string{},
		},
	},

	SRPCHandler: HashGetValuesNumberSRPC,
}

// HashGetValuesNumber get number of recordsID register for specific IndexHash
func HashGetValuesNumber(req *HashGetValuesNumberReq) (res *HashGetValuesNumberRes, err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByStorage(req.MediaTypeID, req.IndexKey)
	if err != nil {
		return
	}

	if node.Node.State == protocol.ApplicationState_LocalNode {
		return HashGetValuesNumber(req)
	}

	var st protocol.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &HashGetValuesNumberService
	st.OutcomePayload = req.ToSyllab()

	err = node.Conn.Send(st)
	if err != nil {
		return
	}

	res = &HashGetValuesNumberRes{}
	res.FromSyllab(srpc.GetPayload(st.IncomePayload))
	return
}

// HashGetValuesNumberSRPC is sRPC handler of HashGetValuesNumber service.
func HashGetValuesNumberSRPC(st protocol.Stream) {
	if st.Connection.UserID != protocol.OS.AppManifest().AppUUID() {
		// TODO::: Attack??
		err = authorization.ErrUserNotAllow
		return
	}

	var req = &HashGetValuesNumberReq{}
	req.FromSyllab(srpc.GetPayload(st.IncomePayload))

	var res *HashGetValuesNumberRes
	res, err = HashGetValuesNumber(req)
	if err != nil {
		return
	}

	st.OutcomePayload = res.ToSyllab()
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
func HashGetValuesNumber(req *HashGetValuesNumberReq) (res *HashGetValuesNumberRes, err protocol.Error) {
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

// FromSyllab decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashGetValuesNumberReq) FromSyllab(payload []byte, stackIndex uint32) {
	copy(req.IndexKey[:], buf[:])
	return
}

// ToSyllab encode req to buf
func (req *HashGetValuesNumberReq) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	buf = make([]byte, req.LenAsSyllab()+4) // +4 for sRPC ID instead get offset argument
	copy(buf[4:], req.IndexKey[:])
	return
}

func (req *HashGetValuesNumberReq) LenOfSyllabStack() uint32 {
	return 32
}

func (req *HashGetValuesNumberReq) LenOfSyllabHeap() (ln uint32) {
	return
}

func (req *HashGetValuesNumberReq) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}

// FromSyllab decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *HashGetValuesNumberRes) FromSyllab(payload []byte, stackIndex uint32) {
	res.IndexValuesNumber = syllab.GetUInt64(buf, 0)
	return
}

// ToSyllab encode req to buf
func (res *HashGetValuesNumberRes) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	syllab.SetUInt64(buf, 4, res.IndexValuesNumber)
	return
}

func (res *HashGetValuesNumberRes) LenOfSyllabStack() uint32 {
	return 8
}

func (res *HashGetValuesNumberRes) LenOfSyllabHeap() (ln uint32) {
	return
}

func (res *HashGetValuesNumberRes) LenAsSyllab() uint64 {
	return uint64(res.LenOfSyllabStack() + res.LenOfSyllabHeap())
}
