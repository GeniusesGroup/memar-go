/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../achaemenid"
	"../log"
)

// Cluster is a record that store all cluster data, that each node in cluster
// and SDK in authorized app must have full sync of it to work probably!!
type Cluster struct {
	ClusterCapacity    [16]byte // In bytes, Max 2^128 Byte, Is it enough!!??
	TotalReplications  uint8
	Replications       []replication
	TotalNodes         uint32   // not count replicated nodes, just one of them count.
	PrimaryIndexRanges []uint64 // First left 64bit of record ID!
	HashIndexRanges    []uint64 // First left 64bit of hash index!
	OldCluster         *Cluster // Just exist in Re-Allocate proccess like add node or replication!
	Server             *achaemenid.Server
}

// Init use to initialize a nil cluster to get or make a cluster!
func (c *Cluster) Init(s *achaemenid.Server) (err error) {
	c = &Cluster{
		Replications: make([]replication, s.Manifest.TechnicalInfo.ReplicationNumber),
		Server:       s,
	}

	if s.Manifest.TechnicalInfo.ReplicationNumber < 3 {
		log.Warn("ReplicationNumber set below 3! Loose write ability until two replication available again on replication failure!")
	}

	// s.Manifest.TechnicalInfo.ReplicationNumber
	// s.Manifest.TechnicalInfo.NodeNumber

	// TODO::: order Replications by near to far from logic layer!

	return
}
