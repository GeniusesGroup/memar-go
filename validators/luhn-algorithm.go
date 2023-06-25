/* For license and copyright information please see the LEGAL file in the code repository */

package validators

// LuhnAlgorithm validate given number with the luhn (mod-10) algorithm
// https://en.wikipedia.org/wiki/Luhn_algorithm
func LuhnAlgorithm(pan string) bool {
	/* Validate string with Luhn (mod-10) */
	var alter bool
	var checksum int

	for position := len(pan) - 1; position > -1; position-- {
		var digit = int(pan[position] - '0') // convert characters to its numerical values
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

// LuhnAlgorithm validate given byte slice with the luhn (mod-10) algorithm
// https://en.wikipedia.org/wiki/Luhn_algorithm
func LuhnAlgorithmChecksum(num []byte) (valid bool, checksum byte) {
	var checksumLoc = len(num) - 1
	var buf = num[0:checksumLoc]
	var sum uint

	var numbersLen = len(buf)
	for i := numbersLen - 1; i > -1; i-- {
		var d = buf[i] - '0' // convert characters to its numerical values

		if i&0x1 == 0 {
			d <<= 1 // double the value of every second digit

			if d > 9 {
				d -= 9 // sum the digits of the numbers e.g. 11->2, 14->5, 18->9 etc.
			}
		}
		sum += uint(d)
	}

	checksum = byte((10-sum%10)%10 + '0') // 10 - last digit and convert number to a digit character
	valid = checksum == byte(num[checksumLoc])
	return
}
