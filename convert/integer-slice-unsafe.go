/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"reflect"
	"unsafe"
)

// UnsafeUInt8SliceToByteSlice returns ...
func UnsafeUInt8SliceToByteSlice(req []uint8) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len,
		Cap:  reqStruct.Cap,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceToUInt8Slice returns ...
func UnsafeByteSliceToUInt8Slice(req []byte) (res []uint8) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len,
		Cap:  reqStruct.Cap,
	}

	res = *(*[]uint8)(unsafe.Pointer(&resStruct))
	// TODO::: really need do this here??
	// runtime.KeepAlive(req)
	return
}

// UnsafeUInt16SliceToByteSlice returns ...
func UnsafeUInt16SliceToByteSlice(req []uint16) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len * 2,
		Cap:  reqStruct.Cap * 2,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceToUInt16Slice returns ...
func UnsafeByteSliceToUInt16Slice(req []byte) (res []uint16) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len / 2,
		Cap:  reqStruct.Cap / 2,
	}

	res = *(*[]uint16)(unsafe.Pointer(&resStruct))
	// TODO::: really need do this here??
	// runtime.KeepAlive(req)
	return
}

// UnsafeUInt32SliceToByteSlice returns ...
func UnsafeUInt32SliceToByteSlice(req []uint32) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len * 4,
		Cap:  reqStruct.Cap * 4,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceToUInt32Slice returns ...
func UnsafeByteSliceToUInt32Slice(req []byte) (res []uint32) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len / 4,
		Cap:  reqStruct.Cap / 4,
	}

	res = *(*[]uint32)(unsafe.Pointer(&resStruct))
	// TODO::: really need do this here??
	// runtime.KeepAlive(req)
	return
}

// UnsafeUInt64SliceToByteSlice returns ...
func UnsafeUInt64SliceToByteSlice(req []uint64) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len * 8,
		Cap:  reqStruct.Cap * 8,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceToUInt64Slice returns ...
func UnsafeByteSliceToUInt64Slice(req []byte) (res []uint64) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len / 8,
		Cap:  reqStruct.Cap / 8,
	}

	res = *(*[]uint64)(unsafe.Pointer(&resStruct))
	// TODO::: really need do this here??
	// runtime.KeepAlive(req)
	return
}
