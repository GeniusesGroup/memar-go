/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// CLU is default global protocol.Cluster like window global variable in browsers.
// You must assign to it by any object implement protocol.Cluster on your main.go file. Suggestion:
// protocol.STG = &esteghrar.Cluster
var CLU Cluster

// Cluster store application nodes data and ready to return in many ways.
type Cluster interface {
	LocalNode() ApplicationNode
	ReplicatedLocalNode() []ApplicationNode

	GetReplicatedNodes(node ApplicationNode) []ApplicationNode
	GetNodeByID(nodeID NodeID) ApplicationNode
	GetNodeByStorage(mediaTypeID MediaTypeID, uuid UUID) (node ApplicationNode, err Error)
}

