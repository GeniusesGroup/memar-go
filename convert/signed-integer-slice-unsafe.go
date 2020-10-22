/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"reflect"
	"unsafe"
)

// UnsafeInt8SliceToByteSlice returns ...
func UnsafeInt8SliceToByteSlice(req []int8) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len,
		Cap:  reqStruct.Cap,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceToInt8Slice returns ...
func UnsafeByteSliceToInt8Slice(req []byte) (res []int8) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len,
		Cap:  reqStruct.Cap,
	}

	res = *(*[]int8)(unsafe.Pointer(&resStruct))
	// TODO::: really need do this here??
	// runtime.KeepAlive(req)
	return
}

// UnsafeInt16SliceToByteSlice returns ...
func UnsafeInt16SliceToByteSlice(req []int16) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len * 2,
		Cap:  reqStruct.Cap * 2,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceToInt16Slice returns ...
func UnsafeByteSliceToInt16Slice(req []byte) (res []int16) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len / 2,
		Cap:  reqStruct.Cap / 2,
	}

	res = *(*[]int16)(unsafe.Pointer(&resStruct))
	// TODO::: really need do this here??
	// runtime.KeepAlive(req)
	return
}

// UnsafeInt32SliceToByteSlice returns ...
func UnsafeInt32SliceToByteSlice(req []int32) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len * 4,
		Cap:  reqStruct.Cap * 4,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceToInt32Slice returns ...
func UnsafeByteSliceToInt32Slice(req []byte) (res []int32) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len / 4,
		Cap:  reqStruct.Cap / 4,
	}

	res = *(*[]int32)(unsafe.Pointer(&resStruct))
	// TODO::: really need do this here??
	// runtime.KeepAlive(req)
	return
}

// UnsafeInt64SliceToByteSlice returns ...
func UnsafeInt64SliceToByteSlice(req []int64) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len * 8,
		Cap:  reqStruct.Cap * 8,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceToInt64Slice returns ...
func UnsafeByteSliceToInt64Slice(req []byte) (res []int64) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len / 8,
		Cap:  reqStruct.Cap / 8,
	}

	res = *(*[]int64)(unsafe.Pointer(&resStruct))
	// TODO::: really need do this here??
	// runtime.KeepAlive(req)
	return
}
