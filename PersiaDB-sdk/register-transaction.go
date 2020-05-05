/* For license and copyright information please see LEGAL file in repository */

package persiadb

import (
	chaparkhane "../ChaparKhane"
	persiadb "../persiaDB"
)

// RegisterTransactionReq is request structure of RegisterTransaction()
type RegisterTransactionReq struct {
	IndexHash [32]byte
	RecordID  [16]byte
}

// RegisterTransactionRes is response structure of RegisterTransaction()
type RegisterTransactionRes struct {
	Record []byte
}

// RegisterTransaction use read some part of the specific record by its ID!
func RegisterTransaction(s *chaparkhane.Server, c *persiadb.Cluster, req *RegisterTransactionReq) (res *RegisterTransactionRes, err error) {
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
		return nil, err
	}

	// Make new request-response streams
	var reqStream, resStream *chaparkhane.Stream
	reqStream, resStream, err = conn.MakeBidirectionalStream(0)

	// Set RegisterTransaction ServiceID
	reqStream.ServiceID = 3840530512

	req.syllabEncoder(reqStream.Payload[4:])
	err = reqStream.SrpcOutcomeRequestHandler(s)
	if err == nil {
		return nil, err
	}

	// Listen to response stream and decode error ID and return it to caller
	var responseStatus uint8 = <-resStream.StatusChannel
	if responseStatus == chaparkhane.StreamStateReady {
	}

	res.syllabDecoder(resStream.Payload[4:])

	return res, resStream.Err
}

func (req *RegisterTransactionReq) syllabEncoder(buf []byte) error {
	copy(buf[:], req.IndexHash[:])
	copy(buf[32:], req.RecordID[:])
	return nil
}

func (res *RegisterTransactionRes) syllabDecoder(buf []byte) error {
	return nil
}
