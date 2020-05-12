/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// DeleteRecordReq is request structure of DeleteRecord()
type DeleteRecordReq struct {
	RecordID [16]byte
}

// DeleteRecord use to delete specific record by given ID in all cluster!
// We don't suggest use this service, due to we strongly suggest think about data as immutable entity(stream and time)
// It won't delete record history or indexes associate to it!
func DeleteRecord(s *achaemenid.Server, c *ganjine.Cluster, req *DeleteRecordReq) (err error) {
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
	
	// Set DeleteRecord ServiceID
	reqStream.ServiceID = 1758631843

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

func (req *DeleteRecordReq) syllabEncoder(buf []byte) {
	copy(buf[:], req.RecordID[:])
}
