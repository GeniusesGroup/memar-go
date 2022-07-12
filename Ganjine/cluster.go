/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../object"
	"../protocol"
)

// Cluster is the base datastore object that use by other part of app and platforms!
// Cluster implements protocol.ObjectDirectory && protocol.
// Each node in cluster and SDK in authorized app must have it to work probably.
// First node of cluster is cordiantor of cluster.
type Cluster struct {
	State int

	objects object.Directory        // Distributed object storage
	files   protocol.FileDirectory  // Distributed file storage
	records protocol.StorageRecords // Distributed records storage

	TransactionManager TransactionManager
}

// Cluster State
const (
	ClusterStateStop int = iota
	ClusterStateRunning
	ClusterStateStopping
	ClusterStateStarting // plan to start
)

func (c *Cluster) LocalObjects() protocol.StorageObjects      { return }
func (c *Cluster) LocalObjectsCache() protocol.StorageObjects { return }
func (c *Cluster) LocalFiles() protocol.FileDirectory         { return }
func (c *Cluster) LocalRecords() protocol.StorageRecords      { return }
func (c *Cluster) LocalRecordsCache() protocol.StorageRecords { return }

func (c *Cluster) Objects() protocol.Objects        { return &c.objects }
func (c *Cluster) Files() protocol.FileDirectory    { return c.files }
func (c *Cluster) Records() protocol.StorageRecords { return c.records }

// Init initialize an exiting cluster to get or make a cluster!
func (c *Cluster) Init() (err protocol.Error) {
	registerServices()
	object.Init()

	c.TransactionManager.init()

	return
}

// Shutdown the cluster!
func (c *Cluster) Shutdown() (err protocol.Error) {
	// Close Indexes
	// Notify other nodes
	return
}
