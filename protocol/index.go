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
	IndexTypeList IndexType = iota
	IndexTypeHash
	IndexTypeUniqueHash
	IndexTypeBTree
	IndexTypeUniqueBTree
	IndexTypeText
	IndexTypeSpatial
	// ...
)
