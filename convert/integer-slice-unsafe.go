/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"reflect"
	"unsafe"
)

// UnsafeUInt8SliceToByteSlice returns ...
func UnsafeUInt8SliceToByteSlice(req []uint8) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len
	resStruct.Cap = reqStruct.Cap
	return
}

// UnsafeByteSliceToUInt8Slice returns ...
func UnsafeByteSliceToUInt8Slice(req []byte) (res []uint8) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len
	resStruct.Cap = reqStruct.Cap
	return
}

// UnsafeUInt16SliceToByteSlice returns ...
func UnsafeUInt16SliceToByteSlice(req []uint16) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len * 2
	resStruct.Cap = reqStruct.Cap * 2
	return
}

// UnsafeByteSliceToUInt16Slice returns ...
func UnsafeByteSliceToUInt16Slice(req []byte) (res []uint16) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len / 2
	resStruct.Cap = reqStruct.Cap / 2
	return
}

// UnsafeUInt32SliceToByteSlice returns ...
func UnsafeUInt32SliceToByteSlice(req []uint32) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len * 4
	resStruct.Cap = reqStruct.Cap * 4
	return
}

// UnsafeByteSliceToUInt32Slice returns ...
func UnsafeByteSliceToUInt32Slice(req []byte) (res []uint32) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len / 4
	resStruct.Cap = reqStruct.Cap / 4
	return
}

// UnsafeUInt64SliceToByteSlice returns ...
func UnsafeUInt64SliceToByteSlice(req []uint64) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len * 8
	resStruct.Cap = reqStruct.Cap * 8
	return
}

// UnsafeByteSliceToUInt64Slice returns ...
func UnsafeByteSliceToUInt64Slice(req []byte) (res []uint64) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len / 8
	resStruct.Cap = reqStruct.Cap / 8
	return
}
