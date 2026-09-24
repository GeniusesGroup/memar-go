/* For license and copyright information please see the LEGAL file in the code repository */

package num

import (
	"testing"
)

type test struct {
	numbers       []int
	first, second int
}

var tests = []test{
	{
		numbers: []int{5, 4, 3, 2, 1},
		first:   5,
		second:  4,
	}, {
		numbers: []int{1, 2, 3, 4, 5},
		first:   5,
		second:  4,
	}, {
		numbers: []int{2, 3, 1, 4, 5},
		first:   5,
		second:  4,
	}, {
		numbers: []int{3, 1, 2, 4, 5},
		first:   5,
		second:  4,
	},
}

func TestURI_Unmarshal(t *testing.T) {
	for _, test := range tests {
		var first, second = GetTwoMax(test.numbers...)
		// fmt.Println(first, second)
		if first != test.first || second != test.second {
			t.Errorf("GetTwoMax() = %v, want %v", first, test.first)
		}
	}
}
