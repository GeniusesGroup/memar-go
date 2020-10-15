/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	persiaos "../PersiaOS-sdk"
	"../achaemenid"
	"../ganjine"
	lang "../language"
	"../srpc"
	"../syllab"
)

// SetRecordService store details about SetRecord service
var SetRecordService = achaemenid.Service{
	ID:                10488062,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "",
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "SetRecord",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: "Write a whole record or replace old record if it is exist!",
	},
	TAGS: []string{""},

	SRPCHandler: SetRecordSRPC,
}

// SetRecordSRPC is sRPC handler of SetRecord service.
func SetRecordSRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var req = &SetRecordReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	st.Err = SetRecord(req)
}

// SetRecordReq is request structure of SetRecord()
type SetRecordReq struct {
	Type   requestType
	Record []byte
}

// SetRecord respect all data in record and don't change something like RecordID or WriteTime!
// If data like OwnerAppID is wrong you can't get record anymore!
func SetRecord(req *SetRecordReq) (err error) {
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

			// Set SetRecord ServiceID
			st.Service = &achaemenid.Service{ID: 10488062}
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
	err = persiaos.SetStorageRecord(req.Record)
	return
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *SetRecordReq) SyllabDecoder(buf []byte) {
	req.Type = requestType(syllab.GetUInt8(buf, 0))
	// Due to just have one field in res structure we break syllab rules and skip get address and size of res.Record from buf
	req.Record = buf[1:]
	return
}

// SyllabEncoder encode req to buf
func (req *SetRecordReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	// Due to just have one field in res structure we break syllab rules and skip set address and size of res.Record in buf
	// syllab.SetUInt32(buf, 5, res.syllabStackLen())
	// syllab.SetUInt32(buf, 9, uint32(len(res.Record))))
	copy(buf[5:], req.Record)
	return
}

func (req *SetRecordReq) syllabStackLen() (ln uint32) {
	return 1
}

func (req *SetRecordReq) syllabHeapLen() (ln uint32) {
	ln = uint32(len(req.Record))
	return
}

func (req *SetRecordReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}
