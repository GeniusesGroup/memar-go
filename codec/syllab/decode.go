/* For license and copyright information please see the LEGAL file in the code repository */

package syllab

import (
	"memar/codec/binary"
	"memar/computer/buffer/byteslice/convert"
)

/*
**************************************************************************************************
*****************************************Fixed Size Data******************************************
**************************************************************************************************

	Read more about this codec : https://github.com/GeniusesGroup/memar/blob/main/Syllab.md
*/

func GetFixedByteArray(p []byte, stackIndex uint32, array []byte) { copy(array[0:], p[stackIndex:]) }
func GetByte(p []byte, stackIndex uint32) byte                    { return p[stackIndex] }
func GetInt8(p []byte, stackIndex uint32) int8                    { return int8(p[stackIndex]) }
func GetUInt8(p []byte, stackIndex uint32) uint8                  { return uint8(p[stackIndex]) }
func GetBool(p []byte, stackIndex uint32) bool                    { return p[stackIndex] != 0 }
func GetInt16(p []byte, stackIndex uint32) int16 {
	return int16(binary.LittleEndian.Uint16(p[stackIndex:]))
}
func GetUInt16(p []byte, stackIndex uint32) uint16 { return binary.LittleEndian.Uint16(p[stackIndex:]) }
func GetInt32(p []byte, stackIndex uint32) int32 {
	return int32(binary.LittleEndian.Uint32(p[stackIndex:]))
}
func GetUInt32(p []byte, stackIndex uint32) uint32   { return binary.LittleEndian.Uint32(p[stackIndex:]) }
func GetFloat32(p []byte, stackIndex uint32) float32 { return float32(GetUInt32(p, stackIndex)) }
func GetInt64(p []byte, stackIndex uint32) int64 {
	return int64(binary.LittleEndian.Uint64(p[stackIndex:]))
}
func GetUInt64(p []byte, stackIndex uint32) uint64   { return binary.LittleEndian.Uint64(p[stackIndex:]) }
func GetFloat64(p []byte, stackIndex uint32) float64 { return float64(GetUInt64(p, stackIndex)) }

// GetComplex64 decodes COMPLEX64 from the buffer.
func GetComplex64(p []byte, stackIndex uint32) complex64 {
	return complex(GetFloat32(p, stackIndex), GetFloat32(p, stackIndex+4))
}

// GetComplex128 decodes COMPLEX128 from the buffer.
func GetComplex128(p []byte, stackIndex uint32) complex128 {
	return complex(GetFloat64(p, stackIndex), GetFloat64(p, stackIndex+8))
}

/*
**************************************************************************************************
**************************************Dynamically size ARRAY**************************************
**************************************************************************************************
 */

// GetString decodes string from the buffer in unsafe manner!
func GetString(p []byte, stackIndex uint32) string {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToString(p[add : add+ln])
}

// GetByteArray decodes byte slice from the buffer in unsafe manner!
func GetByteArray(p []byte, stackIndex uint32) []byte {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return p[add : add+ln]
}

// GetInt8Array decodes int8 slice from the buffer in unsafe manner!
func GetInt8Array(p []byte, stackIndex uint32) []int8 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToInt8Slice(p[add : add+ln])
}

// GetBoolArray decodes bool array from the buffer
func GetBoolArray(p []byte, stackIndex uint32) (slice []bool) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToBoolSlice(p[add : add+ln])
}

// GetInt16Array decode Int16 array from the buffer
func GetInt16Array(p []byte, stackIndex uint32) (slice []int16) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToInt16Slice(p[add : add+(ln*2)])
}

// GetUInt16Array decode UInt16 array from the buffer
func GetUInt16Array(p []byte, stackIndex uint32) (slice []uint16) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToUInt16Slice(p[add : add+(ln*2)])
}

// GetInt32Array decode fixed size Int32 array from the buffer
func GetInt32Array(p []byte, stackIndex uint32) (slice []int32) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToInt32Slice(p[add : add+(ln*4)])
}

// GetUInt32Array decode fixed size UInt32 array from the buffer
func GetUInt32Array(p []byte, stackIndex uint32) (slice []uint32) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToUInt32Slice(p[add : add+(ln*4)])
}

// GetInt64Array decode fixed size Int64 array from the buffer
func GetInt64Array(p []byte, stackIndex uint32) (slice []int64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToInt64Slice(p[add : add+(ln*8)])
}

// GetUInt64Array decode fixed size UInt64 array from the buffer
func GetUInt64Array(p []byte, stackIndex uint32) (slice []uint64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToUInt64Slice(p[add : add+(ln*8)])
}

// GetFloat32Array decode fixed size Float32 array from the buffer
func GetFloat32Array(p []byte, stackIndex uint32) (slice []float32) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToFloat32Slice(p[add : add+(ln*4)])
}

// GetFloat64Array decode fixed size Float64 array from the buffer
func GetFloat64Array(p []byte, stackIndex uint32) (slice []float64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToFloat64Slice(p[add : add+(ln*8)])
}

// GetComplex64Array decode fixed size Complex64 array from the buffer
func GetComplex64Array(p []byte, stackIndex uint32) (slice []complex64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToComplex64Slice(p[add : add+(ln*8)])
}

// GetComplex128Array decode fixed size Complex128 array from the buffer
func GetComplex128Array(p []byte, stackIndex uint32) (slice []complex128) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToComplex128Slice(p[add : add+(ln*16)])
}

/*
**************************************************************************************************
*******************Dynamically size ARRAY inside other Dynamically size Array*******************
**************************************************************************************************
 */

// GetStringArray encode string array from the buffer
func GetStringArray(p []byte, stackIndex uint32) (slice []string) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]string, ln)

	for i := 0; i < int(ln); i++ {
		var eachAdd uint32 = GetUInt32(p, add)
		var eachLn uint32 = GetUInt32(p, add+4)
		slice[i] = convert.UnsafeByteSliceToString(p[eachAdd : eachAdd+eachLn])
		add += 8
	}
	return
}

// Get16ByteArrayArray decode [16]byte array from the buffer
func Get16ByteArrayArray(p []byte, stackIndex uint32) (slice [][16]byte) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceTo16ByteArraySlice(p[add : add+(ln*16)])
}

// Get32ByteArraySlice decode [32]byte array from the buffer
func Get32ByteArraySlice(p []byte, stackIndex uint32) (slice [][32]byte) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceTo32ByteArraySlice(p[add : add+(ln*32)])
}
