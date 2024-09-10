/* For license and copyright information please see the LEGAL file in the code repository */

package primitive_p

// Boolean algebra is a branch of mathematics that deals with operations on logical values with binary variables.
// The Boolean variables are represented as binary numbers to represent truths: 1 = true and 0 = false.
// https://en.wikipedia.org/wiki/Boolean_datatype
// https://en.wikipedia.org/wiki/Boolean_algebra
// https://en.wikipedia.org/wiki/Boolean_algebra_(structure)
//
// Some other languages:
// - https://doc.rust-lang.org/stable/std/primitive.bool.html
type Boolean[B any /*Boolean*/] interface {
	Conjunction[B, B]
	Disjunction[B, B]
	Equivalence[B]
	Negation[B]
}
