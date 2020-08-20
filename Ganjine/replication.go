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

// GetNodeByRecordID returns the node have desire record in best replication.
func (c *Cluster) GetNodeByRecordID(recordID [32]byte) (node *Node) {
	var nodeID uint32 = c.FindNodeIDByRecordID(recordID)

	var i uint8
	for i = 0; i < c.Replications.TotalZones; i++ {
		// Indicate conn! Maybe closest Ganjine node not response recently
		if c.Replications.Zones[i].Nodes[nodeID].Conn.State == achaemenid.ConnectionStateOpen {
			return &c.Replications.Zones[i].Nodes[nodeID]
		}
	}

	return
}

func (c *Cluster) addNode(an *achaemenid.Node) {
	var err error
	var n = Node{
		Node: *an,
	}

	if an.Conn == nil {
		c.Node = &n
	} else {
		err = n.getNodeDetails(c.Server, an)
		if err != nil {

		}

		// Add n.primaryIndexStartRange to c.PrimaryIndexRanges
		// Add n.hashIndexStartRange to c.HashIndexRanges
	}

	// Update total node
	c.Replications.TotalNodes = uint32(len(c.Replications.Zones[0].Nodes))
}
