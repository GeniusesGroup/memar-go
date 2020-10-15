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

// HashIndexDeleteValueService store details about HashIndexDeleteValue service
var HashIndexDeleteValueService = achaemenid.Service{
	ID:                2091454202,
	CRUD:              authorization.CRUDDelete,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // English name of favor service just to show off!
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "HashIndexDeleteValue",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: "Delete the value from exiting index key",
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashIndexDeleteValueSRPC,
}

// HashIndexDeleteValueSRPC is sRPC handler of HashIndexDeleteValue service.
func HashIndexDeleteValueSRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var req = &HashIndexDeleteValueReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	st.Err = HashIndexDeleteValue(req)
	st.OutcomePayload = make([]byte, 4)
}

// HashIndexDeleteValueReq is request structure of HashIndexDeleteValue()
type HashIndexDeleteValueReq struct {
	Type       requestType
	IndexKey   [32]byte
	IndexValue [32]byte
}

// HashIndexDeleteValue delete the value from exiting index key
func HashIndexDeleteValue(req *HashIndexDeleteValueReq) (err error) {
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

			st.Service = &achaemenid.Service{ID: 2091454202}
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
	var hashIndex = ganjine.HashIndex{
		RecordID: req.IndexKey,
	}
	err = hashIndex.Delete(req.IndexValue)
	return
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashIndexDeleteValueReq) SyllabDecoder(buf []byte) {
	req.Type = requestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	copy(req.IndexValue[:], buf[33:])
	return
}

// SyllabEncoder encode req to buf
func (req *HashIndexDeleteValueReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])
	copy(buf[37:], req.IndexValue[:])
	return
}

func (req *HashIndexDeleteValueReq) syllabStackLen() (ln uint32) {
	return 65
}

func (req *HashIndexDeleteValueReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashIndexDeleteValueReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}
