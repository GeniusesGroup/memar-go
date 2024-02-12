/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Slice is dynamically size array
// https://en.wikipedia.org/wiki/Dynamic_array
type ADT_Array_Dynamic[ELEMENT any] interface {
	// Data() T

	// Append(v T)
	// Copy(d T)

	Capacity
	OccupiedLength
}

// Growth factor
