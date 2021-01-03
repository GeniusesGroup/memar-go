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

// HashDeleteKeyService store details about HashDeleteKey service
var HashDeleteKeyService = achaemenid.Service{
	ID:                4172448155,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // English name of favor service just to show off!
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Authorization: authorization.Service{
		CRUD:     authorization.CRUDDelete,
		UserType: authorization.UserTypeApp,
	},

	Name: map[lang.Language]string{
		lang.LanguageEnglish: "Index Hash - Delete Key",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: `Delete just exiting index hash without any related record!
It wouldn't delete related records! Use DeleteIndexHistory() instead if you want delete all records too!`,
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashDeleteKeySRPC,
}

// HashDeleteKeySRPC is sRPC handler of HashDeleteKey service.
func HashDeleteKeySRPC(st *achaemenid.Stream) {
	if st.Connection.UserID != achaemenid.Server.AppID {
		// TODO::: Attack??
		st.Err = ganjine.ErrNotAuthorizeRequest
		return
	}

	var req = &HashDeleteKeyReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	st.Err = HashDeleteKey(req)
	st.OutcomePayload = make([]byte, 4)
}

// HashDeleteKeyReq is request structure of HashDeleteKey()
type HashDeleteKeyReq struct {
	Type     gs.RequestType
	IndexKey [32]byte
}

// HashDeleteKey delete just exiting index hash without any related record.
func HashDeleteKey(req *HashDeleteKeyReq) (err *er.Error) {
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

			// Set HashDeleteKey ServiceID
			st.Service = &achaemenid.Service{ID: 3411747355}
			st.OutcomePayload = reqEncoded

			err = achaemenid.SrpcOutcomeRequestHandler(st)
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
	err = hashIndex.DeleteRecord()
	return
}

/*
	-- Syllab Encoder & Decoder --
*/

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashDeleteKeyReq) SyllabDecoder(buf []byte) {
	req.Type = gs.RequestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	return
}

// SyllabEncoder encode req to buf
func (req *HashDeleteKeyReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])

	return
}

func (req *HashDeleteKeyReq) syllabStackLen() (ln uint32) {
	return 33
}

func (req *HashDeleteKeyReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashDeleteKeyReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}
