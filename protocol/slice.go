/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Slice is dynamically size array
type Slice interface {
	// Data() T
	Cap() int

	// Append(v T)
	// Copy(d T)

	Len
}
