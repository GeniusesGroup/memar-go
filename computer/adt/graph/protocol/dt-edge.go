/* For license and copyright information please see the LEGAL file in the code repository */

package graph_p

type Field_Edge[ID any] interface {
	// Edge usually is a label.
	// labelID. Edge in graph data type
	Edge() ID
}
