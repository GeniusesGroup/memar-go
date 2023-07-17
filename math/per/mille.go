/* For license and copyright information please see the LEGAL file in the code repository */

package per

// Mille equal to per thousand(1000) with ‰ sign
type Mille uint16

// Calculate return PerMille of given number.
func (b Mille) Calculate(num uint64) (per uint64) {
	return (num * uint64(b)) / 1000
}
