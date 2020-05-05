/* For license and copyright information please see LEGAL file in repository */

package persiadb

import (
	chaparkhane "../ChaparKhane"
	persiadb "../persiaDB"
)

// DeleteIndexHashReq is request structure of DeleteIndexHash()
type DeleteIndexHashReq struct {
	IndexHash [32]byte
}

// DeleteIndexHash use to delete exiting index hash with all related records IDs!
// It wouldn't delete related records! Use DeleteIndexHistory() instead if you want delete all records too!
func DeleteIndexHash(s *chaparkhane.Server, c *persiadb.Cluster, req *DeleteIndexHashReq) (err error) {
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
	reqStream, resStream, err = conn.MakeBidirectionalStream(0)

	// Set DeleteIndexHash ServiceID
	reqStream.ServiceID = 3411747355

	req.syllabEncoder(reqStream.Payload[4:])
	err = reqStream.SrpcOutcomeRequestHandler(s)
	if err == nil {
		return err
	}

	// Listen to response stream and decode error ID and return it to caller
	var responseStatus uint8 = <-resStream.StatusChannel
	if responseStatus == chaparkhane.StreamStateReady {
	}

	return resStream.Err
}

func (req *DeleteIndexHashReq) syllabEncoder(buf []byte) {
	copy(buf[:], req.IndexHash[:])
}
