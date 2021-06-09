/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"reflect"
	"unsafe"
)

// UnsafeInt8SliceToByteSlice returns ...
func UnsafeInt8SliceToByteSlice(req []int8) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len
	resStruct.Cap = reqStruct.Cap
	return
}

// UnsafeByteSliceToInt8Slice returns ...
func UnsafeByteSliceToInt8Slice(req []byte) (res []int8) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len
	resStruct.Cap = reqStruct.Cap
	return
}

// UnsafeInt16SliceToByteSlice returns ...
func UnsafeInt16SliceToByteSlice(req []int16) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len * 2
	resStruct.Cap = reqStruct.Cap * 2
	return
}

// UnsafeByteSliceToInt16Slice returns ...
func UnsafeByteSliceToInt16Slice(req []byte) (res []int16) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len / 2
	resStruct.Cap = reqStruct.Cap / 2
	return
}

// UnsafeInt32SliceToByteSlice returns ...
func UnsafeInt32SliceToByteSlice(req []int32) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len * 4
	resStruct.Cap = reqStruct.Cap * 4
	return
}

// UnsafeByteSliceToInt32Slice returns ...
func UnsafeByteSliceToInt32Slice(req []byte) (res []int32) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len / 4
	resStruct.Cap = reqStruct.Cap / 4
	return
}

// UnsafeInt64SliceToByteSlice returns ...
func UnsafeInt64SliceToByteSlice(req []int64) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len * 8
	resStruct.Cap = reqStruct.Cap * 8
	return
}

// UnsafeByteSliceToInt64Slice returns ...
func UnsafeByteSliceToInt64Slice(req []byte) (res []int64) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len / 8
	resStruct.Cap = reqStruct.Cap / 8
	return
}
