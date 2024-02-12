/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Stack is dynamically size array
// https://en.wikipedia.org/wiki/Container_(abstract_data_type)
// https://en.wikipedia.org/wiki/Collection_(abstract_data_type)
type ADT_Container[ELEMENT any] interface {
	ADT_Container_Accessor[ELEMENT]
	Capacity
	OccupiedLength
	AvailableLength
}

// ADT_Container_Accessor is the interface that wraps the Accessor methods.
//
// Implementations must not retain p after `Buffer.Reinit` or `Buffer.Deinit` called.
type ADT_Container_Accessor[ELEMENT any] interface {
	// When `Get` returns limit > len(p), it returns a non-nil error explaining why more bytes were not returned.
	Get(offset ElementIndex, limit NumberOfElement) (el ELEMENT, err Error)
	// GetByte provides an efficient interface for byte-at-time processing.
	GetByte(offset ElementIndex) (p byte, err Error)
	GetLast(limit NumberOfElement) (el ELEMENT, err Error)

	// `Set` writes len(p) bytes from p to the underlying data stream at given offset.
	// Clients can execute parallel `Set` calls on the same destination if the ranges do not overlap.
	// If p is a `Get` result, no copy action need and just increase buffer write index.
	Set(el ELEMENT, offset ElementIndex) (err Error)
	SetByte(p byte, offset ElementIndex) (err Error)

	// String() string
	// WriteString(s string) (n NumberOfElement, err Error)

	ADT_Push[ELEMENT]
	ADT_Pop[ELEMENT]
	ADT_Peek[ELEMENT]
	ADT_Insert[ELEMENT]
	ADT_Append[ELEMENT]
	ADT_Prepend[ELEMENT]
	ADT_Clear
}
