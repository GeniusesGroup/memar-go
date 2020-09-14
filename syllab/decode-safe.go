/* For license and copyright information please see LEGAL file in repository */

package syllab

/*
	************************************************************
	**********************Fixed Size Data**********************
	************************************************************
	***********************PAY ATTENTION************************
	By use below helper functions you can't achieve max performance!
	Use code generation to prevent unneeded memory alloc by CompleteMethods()!
*/

// GetFixedByteArray decodes fixed sized byte array from the payload buffer.
// If you want array instead of slice from function below, You can copy function and edit it for your usage! e.g.
// for get a [2]byte  ```var array [2]byte = copy(array[:], p[offset:])```
// for get a [32]byte  ```var array [32]byte = copy(array[:], p[offset:])```
func GetFixedByteArray(p []byte, n uint32) (array []byte) {
	// TODO::: check copy(this func) vs assignment performance
	copy(array[:], p[:n])
	return
}

// GetByte decodes BYTE from the payload buffer.
func GetByte(p []byte) byte {
	return p[0]
}

// GetInt8 decodes INT8 from the payload buffer.
func GetInt8(p []byte) int8 {
	return int8(p[0])
}

// GetUInt8 decodes UINT8 from the payload buffer.
func GetUInt8(p []byte) uint8 {
	return uint8(p[0])
}

// GetBool decodes BOOL from the payload buffer.
func GetBool(p []byte) bool {
	return p[0] == 1
}

// GetInt16 decodes INT16 from the payload buffer.
func GetInt16(p []byte) int16 {
	return int16(p[0]) | int16(p[1])<<8
}

// GetUInt16 decodes UINT16 from the payload buffer.
func GetUInt16(p []byte) uint16 {
	return uint16(p[0]) | uint16(p[1])<<8
}

// GetInt32 decodes INT32 from the payload buffer.
func GetInt32(p []byte) int32 {
	return int32(p[0]) | int32(p[1])<<8 | int32(p[2])<<16 | int32(p[3])<<24
}

// GetUInt32 decodes UINT32 from the payload buffer.
func GetUInt32(p []byte) uint32 {
	return uint32(p[0]) | uint32(p[1])<<8 | uint32(p[2])<<16 | uint32(p[3])<<24
}

// GetFloat32 decodes FLOAT32 from the payload buffer.
func GetFloat32(p []byte) float32 {
	return float32(GetUInt32(p))
}

// GetInt64 decodes INT64 from the payload buffer.
func GetInt64(p []byte) int64 {
	return int64(p[0]) | int64(p[1])<<8 | int64(p[2])<<16 | int64(p[3])<<24 |
		int64(p[4])<<32 | int64(p[5])<<40 | int64(p[6])<<48 | int64(p[7])<<56
}

// GetUInt64 decodes UINT64 from the payload buffer.
func GetUInt64(p []byte) uint64 {
	return uint64(p[0]) | uint64(p[1])<<8 | uint64(p[2])<<16 | uint64(p[3])<<24 |
		uint64(p[4])<<32 | uint64(p[5])<<40 | uint64(p[6])<<48 | uint64(p[7])<<56
}

// GetFloat64 decodes FLOAT64 from the payload buffer.
func GetFloat64(p []byte) float64 {
	return float64(GetUInt64(p))
}

// GetComplex64 decodes COMPLEX64 from the payload buffer.
func GetComplex64(p []byte) complex64 {
	return complex(GetFloat32(p), GetFloat32(p[4:]))
}

// GetComplex128 decodes COMPLEX128 from the payload buffer.
func GetComplex128(p []byte) complex128 {
	return complex(GetFloat64(p), GetFloat64(p[8:]))
}

/*
************************************************************
*******************Dynamically size Data*******************
************************************************************
 */

// GetString decodes string from the payload buffer!
func GetString(p []byte, offset uint32) string {
	return string(GetByteArray(p, offset))
}

// GetByteArray decodes byte array from the payload buffer!
func GetByteArray(p []byte, offset uint32) []byte {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	return p[add : add+len]
}

// GetInt8Array decodes int8 array from the payload buffer!
func GetInt8Array(p []byte, offset uint32) []int8 {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	var int8Array = make([]int8, len)
	p = p[add:]
	var i uint32
	for i = 0; i <= len; i++ {
		int8Array[i] = GetInt8(p)
		p = p[1:]
	}
	return int8Array
}

// GetUInt8Array decodes uint8 array from the payload buffer!
func GetUInt8Array(p []byte, offset uint32) []uint8 {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	var uint8Array = make([]uint8, len)
	p = p[add:]
	var i uint32
	for i = 0; i <= len; i++ {
		uint8Array[i] = GetUInt8(p)
		p = p[1:]
	}
	return uint8Array
}

// GetBoolArray decodes bool array from the payload buffer!
func GetBoolArray(p []byte, offset uint32) []bool {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	var boolArray = make([]bool, len)
	p = p[add:]
	var i uint32
	for i = 0; i <= len; i++ {
		boolArray[i] = GetBool(p)
		p = p[1:]
	}
	return boolArray
}

// GetInt16Array decode Int16 array from the payload buffer
func GetInt16Array(p []byte, offset uint32) []int16 {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	var int16Array = make([]int16, len)
	p = p[add:]
	var i uint32
	for i = 0; i <= len; i++ {
		int16Array[i] = GetInt16(p)
		p = p[2:]
	}
	return int16Array
}

// GetUInt16Array decode UInt16 array from the payload buffer
func GetUInt16Array(p []byte, offset uint32) []uint16 {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	var uint16Array = make([]uint16, len)
	p = p[add:]
	var i uint32
	for i = 0; i <= len; i++ {
		uint16Array[i] = GetUInt16(p)
		p = p[2:]
	}
	return uint16Array
}

// GetInt32Array decode fixed size Int32 array from the payload buffer
func GetInt32Array(p []byte, offset uint32) []int32 {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	var int32Array = make([]int32, len)
	p = p[add:]
	var i uint32
	for i = 0; i <= len; i++ {
		int32Array[i] = GetInt32(p)
		p = p[4:]
	}
	return int32Array
}

// GetUInt32Array decode fixed size UInt32 array from the payload buffer
func GetUInt32Array(p []byte, offset uint32) []uint32 {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	var uint32Array = make([]uint32, len)
	p = p[add:]
	var i uint32
	for i = 0; i <= len; i++ {
		uint32Array[i] = GetUInt32(p)
		p = p[4:]
	}
	return uint32Array
}

// GetFloat32Array decode fixed size Float32 array from the payload buffer
func GetFloat32Array(p []byte, offset uint32) []float32 {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	var float32Array = make([]float32, len)
	p = p[add:]
	var i uint32
	for i = 0; i <= len; i++ {
		float32Array[i] = GetFloat32(p)
		p = p[4:]
	}
	return float32Array
}

// GetInt64Array decode fixed size Int64 array from the payload buffer
func GetInt64Array(p []byte, offset uint32) []int64 {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	var int64Array = make([]int64, len)
	p = p[add:]
	var i uint32
	for i = 0; i <= len; i++ {
		int64Array[i] = GetInt64(p)
		p = p[8:]
	}
	return int64Array
}

// GetUInt64Array decode fixed size UInt64 array from the payload buffer
func GetUInt64Array(p []byte, offset uint32) []uint64 {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	var uint64Array = make([]uint64, len)
	p = p[add:]
	var i uint32
	for i = 0; i <= len; i++ {
		uint64Array[i] = GetUInt64(p)
		p = p[8:]
	}
	return uint64Array
}

// GetFloat64Array decode fixed size Float64 array from the payload buffer
func GetFloat64Array(p []byte, offset uint32) []float64 {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	var float64Array = make([]float64, len)
	p = p[add:]
	var i uint32
	for i = 0; i <= len; i++ {
		float64Array[i] = GetFloat64(p)
		p = p[8:]
	}
	return float64Array
}

// GetComplex64Array decode fixed size Complex64 array from the payload buffer
func GetComplex64Array(p []byte, offset uint32) []complex64 {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	var complex64Array = make([]complex64, len)
	p = p[add:]
	var i uint32
	for i = 0; i <= len; i++ {
		complex64Array[i] = GetComplex64(p)
		p = p[8:]
	}
	return complex64Array
}

// GetComplex128Array decode fixed size Complex128 array from the payload buffer
func GetComplex128Array(p []byte, offset uint32) []complex128 {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	var complex128Array = make([]complex128, len)
	p = p[add:]
	var i uint32
	for i = 0; i <= len; i++ {
		complex128Array[i] = GetComplex128(p)
		p = p[16:]
	}
	return complex128Array
}
