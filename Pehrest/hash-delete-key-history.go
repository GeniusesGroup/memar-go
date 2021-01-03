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

// HashDeleteKeyHistoryService store details about HashDeleteKeyHistory service
var HashDeleteKeyHistoryService = achaemenid.Service{
	ID:                2599603093,
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
		lang.LanguageEnglish: "Index Hash - Delete Key History",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: "Delete all record associate to given index and delete index itself!",
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashDeleteKeyHistorySRPC,
}

// HashDeleteKeyHistorySRPC is sRPC handler of HashDeleteKeyHistory service.
func HashDeleteKeyHistorySRPC(st *achaemenid.Stream) {
	if st.Connection.UserID != achaemenid.Server.AppID {
		// TODO::: Attack??
		st.Err = ganjine.ErrNotAuthorizeRequest
		return
	}

	var req = &HashDeleteKeyHistoryReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	st.Err = HashDeleteKeyHistory(req)
	st.OutcomePayload = make([]byte, 4)
}

// HashDeleteKeyHistoryReq is request structure of HashDeleteKeyHistory()
type HashDeleteKeyHistoryReq struct {
	Type     gs.RequestType
	IndexKey [32]byte
}

// HashDeleteKeyHistory delete all records associate to given IndexKey and delete indexKey itself!
func HashDeleteKeyHistory(req *HashDeleteKeyHistoryReq) (err *er.Error) {
	var hashIndex = IndexHash{
		RecordID: req.IndexKey,
	}
	var recordsID [][32]byte
	recordsID, err = hashIndex.Get(0, 0)
	var ln = len(recordsID)
	for i := 0; i < ln; i++ {
		err = gs.DeleteRecord(&gs.DeleteRecordReq{Type: req.Type, RecordID: recordsID[i]})
		if err != nil {
			// TODO::: Can we easily return error if two nodes do their job and just one node connection lost??
			return
		}
	}

	err = HashDeleteKey(&HashDeleteKeyReq{Type: req.Type, IndexKey: req.IndexKey})
	if err != nil {
		// TODO::: Can we easily return error if two nodes do their job and just one node connection lost??
		return
	}

	return
}

/*
	-- Syllab Encoder & Decoder --
*/

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashDeleteKeyHistoryReq) SyllabDecoder(buf []byte) {
	req.Type = gs.RequestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	return
}

// SyllabEncoder encode req to buf
func (req *HashDeleteKeyHistoryReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])
	return
}

func (req *HashDeleteKeyHistoryReq) syllabStackLen() (ln uint32) {
	return 33
}

func (req *HashDeleteKeyHistoryReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashDeleteKeyHistoryReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}
