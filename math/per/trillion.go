/* For license and copyright information please see the LEGAL file in the code repository */

package per

// Trillion equal to per trillion(1,000,000,000,000) with (ppm) sign
type Trillion uint64

// Calculate return PerTrillion of given number.
func (b Trillion) Calculate(num uint64) (per uint64) {
	return (num * uint64(b)) / 1000000
}
