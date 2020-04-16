/* For license and copyright information please see LEGAL file in repository */

package syllab

/*
	************************************************************
	**********************Fixed Size ARRAY**********************
	************************************************************
	***********************PAY ATTENTION************************
	If you want fixed sized array other than standard golang types use first function and edit it for your usage!
	Use code generation to make specific size array in return by ChaparKhane!
*/

// SetnByte encode n BYTE to the payload buffer.
// If you want to set fixed sized array, You can copy function and edit it for your usage! e.g.
// for set a [2]byte  ```copy(p[offset:], b[:])```
// for get a [32]byte  ```copy(p[offset:], b[:])```
func SetnByte(p []byte, offset uint32, b []byte) {
	copy(p[offset:], b[:])
	// TODO : check assignment performance!
	// p[offset] = b[0]
	// p[offset+1] = b[1]
	// ...
}

// SetByte encode BYTE to the payload buffer.
func SetByte(p []byte, offset uint32, b byte) {
	p[offset] = b
}

// SetInt8 encode INT8 to the payload buffer.
func SetInt8(p []byte, offset uint32, n int8) {
	p[offset] = byte(n)
	return
}

// SetUInt8 encode UINT8 to the payload buffer.
func SetUInt8(p []byte, offset uint32, n uint8) {
	p[offset] = byte(n)
}

// SetBool encode BOOL to the payload buffer.
func SetBool(p []byte, offset uint32, b bool) {
	if b {
		p[offset] = 1
	} else {
		p[offset] = 0
	}
}

// SetInt16 encode INT16 to the payload buffer.
func SetInt16(p []byte, offset uint32, n int16) {
	p[offset] = byte(n)
	p[offset+1] = byte(n >> 8)
}

// SetUInt16 encode UINT16 to the payload buffer.
func SetUInt16(p []byte, offset uint32, n uint16) {
	p[offset] = byte(n)
	p[offset+1] = byte(n >> 8)
}

// SetInt32 encode INT32 to the payload buffer.
func SetInt32(p []byte, offset uint32, n int32) {
	p[offset] = byte(n)
	p[offset+1] = byte(n >> 8)
	p[offset+2] = byte(n >> 16)
	p[offset+3] = byte(n >> 24)
}

// SetUInt32 encode UINT32 to the payload buffer.
func SetUInt32(p []byte, offset uint32, n uint32) {
	p[offset] = byte(n)
	p[offset+1] = byte(n >> 8)
	p[offset+2] = byte(n >> 16)
	p[offset+3] = byte(n >> 24)
}

// SetFloat32 encode FLOAT32 to the payload buffer.
func SetFloat32(p []byte, offset uint32, n float32) {
	SetUInt32(p, offset, uint32(n))
}

// SetInt64 encode INT64 to the payload buffer.
func SetInt64(p []byte, offset uint32, n int64) {
	p[offset] = byte(n)
	p[offset+1] = byte(n >> 8)
	p[offset+2] = byte(n >> 16)
	p[offset+3] = byte(n >> 24)
	p[offset+4] = byte(n >> 32)
	p[offset+5] = byte(n >> 40)
	p[offset+6] = byte(n >> 48)
	p[offset+7] = byte(n >> 56)
}

// SetUInt64 encode UINT64 to the payload buffer.
func SetUInt64(p []byte, offset uint32, n uint64) {
	p[offset] = byte(n)
	p[offset+1] = byte(n >> 8)
	p[offset+2] = byte(n >> 16)
	p[offset+3] = byte(n >> 24)
	p[offset+4] = byte(n >> 32)
	p[offset+5] = byte(n >> 40)
	p[offset+6] = byte(n >> 48)
	p[offset+7] = byte(n >> 56)
}

// SetFloat64 encode FLOAT64 to the payload buffer.
func SetFloat64(p []byte, offset uint32, n float64) {
	SetUInt64(p, offset, uint64(n))
}

// SetComplex64 encode COMPLEX64 to the payload buffer.
func SetComplex64(p []byte, offset uint32, n complex64) {
	SetFloat32(p, offset, real(n))
	SetFloat32(p, offset+3, imag(n))
}

// SetComplex128 encode COMPLEX128 to the payload buffer.
func SetComplex128(p []byte, offset uint32, n complex128) {
	SetFloat64(p, offset, real(n))
	SetFloat64(p, offset+7, imag(n))
}

/*
************************************************************
*******************Dynamically size ARRAY*******************
************************************************************
 */

// SetArrayAddress encode array address to the payload buffer.
func SetArrayAddress(p []byte, offset uint32, n uint32) {
	SetUInt32(p, offset, n)
}

// SetArrayLength encode array length to the payload buffer.
func SetArrayLength(p []byte, offset uint32, n uint32) {
	SetUInt32(p, offset+3, n)
}

// SetString encode string to the payload buffer!
func SetString(p []byte, s string, offset uint32, arrayAddress uint32) {
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, uint32(len(s)))
	copy(p[arrayAddress:], s)
}

// SetByteArray encode byte array to the payload buffer!
func SetByteArray(p []byte, s []byte, offset uint32, arrayAddress uint32) {
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, uint32(len(s)))
	copy(p[arrayAddress:], s)
}

// SetInt8Array encode int8 array to the payload buffer!
func SetInt8Array(p []byte, s []int8, offset uint32, arrayAddress uint32) {
	var len uint32 = uint32(len(s))
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, len)
	var i uint32
	for i = 0; i <= len; i++ {
		p[arrayAddress+i] = byte(s[i])
	}
}

// SetUInt8Array encode uint8 array to the payload buffer!
func SetUInt8Array(p []byte, s []uint8, offset uint32, arrayAddress uint32) {
	var len uint32 = uint32(len(s))
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, len)
	var i uint32
	for i = 0; i <= len; i++ {
		p[arrayAddress+i] = byte(s[i])
	}
}

// SetBoolArray encode bool array to the payload buffer!
func SetBoolArray(p []byte, s []bool, offset uint32, arrayAddress uint32) {
	var len uint32 = uint32(len(s))
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, len)
	var i uint32
	for i = 0; i <= len; i++ {
		SetBool(p, arrayAddress+i, s[i])
	}
}

// SetInt16Array encode int16 array to the payload buffer!
func SetInt16Array(p []byte, s []int16, offset uint32, arrayAddress uint32) {
	var len uint32 = uint32(len(s))
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, len)
	var i uint32
	for i = 0; i <= len; i++ {
		SetInt16(p, arrayAddress+(i*2), s[i])
	}
}

// SetUInt16Array encode uint16 array to the payload buffer!
func SetUInt16Array(p []byte, s []uint16, offset uint32, arrayAddress uint32) {
	var len uint32 = uint32(len(s))
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, len)
	var i uint32
	for i = 0; i <= len; i++ {
		SetUInt16(p, arrayAddress+(i*2), s[i])
	}
}

// SetInt32Array encode int32 array to the payload buffer!
func SetInt32Array(p []byte, s []int32, offset uint32, arrayAddress uint32) {
	var len uint32 = uint32(len(s))
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, len)
	var i uint32
	for i = 0; i <= len; i++ {
		SetInt32(p, arrayAddress+(i*4), s[i])
	}
}

// SetUInt32Array encode uint32 array to the payload buffer!
func SetUInt32Array(p []byte, s []uint32, offset uint32, arrayAddress uint32) {
	var len uint32 = uint32(len(s))
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, len)
	var i uint32
	for i = 0; i <= len; i++ {
		SetUInt32(p, arrayAddress+(i*4), s[i])
	}
}

// SetFloat32Array encode float32 array to the payload buffer!
func SetFloat32Array(p []byte, s []float32, offset uint32, arrayAddress uint32) {
	var len uint32 = uint32(len(s))
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, len)
	var i uint32
	for i = 0; i <= len; i++ {
		SetFloat32(p, arrayAddress+(i*4), s[i])
	}
}

// SetInt64Array encode int64 array to the payload buffer!
func SetInt64Array(p []byte, s []int64, offset uint32, arrayAddress uint32) {
	var len uint32 = uint32(len(s))
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, len)
	var i uint32
	for i = 0; i <= len; i++ {
		SetInt64(p, arrayAddress+(i*8), s[i])
	}
}

// SetUInt64Array encode uint64 array to the payload buffer!
func SetUInt64Array(p []byte, s []uint64, offset uint32, arrayAddress uint32) {
	var len uint32 = uint32(len(s))
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, len)
	var i uint32
	for i = 0; i <= len; i++ {
		SetUInt64(p, arrayAddress+(i*8), s[i])
	}
}

// SetFloat64Array encode float64 array to the payload buffer!
func SetFloat64Array(p []byte, s []float64, offset uint32, arrayAddress uint32) {
	var len uint32 = uint32(len(s))
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, len)
	var i uint32
	for i = 0; i <= len; i++ {
		SetFloat64(p, arrayAddress+(i*8), s[i])
	}
}

// SetComplex64Array encode complex64 array to the payload buffer!
func SetComplex64Array(p []byte, s []complex64, offset uint32, arrayAddress uint32) {
	var len uint32 = uint32(len(s))
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, len)
	var i uint32
	for i = 0; i <= len; i++ {
		SetComplex64(p, arrayAddress+(i*8), s[i])
	}
}

// SetComplex128Array encode complex128 array to the payload buffer!
func SetComplex128Array(p []byte, s []complex128, offset uint32, arrayAddress uint32) {
	var len uint32 = uint32(len(s))
	SetArrayAddress(p, offset, arrayAddress)
	SetArrayLength(p, offset+3, len)
	var i uint32
	for i = 0; i <= len; i++ {
		SetComplex128(p, arrayAddress+(i*8), s[i])
	}
}
