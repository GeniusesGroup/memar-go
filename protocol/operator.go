/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

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

// Equal provide comparable two items.
type Equal[T any] interface {
	Equal(v1, v2 T) bool
}
