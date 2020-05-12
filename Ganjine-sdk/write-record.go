/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// WriteRecordReq is request structure of WriteRecord()
type WriteRecordReq struct {
	RecordID [16]byte
	Offset   uint64 // Do something like block storage API
	Limit    uint64 // Do something like block storage API
	Data     []byte
}

// WriteRecord use to write some part of a record!
// Don't use this service until you force to use!
func WriteRecord(s *achaemenid.Server, c *ganjine.Cluster, req *WriteRecordReq) (err error) {
	var nodeID uint32 = c.FindNodeIDByRecordID(req.RecordID)

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
		return err
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = conn.MakeBidirectionalStream(0)

	// Set WriteRecord ServiceID
	reqStream.ServiceID = 3836795965

	req.syllabEncoder(reqStream.Payload[4:])
	err = reqStream.SrpcOutcomeRequestHandler(s)
	if err == nil {
		return err
	}

	// Listen to response stream and decode error ID and return it to caller
	var responseStatus uint8 = <-resStream.StatusChannel
	if responseStatus == achaemenid.StreamStateReady {
	}

	return resStream.Err
}

func (req *WriteRecordReq) syllabEncoder(buf []byte) error {
	// RecordID
	copy(buf[:], req.RecordID[:])
	// Offset
	buf[16] = byte(req.Offset)
	buf[17] = byte(req.Offset >> 8)
	buf[18] = byte(req.Offset >> 16)
	buf[19] = byte(req.Offset >> 24)
	buf[20] = byte(req.Offset >> 32)
	buf[21] = byte(req.Offset >> 40)
	buf[22] = byte(req.Offset >> 48)
	buf[23] = byte(req.Offset >> 56)
	// Limit
	buf[24] = byte(req.Limit)
	buf[25] = byte(req.Limit >> 8)
	buf[26] = byte(req.Limit >> 16)
	buf[27] = byte(req.Limit >> 24)
	buf[28] = byte(req.Limit >> 32)
	buf[29] = byte(req.Limit >> 40)
	buf[30] = byte(req.Limit >> 48)
	buf[31] = byte(req.Limit >> 56)
	// Data
	copy(buf[32:], req.Data[:])
	return nil
}
