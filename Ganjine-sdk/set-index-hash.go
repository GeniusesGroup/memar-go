/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// SetIndexHashReq is request structure of SetIndexHash()
type SetIndexHashReq struct {
	IndexHash [32]byte
	RecordID  [16]byte
}

// SetIndexHash use to set a record ID to new||exiting index hash!
func SetIndexHash(s *achaemenid.Server, c *ganjine.Cluster, req *SetIndexHashReq) (err error) {
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
		return err
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = conn.MakeBidirectionalStream(0)

	// Set SetIndexHash ServiceID
	reqStream.ServiceID = 1881585857

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

func (req *SetIndexHashReq) syllabEncoder(buf []byte) error {
	copy(buf[:], req.IndexHash[:])
	copy(buf[32:], req.RecordID[:])
	return nil
}
