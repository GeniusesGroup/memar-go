/* For license and copyright information please see LEGAL file in repository */

package persiadb

import (
	chaparkhane "../ChaparKhane"
	persiadb "../persiaDB"
)

// ReadRecordReq is request structure of ReadRecord()
type ReadRecordReq struct {
	RecordID [16]byte
	Offset   uint64 // Do something like block storage API
	Limit    uint64 // Do something like block storage API
}

// ReadRecordRes is response structure of ReadRecord()
type ReadRecordRes struct {
	Data []byte
}

// ReadRecord use read some part of the specific record by its ID!
func ReadRecord(s *chaparkhane.Server, c *persiadb.Cluster, req *ReadRecordReq) (res *ReadRecordRes, err error) {
	var nodeID uint32 = c.FindNodeIDByRecordID(req.RecordID)

	var ok bool
	var i uint8
	var conn *chaparkhane.Connection
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
	var reqStream, resStream *chaparkhane.Stream
	reqStream, resStream = conn.MakeBidirectionalStream(0)

	// Set ReadRecord ServiceID
	reqStream.ServiceID = 108857663

	req.syllabEncoder(reqStream.Payload[4:])
	err = reqStream.SrpcOutcomeRequestHandler(s)
	if err == nil {
		return nil, err
	}

	// Listen to response stream and decode error ID and return it to caller
	var responseStatus = <-resStream.Status
	if responseStatus == chaparkhane.StreamStateReady {
	}

	res.syllabDecoder(resStream.Payload[4:])

	return res, resStream.Err
}

func (req *ReadRecordReq) syllabEncoder(buf []byte) error {
	copy(buf[:], req.RecordID[:])
	return nil
}

func (res *ReadRecordRes) syllabDecoder(buf []byte) error {
	return nil
}
