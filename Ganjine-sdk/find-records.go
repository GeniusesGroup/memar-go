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
func FindRecords(s *achaemenid.Server, c *ganjine.Cluster, req *FindRecordsReq) (res *FindRecordsRes, err error) {
	var nodeID uint32 = c.FindNodeIDByIndexHash(req.IndexHash)

	var ok bool
	var i uint8
	var conn *achaemenid.Connection
	// Indicate conn! Maybe closest PersiaDB node not response recently
	for i = 0; i < c.TotalReplications; i++ {
		var domainID = c.Replications[i].Nodes[nodeID].DomainID
		conn, ok = s.Connections.GetConnectionByDomainID(domainID)
		if !ok {
			conn, err = s.Connections.MakeNewConnectionByDomainID(domainID)
			if err == nil {
				break
			}
		} else {
			break
		}
	}

	// Check if no connection exist to use!
	if conn == nil {
		return nil, err
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = conn.MakeBidirectionalStream(0)
	// Set DeleteIndexHashHistory ServiceID
	reqStream.ServiceID = 1992558377

	req.syllabEncoder(reqStream.Payload[4:])
	err = reqStream.SrpcOutcomeRequestHandler(s)
	if err == nil {
		return nil, err
	}

	// Listen to response stream and decode error ID and return it to caller
	var responseStatus uint8 = <-resStream.StatusChannel
	if responseStatus == achaemenid.StreamStateReady {
	}

	res.syllabDecoder(resStream.Payload[4:])

	return res, resStream.Err
}

func (req *FindRecordsReq) syllabEncoder(buf []byte) error {
	// Index Hash
	copy(buf[:], req.IndexHash[:])
	// Offset
	buf[32] = byte(req.Offset)
	buf[33] = byte(req.Offset >> 8)
	buf[34] = byte(req.Offset >> 16)
	buf[35] = byte(req.Offset >> 24)
	buf[36] = byte(req.Offset >> 32)
	buf[37] = byte(req.Offset >> 40)
	buf[38] = byte(req.Offset >> 48)
	buf[39] = byte(req.Offset >> 56)
	// Limit
	buf[40] = byte(req.Limit)
	buf[41] = byte(req.Limit >> 8)
	buf[42] = byte(req.Limit >> 16)
	buf[43] = byte(req.Limit >> 24)
	buf[44] = byte(req.Limit >> 32)
	buf[45] = byte(req.Limit >> 40)
	buf[46] = byte(req.Limit >> 48)
	buf[47] = byte(req.Limit >> 56)

	return nil
}

func (res *FindRecordsRes) syllabDecoder(buf []byte) error {
	return nil
}
