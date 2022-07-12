/* For license and copyright information please see LEGAL file in repository */

package matn

// WordLabel indicate any word||phrase label
type WordLabel uint64

// Word||phrases lables
const (
	WordLabelUnset  WordLabel = iota
	WordLabelPERSON           // People
	WordLabelGPE              // geographical/political Entities (GPE
)
