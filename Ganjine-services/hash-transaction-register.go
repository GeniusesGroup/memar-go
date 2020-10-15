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

// HashTransactionRegisterService store details about HashTransactionRegister service
var HashTransactionRegisterService = achaemenid.Service{
	ID:                2613523897,
	CRUD:              authorization.CRUDCreate,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // English name of favor service just to show off!
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "HashTransactionRegister",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: `Register new transaction on queue and get last record when transaction ready for this one!
Requester must send FinishTransaction() immediately, otherwise Transaction manager will drop this request from queue and chain!
transaction write can be on secondary indexes not primary indexes, due to primary index must always unique!
transaction manager on any node in a replication must sync with master replication corresponding node manager!
Get a record by ID when record ready to submit! Usually use in transaction queue to act when record ready to read!
Must send this request to specific node that handle that range!!`,
	},
	TAGS: []string{"transactional authority", "index lock ticket"},

	SRPCHandler: HashTransactionRegisterSRPC,
}

// HashTransactionRegisterSRPC is sRPC handler of HashTransactionRegister service.
func HashTransactionRegisterSRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var req = &HashTransactionRegisterReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	var res *HashTransactionRegisterRes
	res, st.Err = HashTransactionRegister(req)
	if st.Err != nil {
		return
	}

	st.OutcomePayload = res.SyllabEncoder()
}

// HashTransactionRegisterReq is request structure of HashTransactionRegister()
type HashTransactionRegisterReq struct {
	Type     requestType
	IndexKey [32]byte
	RecordID [32]byte
}

// HashTransactionRegisterRes is response structure of HashTransactionRegister()
type HashTransactionRegisterRes struct {
	Record []byte
}

// HashTransactionRegister register new transaction on queue and get last record when transaction ready for this one!
func HashTransactionRegister(req *HashTransactionRegisterReq) (res *HashTransactionRegisterRes, err error) {
	res = &HashTransactionRegisterRes{}

	if req.Type == RequestTypeBroadcast {
		// tell other node that this node handle request and don't send this request to other nodes!
		req.Type = RequestTypeStandalone
		var reqEncoded = req.SyllabEncoder()

		// send request to other related nodes
		var i uint8
		for i = 1; i < cluster.Manifest.TotalZones; i++ {
			var st *achaemenid.Stream
			st, err = cluster.Replications.Zones[i].Nodes[cluster.Node.ID].Conn.MakeOutcomeStream(0)
			if err != nil {
				// TODO::: Can we easily return error if two nodes did their job and not have enough resource to send request to final node??
				return
			}

			st.Service = &achaemenid.Service{ID: 2613523897}
			st.OutcomePayload = reqEncoded

			err = achaemenid.SrpcOutcomeRequestHandler(server, st)
			if err != nil {
				// TODO::: Can we easily return error if two nodes do their job and just one node connection lost??
				return
			}

			// TODO::: Can we easily return response error without handle some known situations??
			err = st.Err
		}

		// Do for i=0 as local node
		res.Record, err = cluster.TransactionManager.RegisterTransaction(req.IndexKey, req.RecordID)
	} else {
		// Don't send last record due to Master node will give it to requester!
		_, err = cluster.TransactionManager.RegisterTransaction(req.IndexKey, req.RecordID)
	}

	return
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashTransactionRegisterReq) SyllabDecoder(buf []byte) {
	req.Type = requestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	copy(req.RecordID[:], buf[33:])
	return
}

// SyllabEncoder encode req to buf
func (req *HashTransactionRegisterReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])
	copy(buf[37:], req.RecordID[:])
	return
}

func (req *HashTransactionRegisterReq) syllabStackLen() (ln uint32) {
	return 65
}

func (req *HashTransactionRegisterReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashTransactionRegisterReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (res *HashTransactionRegisterRes) SyllabDecoder(buf []byte) {
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	res.Record = buf
	return
}

// SyllabEncoder encode req to buf
func (res *HashTransactionRegisterRes) SyllabEncoder() (buf []byte) {
	buf = make([]byte, res.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Record in buf
	// syllab.SetUInt32(buf, 4, res.syllabStackLen())
	// syllab.SetUInt32(buf, 8, uint32(len(res.Record)))
	copy(buf[4:], res.Record)
	return
}

func (res *HashTransactionRegisterRes) syllabStackLen() (ln uint32) {
	return 0
}

func (res *HashTransactionRegisterRes) syllabHeapLen() (ln uint32) {
	ln = uint32(len(res.Record))
	return
}

func (res *HashTransactionRegisterRes) syllabLen() (ln uint64) {
	return uint64(res.syllabStackLen() + res.syllabHeapLen())
}
