/* For license and copyright information please see the LEGAL file in the code repository */

package cl

type Field_Conditional interface {
	Conditional() Conditional
}

//
// https://en.wikipedia.org/wiki/Conditional_(computer_programming)
type Conditional uint64

const (
	Conditional_Unset Conditional = iota
	Conditional_Then
	Conditional_ActionAndContinue
	Conditional_Continue
	Conditional_Break
	Conditional_Return
)
