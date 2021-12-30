/* For license and copyright information please see LEGAL file in repository */

package protocol

// Index store index data for data storage structure.
// When "For" point to PrimaryKey it means "Index" and when point to other fields means "List"
type Index struct {
	How      IndexType
	What     string   // field name
	For      string   // field name
	Pair     []string // to have composite index. other filed names that will combine with "For"
	If       string   // field name that must not nul or empty or zero
	Interval string   // Usually daily but can be hourly or minutely
}

type IndexType uint8

const (
	IndexType_Unset           IndexType = iota
	IndexType_PrimaryKey                // auto select OneToOne type by storage engine e.g. hash, btree,...
	IndexType_Unique                    // auto select OneToOne type by storage engine e.g. hash, btree,...
	IndexType_List                      // auto select OneToMany type by storage engine e.g. hash, btree,...
	IndexType_Hash_OneToOne             // unique key to unique value
	IndexType_Hash_OneToMany            // unique key to multi value
	IndexType_BTree_OneToOne            // unique key to unique value
	IndexType_BTree_OneToMany           // unique key to multi value
	IndexType_Text
	IndexType_Spatial // maps, ...
	IndexType_Vector  // N-dimensional space use in facial recognition (min euclidean-distance),	https://en.wikipedia.org/wiki/Locality-sensitive_hashing
	// ...
)
