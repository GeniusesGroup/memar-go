/* For license and copyright information please see LEGAL file in repository */

package protocol

type Index struct {
	How  IndexType
	What string   // field name
	For  string   // field name
	Pair []string // to have composite index. other filed names that will combine with For
	If   string   // field name that must not nul or empty or zero
	When string   // e.g. daily
}

type IndexType uint8

const (
	IndexType_Unset IndexType = iota
	IndexType_List
	IndexType_Hash_OneToOne   // unique key to unique value
	IndexType_Hash_OneToMany  // unique key to multi value
	IndexType_BTree_OneToOne  // unique key to unique value
	IndexType_BTree_OneToMany // unique key to multi value
	IndexType_Text
	IndexType_Spatial // maps, ...
	IndexType_Vector  // N-dimensional space use in facial recognition (min euclidean-distance),	https://en.wikipedia.org/wiki/Locality-sensitive_hashing
	// ...
)
