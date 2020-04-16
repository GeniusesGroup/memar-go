/* For license and copyright information please see LEGAL file in repository */

package validators

// LuhnAlgorithm validate strings with the luhn (mod-10) algorithm!
// https://en.wikipedia.org/wiki/Luhn_algorithm
func LuhnAlgorithm(pan string) bool {
	/* Validate string with Luhn (mod-10) */
	var alter bool
	var checksum int

	for position := len(pan) - 1; position > -1; position-- {
		digit := int(pan[position] - 48)
		if alter {
			digit = digit * 2
			if digit > 9 {
				digit = (digit % 10) + 1
			}
		}
		alter = !alter
		checksum += digit
	}
	return checksum%10 == 0
}
