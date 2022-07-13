/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../authorization"
	"../ganjine"
	"../protocol"
	"../srpc"
	"../syllab"
)

// HashDeleteValueService store details about HashDeleteValue service
var HashDeleteValueService = service.Service{
	URN:                "urn:giti:index.protocol:service:has-delete-value",
	Domain:             DomainName,
	ID:                 4811402533981149999,
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
			Name:        "Index Hash - Delete Value",
			Description: "Delete the value from exiting index hash key",
			TAGS:        []string{},
		},
	},

	SRPCHandler: HashDeleteValueSRPC,
}

// HashDeleteValue delete the value from exiting index key
func HashDeleteValue(req *HashDeleteValueReq) (err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByStorage(req.MediaTypeID, req.IndexKey)
	if err != nil {
		return
	}

	if node.Node.State == protocol.ApplicationState_LocalNode {
		return HashDeleteValue(req)
	}

	var st protocol.Stream
	st, err = node.Conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	st.Service = &HashDeleteValueService
	st.OutcomePayload = req.ToSyllab()

	err = node.Conn.Send(st)
	if err != nil {
		return
	}
	return
}

// HashDeleteValueSRPC is sRPC handler of HashDeleteValue service.
func HashDeleteValueSRPC(st protocol.Stream) {
	if st.Connection.UserID != protocol.OS.AppManifest().AppUUID() {
		// TODO::: Attack??
		err = authorization.ErrUserNotAllow
		return
	}

	var req = &HashDeleteValueReq{}
	req.FromSyllab(srpc.GetPayload(st.IncomePayload))

	err = HashDeleteValue(req)
	st.OutcomePayload = make([]byte, srpc.MinLength)
}

// HashDeleteValueReq is request structure of HashDeleteValue()
type HashDeleteValueReq struct {
	Type       ganjine.RequestType
	IndexKey   [32]byte
	IndexValue [32]byte
}

// HashDeleteValue delete the value from exiting index key
func HashDeleteValue(req *HashDeleteValueReq) (err protocol.Error) {
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

			st.Service = &service.Service{ID: 4811402533981149999}
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
	err = hashIndex.Delete(req.IndexValue)
	return
}

/*
	-- Syllab Encoder & Decoder --
*/

// FromSyllab decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashDeleteValueReq) FromSyllab(payload []byte, stackIndex uint32) {
	req.Type = ganjine.RequestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	copy(req.IndexValue[:], buf[33:])
	return
}

// ToSyllab encode req to buf
func (req *HashDeleteValueReq) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	buf = make([]byte, req.LenAsSyllab()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])
	copy(buf[37:], req.IndexValue[:])
	return
}

func (req *HashDeleteValueReq) LenOfSyllabStack() uint32 {
	return 65
}

func (req *HashDeleteValueReq) LenOfSyllabHeap() (ln uint32) {
	return
}

func (req *HashDeleteValueReq) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}
