/* For license and copyright information please see the LEGAL file in the code repository */

package container_p

// Accessor is the interface that wraps the Accessor methods.
type Accessor[ELEMENT Element] interface {
	GetElement[ELEMENT]

	Peek[ELEMENT]

	Index[ELEMENT]
	Count[ELEMENT]
	Contain[ELEMENT]
	
	Iteration[ELEMENT]
}
