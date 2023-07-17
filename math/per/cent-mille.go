/* For license and copyright information please see the LEGAL file in the code repository */

package per

// CentMille equal to per Hundred Thousand(100,000) with (pcm) sign
type CentMille uint32

// Calculate return PerCentMille of given number.
func (b CentMille) Calculate(num uint64) (per uint64) {
	return (num * uint64(b)) / 100000
}
