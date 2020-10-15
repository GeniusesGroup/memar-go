/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	persiaos "../PersiaOS-sdk"
	"../achaemenid"
	"../authorization"
	"../ganjine"
	lang "../language"
	"../srpc"
	"../syllab"
)

// HashIndexDeleteKeyService store details about HashIndexDeleteKey service
var HashIndexDeleteKeyService = achaemenid.Service{
	ID:                4081206292,
	CRUD:              authorization.CRUDDelete,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // English name of favor service just to show off!
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "HashIndexDeleteKey",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: `Delete just exiting index hash without any related record!
It wouldn't delete related records! Use DeleteIndexHistory() instead if you want delete all records too!`,
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashIndexDeleteKeySRPC,
}

// HashIndexDeleteKeySRPC is sRPC handler of HashIndexDeleteKey service.
func HashIndexDeleteKeySRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var req = &HashIndexDeleteKeyReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	st.Err = HashIndexDeleteKey(req)
	st.OutcomePayload = make([]byte, 4)
}

// HashIndexDeleteKeyReq is request structure of HashIndexDeleteKey()
type HashIndexDeleteKeyReq struct {
	Type     requestType
	IndexKey [32]byte
}

// HashIndexDeleteKey delete just exiting index hash without any related record.
func HashIndexDeleteKey(req *HashIndexDeleteKeyReq) (err error) {
	if req.Type == RequestTypeBroadcast {
		// tell other node that this node handle request and don't send this request to other nodes!
		req.Type = RequestTypeStandalone
		var reqEncoded = req.SyllabEncoder()

		// send request to other related nodes
		var i uint8
		for i = 1; i < cluster.Manifest.TotalZones; i++ {
			// Make new request-response streams
			var st *achaemenid.Stream
			st, err = cluster.Replications.Zones[i].Nodes[cluster.Node.ID].Conn.MakeOutcomeStream(0)
			if err != nil {
				// TODO::: Can we easily return error if two nodes did their job and not have enough resource to send request to final node??
				return
			}

			// Set HashIndexDeleteKey ServiceID
			st.Service = &achaemenid.Service{ID: 3411747355}
			st.OutcomePayload = reqEncoded

			err = achaemenid.SrpcOutcomeRequestHandler(server, st)
			if err != nil {
				// TODO::: Can we easily return error if two nodes do their job and just one node connection lost??
				return
			}

			// TODO::: Can we easily return response error without handle some known situations??
			err = st.Err
		}
	}

	// Do for i=0 as local node
	err = persiaos.DeleteStorageRecord(req.IndexKey)
	return
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashIndexDeleteKeyReq) SyllabDecoder(buf []byte) {
	req.Type = requestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	return
}

// SyllabEncoder encode req to buf
func (req *HashIndexDeleteKeyReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])

	return
}

func (req *HashIndexDeleteKeyReq) syllabStackLen() (ln uint32) {
	return 33
}

func (req *HashIndexDeleteKeyReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashIndexDeleteKeyReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}
