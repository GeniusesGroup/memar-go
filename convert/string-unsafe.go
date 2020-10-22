/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"reflect"
	"unsafe"
)

// UnsafeStringToByteSlice returns ...
func UnsafeStringToByteSlice(req string) (res []byte) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.SliceHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len,
		Cap:  reqStruct.Cap,
	}

	res = *(*[]byte)(unsafe.Pointer(&resStruct))
	return
}

// UnsafeByteSliceToString returns ...
func UnsafeByteSliceToString(req []byte) (res string) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = reflect.StringHeader{
		Data: reqStruct.Data,
		Len:  reqStruct.Len,
	}

	res = *(*string)(unsafe.Pointer(&resStruct))
	return
}
