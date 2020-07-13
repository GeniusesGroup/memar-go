/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// FindRecordsReq is request structure of FindRecords()
type FindRecordsReq struct {
	IndexHash [32]byte
	Offset    uint64
	Limit     uint64 // It is better to be modulus of 64 or even 256 if storage devices use 4K clusters!
}

// FindRecordsRes is response structure of FindRecords()
type FindRecordsRes struct {
	RecordIDs [][16]byte
}

// FindRecords will get related RecordsID that set to given indexHash before!
// get 64 related index to given IndexHash even if just one of them use!
func FindRecords(c *ganjine.Cluster, req *FindRecordsReq) (res *FindRecordsRes, err error) {
	var node *ganjine.Node = c.GetNodeByIndexHash(req.IndexHash)
	if node == nil {
		return nil, ErrNoNodeAvailableToHandleRequests
	}

	// Make new request-response streams
	var conn *achaemenid.Connection = node.GetConnection()
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = conn.MakeBidirectionalStream(0)
	if err != nil {
		return nil, err
	}
	
	// Set FindRecords ServiceID
	reqStream.ServiceID = 1992558377
	reqStream.Payload = req.syllabEncoder()

	err = node.SendStream(reqStream)
	if err != nil {
		return nil, err
	}

	res = &FindRecordsRes{}
	err = res.syllabDecoder(resStream.Payload[4:])
	if err != nil {
		return nil, err
	}

	return res, resStream.Err
}

func (req *FindRecordsReq) syllabEncoder() (buf []byte) {
	buf = make([]byte, 48+4) // +4 for sRPC ID instead get offset argument

	// Index Hash
	copy(buf[4:], req.IndexHash[:])
	// Offset
	buf[36] = byte(req.Offset)
	buf[37] = byte(req.Offset >> 8)
	buf[38] = byte(req.Offset >> 16)
	buf[39] = byte(req.Offset >> 24)
	buf[40] = byte(req.Offset >> 32)
	buf[41] = byte(req.Offset >> 40)
	buf[42] = byte(req.Offset >> 48)
	buf[43] = byte(req.Offset >> 56)
	// Limit
	buf[44] = byte(req.Limit)
	buf[45] = byte(req.Limit >> 8)
	buf[46] = byte(req.Limit >> 16)
	buf[47] = byte(req.Limit >> 24)
	buf[48] = byte(req.Limit >> 32)
	buf[49] = byte(req.Limit >> 40)
	buf[50] = byte(req.Limit >> 48)
	buf[51] = byte(req.Limit >> 56)

	return nil
}

func (res *FindRecordsRes) syllabDecoder(buf []byte) error {
	return nil
}
