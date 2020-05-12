/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// SetRecordReq is request structure of SetRecord()
type SetRecordReq struct {
	RecordID [16]byte
	Record   []byte
}

// SetRecord use to respect all data in record and don't change something like RecordID or WriteTime!
// If data like OwnerAppID is wrong you can't get record anymore!
func SetRecord(s *achaemenid.Server, c *ganjine.Cluster, req *SetRecordReq) (err error) {
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

	// Set SetRecord ServiceID
	reqStream.ServiceID = 10488062

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

func (req *SetRecordReq) syllabEncoder(buf []byte) error {
	// Don't need to include RecordID! we just get it from upper due to Go is strongly type
	// and we don't want to use unsafe here in SDK!
	req.Record = buf
	return nil
}
