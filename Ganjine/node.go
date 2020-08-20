/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
)

// Node is an individual machine running Ganjine
// Same replicated node on other replication must have same StartRange, EndRange and HashIndexRecordID!
type Node struct {
	achaemenid.Node
	ID                     uint32 // this ID locate node in replication
	replicationID          uint8
	splitFrom              uint32 // NodeID, Use in splitting status to response request!
	primaryIndexStartRange uint64
	State                  uint8
}

// Node State
const (
	NodeStateSplitting uint8 = iota
	NodeStateReAllocate
	NodeStateAcceptWrite
)

func (n *Node) getNodeDetails(s *achaemenid.Server, node *achaemenid.Node) (err error) {
	// Check if desire node is not a ganjine node!
	if node.Conn == nil {
		// err = 
		return
	}

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = node.Conn.MakeBidirectionalStream(0)
	if err != nil {
		return err
	}

	// Set GetNodeDetails ServiceID
	reqStream.ServiceID = 3707636027

	err = achaemenid.SrpcOutcomeRequestHandler(s, reqStream)
	if err != nil {
		return err
	}

	err = n.SyllabDecoder(resStream.Payload[4:])
	if err != nil {
		return err
	}

	return resStream.Err
}

// SyllabDecoder decode from buf to req
func (n *Node) SyllabDecoder(buf []byte) (err error) {
	return
}

// SyllabEncoder encode req to buf
func (n *Node) SyllabEncoder() (buf []byte) {
	return
}

// GetLocalNodeDetail return local ganjine node detail or related error.
func (c *Cluster) GetLocalNodeDetail() (res *Node, err error) {
	if c.Node == nil {
		err = ErrNodeNotGanjineNode
	} else {
		res = c.Node
	}
	return
}

// FindNodeIDByRecordID use to find responsible node ID for given record node part!
// Nodes in each replication store in sort so nodeID is array location of desire node in any replication!
func (c *Cluster) FindNodeIDByRecordID(recordID [32]byte) (nodeID uint32) {
	var recordNodePart uint64 = uint64(recordID[0]) | uint64(recordID[1])<<8 | uint64(recordID[2])<<16 | uint64(recordID[3])<<24 |
		uint64(recordID[4])<<32 | uint64(recordID[5])<<40 | uint64(recordID[6])<<48 | uint64(recordID[7])<<56

	if c.Replications.TotalNodes == 1 {
		// Due to nodeID == 0, Don't need to assign it again!!
		return
	}

	var primaryIndexRanges = c.PrimaryIndexRanges
	var high uint32 = c.Replications.TotalNodes - 1
	var median uint32
	var diff uint32
	for nodeID < high {
		diff = high - nodeID
		if diff < 3 {
			median = nodeID + 1 // or high - 1
			if primaryIndexRanges[high] <= recordNodePart {
				nodeID = high
			} else if primaryIndexRanges[median] <= recordNodePart {
				nodeID = median
			}
			break
		} else {
			median = nodeID + diff/2
			if primaryIndexRanges[median] < recordNodePart {
				nodeID = median
			} else {
				high = median
			}
		}
	}

	return
}
