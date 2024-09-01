/* For license and copyright information please see the LEGAL file in the code repository */

package adt_p

// Nil is for object pointers, NULL is for non pointers
// TODO::: Is nil same as void??
type Empty interface {
	// an operation for testing whether or not a container is empty(Nil && Null).
	Empty() bool
}
type Nil interface {
	// IsNil report is pointer to the Element valid or not.
	IsNil() bool
}
type Null interface {
	// IsNull report is the Element has value(data set before) or not.
	IsNull() bool
}
