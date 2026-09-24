/* For license and copyright information please see the LEGAL file in the code repository */

package syllab

import (
	"memar/codec/binary"
	"memar/computer/buffer/byteslice/convert"
)

/*
**************************************************************************************************
*****************************************Fixed Size ARRAY*****************************************
**************************************************************************************************
 */

func SetArray(p []byte, stackIndex uint32, a []byte) { copy(p[stackIndex:], a[:]) }
func SetByte(p []byte, stackIndex uint32, b byte)    { p[stackIndex] = b }
func SetInt8(p []byte, stackIndex uint32, n int8)    { p[stackIndex] = byte(n) }
func SetUInt8(p []byte, stackIndex uint32, n uint8)  { p[stackIndex] = byte(n) }
func SetBool(p []byte, stackIndex uint32, b bool) {
	if b {
		p[stackIndex] = 1
	} else {
		p[stackIndex] = 0
	}
}
func SetInt16(p []byte, stackIndex uint32, n int16) {
	binary.LittleEndian.PutUint16(p[stackIndex:], uint16(n))
}
func SetUInt16(p []byte, stackIndex uint32, n uint16) {
	binary.LittleEndian.PutUint16(p[stackIndex:], n)
}
func SetInt32(p []byte, stackIndex uint32, n int32) {
	binary.LittleEndian.PutUint32(p[stackIndex:], uint32(n))
}
func SetUInt32(p []byte, stackIndex uint32, n uint32) {
	binary.LittleEndian.PutUint32(p[stackIndex:], n)
}
func SetFloat32(p []byte, stackIndex uint32, n float32) { SetUInt32(p, stackIndex, uint32(n)) }
func SetInt64(p []byte, stackIndex uint32, n int64) {
	binary.LittleEndian.PutUint64(p[stackIndex:], uint64(n))
}
func SetUInt64(p []byte, stackIndex uint32, n uint64) {
	binary.LittleEndian.PutUint64(p[stackIndex:], n)
}
func SetFloat64(p []byte, stackIndex uint32, n float64) { SetUInt64(p, stackIndex, uint64(n)) }
func SetComplex64(p []byte, stackIndex uint32, n complex64) {
	SetFloat32(p, stackIndex, real(n))
	SetFloat32(p, stackIndex+3, imag(n))
}
func SetComplex128(p []byte, stackIndex uint32, n complex128) {
	SetFloat64(p, stackIndex, real(n))
	SetFloat64(p, stackIndex+7, imag(n))
}

/*
**************************************************************************************************
**************************************Dynamically size ARRAY**************************************
**************************************************************************************************
 */

// SetString encode string to the buffer
func SetString(p []byte, s string, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], s)
	return heapAddr + ln
}

// SetByteArray encode byte array || uint8 array to the buffer
func SetByteArray(p []byte, s []byte, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], s)
	return heapAddr + ln
}

// SetInt8Array encode int8 array to the buffer
func SetInt8Array(p []byte, s []int8, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeInt8SliceToByteSlice(s))
	return heapAddr + ln
}

// SetBoolArray encode bool array to the buffer
func SetBoolArray(p []byte, s []bool, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeBoolSliceToByteSlice(s))
	return heapAddr + ln
}

// SetInt16Array encode int16 array to the buffer
func SetInt16Array(p []byte, s []int16, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeInt16SliceToByteSlice(s))
	return heapAddr + (ln * 2)
}

// SetUInt16Array encode uint16 array to the buffer
func SetUInt16Array(p []byte, s []uint16, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeUInt16SliceToByteSlice(s))
	return heapAddr + (ln * 2)
}

// SetInt32Array encode int32 array to the buffer
func SetInt32Array(p []byte, s []int32, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeInt32SliceToByteSlice(s))
	return heapAddr + (ln * 4)
}

// SetUInt32Array encode uint32 array to the buffer
func SetUInt32Array(p []byte, s []uint32, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeUInt32SliceToByteSlice(s))
	return heapAddr + (ln * 4)
}

// SetInt64Array encode int64 array to the buffer
func SetInt64Array(p []byte, s []int64, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeInt64SliceToByteSlice(s))
	return heapAddr + (ln * 8)
}

// SetUInt64Array encode uint64 array to the buffer
func SetUInt64Array(p []byte, s []uint64, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeUInt64SliceToByteSlice(s))
	return heapAddr + (ln * 8)
}

// SetFloat32Array encode float32 array to the buffer
func SetFloat32Array(p []byte, s []float32, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeFloat32SliceToByteSlice(s))
	return heapAddr + (ln * 4)
}

// SetFloat64Array encode float64 array to the buffer
func SetFloat64Array(p []byte, s []float64, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeFloat64SliceToByteSlice(s))
	return heapAddr + (ln * 8)
}

// SetComplex64Array encode complex64 array to the buffer
func SetComplex64Array(p []byte, s []complex64, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.UnsafeComplex64SliceToByteSlice(s))
	return heapAddr + (ln * 8)
}

// SetComplex128Array encode complex128 array to the buffer
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

// SetStringArray encode string array to the buffer
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

// Set16ByteArrayArray encode [16]byte array to the buffer
func Set16ByteArrayArray(p []byte, s [][16]byte, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.Unsafe16ByteArraySliceToByteSlice(s))
	nextHeapAddr = heapAddr + (ln * 16)
	return
}

// Set32ByteArrayArray encode string array to the buffer
func Set32ByteArrayArray(p []byte, s [][32]byte, stackIndex uint32, heapAddr uint32) (nextHeapAddr uint32) {
	var ln = uint32(len(s))
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, ln)
	copy(p[heapAddr:], convert.Unsafe32ByteArraySliceToByteSlice(s))
	nextHeapAddr = heapAddr + (ln * 32)
	return
}
