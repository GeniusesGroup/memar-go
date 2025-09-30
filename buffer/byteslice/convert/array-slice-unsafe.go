/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"reflect"
	"unsafe"
)

// Unsafe16ByteArraySliceToByteSlice returns ...
func Unsafe16ByteArraySliceToByteSlice(req [][16]byte) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len * 16
	resStruct.Cap = reqStruct.Cap * 16
	return
}

// UnsafeByteSliceTo16ByteArraySlice returns ...
func UnsafeByteSliceTo16ByteArraySlice(req []byte) (res [][16]byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len / 16
	resStruct.Cap = reqStruct.Cap / 16
	return
}

// Unsafe32ByteArraySliceToByteSlice returns ...
func Unsafe32ByteArraySliceToByteSlice(req [][32]byte) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len * 32
	resStruct.Cap = reqStruct.Cap * 32
	return
}

// UnsafeByteSliceTo32ByteArraySlice returns ...
func UnsafeByteSliceTo32ByteArraySlice(req []byte) (res [][32]byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len / 32
	resStruct.Cap = reqStruct.Cap / 32
	return
}
