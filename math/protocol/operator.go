/* For license and copyright information please see the LEGAL file in the code repository */

package math_p

// TODO::: Does comparison circuitry belong in a purely mathematical or computing category??
// https://en.wikipedia.org/wiki/Category:Comparison_(mathematical)
// For example the symbol ">" could imply greater than, better than, ahead of, higher than, etc.

// https://en.wikipedia.org/wiki/Operators_in_C_and_C%2B%2B

type ComparisonOperator uint8

// Comparison operators/relational operators
const (
	ComparisonOperator_Unset  ComparisonOperator = 0
	ComparisonOperator_Negate ComparisonOperator = (1 << iota)
	ComparisonOperator_Equal
	ComparisonOperator_GreaterThan
	ComparisonOperator_LessThan
	// https://en.wikipedia.org/wiki/Three-way_comparison
	ComparisonOperator_ThreeWayComparison
)
