/* For license and copyright information please see the LEGAL file in the code repository */

package container_p

import (
	adt_p "memar/adt/protocol"
)

// Container is a data structure whose instances are collections of other objects.
// In other words, they store objects in an organized way that follows specific access rules.
// https://en.wikipedia.org/wiki/Container_(abstract_data_type)
// https://en.wikipedia.org/wiki/Collection_(abstract_data_type)
// 
// Other frameworks:
// https://learn.microsoft.com/en-us/dotnet/api/system.componentmodel.icontainer
type Container[ELEMENT Element] interface {
	adt_p.ADT

	Accessor[ELEMENT]

	Clear
	Reversed
	Sorted
	Resize

	Capacity
	OccupiedLength
	AvailableLength
	// ExpectedLength
}

type Container_READONLY[ELEMENT Element] interface {
	OccupiedLength
	Iteration[ELEMENT]
	Count[ELEMENT]
	Contain[ELEMENT]
	Compare[ELEMENT]
	GetElement[ELEMENT]
}

// Accessor is the interface that wraps the Accessor methods.
//
// Implementations must not retain p after `Buffer.Reinit` or `Buffer.Deinit` called.
type Accessor[ELEMENT Element] interface {
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
