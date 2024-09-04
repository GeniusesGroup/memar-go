/* For license and copyright information please see the LEGAL file in the code repository */

package cl

type Field_Logical_Connective interface {
	LogicalConnective() Logical_Connective
}

//
// https://en.wikipedia.org/wiki/Logical_connective
type Logical_Connective uint64

const Logical_Connective_Unset Logical_Connective = 0

const (
	// https://en.wikipedia.org/wiki/Logical_conjunction
	Logical_Connective_AND Logical_Connective = (1 << iota)

	// https://en.wikipedia.org/wiki/Logical_biconditional
	Logical_Connective_Equivalent

	// It can indicate by NOT(Equivalent) flags.
	// Logical_Connective_Nonequivalent

	// https://en.wikipedia.org/wiki/Material_conditional
	Logical_Connective_Implies

	// https://en.wikipedia.org/wiki/Sheffer_stroke
	Logical_Connective_NAND

	// https://en.wikipedia.org/wiki/Logical_NOR
	Logical_Connective_NOR

	// https://en.wikipedia.org/wiki/Negation
	Logical_Connective_NOT

	// https://en.wikipedia.org/wiki/Logical_disjunction
	Logical_Connective_OR

	// https://en.wikipedia.org/wiki/XNOR_gate
	Logical_Connective_XNOR

	// https://en.wikipedia.org/wiki/Exclusive_or
	Logical_Connective_XOR

	// https://en.wikipedia.org/wiki/Converse_(logic)
	Logical_Connective_Converse
)
