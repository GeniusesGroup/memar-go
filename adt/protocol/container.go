/* For license and copyright information please see the LEGAL file in the code repository */

package adt_p

// Container is a data structure whose instances are collections of other objects.
// In other words, they store objects in an organized way that follows specific access rules.
// https://en.wikipedia.org/wiki/Container_(abstract_data_type)
// https://en.wikipedia.org/wiki/Collection_(abstract_data_type)
type Container[ELEMENT Element] interface {
	ADT

	Container_Accessor[ELEMENT]

	Clear
	Reversed
	Sorted
	Resize

	Capacity
	OccupiedLength
	AvailableLength
	// ExpectedLength
}

// Container_Accessor is the interface that wraps the Accessor methods.
//
// Implementations must not retain p after `Buffer.Reinit` or `Buffer.Deinit` called.
type Container_Accessor[ELEMENT Element] interface {
	GetElement[ELEMENT]
	SetElements[ELEMENT]

	Push[ELEMENT]
	Pop[ELEMENT]
	Peek[ELEMENT]
	Insert[ELEMENT]
	Add[ELEMENT]
	Append[ELEMENT]
	Prepend[ELEMENT]
	Replace[ELEMENT]

	Index[ELEMENT]
	Count[ELEMENT]
	Iteration[ELEMENT]
}
