/* For license and copyright information please see the LEGAL file in the code repository */

package graph_p

// https://en.wikipedia.org/wiki/Graph_(abstract_data_type)
type Graph[pID, sID, eID any] interface {
	Field_PrimaryNode[pID]
	Field_SecondaryNode[sID]
	Field_Edge[eID]
}
