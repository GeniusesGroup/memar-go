/* For license and copyright information please see LEGAL file in repository */

package syllab

import (
	"unsafe"
)

// UnsafeGetServiceID decodes service ID from payload slice in unsafe manner.
// This would blow up silently if len(p) < 4
func UnsafeGetServiceID(p []byte) uint32 {
	return *(*uint32)(unsafe.Pointer(&p[0]))
}

// UnsafeGetErrorCode decodes error code from payload slice in unsafe manner.
// This would blow up silently if len(p) < 4
func UnsafeGetErrorCode(p []byte) uint32 {
	return *(*uint32)(unsafe.Pointer(&p[3]))
}

/*
	************************************************************
	**********************Fixed Size ARRAY**********************
	************************************************************
	***********************PAY ATTENTION************************
	If you want fixed sized array other than standard golang types use first function and edit it for your usage!
	Use code generation to make specific size array in return by ChaparKhane!
*/

// UnsafeGetnByte decodes n byte from the payload buffer in unsafe manner.
// This would blow up silently if len(p) < n
// If you want fixed sized array from function below, You can copy function and edit it for your usage! e.g.
// for get a [2]byte  ```var array [2]byte = *(*[2]byte)(unsafe.Pointer(&p[offset]))```
// for get a [32]byte  ```var array [32]byte = *(*[32]byte)(unsafe.Pointer(&p[offset]))```
func UnsafeGetnByte(p []byte, offset uint32) []byte {
	return *(*[]byte)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetByte decodes BYTE from the payload buffer in unsafe manner.
// This would blow up silently if len(p) < 1
func UnsafeGetByte(p []byte, offset uint32) byte {
	return *(*byte)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetInt8 decodes INT8 from the payload buffer in unsafe manner.
// This would blow up silently if len(p) < 1
func UnsafeGetInt8(p []byte, offset uint32) int8 {
	return *(*int8)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetUInt8 decodes UINT8 from the payload buffer in unsafe manner.
// This would blow up silently if len(p) < 1
func UnsafeGetUInt8(p []byte, offset uint32) uint8 {
	return *(*uint8)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetBool decodes BOOL from the payload buffer in unsafe manner.
// This would blow up silently if len(p) < 1
func UnsafeGetBool(p []byte, offset uint32) bool {
	return *(*bool)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetInt16 decodes INT16 from the payload buffer in unsafe manner.
// This would blow up silently if len(p) < 2
func UnsafeGetInt16(p []byte, offset uint32) int16 {
	return *(*int16)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetUInt16 decodes UINT16 from the payload buffer in unsafe manner.
// This would blow up silently if len(p) < 2
func UnsafeGetUInt16(p []byte, offset uint32) uint16 {
	return *(*uint16)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetInt32 decodes INT32 from the payload buffer in unsafe manner.
// This would blow up silently if len(p) < 4
func UnsafeGetInt32(p []byte, offset uint32) int32 {
	return *(*int32)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetUInt32 decodes UINT32 from the payload buffer in unsafe manner.
// This would blow up silently if len(p) < 4
func UnsafeGetUInt32(p []byte, offset uint32) uint32 {
	return *(*uint32)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetFloat32 decodes FLOAT32 from the payload buffer in unsafe manner.
// This would blow up silently if len(p) < 4
func UnsafeGetFloat32(p []byte, offset uint32) float32 {
	return *(*float32)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetInt64 decodes INT64 from the payload buffer in unsafe manner.
// This would blow up silently if len(p) < 8
func UnsafeGetInt64(p []byte, offset uint32) int64 {
	return *(*int64)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetUInt64 decodes UINT64 from the payload buffer in unsafe manner.
// This would blow up silently if len(p) < 8
func UnsafeGetUInt64(p []byte, offset uint32) uint64 {
	return *(*uint64)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetFloat64 decodes FLOAT64 from the payload buffer in unsafe manner.
// This would blow up silently if len(p) < 8
func UnsafeGetFloat64(p []byte, offset uint32) float64 {
	return *(*float64)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetComplex64 decodes COMPLEX64 from the payload buffer in unsafe manner.
func UnsafeGetComplex64(p []byte, offset uint32) complex64 {
	return complex(UnsafeGetFloat32(p, offset), UnsafeGetFloat32(p, offset+3))
}

// UnsafeGetComplex128 decodes COMPLEX128 from the payload buffer in unsafe manner.
func UnsafeGetComplex128(p []byte, offset uint32) complex128 {
	return complex(UnsafeGetFloat64(p, offset), UnsafeGetFloat64(p, offset+7))
}

/*
************************************************************
*******************Dynamically size ARRAY*******************
************************************************************
 */

// UnsafeGetArrayAddress decodes string address from the payload buffer in unsafe manner!
// This would blow up silently if len(p) < 4
func UnsafeGetArrayAddress(p []byte, offset uint32) uint32 {
	return *(*uint32)(unsafe.Pointer(&p[offset]))
}

// UnsafeGetArrayLength decodes string length from the payload buffer in unsafe manner!
// This would blow up silently if len(p) < 8
func UnsafeGetArrayLength(p []byte, offset uint32) uint32 {
	return *(*uint32)(unsafe.Pointer(&p[offset+3]))
}

// UnsafeGetString decodes string from the payload buffer in unsafe manner!
// This would blow up silently if return not limit by make(string, 0, UnsafeGetArrayLength(p))
func UnsafeGetString(p []byte, offset uint32) string {
	return *(*string)(unsafe.Pointer(&p[UnsafeGetArrayAddress(p, offset)]))
}

// UnsafeGetByteArray decodes byte slice from the payload buffer in unsafe manner!
// This would blow up silently if return not limit by make([]byte, 0, UnsafeGetArrayLength(p))
func UnsafeGetByteArray(p []byte, offset uint32) []byte {
	return *(*[]byte)(unsafe.Pointer(&p[UnsafeGetArrayAddress(p, offset)]))
}
