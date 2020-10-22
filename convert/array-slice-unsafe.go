/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"reflect"
	"unsafe"
)

// Unsafe16ByteArraySliceToByteSlice returns ...
func Unsafe16ByteArraySliceToByteSlice(req [][16]byte) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len * 16,
		Cap:  reqStruct.Cap * 16,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceTo16ByteArraySlice returns ...
func UnsafeByteSliceTo16ByteArraySlice(req []byte) (res [][16]byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len / 16,
		Cap:  reqStruct.Cap / 16,
	}

	res = *(*[][16]byte)(unsafe.Pointer(&resStruct))
	// TODO::: really need do this here??
	// runtime.KeepAlive(req)
	return
}

// Unsafe32ByteArraySliceToByteSlice returns ...
func Unsafe32ByteArraySliceToByteSlice(req [][32]byte) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len * 32,
		Cap:  reqStruct.Cap * 32,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceTo32ByteArraySlice returns ...
func UnsafeByteSliceTo32ByteArraySlice(req []byte) (res [][32]byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len / 32,
		Cap:  reqStruct.Cap / 32,
	}

	res = *(*[][32]byte)(unsafe.Pointer(&resStruct))
	// TODO::: really need do this here??
	// runtime.KeepAlive(req)
	return
}
