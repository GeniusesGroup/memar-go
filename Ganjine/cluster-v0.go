/* For license and copyright information please see LEGAL file in repository */

package ganjine

// Cluster is a record that store all cluster data, that each node in cluster
// and SDK in authorized app must have full sync of it to work probably!!
type Cluster struct {
	/* Common header data */
	Checksum          [32]byte
	RecordID          [16]byte
	RecordSize        uint64
	RecordStructureID uint64
	WriteTime         int64
	OwnerAppID        [16]byte
	/* Unique data */
	AppID              [16]byte // Unique ID to listen it in any node for new record!
	ClusterCapacity    [16]byte // In bytes, Max 2^128 Byte, Is it enough!!??
	TotalReplications  uint8
	Replications       []ReplicationZone
	TotalNodes         uint32   // we just support max nodes by uint32 limit!
	PrimaryIndexRanges []uint64 // First left 64bit of record ID!
	HashIndexRanges    []uint64 // First left 64bit of hash index!
	TransactionTimeOut uint16   // in ms, default 500ms, Max 65.535s timeout
	NodeFailureTimeOut uint16   // in minute, default 60m, other corresponding node same range will replace failed node! not use in network failure, it is handy proccess!
	OldCluster         *Cluster // Just exist in Re-Allocate proccess like add node or replication!
}

// Init use to initialize a multi-node cluster.
func (c *Cluster) Init(replicationNumber uint8, nodeNumber uint32) {
	// if replicationNumber < 3, warn dev that loose write ability until two replication available again on replication failure!
}

// InitSingleNode use to initialize a single-node cluster!
func (c *Cluster) InitSingleNode() {
	c.Init(1, 1)
}

// Get use to get last version of exiting cluster record if exist
func (c *Cluster) Get(AppID [16]byte) error { return nil }

// GetBySDK use to get cluster and ready it to use by SDK.
func (c *Cluster) GetBySDK() {
	// order Replications by near to far from logic layer!
}

// Set use to write cluster record to storage!
func (c *Cluster) Set() error { return nil }

// ChangeReplicationNumber call it by dev to change replication number!
func (c *Cluster) ChangeReplicationNumber(replicationNumber uint8) {
	// first check exiting number of replications!
	if c.TotalReplications == replicationNumber {
		return
	}
}

// AddNewNode just call from master responsible node to split node range to two or more node!
func (c *Cluster) AddNewNode() {}

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
			median = nodeID + 1  // or high - 1
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

// ReplicationZone :
// Master use to consistency of read and write and detect active master transaction manager
type ReplicationZone struct {
	ReplicationID uint8
	DataCenterID  [16]byte
	Status        uint8  // Master, Read/Write , Read , Stop ,
	Nodes         []Node // sort by index ranges!
}

// Node is an individual machine running PersiaDB
// Same replicated node on other replication must have same StartRange, EndRange and SecondaryIndexRecord!
// Rarely Nodes can add for secondary index and transaction manager!! These types of node can start and shutdown auto!
type Node struct {
	ReplicationID            uint8
	DomainID                 [16]byte // Usually ID of subdomain of platform domain! like node12.db.sabz.city
	SecondaryIndexRecordID   [16]byte
	SplitFrom                uint32 // NodeID, Use in splitting status to response request!
	StartPrimaryIndexRange   [8]byte
	StartSecondaryIndexRange [8]byte
	NodeCapacity             uint64 // In bytes, Max 16EB(Exabyte) that more enough for one node capacity
	AddTime                  uint64
	ExitTime                 uint64
	Status                   uint8 // Stable, Splitting, Re-Allocate , AcceptWrite , Stop, NotResponse,
}
