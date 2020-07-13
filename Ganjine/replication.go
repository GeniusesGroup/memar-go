/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
)

type replication struct {
	ReplicationID uint8
	State         uint8
	Nodes         []Node // sort by index ranges!
}

// Replication State
const (
	// Master use to consistency of read and write and detect active master transaction manager
	ReplicationStateMaster uint8 = iota
	ReplicationStateReadWrite
	ReplicationStateRead
	ReplicationStateStop
)

// GetNodeByIndexHash returns the node have desire index in best replication.
func (c *Cluster) GetNodeByIndexHash(indexHash [32]byte) (node *Node) {
	var nodeID uint32 = c.FindNodeIDByIndexHash(indexHash)

	var i uint8
	var conn *achaemenid.Connection
	// Maybe closest Ganjine node not response recently, so check all replications
	for i = 0; i < c.TotalReplications; i++ {
		conn = c.Replications[i].Nodes[nodeID].GetConnection()
		if conn.State == achaemenid.ConnectionStateOpen {
			return &c.Replications[i].Nodes[nodeID]
		}
	}

	return
}

// GetNodeByRecordID returns the node have desire record in best replication.
func (c *Cluster) GetNodeByRecordID(recordID [16]byte) (node *Node) {
	var nodeID uint32 = c.FindNodeIDByRecordID(recordID)

	var i uint8
	var conn *achaemenid.Connection
	// Indicate conn! Maybe closest Ganjine node not response recently
	for i = 0; i < c.TotalReplications; i++ {
		conn = c.Replications[i].Nodes[nodeID].GetConnection()
		if conn != nil {
			return &c.Replications[i].Nodes[nodeID]
		}
	}

	return
}

// GetNodeByReplicationID returns the node in desire replication.
func (c *Cluster) GetNodeByReplicationID(repID uint8, nodeLoc uint32) (node *Node) {
	return &c.Replications[repID].Nodes[nodeLoc]
}
