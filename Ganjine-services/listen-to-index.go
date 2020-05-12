/* For license and copyright information please see LEGAL file in repository */

package services

import "../achaemenid"

var listenToIndexService = achaemenid.Service{
	ID:              2145882122,
	Name:            "ListenToIndex",
	IssueDate:       1587282740,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          achaemenid.ServiceStatePreAlpha,
	Handler:         ListenToIndex,
	Description: []string{
		`get records to given index hash when new record set!
		Request Must send to specific node that handle that hash index range!!
		This service has a lot of use cases like:
		- any geospatial usage e.g. tracking device or user, ...
		- following content author like telegram channels or instagram live video!`,
	},
	TAGS: []string{""},
}

// ListenToIndex use to get records to given index hash when new record set!
func ListenToIndex(s *achaemenid.Server, st *achaemenid.Stream) {}

type listenToIndexReq struct {
	IndexHash [32]byte
}

type listenToIndexRes struct {
	// Record []byte TODO::: it can't be simple byte, maybe channel
}

func listenToIndex(st *achaemenid.Stream, req *listenToIndexReq) (res *listenToIndexRes, err error) {
	return res, nil
}

func (req *listenToIndexReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *listenToIndexRes) syllabEncoder(buf []byte) error {
	return nil
}
