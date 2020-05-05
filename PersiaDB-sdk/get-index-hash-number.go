/* For license and copyright information please see LEGAL file in repository */

package persiadb

import (
	chaparkhane "../ChaparKhane"
	persiadb "../persiaDB"
)

// GetIndexHashNumberReq is request structure of GetIndexHashNumber()
type GetIndexHashNumberReq struct {
	IndexHash [32]byte
}

// GetIndexHashNumberRes is response structure of GetIndexHashNumber()
type GetIndexHashNumberRes struct {
	RecordNumber uint64
}

// GetIndexHashNumber use to get number of recordsID register for specific IndexHash
func GetIndexHashNumber(s *chaparkhane.Server, c *persiadb.Cluster, req *GetIndexHashNumberReq) (res *GetIndexHashNumberRes, err error) {
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

	// Set GetIndexHashNumber ServiceID
	reqStream.ServiceID = 222077451

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

func (req *GetIndexHashNumberReq) syllabEncoder(buf []byte) error {
	// Index Hash
	copy(buf[:], req.IndexHash[:])

	return nil
}

func (res *GetIndexHashNumberRes) syllabDecoder(buf []byte) error {
	return nil
}
