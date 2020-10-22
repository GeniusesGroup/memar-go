/* For license and copyright information please see LEGAL file in repository */

package syllab

import "../convert"

/*
**************************************************************************************************
*****************************************Fixed Size ARRAY*****************************************
**************************************************************************************************
 */

// Unsafe don't give any efficient in fixed size data! Use decode safe functions!

/*
**************************************************************************************************
**************************************Dynamically size ARRAY**************************************
**************************************************************************************************
 */

// UnsafeGetString decodes string from the payload buffer in unsafe manner!
func UnsafeGetString(p []byte, stackIndex uint32) string {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToString(p[add : add+ln])
}

// UnsafeGetByteArray decodes byte slice from the payload buffer in unsafe manner!
func UnsafeGetByteArray(p []byte, stackIndex uint32) []byte {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return p[add : add+ln]
}

// UnsafeGetInt8Array decodes int8 slice from the payload buffer in unsafe manner!
func UnsafeGetInt8Array(p []byte, stackIndex uint32) []int8 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToInt8Slice(p[add : add+ln])
}

// UnsafeGetBoolArray decodes bool array from the payload buffer!
func UnsafeGetBoolArray(p []byte, stackIndex uint32) (slice []bool) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToBoolSlice(p[add : add+ln])
}

// UnsafeGetInt16Array decode Int16 array from the payload buffer
func UnsafeGetInt16Array(p []byte, stackIndex uint32) (slice []int16) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToInt16Slice(p[add : add+(ln*2)])
}

// UnsafeGetUInt16Array decode UInt16 array from the payload buffer
func UnsafeGetUInt16Array(p []byte, stackIndex uint32) (slice []uint16) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToUInt16Slice(p[add : add+(ln*2)])
}

// UnsafeGetInt32Array decode fixed size Int32 array from the payload buffer
func UnsafeGetInt32Array(p []byte, stackIndex uint32) (slice []int32) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToInt32Slice(p[add : add+(ln*4)])
}

// UnsafeGetUInt32Array decode fixed size UInt32 array from the payload buffer
func UnsafeGetUInt32Array(p []byte, stackIndex uint32) (slice []uint32) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToUInt32Slice(p[add : add+(ln*4)])
}

// UnsafeGetInt64Array decode fixed size Int64 array from the payload buffer
func UnsafeGetInt64Array(p []byte, stackIndex uint32) (slice []int64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToInt64Slice(p[add : add+(ln*8)])
}

// UnsafeGetUInt64Array decode fixed size UInt64 array from the payload buffer
func UnsafeGetUInt64Array(p []byte, stackIndex uint32) (slice []uint64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToUInt64Slice(p[add : add+(ln*8)])
}

// UnsafeGetFloat32Array decode fixed size Float32 array from the payload buffer
func UnsafeGetFloat32Array(p []byte, stackIndex uint32) (slice []float32) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToFloat32Slice(p[add : add+(ln*4)])
}

// UnsafeGetFloat64Array decode fixed size Float64 array from the payload buffer
func UnsafeGetFloat64Array(p []byte, stackIndex uint32) (slice []float64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToFloat64Slice(p[add : add+(ln*8)])
}

// UnsafeGetComplex64Array decode fixed size Complex64 array from the payload buffer
func UnsafeGetComplex64Array(p []byte, stackIndex uint32) (slice []complex64) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToComplex64Slice(p[add : add+(ln*8)])
}

// UnsafeGetComplex128Array decode fixed size Complex128 array from the payload buffer
func UnsafeGetComplex128Array(p []byte, stackIndex uint32) (slice []complex128) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	return convert.UnsafeByteSliceToComplex128Slice(p[add : add+(ln*16)])
}

/*
**************************************************************************************************
*******************Dynamically size ARRAY inside other Dynamically size Array*******************
**************************************************************************************************
 */

// UnsafeGetStringArray encode string array to the payload buffer!
func UnsafeGetStringArray(p []byte, stackIndex uint32) (slice []string) {
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
