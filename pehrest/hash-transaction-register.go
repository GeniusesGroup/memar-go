/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../authorization"
	"../ganjine"
	"../protocol"
	"../srpc"
	"../syllab"
)

// HashTransactionRegisterService store details about HashTransactionRegister service
var HashTransactionRegisterService = service.Service{
	URN:                "urn:giti:index.protocol:service:hash-transaction-register",
	Domain:             DomainName,
	ID:                 8526423448282596109,
	IssueDate:          1587282740,
	ExpiryDate:         0,
	ExpireInFavorOfURN: "",
	ExpireInFavorOfID:  0,
	Status:             protocol.Software_PreAlpha,

	Authorization: authorization.Service{
		CRUD:     authorization.CRUDCreate,
		UserType: protocol.UserType_App,
	},

	Detail: map[protocol.LanguageID]service.ServiceDetail{
		protocol.LanguageEnglish: {
			Name: "Index Hash - Transaction Register",
			Description: `Register new transaction on queue and get last record when transaction ready for this one!
Requester must send FinishTransaction() immediately, otherwise Transaction manager will drop this request from queue and chain!
transaction write can be on secondary indexes not primary indexes, due to primary index must always unique!
transaction manager on any node in a replication must sync with master replication corresponding node manager!
Get a record by ID when record ready to submit! Usually use in transaction queue to act when record ready to read!
Must send this request to specific node that handle that range!!`,
			TAGS: []string{"transactional authority", "index lock ticket"},
		},
	},

	SRPCHandler: HashTransactionRegisterSRPC,
}

// HashTransactionRegister register new transaction on queue and get last record when transaction ready for this one!
func HashTransactionRegister(req *HashTransactionRegisterReq) (res *HashTransactionRegisterRes, err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByStorage(req.MediaTypeID, req.IndexKey)
	if err != nil {
		return
	}

	if node.Node.State == protocol.ApplicationState_LocalNode {
		res, err = HashTransactionRegister(req)
		return
	}

	var st protocol.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &HashTransactionRegisterService
	st.OutcomePayload = req.ToSyllab()

	err = node.Conn.Send(st)
	if err != nil {
		return
	}

	res = &HashTransactionRegisterRes{}
	res.FromSyllab(srpc.GetPayload(st.IncomePayload))
	return
}

// HashTransactionRegisterSRPC is sRPC handler of HashTransactionRegister service.
func HashTransactionRegisterSRPC(st protocol.Stream) {
	if st.Connection.UserID != protocol.OS.AppManifest().AppUUID() {
		// TODO::: Attack??
		err = authorization.ErrUserNotAllow
		return
	}

	var req = &HashTransactionRegisterReq{}
	req.FromSyllab(srpc.GetPayload(st.IncomePayload))

	var res *HashTransactionRegisterRes
	res, err = HashTransactionRegister(req)
	if err != nil {
		return
	}

	st.OutcomePayload = res.ToSyllab()
}

// HashTransactionRegisterReq is request structure of HashTransactionRegister()
type HashTransactionRegisterReq struct {
	Type     ganjine.RequestType
	IndexKey [32]byte
	RecordID [32]byte
}

// HashTransactionRegisterRes is response structure of HashTransactionRegister()
type HashTransactionRegisterRes struct {
	Record []byte
}

// HashTransactionRegister register new transaction on queue and get last record when transaction ready for this one!
func HashTransactionRegister(req *HashTransactionRegisterReq) (res *HashTransactionRegisterRes, err protocol.Error) {
	res = &HashTransactionRegisterRes{}

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

			st.Service = &service.Service{ID: 8526423448282596109}
			st.OutcomePayload = reqEncoded

			err = conn.Send(st)
			if err != nil {
				// TODO::: Can we easily return error if two nodes do their job and just one node connection lost??
				return
			}

			// TODO::: Can we easily return response error without handle some known situations??
			err = err
		}

		// Do for i=0 as local node
		res.Record, err = ganjine.Cluster.TransactionManager.RegisterTransaction(req.IndexKey, req.RecordID)
	} else {
		// Don't send last record due to Master node will give it to requester!
		_, err = ganjine.Cluster.TransactionManager.RegisterTransaction(req.IndexKey, req.RecordID)
	}

	return
}

/*
	-- Syllab Encoder & Decoder --
*/

// FromSyllab decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashTransactionRegisterReq) FromSyllab(payload []byte, stackIndex uint32) {
	req.Type = ganjine.RequestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	copy(req.RecordID[:], buf[33:])
	return
}

// ToSyllab encode req to buf
func (req *HashTransactionRegisterReq) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	buf = make([]byte, req.LenAsSyllab()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])
	copy(buf[37:], req.RecordID[:])
	return
}

func (req *HashTransactionRegisterReq) LenOfSyllabStack() uint32 {
	return 65
}

func (req *HashTransactionRegisterReq) LenOfSyllabHeap() (ln uint32) {
	return
}

func (req *HashTransactionRegisterReq) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}

// FromSyllab decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *HashTransactionRegisterRes) FromSyllab(payload []byte, stackIndex uint32) {
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	res.Record = buf
	return
}

// ToSyllab encode req to buf
func (res *HashTransactionRegisterRes) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Record in buf
	// syllab.SetUInt32(buf, 4, res.LenOfSyllabStack())
	// syllab.SetUInt32(buf, 8, uint32(len(res.Record)))
	copy(buf[4:], res.Record)
	return
}

func (res *HashTransactionRegisterRes) LenOfSyllabStack() uint32 {
	return 0
}

func (res *HashTransactionRegisterRes) LenOfSyllabHeap() (ln uint32) {
	ln = uint32(len(res.Record))
	return
}

func (res *HashTransactionRegisterRes) LenAsSyllab() uint64 {
	return uint64(res.LenOfSyllabStack() + res.LenOfSyllabHeap())
}
