/* For license and copyright information please see the LEGAL file in the code repository */

package per

// Billion equal to per billion(1,000,000,000) with (ppm) sign
type Billion uint32

// Calculate return PerBillion of given number.
func (b Billion) Calculate(num uint64) (per uint64) {
	return (num * uint64(b)) / 1000000000
}
