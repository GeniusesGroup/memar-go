/* For license and copyright information please see the LEGAL file in the code repository */

package convert

import (
	"fmt"
	"testing"
)

func Test_Unsafe32ByteArraySliceToByteSlice(t *testing.T) {
	var hi = [][32]byte{
		{1, 2, 3, 4},
		{10, 9, 8, 7},
	}

	var temp = Unsafe32ByteArraySliceToByteSlice(hi)

	fmt.Println(cap(hi), len(hi))
	fmt.Println(hi)
	fmt.Println(cap(temp), len(temp))
	fmt.Println(temp)
}

func Test_UnsafeByteSliceTo32ByteArraySlice(t *testing.T) {
	var hi = [][32]byte{
		{1, 2, 3, 4},
		{10, 9, 8, 7},
	}

	var slice = Unsafe32ByteArraySliceToByteSlice(hi)

	var temp = UnsafeByteSliceTo32ByteArraySlice(slice)

	fmt.Println(cap(slice), len(slice))
	fmt.Println(slice)
	fmt.Println(cap(temp), len(temp))
	fmt.Println(temp)

	if hi[0] != temp[0] || hi[1] != temp[1] {
		t.Error("UnsafeByteSliceTo32ByteArraySlice occur problem")
	}
}
