/* For license and copyright information please see the LEGAL file in the code repository */

package num

func GetTwoMax(numbers ...int) (first, second int) {
	first = numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > second {
			second = numbers[i]
		} else if numbers[i] > first {
			first = numbers[i]
		}

		if first < second {
			first, second = second, first
		}
	}
	return
}
