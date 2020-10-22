/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"reflect"
	"unsafe"
)

// UnsafeFloat32SliceToByteSlice returns ...
func UnsafeFloat32SliceToByteSlice(req []float32) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len * 4,
		Cap:  reqStruct.Cap * 4,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceToFloat32Slice returns ...
func UnsafeByteSliceToFloat32Slice(req []byte) (res []float32) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len / 4,
		Cap:  reqStruct.Cap / 4,
	}

	res = *(*[]float32)(unsafe.Pointer(&resStruct))
	// TODO::: really need do this here??
	// runtime.KeepAlive(req)
	return
}

// UnsafeFloat64SliceToByteSlice returns ...
func UnsafeFloat64SliceToByteSlice(req []float64) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len * 8,
		Cap:  reqStruct.Cap * 8,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceToFloat64Slice returns ...
func UnsafeByteSliceToFloat64Slice(req []byte) (res []float64) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len / 8,
		Cap:  reqStruct.Cap / 8,
	}

	res = *(*[]float64)(unsafe.Pointer(&resStruct))
	// TODO::: really need do this here??
	// runtime.KeepAlive(req)
	return
}
