/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Container is a data structure whose instances are collections of other objects.
// In other words, they store objects in an organized way that follows specific access rules.
// https://en.wikipedia.org/wiki/Container_(abstract_data_type)
// https://en.wikipedia.org/wiki/Collection_(abstract_data_type)
type Container[ELEMENT any] interface {
	ADT_Container_Accessor[ELEMENT]

	ADT_Clear
	ADT_Empty
	ADT_Reversed
	ADT_Sorted
	ADT_Resize

	Capacity
	OccupiedLength
	AvailableLength
}

// ADT_Container_Accessor is the interface that wraps the Accessor methods.
//
// Implementations must not retain p after `Buffer.Reinit` or `Buffer.Deinit` called.
type ADT_Container_Accessor[ELEMENT any] interface {
	ADT_GetElement[ELEMENT]
	ADT_SetElements[ELEMENT]

	ADT_Push[ELEMENT]
	ADT_Pop[ELEMENT]
	ADT_Peek[ELEMENT]
	ADT_Insert[ELEMENT]
	ADT_Add[ELEMENT]
	ADT_Append[ELEMENT]
	ADT_Prepend[ELEMENT]
	ADT_Replace[ELEMENT]

	ADT_Index[ELEMENT]
	ADT_Count[ELEMENT]
	ADT_Iteration[ELEMENT]
}
