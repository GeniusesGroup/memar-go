/* For license and copyright information please see LEGAL file in repository */

package gs

import "../achaemenid"

var listenToIndexService = achaemenid.Service{
	ID:              2145882122,
	Name:            "ListenToIndex",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Description: []string{
		`get records to given index hash when new record set!
		Request Must send to specific node that handle that hash index range!!
		This service has a lot of use cases like:
		- any geospatial usage e.g. tracking device or user, ...
		- following content author like telegram channels or instagram live video!`,
	},
	TAGS:        []string{""},
	SRPCHandler: ListenToIndexSRPC,
}

// ListenToIndexSRPC is sRPC handler of ListenToIndex service.
func ListenToIndexSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var req = &ListenToIndexReq{}
	st.ReqRes.Err = req.SyllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	var res *ListenToIndexRes
	res, st.ReqRes.Err = ListenToIndex(req)
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Payload = res.SyllabEncoder()
}

// ListenToIndexReq is request structure of listenToIndex()
type ListenToIndexReq struct {
	IndexHash [32]byte
}

// ListenToIndexRes is response structure of listenToIndex()
type ListenToIndexRes struct {
	// Record []byte TODO::: it can't be simple byte, maybe channel
}

// ListenToIndex get the recordID by index hash when new record set!
func ListenToIndex(req *ListenToIndexReq) (res *ListenToIndexRes, err error) {
	return
}

// SyllabDecoder decode from buf to req
func (req *ListenToIndexReq) SyllabDecoder(buf []byte) (err error) {
	return
}

// SyllabEncoder encode req to buf
func (req *ListenToIndexReq) SyllabEncoder() (buf []byte) {
	buf = make([]byte, 32+4) // +4 for sRPC ID instead get offset argument

	copy(buf[4:], req.IndexHash[:])

	return
}

// SyllabDecoder decode from buf to req
func (res *ListenToIndexRes) SyllabDecoder(buf []byte) (err error) {
	return
}

// SyllabEncoder encode req to buf
func (res *ListenToIndexRes) SyllabEncoder() (buf []byte) {
	return
}
