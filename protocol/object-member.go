/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Object_Member can use for any data like CLA flags or json fields or any other data structures
// even in compilers or runtime packages
type Object_Member /*[T any]*/ interface {
	// Get() T
	// Set(value T)
	// Default() T

	Name() string
	Abbreviation() string
	Type() Object_Type
	Access() Object_Access
	Optional() bool
	// TODO::: add more
	Validate() Error

	SetDefault() // default value

	Object_Member_Len
	Details
	Stringer
}

type Object_Member_Len interface {
	// Expected length
	MinLen() int // in byte or 8bit
	MaxLen() int // in byte or 8bit

	// existence length
	Len
}
