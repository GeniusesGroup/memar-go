/* For license and copyright information please see the LEGAL file in the code repository */

package container_p

// Assignment is the interface that wraps the assignment methods.
type Assignment[ELEMENT Element] interface {
	SetElements[ELEMENT]

	Push[ELEMENT]
	Pop[ELEMENT]
	Insert[ELEMENT]
	Add[ELEMENT]
	Append[ELEMENT]
	Prepend[ELEMENT]
	Replace[ELEMENT]
}
