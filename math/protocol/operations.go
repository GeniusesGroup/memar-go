/* For license and copyright information please see the LEGAL file in the code repository */

package math_p

// addition, subtraction, multiplication, division (with care for division by zero), comparison, etc.
// Increment() { i++ }
// Decrement() { i-- }

type Operations interface {
	Mul() // multiple
	Add()
	Sub()
	Div()
}
