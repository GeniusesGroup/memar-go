/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"reflect"
	"unsafe"
)

// UnsafeComplex64SliceToByteSlice returns ...
func UnsafeComplex64SliceToByteSlice(req []complex64) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len * 8
	resStruct.Cap = reqStruct.Cap * 8
	return
}

// UnsafeByteSliceToComplex64Slice returns ...
func UnsafeByteSliceToComplex64Slice(req []byte) (res []complex64) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len / 8
	resStruct.Cap = reqStruct.Cap / 8
	return
}

// UnsafeComplex128SliceToByteSlice returns ...
func UnsafeComplex128SliceToByteSlice(req []complex128) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len * 16
	resStruct.Cap = reqStruct.Cap * 16
	return
}

// UnsafeByteSliceToComplex128Slice returns ...
func UnsafeByteSliceToComplex128Slice(req []byte) (res []complex128) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len / 16
	resStruct.Cap = reqStruct.Cap / 16
	return
}
