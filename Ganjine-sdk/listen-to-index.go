/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../ganjine"
)

// listenToIndexReq is request structure of listenToIndex()
type listenToIndexReq struct {
	IndexHash [32]byte
}

// listenToIndexRes is response structure of listenToIndex()
type listenToIndexRes struct {
	// Record []byte TODO::: it can't be simple byte, maybe channel
}

// listenToIndex use to get the recordID by index hash when new record set!
// Must send this request to specific node that handle that range!!
func listenToIndex(s *achaemenid.Server, c *ganjine.Cluster, req *listenToIndexReq) (res *listenToIndexRes, err error) {
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

	// Set listenToIndex ServiceID
	reqStream.ServiceID = 2145882122

	req.syllabEncoder(reqStream.Payload[4:])
	err = reqStream.SrpcOutcomeRequestHandler(s)
	if err == nil {
		return nil, err
	}

	// Listen to response stream and decode error ID and return it to caller
	var responseStatus uint8 = <-resStream.StatusChannel
	if responseStatus == achaemenid.StreamStateReady {
	}

	// Sender can reuse exiting stream to send new record

	res.syllabDecoder(resStream.Payload[4:])

	return res, resStream.Err
}

func (req *listenToIndexReq) syllabEncoder(buf []byte) error {
	copy(buf[:], req.IndexHash[:])
	return nil
}

func (res *listenToIndexRes) syllabDecoder(buf []byte) error {
	return nil
}
