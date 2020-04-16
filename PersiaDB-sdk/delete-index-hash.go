/* For license and copyright information please see LEGAL file in repository */

package persiadb

import (
	chaparkhane "../ChaparKhane"
	persiadb "../persiaDB"
	srpc "../srpc"
)

// DeleteIndexHashReq is request structure of DeleteIndexHash()
type DeleteIndexHashReq struct {
	IndexHash [32]byte
}

// DeleteIndexHash use to delete exiting index hash with all related records IDs!
func DeleteIndexHash(s *chaparkhane.Server, c *persiadb.Cluster, req *DeleteIndexHashReq) (err error) {
	// Make new request stream
	var st = chaparkhane.NewStream()

	// Set DeleteIndexHash ServiceID
	srpc.SetID(st.Payload, 0)

	req.syllabEecoder(st.Payload[4:])

	var nodeID = c.FindIndexNodeID(req.IndexHash)

	var ok bool
	var i uint8
	// Maybe closest PersiaDB node not response recently
	for i = 0; i < c.TotalReplications; i++ {
		st.Connection, ok = s.Connections.GetConnectionByDomainID(c.Replications[i].Nodes[nodeID].DomainID)
		if !ok {
			// TODO : Ask to make the connection!
		}
		err = s.WorkerPool.RegisterStreamToSend(st)
		if err == nil {
			break
		}
	}

	// TODO : Listen to response stream and decode error ID and return it to caller

	return nil
}

func (req *DeleteIndexHashReq) syllabEecoder(buf []byte) {
	copy(buf[:], req.IndexHash[:])
}
