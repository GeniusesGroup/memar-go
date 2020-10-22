/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"reflect"
	"unsafe"
)

// UnsafeBoolSliceToByteSlice returns ...
func UnsafeBoolSliceToByteSlice(req []bool) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len,
		Cap:  reqStruct.Cap,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceToBoolSlice returns ...
func UnsafeByteSliceToBoolSlice(req []byte) (res []bool) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len,
		Cap:  reqStruct.Cap,
	}

	res = *(*[]bool)(unsafe.Pointer(&resStruct))
	// TODO::: really need do this here??
	// runtime.KeepAlive(req)
	return
}
