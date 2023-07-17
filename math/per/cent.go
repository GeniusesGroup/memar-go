/* For license and copyright information please see the LEGAL file in the code repository */

package per

// Cent equal to per hundred(100) with % sign
type Cent uint8

// Calculate return PerCent of given number.
func (b Cent) Calculate(num uint64) (per uint64) {
	return (num * uint64(b)) / 100
}
