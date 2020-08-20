/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../log"
)

// Cluster store all cluster data, that each node in cluster
// and SDK in authorized app must have full sync of it to work probably!!
type Cluster struct {
	State              int // States locate in const of this file.
	Capacity           [16]byte // In bytes, Max 2^128 Byte, Is it enough!!??
	Manifest           Manifest
	Replications       replications
	Node               *Node    // This active node, nil if node just be service node!
	PrimaryIndexRanges []uint64 // First left 64bit of record ID!
	TransactionManager TransactionManager
	OldCluster         *Cluster // Just exist in Re-Allocate proccess like add node or replication zone!
	Server             *achaemenid.Server
}

// Server State
const (
	ClusterStateStop int = iota
	ClusterStateRunning
	ClusterStateStopping
	ClusterStateStarting // plan to start
)

// Init initialize an exiting cluster to get or make a cluster!
func (c *Cluster) Init(s *achaemenid.Server) (err error) {
	c.Replications.TotalZones = s.Manifest.TechnicalInfo.ReplicationNumber
	c.Replications.Zones = make([]replication, c.Replications.TotalZones)
	c.Server = s

	if c.Replications.TotalZones < 3 {
		log.Warn("ReplicationNumber set below 3! Loose write ability until two replication available again on replication failure!")
	}

	var ln = len(s.Nodes.Nodes)
	for i := 0; i < ln; i++ {
		go c.addNode(&s.Nodes.Nodes[i])
	}

	go c.orderZones()

	c.TransactionManager.init()

	return
}

// Shutdown the cluster!
func (c *Cluster) Shutdown(s *achaemenid.Server) (err error) {
	// Close Indexes
	// Notify other nodes
	return
}
