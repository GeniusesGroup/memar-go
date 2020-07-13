/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
)

// Node is an individual machine running Ganjine
// Same replicated node on other replication must have same StartRange, EndRange and HashIndexRecordID!
type Node struct {
	achaemenid.Node
	replicationID          uint8
	hashIndexRecordID      [16]byte
	splitFrom              uint32 // NodeID, Use in splitting status to response request!
	primaryIndexStartRange uint64
	hashIndexStartRange    uint64
	state                  uint8
}

// Node State
const (
	NodeStateSplitting uint8 = iota
	NodeStateReAllocate
	NodeStateAcceptWrite
)

// FindNodeIDByRecordID use to find responsible node ID for given record node part!
// Nodes in each replication store in sort so nodeID is array location of desire node in any replication!
func (c *Cluster) FindNodeIDByRecordID(recordID [16]byte) (nodeID uint32) {
	var recordNodePart uint64 = uint64(recordID[0]) | uint64(recordID[1])<<8 | uint64(recordID[2])<<16 | uint64(recordID[3])<<24 |
		uint64(recordID[4])<<32 | uint64(recordID[5])<<40 | uint64(recordID[6])<<48 | uint64(recordID[7])<<56

	if c.TotalNodes == 1 {
		// Due to nodeID == 0, Don't need to assign it again!!
		return
	}

	var high uint32 = c.TotalNodes - 1
	var median uint32
	var diff uint32
	for nodeID < high {
		diff = high - nodeID
		if diff < 3 {
			median = nodeID + 1 // or high - 1
			if c.PrimaryIndexRanges[high] <= recordNodePart {
				nodeID = high
			} else if c.PrimaryIndexRanges[median] <= recordNodePart {
				nodeID = median
			}
			break
		} else {
			median = nodeID + diff/2
			if c.PrimaryIndexRanges[median] < recordNodePart {
				nodeID = median
			} else {
				high = median
			}
		}
	}

	return
}

// FindNodeIDByIndexHash use to find responsible node ID for given index hash!
// Nodes in each replication store in sort so nodeID is array location of desire node in any replication!
func (c *Cluster) FindNodeIDByIndexHash(indexHash [32]byte) (nodeID uint32) {
	var indexNodePart uint64 = uint64(indexHash[0]) | uint64(indexHash[1])<<8 | uint64(indexHash[2])<<16 | uint64(indexHash[3])<<24 |
		uint64(indexHash[4])<<32 | uint64(indexHash[5])<<40 | uint64(indexHash[6])<<48 | uint64(indexHash[7])<<56

	if c.TotalNodes == 1 {
		// Due to nodeID == 0, Don't need to assign it again!!
		return
	}

	var high uint32 = c.TotalNodes - 1
	var median uint32
	var diff uint32
	for nodeID < high {
		diff = high - nodeID
		if diff < 3 {
			median = nodeID + 1 // or high - 1
			if c.HashIndexRanges[high] <= indexNodePart {
				nodeID = high
			} else if c.HashIndexRanges[median] <= indexNodePart {
				nodeID = median
			}
			break
		} else {
			median = nodeID + diff/2
			if c.HashIndexRanges[median] < indexNodePart {
				nodeID = median
			} else {
				high = median
			}
		}
	}

	return
}
