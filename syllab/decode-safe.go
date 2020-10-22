/* For license and copyright information please see LEGAL file in repository */

package syllab

import "../convert"

/*
**************************************************************************************************
*****************************************Fixed Size Data******************************************
**************************************************************************************************
 */

// GetFixedByteArray decodes fixed sized byte array from the payload buffer.
// If you want array instead of slice from function below, You can copy function and edit it for your usage! e.g.
// for get a [2]byte  ```var array [2]byte = copy(array[:], p[stackIndex:])```
// for get a [32]byte  ```var array [32]byte = copy(array[:], p[stackIndex:])```
func GetFixedByteArray(p []byte, stackIndex uint32, n uint32) (array []byte) {
	copy(array[:], p[stackIndex:stackIndex+n])
	return
}

// GetByte decodes BYTE from the payload buffer.
func GetByte(p []byte, stackIndex uint32) byte {
	return p[stackIndex]
}

// GetInt8 decodes INT8 from the payload buffer.
func GetInt8(p []byte, stackIndex uint32) int8 {
	return int8(p[stackIndex])
}

// GetUInt8 decodes UINT8 from the payload buffer.
func GetUInt8(p []byte, stackIndex uint32) uint8 {
	return uint8(p[stackIndex])
}

// GetBool decodes BOOL from the payload buffer.
func GetBool(p []byte, stackIndex uint32) bool {
	return p[stackIndex] == 1
}

// GetInt16 decodes INT16 from the payload buffer.
func GetInt16(p []byte, stackIndex uint32) int16 {
	return int16(p[stackIndex]) | int16(p[stackIndex+1])<<8
}

// GetUInt16 decodes UINT16 from the payload buffer.
func GetUInt16(p []byte, stackIndex uint32) uint16 {
	return uint16(p[stackIndex]) | uint16(p[stackIndex+1])<<8
}

// GetInt32 decodes INT32 from the payload buffer.
func GetInt32(p []byte, stackIndex uint32) int32 {
	return int32(p[stackIndex]) | int32(p[stackIndex+1])<<8 | int32(p[stackIndex+2])<<16 | int32(p[stackIndex+3])<<24
}

// GetUInt32 decodes UINT32 from the payload buffer.
func GetUInt32(p []byte, stackIndex uint32) uint32 {
	return uint32(p[stackIndex]) | uint32(p[stackIndex+1])<<8 | uint32(p[stackIndex+2])<<16 | uint32(p[stackIndex+3])<<24
}

// GetFloat32 decodes FLOAT32 from the payload buffer.
func GetFloat32(p []byte, stackIndex uint32) float32 {
	return float32(GetUInt32(p, stackIndex))
}

// GetInt64 decodes INT64 from the payload buffer.
func GetInt64(p []byte, stackIndex uint32) int64 {
	return int64(p[stackIndex]) | int64(p[stackIndex+1])<<8 | int64(p[stackIndex+2])<<16 | int64(p[stackIndex+3])<<24 |
		int64(p[stackIndex+4])<<32 | int64(p[stackIndex+5])<<40 | int64(p[stackIndex+6])<<48 | int64(p[stackIndex+7])<<56
}

// GetUInt64 decodes UINT64 from the payload buffer.
func GetUInt64(p []byte, stackIndex uint32) uint64 {
	return uint64(p[stackIndex]) | uint64(p[stackIndex+1])<<8 | uint64(p[stackIndex+2])<<16 | uint64(p[stackIndex+3])<<24 |
		uint64(p[stackIndex+4])<<32 | uint64(p[stackIndex+5])<<40 | uint64(p[stackIndex+6])<<48 | uint64(p[stackIndex+7])<<56
}

// GetFloat64 decodes FLOAT64 from the payload buffer.
func GetFloat64(p []byte, stackIndex uint32) float64 {
	return float64(GetUInt64(p, stackIndex))
}

// GetComplex64 decodes COMPLEX64 from the payload buffer.
func GetComplex64(p []byte, stackIndex uint32) complex64 {
	return complex(GetFloat32(p, stackIndex), GetFloat32(p, stackIndex+4))
}

// GetComplex128 decodes COMPLEX128 from the payload buffer.
func GetComplex128(p []byte, stackIndex uint32) complex128 {
	return complex(GetFloat64(p, stackIndex), GetFloat64(p, stackIndex+8))
}

/*
**************************************************************************************************
**************************************Dynamically size Data**************************************
**************************************************************************************************
 */

// GetString decodes string from the payload buffer!
func GetString(p []byte, stackIndex uint32) string {
	return string(GetByteArray(p, stackIndex))
}

// GetByteArray decodes byte||uint8 array from the payload buffer!
func GetByteArray(p []byte, stackIndex uint32) (slice []byte) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]byte, ln)
	copy(slice, p[add:])
	return
}

// GetInt8Array decodes int8 array from the payload buffer!
func GetInt8Array(p []byte, stackIndex uint32) (slice []int8) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]int8, ln)
	copy(slice, convert.UnsafeByteSliceToInt8Slice(p[add:]))
	return
}

// GetBoolArray decodes bool array from the payload buffer!
func GetBoolArray(p []byte, stackIndex uint32) (slice []bool) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]bool, ln)
	copy(slice, convert.UnsafeByteSliceToBoolSlice(p[add:]))
	return
}

// GetInt16Array decode Int16 array from the payload buffer
func GetInt16Array(p []byte, stackIndex uint32) (slice []int16) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]int16, ln)
	copy(slice, convert.UnsafeByteSliceToInt16Slice(p[add:]))
	return
}

// GetUInt16Array decode UInt16 array from the payload buffer
func GetUInt16Array(p []byte, stackIndex uint32) (slice []uint16) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]uint16, ln)
	copy(slice, convert.UnsafeByteSliceToUInt16Slice(p[add:]))
	return
}

// GetInt32Array decode fixed size Int32 array from the payload buffer
func GetInt32Array(p []byte, stackIndex uint32) (slice []int32) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]int32, ln)
	copy(slice, convert.UnsafeByteSliceToInt32Slice(p[add:]))
	return
}

// GetUInt32Array decode fixed size UInt32 array from the payload buffer
func GetUInt32Array(p []byte, stackIndex uint32) (slice []uint32) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]uint32, ln)
	copy(slice, convert.UnsafeByteSliceToUInt32Slice(p[add:]))
	return
}

// GetInt64Array decode fixed size Int64 array from the payload buffer
func GetInt64Array(p []byte, stackIndex uint32) (slice []int64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]int64, ln)
	copy(slice, convert.UnsafeByteSliceToInt64Slice(p[add:]))
	return
}

// GetUInt64Array decode fixed size UInt64 array from the payload buffer
func GetUInt64Array(p []byte, stackIndex uint32) (slice []uint64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]uint64, ln)
	copy(slice, convert.UnsafeByteSliceToUInt64Slice(p[add:]))
	return
}

// GetFloat32Array decode fixed size Float32 array from the payload buffer
func GetFloat32Array(p []byte, stackIndex uint32) (slice []float32) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]float32, ln)
	copy(slice, convert.UnsafeByteSliceToFloat32Slice(p[add:]))
	return
}

// GetFloat64Array decode fixed size Float64 array from the payload buffer
func GetFloat64Array(p []byte, stackIndex uint32) (slice []float64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]float64, ln)
	copy(slice, convert.UnsafeByteSliceToFloat64Slice(p[add:]))
	return
}

// GetComplex64Array decode fixed size Complex64 array from the payload buffer
func GetComplex64Array(p []byte, stackIndex uint32) (slice []complex64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]complex64, ln)
	copy(slice, convert.UnsafeByteSliceToComplex64Slice(p[add:]))
	return
}

// GetComplex128Array decode fixed size Complex128 array from the payload buffer
func GetComplex128Array(p []byte, stackIndex uint32) (slice []complex128) {
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

// GetStringArray encode string array to the payload buffer!
func GetStringArray(p []byte, stackIndex uint32) (slice []string) {
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
