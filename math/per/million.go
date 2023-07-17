/* For license and copyright information please see the LEGAL file in the code repository */

package per

// Million equal to per million(1,000,000) with (ppm) sign
type Million uint32

// Calculate return PerMillion of given number.
func (b Million) Calculate(num uint64) (per uint64) {
	return (num * uint64(b)) / 1000000
}
