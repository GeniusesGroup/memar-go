/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// FindRecordsConsistentlyReq is request structure of FindRecordsConsistently()
type FindRecordsConsistentlyReq struct {
	IndexHash [32]byte
}

// FindRecordsConsistentlyRes is response structure of FindRecordsConsistently()
type FindRecordsConsistentlyRes struct {
	RecordID [][16]byte
}

// FindRecordsConsistently use to find records by indexes that store before in consistently!
// It will get index from transaction managers not indexes nodes!
func FindRecordsConsistently(s *achaemenid.Server, c *ganjine.Cluster, req *FindRecordsConsistentlyReq) (res *FindRecordsConsistentlyRes, err error) {
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

	// Set FindRecordsConsistently ServiceID
	reqStream.ServiceID = 480215407

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

func (req *FindRecordsConsistentlyReq) syllabEncoder(buf []byte) error {
	copy(buf[:], req.IndexHash[:])
	return nil
}

func (res *FindRecordsConsistentlyRes) syllabDecoder(buf []byte) error {
	return nil
}
