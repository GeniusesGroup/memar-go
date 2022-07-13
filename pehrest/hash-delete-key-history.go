/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../authorization"
	"../ganjine"
	"../protocol"
	"../service"
	"../srpc"
	"../syllab"
)

var HashDeleteKeyHistoryService = hashDeleteKeyHistoryService{
	Service: service.New("urn:giti:index.protocol:service:hash-delete-key-history", "", protocol.Software_PreAlpha, 1587282740).
		SetDetail(protocol.LanguageEnglish, "Index Hash - Delete Key History",
			"Delete all records associate to given IndexKey and delete indexKey itself!",
			[]string{}).
		SetAuthorization(protocol.CRUDDelete, protocol.UserType_App).Expired(0, ""),
}

type hashDeleteKeyHistoryService struct {
	service.Service
}

func (ser *hashDeleteKeyHistoryService) Process(req HashDeleteKeyHistoryReq) (err protocol.Error) {
	var hashIndex = IndexHash{
		RecordID: req.IndexKey,
	}
	var recordsID [][32]byte
	recordsID, err = hashIndex.Get(0, 0)
	var ln = len(recordsID)
	for i := 0; i < ln; i++ {
		err = ganjine.DeleteRecord(&ganjine.DeleteRecordReq{Type: req.Type, RecordID: recordsID[i]})
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

func (ser *hashDeleteKeyHistoryService) ServeSRPC(st protocol.Stream) (err protocol.Error) {
	if st.Connection.UserID != protocol.OS.AppManifest().AppUUID() {
		// TODO::: Attack??
		err = authorization.ErrUserNotAllow
		return
	}

	var req = &HashDeleteKeyHistoryReq{}
	req.FromSyllab(srpc.GetPayload(st.IncomePayload))

	err = ser.Process(req)
	st.OutcomePayload = make([]byte, srpc.MinLength)
	return
}

func (ser *hashDeleteKeyHistoryService) Do(req HashDeleteKeyHistoryReq) (err protocol.Error) {
	return
}
func (ser *hashDeleteKeyHistoryService) doSRPC(req HashDeleteKeyHistoryReq) (err protocol.Error) {
	var node protocol.ApplicationNode
	node, err = protocol.App.GetNodeByObjectID(req.IndexKey)
	if err != nil {
		return
	}

	if node.Status() == protocol.ApplicationState_LocalNode {
		return hashDeleteKeyHistory(req)
	}

	_, err = srpc.SendBidirectionalRequest(node.Conn(), ser, &req)
	return
}

// HashDeleteKeyHistoryReq is request structure of HashDeleteKeyHistory()
type HashDeleteKeyHistoryReq struct {
	Type     ganjine.RequestType
	IndexKey [32]byte
}

/*
	-- Syllab Encoder & Decoder --
*/

// FromSyllab decode from buf to req
// Due to this service just use internally, It skip check buf size syllab rule! Panic occur if bad request received!
func (req *HashDeleteKeyHistoryReq) FromSyllab(payload []byte, stackIndex uint32) {
	req.Type = ganjine.RequestType(syllab.GetUInt8(buf, 0))
	copy(req.IndexKey[:], buf[1:])
	return
}

// ToSyllab encode req to buf
func (req *HashDeleteKeyHistoryReq) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	buf = make([]byte, req.LenAsSyllab()+4) // +4 for sRPC ID instead get offset argument
	syllab.SetUInt8(buf, 4, uint8(req.Type))
	copy(buf[5:], req.IndexKey[:])
	return
}

func (req *HashDeleteKeyHistoryReq) LenOfSyllabStack() uint32 {
	return 33
}

func (req *HashDeleteKeyHistoryReq) LenOfSyllabHeap() (ln uint32) {
	return
}

func (req *HashDeleteKeyHistoryReq) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}
