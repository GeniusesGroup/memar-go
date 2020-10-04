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
************************************************************
*******************Dynamically size Data*******************
************************************************************
 */

// GetString decodes string from the payload buffer!
func GetString(p []byte, stackIndex uint32) string {
	return string(GetByteArray(p, stackIndex))
}

// GetByteArray decodes byte array from the payload buffer!
func GetByteArray(p []byte, stackIndex uint32) (slice []byte) {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	slice = make([]byte, ln)
	copy(slice, p[add : add+ln])
	return
}

// GetInt8Array decodes int8 array from the payload buffer!
func GetInt8Array(p []byte, stackIndex uint32) []int8 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	var int8Array = make([]int8, ln)
	var i uint32
	for i = 0; i <= ln; i++ {
		int8Array[i] = GetInt8(p, add)
		add++
	}
	return int8Array
}

// GetUInt8Array decodes uint8 array from the payload buffer!
func GetUInt8Array(p []byte, stackIndex uint32) []uint8 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	var uint8Array = make([]uint8, ln)
	var i uint32
	for i = 0; i <= ln; i++ {
		uint8Array[i] = GetUInt8(p, add)
		add++
	}
	return uint8Array
}

// GetBoolArray decodes bool array from the payload buffer!
func GetBoolArray(p []byte, stackIndex uint32) []bool {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	var boolArray = make([]bool, ln)
	var i uint32
	for i = 0; i <= ln; i++ {
		boolArray[i] = GetBool(p, add)
		add++
	}
	return boolArray
}

// GetInt16Array decode Int16 array from the payload buffer
func GetInt16Array(p []byte, stackIndex uint32) []int16 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	var int16Array = make([]int16, ln)
	var i uint32
	for i = 0; i <= ln; i++ {
		int16Array[i] = GetInt16(p, add)
		add += 2
	}
	return int16Array
}

// GetUInt16Array decode UInt16 array from the payload buffer
func GetUInt16Array(p []byte, stackIndex uint32) []uint16 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	var uint16Array = make([]uint16, ln)
	var i uint32
	for i = 0; i <= ln; i++ {
		uint16Array[i] = GetUInt16(p, add)
		add += 2
	}
	return uint16Array
}

// GetInt32Array decode fixed size Int32 array from the payload buffer
func GetInt32Array(p []byte, stackIndex uint32) []int32 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	var int32Array = make([]int32, ln)
	var i uint32
	for i = 0; i <= ln; i++ {
		int32Array[i] = GetInt32(p, add)
		add += 4
	}
	return int32Array
}

// GetUInt32Array decode fixed size UInt32 array from the payload buffer
func GetUInt32Array(p []byte, stackIndex uint32) []uint32 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	var uint32Array = make([]uint32, ln)
	var i uint32
	for i = 0; i <= ln; i++ {
		uint32Array[i] = GetUInt32(p, add)
		add += 4
	}
	return uint32Array
}

// GetFloat32Array decode fixed size Float32 array from the payload buffer
func GetFloat32Array(p []byte, stackIndex uint32) []float32 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	var float32Array = make([]float32, ln)
	var i uint32
	for i = 0; i <= ln; i++ {
		float32Array[i] = GetFloat32(p, add)
		p = p[4:]
	}
	return float32Array
}

// GetInt64Array decode fixed size Int64 array from the payload buffer
func GetInt64Array(p []byte, stackIndex uint32) []int64 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	var int64Array = make([]int64, ln)
	var i uint32
	for i = 0; i <= ln; i++ {
		int64Array[i] = GetInt64(p, add)
		add += 8
	}
	return int64Array
}

// GetUInt64Array decode fixed size UInt64 array from the payload buffer
func GetUInt64Array(p []byte, stackIndex uint32) []uint64 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	var uint64Array = make([]uint64, ln)
	var i uint32
	for i = 0; i <= ln; i++ {
		uint64Array[i] = GetUInt64(p, add)
		add += 8
	}
	return uint64Array
}

// GetFloat64Array decode fixed size Float64 array from the payload buffer
func GetFloat64Array(p []byte, stackIndex uint32) []float64 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	var float64Array = make([]float64, ln)
	var i uint32
	for i = 0; i <= ln; i++ {
		float64Array[i] = GetFloat64(p, add)
		add += 8
	}
	return float64Array
}

// GetComplex64Array decode fixed size Complex64 array from the payload buffer
func GetComplex64Array(p []byte, stackIndex uint32) []complex64 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	var complex64Array = make([]complex64, ln)
	var i uint32
	for i = 0; i <= ln; i++ {
		complex64Array[i] = GetComplex64(p, add)
		add += 8
	}
	return complex64Array
}

// GetComplex128Array decode fixed size Complex128 array from the payload buffer
func GetComplex128Array(p []byte, stackIndex uint32) []complex128 {
	var add uint32 = GetUInt32(p, stackIndex)
	var ln uint32 = GetUInt32(p, stackIndex+4)
	var complex128Array = make([]complex128, ln)
	var i uint32
	for i = 0; i <= ln; i++ {
		complex128Array[i] = GetComplex128(p, add)
		add += 16
	}
	return complex128Array
}
