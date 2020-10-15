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

// DeleteRecordService store details about DeleteRecord service
var DeleteRecordService = achaemenid.Service{
	ID:                1758631843,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "",
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "DeleteRecord",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: `Delete specific record by given ID in all cluster!
We don't suggest use this service, due to we strongly suggest think about data as immutable entity(stream and time)
It won't delete record history or indexes associate to it!`,
	},
	TAGS: []string{""},

	SRPCHandler: DeleteRecordSRPC,
}

// DeleteRecordSRPC is sRPC handler of DeleteRecord service.
func DeleteRecordSRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var req = &DeleteRecordReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	st.Err = DeleteRecord(req)
}

// DeleteRecordReq is request structure of DeleteRecord()
type DeleteRecordReq struct {
	Type     requestType
	RecordID [32]byte
}

// DeleteRecord delete specific record by given ID in all cluster!
func DeleteRecord(req *DeleteRecordReq) (err error) {
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

			// Set DeleteRecord ServiceID
			st.Service = &achaemenid.Service{ID: 1758631843}
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
	err = persiaos.DeleteStorageRecord(req.RecordID)
	return
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *DeleteRecordReq) SyllabDecoder(buf []byte) {
	req.Type = requestType(syllab.GetUInt8(buf, 0))
	copy(req.RecordID[:], buf[1:])
	return
}

// SyllabEncoder encode req to buf
func (req *DeleteRecordReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.RecordID[:])

	return
}

func (req *DeleteRecordReq) syllabStackLen() (ln uint32) {
	return 33
}

func (req *DeleteRecordReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *DeleteRecordReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}
