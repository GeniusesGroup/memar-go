/* For license and copyright information please see LEGAL file in repository */

package protocol

// Save in on device node as PrimaryNode() but with PrimaryKey as sha3.256(PrimaryNode, SecondaryNode) in related GraphMediaTypeID.
// - One secondary key can exist to list all SecondaryNode for PrimaryNode. But in some cases it isn't practical e.g. all relation to a city.
// - On some requirements may need more secondary indexes e.g. all "act" relation for an actor to get all movies or vice versa
// But is it true to save these type of data as graph structure? Isn't it so simple to has dedicated media type?
// We think we must graph structure to store time sensitive data that a relation change over times.
// https://en.wikipedia.org/wiki/Graph_database
// https://stackoverflow.com/questions/43712490/time-based-graph-data-modeling
// https://www.researchgate.net/publication/324435978_Graph_based_Platform_for_Electricity_Market_Study_Education_and_Training
// https://github.com/milvus-io/milvus
type Graph interface {
	PrimaryNode() [16]byte
	SecondaryNode() [16]byte
	Relation() GraphRelationLabel // Edge in graph data type
}

type GraphRelationLabel uint32

const (
	GraphRelationLabel_Unset GraphRelationLabel = iota
	GraphRelationLabel_Knowns
	// friend, FriendCircleOne(best-friend), brother, LIVES_IN, NATIONAL_OF, ...
)
