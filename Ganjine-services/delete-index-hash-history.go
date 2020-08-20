/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	"../achaemenid"
	"../ganjine"
)

var deleteIndexHashHistoryService = achaemenid.Service{
	ID:              691384835,
	Name:            "DeleteIndexHashHistory",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		"Delete all record associate to given index and delete index itself!",
	},
	TAGS:        []string{""},
	SRPCHandler: DeleteIndexHashHistorySRPC,
}

// DeleteIndexHashHistorySRPC is sRPC handler of DeleteIndexHashHistory service.
func DeleteIndexHashHistorySRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &DeleteIndexHashHistoryReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Err = DeleteIndexHashHistory(req)
}

// DeleteIndexHashHistoryReq is request structure of DeleteIndexHashHistory()
type DeleteIndexHashHistoryReq struct {
	Type      requestType
	IndexHash [32]byte
}

// DeleteIndexHashHistory delete all record associate to given index and delete index itself!
func DeleteIndexHashHistory(req *DeleteIndexHashHistoryReq) (err error) {
	err = DeleteIndexHash(&DeleteIndexHashReq{Type: req.Type, IndexHash: req.IndexHash})
	if err != nil {
		// TODO::: Can we easily return error if two nodes do their job and just one node connection lost??
		return
	}

	var hashIndex = ganjine.HashIndex{
		RecordID: req.IndexHash,
	}
	var recordsID [][32]byte
	recordsID, err = hashIndex.GetIndexRecords(0, 0)
	var ln = len(recordsID)
	for i := 0; i < ln; i++ {
		err = DeleteRecord(&DeleteRecordReq{Type: req.Type, RecordID: recordsID[i]})
		if err != nil {
			// TODO::: Can we easily return error if two nodes do their job and just one node connection lost??
			return
		}
	}

	return
}

// SyllabDecoder decode from buf to req
func (req *DeleteIndexHashHistoryReq) SyllabDecoder(buf []byte) (err error) {
	req.Type = requestType(buf[0])
	copy(req.IndexHash[:], buf[1:])
	return
}

// SyllabEncoder encode req to buf
func (req *DeleteIndexHashHistoryReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, 37) // 37=4+1+32 >> first 4+ for sRPC ID instead get offset argument

	buf[4] = byte(req.Type)
	copy(buf[5:], req.IndexHash[:])

	return
}
