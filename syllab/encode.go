/* For license and copyright information please see LEGAL file in repository */

package syllab

/*
		************************************************************
		**********************Fixed Size ARRAY**********************
		************************************************************
		***********************PAY ATTENTION************************
	By use below helper functions you can't achieve max performance!
	Use code generation to prevent unneeded memory alloc by CompleteMethods()!
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
************************************************************
*******************Dynamically size ARRAY*******************
************************************************************
 */

// SetString encode string to the payload buffer!
func SetString(p []byte, s string, stackIndex uint32, heapAddr uint32) {
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(len(s)))
	copy(p[heapAddr:], s)
}

// SetByteArray encode byte array to the payload buffer!
func SetByteArray(p []byte, s []byte, stackIndex uint32, heapAddr uint32) {
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(len(s)))
	copy(p[heapAddr:], s)
}

// SetInt8Array encode int8 array to the payload buffer!
func SetInt8Array(p []byte, s []int8, stackIndex uint32, heapAddr uint32) {
	var ln = len(s)
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(ln))
	for i := 0; i < ln; i++ {
		p[heapAddr] = byte(s[i])
		heapAddr++
	}
}

// SetUInt8Array encode uint8 array to the payload buffer!
func SetUInt8Array(p []byte, s []uint8, stackIndex uint32, heapAddr uint32) {
	var ln = len(s)
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(ln))
	for i := 0; i < ln; i++ {
		p[heapAddr] = byte(s[i])
		heapAddr++
	}
}

// SetBoolArray encode bool array to the payload buffer!
func SetBoolArray(p []byte, s []bool, stackIndex uint32, heapAddr uint32) {
	var ln = len(s)
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(ln))
	for i := 0; i < ln; i++ {
		SetBool(p, heapAddr, s[i])
		heapAddr += 2
	}
}

// SetInt16Array encode int16 array to the payload buffer!
func SetInt16Array(p []byte, s []int16, stackIndex uint32, heapAddr uint32) {
	var ln = len(s)
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(ln))
	for i := 0; i <= ln; i++ {
		SetInt16(p, heapAddr, s[i])
		heapAddr += 2
	}
}

// SetUInt16Array encode uint16 array to the payload buffer!
func SetUInt16Array(p []byte, s []uint16, stackIndex uint32, heapAddr uint32) {
	var ln = len(s)
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(ln))
	for i := 0; i <= ln; i++ {
		SetUInt16(p, heapAddr, s[i])
		heapAddr += 2
	}
}

// SetInt32Array encode int32 array to the payload buffer!
func SetInt32Array(p []byte, s []int32, stackIndex uint32, heapAddr uint32) {
	var ln = len(s)
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(ln))
	for i := 0; i <= ln; i++ {
		SetInt32(p, heapAddr, s[i])
		heapAddr += 4
	}
}

// SetUInt32Array encode uint32 array to the payload buffer!
func SetUInt32Array(p []byte, s []uint32, stackIndex uint32, heapAddr uint32) {
	var ln = len(s)
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(ln))
	for i := 0; i <= ln; i++ {
		SetUInt32(p, heapAddr, s[i])
		heapAddr += 4
	}
}

// SetFloat32Array encode float32 array to the payload buffer!
func SetFloat32Array(p []byte, s []float32, stackIndex uint32, heapAddr uint32) {
	var ln = len(s)
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(ln))
	for i := 0; i <= ln; i++ {
		SetFloat32(p, heapAddr, s[i])
		heapAddr += 4
	}
}

// SetInt64Array encode int64 array to the payload buffer!
func SetInt64Array(p []byte, s []int64, stackIndex uint32, heapAddr uint32) {
	var ln = len(s)
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(ln))
	for i := 0; i <= ln; i++ {
		SetInt64(p, heapAddr, s[i])
		heapAddr += 8
	}
}

// SetUInt64Array encode uint64 array to the payload buffer!
func SetUInt64Array(p []byte, s []uint64, stackIndex uint32, heapAddr uint32) {
	var ln = len(s)
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(ln))
	for i := 0; i <= ln; i++ {
		SetUInt64(p, heapAddr, s[i])
		heapAddr += 8
	}
}

// SetFloat64Array encode float64 array to the payload buffer!
func SetFloat64Array(p []byte, s []float64, stackIndex uint32, heapAddr uint32) {
	var ln = len(s)
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(ln))
	for i := 0; i <= ln; i++ {
		SetFloat64(p, heapAddr, s[i])
		heapAddr += 8
	}
}

// SetComplex64Array encode complex64 array to the payload buffer!
func SetComplex64Array(p []byte, s []complex64, stackIndex uint32, heapAddr uint32) {
	var ln = len(s)
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(ln))
	for i := 0; i <= ln; i++ {
		SetComplex64(p, heapAddr, s[i])
		heapAddr += 8
	}
}

// SetComplex128Array encode complex128 array to the payload buffer!
func SetComplex128Array(p []byte, s []complex128, stackIndex uint32, heapAddr uint32) {
	var ln = len(s)
	SetUInt32(p, stackIndex, heapAddr)
	SetUInt32(p, stackIndex+4, uint32(ln))
	for i := 0; i <= ln; i++ {
		SetComplex128(p, heapAddr, s[i])
		heapAddr += 16
	}
}
