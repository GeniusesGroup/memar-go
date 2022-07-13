/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../protocol"
	"../authorization"
	"../ganjine"
	"../srpc"
	"../syllab"
)

// HashDeleteKeyService store details about HashDeleteKey service
var HashDeleteKeyService = service.Service{
	URN:                "urn:giti:index.protocol:service:hash-delete-key",
	Domain:             DomainName,
	ID:                 15822279303913373238,
	IssueDate:          1587282740,
	ExpiryDate:         0,
	ExpireInFavorOfURN: "",
	ExpireInFavorOfID:  0,
	Status:             protocol.Software_PreAlpha,

	Authorization: authorization.Service{
		CRUD:     authorization.CRUDDelete,
		UserType: protocol.UserType_App,
	},

	Detail: map[protocol.LanguageID]service.ServiceDetail{
		protocol.LanguageEnglish: {
			Name: "Index Hash - Delete Key",
			Description: `Delete just exiting index hash without any related record!
It wouldn't delete related records! Use DeleteIndexHistory() instead if you want delete all records too!`,
			TAGS: []string{},
		},
	},

	SRPCHandler: HashDeleteKeySRPC,
}

// HashDeleteKey use to delete exiting index hash with all related records IDs!
// It wouldn't delete related records! Use DeleteIndexHistory() instead if you want delete all records too!
func HashDeleteKey(req *HashDeleteKeyReq) (err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByStorage(req.MediaTypeID, req.IndexKey)
	if err != nil {
		return
	}

	if node.Node.State == protocol.ApplicationState_LocalNode {
		return HashDeleteKey(req)
	}

	var st protocol.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &HashDeleteKeyService
	st.OutcomePayload = req.ToSyllab()

	err = node.Conn.Send(st)
	if err != nil {
		return
	}

	return
}

// HashDeleteKeySRPC is sRPC handler of HashDeleteKey service.
func HashDeleteKeySRPC(st protocol.Stream) {
	if st.Connection.UserID != protocol.OS.AppManifest().AppUUID() {
		// TODO::: Attack??
		err = authorization.ErrUserNotAllow
		return
	}

	var req = &HashDeleteKeyReq{}
	req.FromSyllab(srpc.GetPayload(st.IncomePayload))

	err = HashDeleteKey(req)
	st.OutcomePayload = make([]byte, srpc.MinLength)
}

// HashDeleteKeyReq is request structure of HashDeleteKey()
type HashDeleteKeyReq struct {
	Type     ganjine.RequestType
	IndexKey [32]byte
}

// HashDeleteKey delete just exiting index hash without any related record.
func HashDeleteKey(req *HashDeleteKeyReq) (err protocol.Error) {
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

			// Set HashDeleteKey ServiceID
			st.Service = &service.Service{ID: 15822279303913373238}
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
	err = hashIndex.DeleteRecord()
	return
}

/*
	-- Syllab Encoder & Decoder --
*/

// FromSyllab decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashDeleteKeyReq) FromSyllab(payload []byte, stackIndex uint32) {
	req.Type = ganjine.RequestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	return
}

// ToSyllab encode req to buf
func (req *HashDeleteKeyReq) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	buf = make([]byte, req.LenAsSyllab()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])

	return
}

func (req *HashDeleteKeyReq) LenOfSyllabStack() uint32 {
	return 33
}

func (req *HashDeleteKeyReq) LenOfSyllabHeap() (ln uint32) {
	return
}

func (req *HashDeleteKeyReq) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}
