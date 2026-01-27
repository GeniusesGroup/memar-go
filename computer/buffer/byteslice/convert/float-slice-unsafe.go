/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"reflect"
	"unsafe"
)

// UnsafeFloat32SliceToByteSlice returns ...
func UnsafeFloat32SliceToByteSlice(req []float32) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len * 4
	resStruct.Cap = reqStruct.Cap * 4
	return
}

// UnsafeByteSliceToFloat32Slice returns ...
func UnsafeByteSliceToFloat32Slice(req []byte) (res []float32) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len / 4
	resStruct.Cap = reqStruct.Cap / 4
	return
}

// UnsafeFloat64SliceToByteSlice returns ...
func UnsafeFloat64SliceToByteSlice(req []float64) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len * 8
	resStruct.Cap = reqStruct.Cap * 8
	return
}

// UnsafeByteSliceToFloat64Slice returns ...
func UnsafeByteSliceToFloat64Slice(req []byte) (res []float64) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len / 8
	resStruct.Cap = reqStruct.Cap / 8
	return
}
