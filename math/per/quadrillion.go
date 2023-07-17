/* For license and copyright information please see the LEGAL file in the code repository */

package per

// Quadrillion equal to per quadrillion(1,000,000,000,000,000) with (ppm) sign
type Quadrillion uint64

// Calculate return PerQuadrillion of given number.
func (b Quadrillion) Calculate(num uint64) (per uint64) {
	return (num * uint64(b)) / 1000000
}
