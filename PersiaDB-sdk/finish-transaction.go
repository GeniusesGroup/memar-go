/* For license and copyright information please see LEGAL file in repository */

package persiadb

import (
	chaparkhane "../ChaparKhane"
	persiadb "../persiaDB"
)

// FinishTransactionReq is request structure of FinishTransaction()
type FinishTransactionReq struct {
	IndexHash [32]byte
	Record    []byte
}

// FinishTransaction use to approve transaction!
// Transaction Manager will set record and index! no further action need after this call!
func FinishTransaction(s *chaparkhane.Server, c *persiadb.Cluster, req *FinishTransactionReq) (err error) {
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
	// Set DeleteIndexHashHistory ServiceID
	reqStream.ServiceID = 3962420401

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

func (req *FinishTransactionReq) syllabEncoder(buf []byte) error {
	copy(buf[:], req.IndexHash[:])
	copy(buf[32:], req.Record[:])
	return nil
}
