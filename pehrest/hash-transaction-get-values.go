/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../authorization"
	"../convert"
	"../ganjine"
	"../protocol"
	"../srpc"
	"../syllab"
)

// HashTransactionGetValuesService store details about HashTransactionGetValues service
var HashTransactionGetValuesService = service.Service{
	URN:                "urn:giti:index.protocol:service:hash-transaction-get-values",
	Domain:             DomainName,
	ID:                 6289164385419840057,
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
			Name: "Index Hash - Transaction Get Values",
			Description: `Find records by indexes that store before in consistently!
It will get index from transaction managers not indexes nodes!`,
			TAGS: []string{},
		},
	},

	SRPCHandler: HashTransactionGetValuesSRPC,
}

// HashTransactionGetValues find records by indexes that store before in consistently!
// It will get index from transaction managers not indexes nodes!
func HashTransactionGetValues(req *HashTransactionGetValuesReq) (res *HashTransactionGetValuesRes, err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByStorage(req.MediaTypeID, req.IndexKey)
	if err != nil {
		return
	}

	if node.Node.State == protocol.ApplicationState_LocalNode {
		return HashTransactionGetValues(req)
	}

	var st protocol.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &HashTransactionGetValuesService
	st.OutcomePayload = req.ToSyllab()

	err = node.Conn.Send(st)
	if err != nil {
		return
	}

	res = &HashTransactionGetValuesRes{}
	res.FromSyllab(srpc.GetPayload(st.IncomePayload))
	return
}

// HashTransactionGetValuesSRPC is sRPC handler of HashTransactionGetValues service.
func HashTransactionGetValuesSRPC(st protocol.Stream) {
	if st.Connection.UserID != protocol.OS.AppManifest().AppUUID() {
		// TODO::: Attack??
		err = authorization.ErrUserNotAllow
		return
	}

	var req = &HashTransactionGetValuesReq{}
	req.FromSyllab(srpc.GetPayload(st.IncomePayload))

	var res *HashTransactionGetValuesRes
	res, err = HashTransactionGetValues(req)
	if err != nil {
		return
	}

	st.OutcomePayload = res.ToSyllab()
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
func HashTransactionGetValues(req *HashTransactionGetValuesReq) (res *HashTransactionGetValuesRes, err protocol.Error) {
	res = &HashTransactionGetValuesRes{
		// get index from transaction managers not indexes nodes
		IndexValues: ganjine.Cluster.TransactionManager.GetIndexRecords(req.IndexKey),
	}
	return
}

/*
	-- Syllab Encoder & Decoder --
*/

// FromSyllab decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashTransactionGetValuesReq) FromSyllab(payload []byte, stackIndex uint32) {
	copy(req.IndexKey[:], buf[:])
	return
}

// ToSyllab encode req to buf
func (req *HashTransactionGetValuesReq) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	buf = make([]byte, req.LenAsSyllab()+4) // +4 for sRPC ID instead get offset argument
	copy(buf[4:], req.IndexKey[:])
	return
}

func (req *HashTransactionGetValuesReq) LenOfSyllabStack() uint32 {
	return 32
}

func (req *HashTransactionGetValuesReq) LenOfSyllabHeap() (ln uint32) {
	return
}

func (req *HashTransactionGetValuesReq) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}

// FromSyllab decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *HashTransactionGetValuesRes) FromSyllab(payload []byte, stackIndex uint32) {
	buf = buf[8:]
	res.IndexValues = convert.UnsafeByteSliceTo32ByteArraySlice(buf)
	return
}

// ToSyllab encode req to buf
func (res *HashTransactionGetValuesRes) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	// Due to just have one field in res structure we skip set address of res.IndexValues in buf
	// syllab.SetUInt32(buf, 4, res.LenOfSyllabStack())
	syllab.SetUInt32(buf, 8, uint32(len(res.IndexValues)))
	copy(buf[res.LenOfSyllabStack():], convert.Unsafe32ByteArraySliceToByteSlice(res.IndexValues))
	return
}

func (res *HashTransactionGetValuesRes) LenOfSyllabStack() uint32 {
	return 8
}

func (res *HashTransactionGetValuesRes) LenOfSyllabHeap() (ln uint32) {
	return
}

func (res *HashTransactionGetValuesRes) LenAsSyllab() uint64 {
	return uint64(res.LenOfSyllabStack() + res.LenOfSyllabHeap())
}
