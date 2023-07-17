/* For license and copyright information please see the LEGAL file in the code repository */

package per

// Myriad equal to per ten-thousand(10,000) with ‱ sign
type Myriad uint16

// Calculate return PerMyriad of given number.
func (b Myriad) Calculate(num uint64) (per uint64) {
	return (num * uint64(b)) / 10000
}
