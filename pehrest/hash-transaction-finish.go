/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../authorization"
	"../ganjine"
	"../protocol"
	"../srpc"
	"../syllab"
)

// HashTransactionFinishService store details about HashTransactionFinish service
var HashTransactionFinishService = service.Service{
	URN:                "urn:giti:index.protocol:service:hash-transaction-finish",
	Domain:             DomainName,
	ID:                 15819904868420503209,
	IssueDate:          1587282740,
	ExpiryDate:         0,
	ExpireInFavorOfURN: "",
	ExpireInFavorOfID:  0,
	Status:             protocol.Software_PreAlpha,

	Authorization: authorization.Service{
		CRUD:     authorization.CRUDUpdate,
		UserType: protocol.UserType_App,
	},

	Detail: map[protocol.LanguageID]service.ServiceDetail{
		protocol.LanguageEnglish: {
			Name: "Index Hash - Transaction Finish",
			Description: `use to approve transaction!
Transaction Manager will set record and index! no further action need after this call!`,
			TAGS: []string{},
		},
	},

	SRPCHandler: HashTransactionFinishSRPC,
}

// HashTransactionFinish approve transaction!
// Transaction Manager will set record and index! no further action need after this call!
func HashTransactionFinish(req *HashTransactionFinishReq) (err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByStorage(req.MediaTypeID, req.IndexKey)
	if err != nil {
		return
	}

	if node.Node.State == protocol.ApplicationState_LocalNode {
		return HashTransactionFinish(req)
	}

	var st protocol.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &HashTransactionFinishService
	st.OutcomePayload = req.ToSyllab()

	err = node.Conn.Send(st)
	return
}

// HashTransactionFinishSRPC is sRPC handler of HashTransactionFinish service.
func HashTransactionFinishSRPC(st protocol.Stream) {
	if st.Connection.UserID != protocol.OS.AppManifest().AppUUID() {
		// TODO::: Attack??
		err = authorization.ErrUserNotAllow
		return
	}

	var req = &HashTransactionFinishReq{}
	req.FromSyllab(srpc.GetPayload(st.IncomePayload))

	err = HashTransactionFinish(req)
	st.OutcomePayload = make([]byte, srpc.MinLength)
}

// HashTransactionFinishReq is request structure of HashTransactionFinish()
type HashTransactionFinishReq struct {
	Type     ganjine.RequestType
	IndexKey [32]byte
	Record   []byte
}

// HashTransactionFinish approve transaction!
func HashTransactionFinish(req *HashTransactionFinishReq) (err protocol.Error) {
	if req.Type == ganjine.RequestTypeBroadcast {
		// tell other node that this node handle request and don't send this request to other nodes!
		req.Type = ganjine.RequestTypeStandalone
		var reqEncoded = req.ToSyllab()

		// send request to other related nodes
		for i := 1; i < len(ganjine.Cluster.Replications.Zones); i++ {
			var conn = ganjine.Cluster.Replications.Zones[i].Nodes[ganjine.Cluster.Node.ID].Conn
			var st protocol.Stream
			st, err = conn.MakeOutcomeStream(0)
			if err != nil {
				// TODO::: Can we easily return error if two nodes did their job and not have enough resource to send request to final node??
				return
			}

			st.Service = &service.Service{
				URN:    "urn:giti:index.protocol:service:---",
				Domain: DomainName, ID: 15819904868420503209}
			st.OutcomePayload = reqEncoded

			err = conn.Send(st)
			if err != nil {
				// TODO::: Can we easily return error if two nodes do their job and just one node connection lost??
				return
			}

			// TODO::: Can we easily return response error without handle some known situations??
			err = err
		}
	}

	// Do for i=0 as local node
	err = ganjine.Cluster.TransactionManager.FinishTransaction(req.IndexKey, req.Record)
	return
}

/*
	-- Syllab Encoder & Decoder --
*/

// FromSyllab decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashTransactionFinishReq) FromSyllab(payload []byte, stackIndex uint32) {
	req.Type = ganjine.RequestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	req.Record = buf[33:]
	return
}

// ToSyllab encode req to buf
func (req *HashTransactionFinishReq) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	buf = make([]byte, req.LenAsSyllab()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])
	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Record in buf
	// syllab.SetUInt32(buf, 37, res.LenOfSyllabStack())
	// syllab.SetUInt32(buf, 45, uint32(len(res.Record)))
	copy(buf[37:], req.Record[:])
	return
}

func (req *HashTransactionFinishReq) LenOfSyllabStack() uint32 {
	return 33
}

func (req *HashTransactionFinishReq) LenOfSyllabHeap() (ln uint32) {
	ln = uint32(len(req.Record))
	return
}

func (req *HashTransactionFinishReq) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}
