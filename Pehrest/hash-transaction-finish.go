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

// HashTransactionFinishService store details about HashTransactionFinish service
var HashTransactionFinishService = achaemenid.Service{
	ID:                3635898242,
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
		lang.LanguageEnglish: "Index Hash - Transaction Finish",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: `use to approve transaction!
Transaction Manager will set record and index! no further action need after this call!`,
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashTransactionFinishSRPC,
}

// HashTransactionFinishSRPC is sRPC handler of HashTransactionFinish service.
func HashTransactionFinishSRPC(st *achaemenid.Stream) {
	if st.Connection.UserID != achaemenid.Server.AppID {
		// TODO::: Attack??
		st.Err = ganjine.ErrNotAuthorizeRequest
		return
	}

	var req = &HashTransactionFinishReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	st.Err = HashTransactionFinish(req)
	st.OutcomePayload = make([]byte, 4)
}

// HashTransactionFinishReq is request structure of HashTransactionFinish()
type HashTransactionFinishReq struct {
	Type     gs.RequestType
	IndexKey [32]byte
	Record   []byte
}

// HashTransactionFinish approve transaction!
func HashTransactionFinish(req *HashTransactionFinishReq) (err *er.Error) {
	if req.Type == gs.RequestTypeBroadcast {
		// tell other node that this node handle request and don't send this request to other nodes!
		req.Type = gs.RequestTypeStandalone
		var reqEncoded = req.SyllabEncoder()

		// send request to other related nodes
		for i := 1; i < len(ganjine.Cluster.Replications.Zones); i++ {
			var st *achaemenid.Stream
			st, err = ganjine.Cluster.Replications.Zones[i].Nodes[ganjine.Cluster.Node.ID].Conn.MakeOutcomeStream(0)
			if err != nil {
				// TODO::: Can we easily return error if two nodes did their job and not have enough resource to send request to final node??
				return
			}

			st.Service = &achaemenid.Service{ID: 3635898242}
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
	err = ganjine.Cluster.TransactionManager.FinishTransaction(req.IndexKey, req.Record)
	return
}

/*
	-- Syllab Encoder & Decoder --
*/

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashTransactionFinishReq) SyllabDecoder(buf []byte) {
	req.Type = gs.RequestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	req.Record = buf[33:]
	return
}

// SyllabEncoder encode req to buf
func (req *HashTransactionFinishReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])
	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Record in buf
	// syllab.SetUInt32(buf, 37, res.syllabStackLen())
	// syllab.SetUInt32(buf, 45, uint32(len(res.Record)))
	copy(buf[37:], req.Record[:])
	return
}

func (req *HashTransactionFinishReq) syllabStackLen() (ln uint32) {
	return 33
}

func (req *HashTransactionFinishReq) syllabHeapLen() (ln uint32) {
	ln = uint32(len(req.Record))
	return
}

func (req *HashTransactionFinishReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}
