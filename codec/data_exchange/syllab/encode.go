/* For license and copyright information please see LEGAL file in repository */

package syllab

import "../convert"

/*
**************************************************************************************************
*****************************************Fixed Size ARRAY*****************************************
**************************************************************************************************
 */

// SetArray encode fixed sized byte array to the payload buffer.
func SetArray(p []byte, stackIndex uint32, a []byte) {
	copy(p[stackIndex:], a[:])
}

// SetByte encode BYTE to the payload buffer.
func SetByte(p []byte, stackIndex uint32, b byte) {
	p[stackIndex] = b
}

// SetInt8 encode INT8 to the payload buffer.
func SetInt8(p []byte, stackIndex uint32, n int8) {
	p[stackIndex] = byte(n)
}

// SetUInt8 encode UINT8 to the payload buffer.
func SetUInt8(p []byte, stackIndex uint32, n uint8) {
	p[stackIndex] = byte(n)
}

// SetBool encode BOOL to the payload buffer.
func SetBool(p []byte, stackIndex uint32, b bool) {
	if b {
		p[stackIndex] = 1
	} // else {
	// 	p[stackIndex] = 0
	// }
}

// SetInt16 encode INT16 to the payload buffer.
func SetInt16(p []byte, stackIndex uint32, n int16) {
	p[stackIndex] = byte(n)
	p[stackIndex+1] = byte(n >> 8)
}

// SetUInt16 encode UINT16 to the payload buffer.
func SetUInt16(p []byte, stackIndex uint32, n uint16) {
	p[stackIndex] = byte(n)
	p[stackIndex+1] = byte(n >> 8)
}

// SetInt32 encode INT32 to the payload buffer.
func SetInt32(p []byte, stackIndex uint32, n int32) {
	p[stackIndex] = byte(n)
	p[stackIndex+1] = byte(n >> 8)
	p[stackIndex+2] = byte(n >> 16)
	p[stackIndex+3] = byte(n >> 24)
}

// SetUInt32 encode UINT32 to the payload buffer.
func SetUInt32(p []byte, stackIndex uint32, n uint32) {
	p[stackIndex] = byte(n)
	p[stackIndex+1] = byte(n >> 8)
	p[stackIndex+2] = byte(n >> 16)
	p[stackIndex+3] = byte(n >> 24)
}

// SetFloat32 encode FLOAT32 to the payload buffer.
func SetFloat32(p []byte, stackIndex uint32, n float32) {
	SetUInt32(p, stackIndex, uint32(n))
}

// SetInt64 encode INT64 to the payload buffer.
func SetInt64(p []byte, stackIndex uint32, n int64) {
	p[stackIndex] = byte(n)
	p[stackIndex+1] = byte(n >> 8)
	p[stackIndex+2] = byte(n >> 16)
	p[stackIndex+3] = byte(n >> 24)
	p[stackIndex+4] = byte(n >> 32)
	p[stackIndex+5] = byte(n >> 40)
	p[stackIndex+6] = byte(n >> 48)
	p[stackIndex+7] = byte(n >> 56)
}

// SetUInt64 encode UINT64 to the payload buffer.
func SetUInt64(p []byte, stackIndex uint32, n uint64) {
	p[stackIndex] = byte(n)
	p[stackIndex+1] = byte(n >> 8)
	p[stackIndex+2] = byte(n >> 16)
	p[stackIndex+3] = byte(n >> 24)
	p[stackIndex+4] = byte(n >> 32)
	p[stackIndex+5] = byte(n >> 40)
	p[stackIndex+6] = byte(n >> 48)
	p[stackIndex+7] = byte(n >> 56)
}

// SetFloat64 encode FLOAT64 to the payload buffer.
func SetFloat64(p []byte, stackIndex uint32, n float64) {
	SetUInt64(p, stackIndex, uint64(n))
	// TODO::: below code instead up func call not allow go compiler to inline this func! WHY???
	// var un = uint64(n)
	// p[stackIndex] = byte(un)
	// p[stackIndex+1] = byte(un >> 8)
	// p[stackIndex+2] = byte(un >> 16)
	// p[stackIndex+3] = byte(un >> 24)
	// p[stackIndex+4] = byte(un >> 32)
	// p[stackIndex+5] = byte(un >> 40)
	// p[stackIndex+6] = byte(un >> 48)
	// p[stackIndex+7] = byte(un >> 56)
}

// SetComplex64 encode COMPLEX64 to the payload buffer.
func SetComplex64(p []byte, stackIndex uint32, n complex64) {
	SetFloat32(p, stackIndex, real(n))
	SetFloat32(p, stackIndex+3, imag(n))
}

// SetComplex128 encode COMPLEX128 to the payload buffer.
func SetComplex128(p []byte, stackIndex uint32, n complex128) {
	SetFloat64(p, stackIndex, real(n))
	SetFloat64(p, stackIndex+7, imag(n))
}

/*
**************************************************************************************************
**************************************Dynamically size ARRAY**************************************
**************************************************************************************************
 */

// SetString encode string to the payload buffer!
func SetString(p []byte, s string, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	// SetUInt32(p, stackIndex, heapAddr)
	p[stackIndex] = byte(heapAddr)
	p[stackIndex+1] = byte(heapAddr >> 8)
	p[stackIndex+2] = byte(heapAddr >> 16)
	p[stackIndex+3] = byte(heapAddr >> 24)
	// SetUInt32(p, stackIndex+4, ln)
	p[stackIndex+4] = byte(ln)
	p[stackIndex+5] = byte(ln >> 8)
	p[stackIndex+6] = byte(ln >> 16)
	p[stackIndex+7] = byte(ln >> 24)
	copy(p[heapAddr:], s)
	return heapAddr + ln
}

// SetByteArray encode byte array || uint8 array to the payload buffer!
func SetByteArray(p []byte, s []byte, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], s)
	return heapAddr + ln
}

// SetInt8Array encode int8 array to the payload buffer!
func SetInt8Array(p []byte, s []int8, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeInt8SliceToByteSlice(s))
	return heapAddr + ln
}

// SetBoolArray encode bool array to the payload buffer!
func SetBoolArray(p []byte, s []bool, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeBoolSliceToByteSlice(s))
	return heapAddr + ln
}

// SetInt16Array encode int16 array to the payload buffer!
func SetInt16Array(p []byte, s []int16, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeInt16SliceToByteSlice(s))
	return heapAddr + (ln * 2)
}

// SetUInt16Array encode uint16 array to the payload buffer!
func SetUInt16Array(p []byte, s []uint16, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeUInt16SliceToByteSlice(s))
	return heapAddr + (ln * 2)
}

// SetInt32Array encode int32 array to the payload buffer!
func SetInt32Array(p []byte, s []int32, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeInt32SliceToByteSlice(s))
	return heapAddr + (ln * 4)
}

// SetUInt32Array encode uint32 array to the payload buffer!
func SetUInt32Array(p []byte, s []uint32, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeUInt32SliceToByteSlice(s))
	return heapAddr + (ln * 4)
}

// SetInt64Array encode int64 array to the payload buffer!
func SetInt64Array(p []byte, s []int64, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeInt64SliceToByteSlice(s))
	return heapAddr + (ln * 8)
}

// SetUInt64Array encode uint64 array to the payload buffer!
func SetUInt64Array(p []byte, s []uint64, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeUInt64SliceToByteSlice(s))
	return heapAddr + (ln * 8)
}

// SetFloat32Array encode float32 array to the payload buffer!
func SetFloat32Array(p []byte, s []float32, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeFloat32SliceToByteSlice(s))
	return heapAddr + (ln * 4)
}

// SetFloat64Array encode float64 array to the payload buffer!
func SetFloat64Array(p []byte, s []float64, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeFloat64SliceToByteSlice(s))
	return heapAddr + (ln * 8)
}

// SetComplex64Array encode complex64 array to the payload buffer!
func SetComplex64Array(p []byte, s []complex64, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeComplex64SliceToByteSlice(s))
	return heapAddr + (ln * 8)
}

// SetComplex128Array encode complex128 array to the payload buffer!
func SetComplex128Array(p []byte, s []complex128, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeComplex128SliceToByteSlice(s))
	return heapAddr + (ln * 16)
}

/*
**************************************************************************************************
*******************Dynamically size ARRAY inside other Dynamically size Array*******************
**************************************************************************************************
 */

// SetStringArray encode string array to the payload buffer!
func SetStringArray(p []byte, s []string, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	nextHeapAddr = heapAddr + (ln * 8)
	var eln uint32
	for i := 0; i < int(ln); i++ {
		eln = uint32(len(s[i]))
		SetUInt32(p, heapAddr, nextHeapAddr)
		SetUInt32(p, heapAddr+4, eln)
		copy(p[nextHeapAddr:], s[i])
		heapAddr += 8
		nextHeapAddr += eln
	}
	return
}
