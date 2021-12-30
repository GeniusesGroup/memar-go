/* For license and copyright information please see LEGAL file in repository */

package protocol

// Cluster store application nodes data and ready to return in many ways.
type Cluster interface {
	LocalNode() ApplicationNode
	ReplicatedLocalNode() []ApplicationNode

	GetReplicatedNodes(node ApplicationNode) []ApplicationNode
	GetNodeByID(nodeID uint64) ApplicationNode
	GetNodeByObjectID(id [32]byte) (node ApplicationNode, err Error)
}

type ApplicationNode interface {
	Type() NodeType
	ID() [16]byte           // as InstanceID
	DataCenterID() [16]byte // orgID
	ThingID() [16]byte
	ReplicationID() uint8
	StorageCapacity() uint64 // In bytes, Max 16EB(Exabyte) that more enough for one node capacity. 0 means service only node.
	Conn() Connection
	Status() ApplicationState
	Stats()
}

type NodeType uint8

const (
	Node_Unset NodeType = iota
	Node_Cloud
	Node_Edge
)
