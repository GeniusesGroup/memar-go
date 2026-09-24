/* For license and copyright information please see the LEGAL file in the code repository */

package syllab

import (
	"memar/computer/buffer/byteslice/convert"
)

/*
**************************************************************************************************
**************************************Dynamically size Data**************************************
**************************************************************************************************
 */

// GetStringSafely decodes string from the buffer
func GetStringSafely(p []byte, stackIndex uint32) string {
	return string(GetByteArray(p, stackIndex))
}

// GetByteArraySafely decodes byte||uint8 array from the buffer
func GetByteArraySafely(p []byte, stackIndex uint32) (slice []byte) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]byte, ln)
	copy(slice, p[add:])
	return
}

// GetInt8ArraySafely decodes int8 array from the buffer
func GetInt8ArraySafely(p []byte, stackIndex uint32) (slice []int8) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]int8, ln)
	copy(slice, convert.UnsafeByteSliceToInt8Slice(p[add:]))
	return
}

// GetBoolArraySafely decodes bool array from the buffer
func GetBoolArraySafely(p []byte, stackIndex uint32) (slice []bool) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]bool, ln)
	copy(slice, convert.UnsafeByteSliceToBoolSlice(p[add:]))
	return
}

// GetInt16ArraySafely decode Int16 array from the buffer
func GetInt16ArraySafely(p []byte, stackIndex uint32) (slice []int16) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]int16, ln)
	copy(slice, convert.UnsafeByteSliceToInt16Slice(p[add:]))
	return
}

// GetUInt16ArraySafely decode UInt16 array from the buffer
func GetUInt16ArraySafely(p []byte, stackIndex uint32) (slice []uint16) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]uint16, ln)
	copy(slice, convert.UnsafeByteSliceToUInt16Slice(p[add:]))
	return
}

// GetInt32ArraySafely decode fixed size Int32 array from the buffer
func GetInt32ArraySafely(p []byte, stackIndex uint32) (slice []int32) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]int32, ln)
	copy(slice, convert.UnsafeByteSliceToInt32Slice(p[add:]))
	return
}

// GetUInt32ArraySafely decode fixed size UInt32 array from the buffer
func GetUInt32ArraySafely(p []byte, stackIndex uint32) (slice []uint32) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]uint32, ln)
	copy(slice, convert.UnsafeByteSliceToUInt32Slice(p[add:]))
	return
}

// GetInt64ArraySafely decode fixed size Int64 array from the  buffer
func GetInt64ArraySafely(p []byte, stackIndex uint32) (slice []int64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]int64, ln)
	copy(slice, convert.UnsafeByteSliceToInt64Slice(p[add:]))
	return
}

// GetUInt64ArraySafely decode fixed size UInt64 array from the  buffer
func GetUInt64ArraySafely(p []byte, stackIndex uint32) (slice []uint64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]uint64, ln)
	copy(slice, convert.UnsafeByteSliceToUInt64Slice(p[add:]))
	return
}

// GetFloat32ArraySafely decode fixed size Float32 array from the  buffer
func GetFloat32ArraySafely(p []byte, stackIndex uint32) (slice []float32) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]float32, ln)
	copy(slice, convert.UnsafeByteSliceToFloat32Slice(p[add:]))
	return
}

// GetFloat64ArraySafely decode fixed size Float64 array from the  buffer
func GetFloat64ArraySafely(p []byte, stackIndex uint32) (slice []float64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]float64, ln)
	copy(slice, convert.UnsafeByteSliceToFloat64Slice(p[add:]))
	return
}

// GetComplex64ArraySafely decode fixed size Complex64 array from the  buffer
func GetComplex64ArraySafely(p []byte, stackIndex uint32) (slice []complex64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]complex64, ln)
	copy(slice, convert.UnsafeByteSliceToComplex64Slice(p[add:]))
	return
}

// GetComplex128ArraySafely decode fixed size Complex128 array from the  buffer
func GetComplex128ArraySafely(p []byte, stackIndex uint32) (slice []complex128) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]complex128, ln)
	copy(slice, convert.UnsafeByteSliceToComplex128Slice(p[add:]))
	return
}

/*
**************************************************************************************************
*******************Dynamically size ARRAY inside other Dynamically size Array*******************
**************************************************************************************************
 */

// GetStringArraySafely encode string array to the  buffer
func GetStringArraySafely(p []byte, stackIndex uint32) (slice []string) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]string, ln)

	for i := 0; i < int(ln); i++ {
		var eachAdd uint32 = GetUInt32(p, add)
		var eachLn uint32 = GetUInt32(p, add+4)
		slice[i] = string(p[eachAdd : eachAdd+eachLn])
		add += 8
	}
	return
}

// Get16ByteArrayArraySafely decode [16]byte array from the  buffer
func Get16ByteArrayArraySafely(p []byte, stackIndex uint32) (slice [][16]byte) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([][16]byte, ln)
	copy(slice, convert.UnsafeByteSliceTo16ByteArraySlice(p[add:add+(ln*16)]))
	return
}

// Get32ByteArrayArraySafely decode [32]byte array from the  buffer
func Get32ByteArrayArraySafely(p []byte, stackIndex uint32) (slice [][32]byte) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([][32]byte, ln)
	copy(slice, convert.UnsafeByteSliceTo32ByteArraySlice(p[add:add+(ln*32)]))
	return
}
