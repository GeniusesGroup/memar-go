/* For license and copyright information please see the LEGAL file in the code repository */

package per

// Ten equal to per ten(10)
type Ten uint8

// Calculate return PerTen of given number.
func (pt Ten) Calculate(num uint64) (per uint64) {
	return (num * uint64(pt)) / 10
}
