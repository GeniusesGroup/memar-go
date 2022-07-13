/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../authorization"
	"../ganjine"
	"../protocol"
	"../srpc"
	"../syllab"
)

// HashSetValueService store details about HashSetValue service
var HashSetValueService = service.Service{
	URN:                "urn:giti:index.protocol:service:hash-set-value",
	Domain:             DomainName,
	ID:                 17266375388745136913,
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
			Name:        "Index Hash - Set Value",
			Description: "set a record ID to new||exiting index hash.",
			TAGS:        []string{},
		},
	},

	SRPCHandler: HashSetValueSRPC,
}

// HashSetValue set a record ID to new||exiting index hash!
func HashSetValue(req *HashSetValueReq) (err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByStorage(req.MediaTypeID, req.IndexKey)
	if err != nil {
		return
	}

	if node.Node.State == protocol.ApplicationState_LocalNode {
		err = HashSetValue(req)
		return
	}

	var st protocol.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &HashSetValueService
	st.OutcomePayload = req.ToSyllab()

	err = node.Conn.Send(st)
	return
}

// HashSetValueSRPC is sRPC handler of HashSetValue service.
func HashSetValueSRPC(st protocol.Stream) {
	if st.Connection.UserID != protocol.OS.AppManifest().AppUUID() {
		// TODO::: Attack??
		err = authorization.ErrUserNotAllow
		return
	}

	var req = &HashSetValueReq{}
	req.FromSyllab(srpc.GetPayload(st.IncomePayload))

	err = HashSetValue(req)
	st.OutcomePayload = make([]byte, srpc.MinLength)
}

// HashSetValueReq is request structure of HashSetValue()
type HashSetValueReq struct {
	Type       ganjine.RequestType
	IndexKey   [32]byte
	IndexValue [32]byte // can be RecordID or any data up to 32 byte length
}

// HashSetValue set a record ID to new||exiting index hash.
func HashSetValue(req *HashSetValueReq) (err protocol.Error) {
	if req.Type == ganjine.RequestTypeBroadcast {
		// tell other node that this node handle request and don't send this request to other nodes!
		req.Type = ganjine.RequestTypeStandalone
		var reqEncoded = req.ToSyllab()

		// send request to other related nodes
		for i := 1; i < len(ganjine.Cluster.Replications.Zones); i++ {
			var conn = ganjine.Cluster.Replications.Zones[i].Nodes[ganjine.Cluster.Node.ID].Conn
			// Make new request-response streams
			var st protocol.Stream
			st, err = conn.MakeOutcomeStream(0)
			if err != nil {
				// TODO::: Can we easily return error if two nodes did their job and not have enough resource to send request to final node??
				return
			}

			// Set HashSetValue ServiceID
			st.Service = &service.Service{ID: 17266375388745136913}
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
	var hashIndex = IndexHash{
		RecordID: req.IndexKey,
	}
	err = hashIndex.Push(req.IndexValue)
	return
}

/*
	-- Syllab Encoder & Decoder --
*/

// FromSyllab decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashSetValueReq) FromSyllab(payload []byte, stackIndex uint32) {
	req.Type = ganjine.RequestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	copy(req.IndexValue[:], buf[33:])
	return
}

// ToSyllab encode req to buf
func (req *HashSetValueReq) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	buf = make([]byte, req.LenAsSyllab()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])
	copy(buf[37:], req.IndexValue[:])
	return
}

func (req *HashSetValueReq) LenOfSyllabStack() uint32 {
	return 65
}

func (req *HashSetValueReq) LenOfSyllabHeap() (ln uint32) {
	return
}

func (req *HashSetValueReq) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}
