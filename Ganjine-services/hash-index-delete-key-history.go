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

// HashIndexDeleteKeyHistoryService store details about HashIndexDeleteKeyHistory service
var HashIndexDeleteKeyHistoryService = achaemenid.Service{
	ID:                1505692432,
	CRUD:              authorization.CRUDDelete,
	IssueDate:         1587282740,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // English name of favor service just to show off!
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "HashIndexDeleteKeyHistory",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: "Delete all record associate to given index and delete index itself!",
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: HashIndexDeleteKeyHistorySRPC,
}

// HashIndexDeleteKeyHistorySRPC is sRPC handler of HashIndexDeleteKeyHistory service.
func HashIndexDeleteKeyHistorySRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var req = &HashIndexDeleteKeyHistoryReq{}
	req.SyllabDecoder(srpc.GetPayload(st.IncomePayload))

	st.Err = HashIndexDeleteKeyHistory(req)
	st.OutcomePayload = make([]byte, 4)
}

// HashIndexDeleteKeyHistoryReq is request structure of HashIndexDeleteKeyHistory()
type HashIndexDeleteKeyHistoryReq struct {
	Type     requestType
	IndexKey [32]byte
}

// HashIndexDeleteKeyHistory delete all records associate to given IndexKey and delete indexKey itself!
func HashIndexDeleteKeyHistory(req *HashIndexDeleteKeyHistoryReq) (err error) {
	var hashIndex = ganjine.HashIndex{
		RecordID: req.IndexKey,
	}
	var recordsID [][32]byte
	recordsID, err = hashIndex.Get(0, 0)
	var ln = len(recordsID)
	for i := 0; i < ln; i++ {
		err = DeleteRecord(&DeleteRecordReq{Type: req.Type, RecordID: recordsID[i]})
		if err != nil {
			// TODO::: Can we easily return error if two nodes do their job and just one node connection lost??
			return
		}
	}

	err = HashIndexDeleteKey(&HashIndexDeleteKeyReq{Type: req.Type, IndexKey: req.IndexKey})
	if err != nil {
		// TODO::: Can we easily return error if two nodes do their job and just one node connection lost??
		return
	}

	return
}

// SyllabDecoder decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashIndexDeleteKeyHistoryReq) SyllabDecoder(buf []byte) {
	req.Type = requestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	return
}

// SyllabEncoder encode req to buf
func (req *HashIndexDeleteKeyHistoryReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, req.syllabLen()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])
	return
}

func (req *HashIndexDeleteKeyHistoryReq) syllabStackLen() (ln uint32) {
	return 33
}

func (req *HashIndexDeleteKeyHistoryReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *HashIndexDeleteKeyHistoryReq) syllabLen() (ln uint64) {
	return uint64(req.syllabStackLen() + req.syllabHeapLen())
}
