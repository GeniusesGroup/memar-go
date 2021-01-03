/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../achaemenid"
	"../authorization"
	er "../error"
	"../ganjine"
	gs "../ganjine-services"
	lang "../language"
	"../srpc"
	"../syllab"
)

// HashSetValueService store details about HashSetValue service
var HashSetValueService = achaemenid.Service{
	ID:                275144870,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // English name of favor service just to show off!
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Authorization: authorization.Service{
		CRUD:     authorization.CRUDUpdate,
		UserType: authorization.UserTypeApp,
	},

	Name: map[lang.Language]string{
		lang.LanguageEnglish: "Index Hash - Set Value",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: "set a record ID to new||exiting index hash.",
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashSetValueSRPC,
}

// HashSetValueSRPC is sRPC handler of HashSetValue service.
func HashSetValueSRPC(st *achaemenid.Stream) {
	if st.Connection.UserID != achaemenid.Server.AppID {
		// TODO::: Attack??
		st.Err = ganjine.ErrNotAuthorizeRequest
		return
	}

	var req = &HashSetValueReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	st.Err = HashSetValue(req)
	st.OutcomePayload = make([]byte, 4)
}

// HashSetValueReq is request structure of HashSetValue()
type HashSetValueReq struct {
	Type       gs.RequestType
	IndexKey   [32]byte
	IndexValue [32]byte // can be RecordID or any data up to 32 byte length
}

// HashSetValue set a record ID to new||exiting index hash.
func HashSetValue(req *HashSetValueReq) (err *er.Error) {
	if req.Type == gs.RequestTypeBroadcast {
		// tell other node that this node handle request and don't send this request to other nodes!
		req.Type = gs.RequestTypeStandalone
		var reqEncoded = req.SyllabEncoder()

		// send request to other related nodes
		for i := 1; i < len(ganjine.Cluster.Replications.Zones); i++ {
			// Make new request-response streams
			var st *achaemenid.Stream
			st, err = ganjine.Cluster.Replications.Zones[i].Nodes[ganjine.Cluster.Node.ID].Conn.MakeOutcomeStream(0)
			if err != nil {
				// TODO::: Can we easily return error if two nodes did their job and not have enough resource to send request to final node??
				return
			}

			// Set HashSetValue ServiceID
			st.Service = &achaemenid.Service{ID: 1881585857}
			st.OutcomePayload = reqEncoded

			err = achaemenid.SrpcOutcomeRequestHandler( st)
			if err != nil {
				// TODO::: Can we easily return error if two nodes do their job and just one node connection lost??
				return
			}

			// TODO::: Can we easily return response error without handle some known situations??
			err = st.Err
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

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashSetValueReq) SyllabDecoder(buf []byte) {
	req.Type = gs.RequestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	copy(req.IndexValue[:], buf[33:])
	return
}

// SyllabEncoder encode req to buf
func (req *HashSetValueReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])
	copy(buf[37:], req.IndexValue[:])
	return
}

func (req *HashSetValueReq) syllabStackLen() (ln uint32) {
	return 65
}

func (req *HashSetValueReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashSetValueReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}
