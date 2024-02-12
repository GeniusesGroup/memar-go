/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Length is a measure of distance. In the International System of Quantities, length is a quantity with dimension distance.
// In most systems of measurement a base unit for length is chosen, from which all other units are derived.
// https://en.wikipedia.org/wiki/Length
//
// Length can refer to any size of memory blocks in byte or 8bit number.
// It can present capacity, usage, buffer, ...
type NumberOfElement int

// Capacity in (computer science) the amount of information (in bytes) that can be stored.
type Capacity interface {
	// Capacity return a length that underlying implementor can store desire elements such as byte.
	Capacity() NumberOfElement
}

type OccupiedLength interface {
	// OccupiedLength return a length that store before this method call.
	OccupiedLength() NumberOfElement
}

type AvailableLength interface {
	// AvailableLength or EmptyLength() or RemainingLength() returns how a length that are unused or can be set.
	AvailableLength() NumberOfElement
}
