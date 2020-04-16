/* For license and copyright information please see LEGAL file in repository */

package services

// This service has a lot of use cases like
// - any geospatial usage e.g. tracking device or user, ...
// - following content author like telegram channels or instagram live video!

import chaparkhane "../ChaparKhane"

var listenToIndexService = chaparkhane.Service{
	Name:            "ListenToIndex",
	IssueDate:       0,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          chaparkhane.ServiceStatePreAlpha,
	Handler:         ListenToIndex,
	Description: []string{
		"",
	},
	TAGS: []string{""},
}

type listenToIndexReq struct{}

type listenToIndexRes struct{}

func listenToIndex(st *chaparkhane.Stream, req *listenToIndexReq) (res *listenToIndexRes, err error) {
	return res, nil
}

// ListenToIndex use to get the recordID by index hash when new record set!
// Must send this request to specific node that handle that range!!
func ListenToIndex(s *chaparkhane.Server, st *chaparkhane.Stream) {}

func (req *listenToIndexReq) validator() error {
	return nil
}

func (req *listenToIndexReq) syllabDecoder(buf []byte) error {
	return nil
}

func (res *listenToIndexRes) syllabEncoder(buf []byte) error {
	return nil
}
