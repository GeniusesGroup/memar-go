/* For license and copyright information please see LEGAL file in repository */

package persiadb

import (
	chaparkhane "../ChaparKhane"
	persiadb "../persiaDB"
)

// DeleteIndexHashHistoryReq is request structure of DeleteIndexHashHistory()
type DeleteIndexHashHistoryReq struct {
	IndexHash [32]byte
}

// DeleteIndexHashHistory use to delete all record associate to given index and delete index itself!
func DeleteIndexHashHistory(s *chaparkhane.Server, c *persiadb.Cluster, req *DeleteIndexHashHistoryReq) (err error) {
	var nodeID uint32 = c.FindNodeIDByIndexHash(req.IndexHash)

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
		return err
	}

	// Make new request-response streams
	var reqStream, resStream *chaparkhane.Stream
	reqStream, resStream = conn.MakeBidirectionalStream(0)

	// Set DeleteIndexHashHistory ServiceID
	reqStream.ServiceID = 691384835

	req.syllabEncoder(reqStream.Payload[4:])
	err = reqStream.SrpcOutcomeRequestHandler(s)
	if err == nil {
		return err
	}

	// Listen to response stream and decode error ID and return it to caller
	var responseStatus = <-resStream.Status
	if responseStatus == chaparkhane.StreamStateReady {
	}

	return resStream.Err
}

func (req *DeleteIndexHashHistoryReq) syllabEncoder(buf []byte) {
	copy(buf[:], req.IndexHash[:])
}
