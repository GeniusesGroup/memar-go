/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"reflect"
	"unsafe"
)

// UnsafeStringToByteSlice returns ...
func UnsafeStringToByteSlice(req string) (res []byte) {
	var reqStruct = (*reflect.StringHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len
	resStruct.Cap = reqStruct.Len
	return
}

// UnsafeByteSliceToString returns ...
func UnsafeByteSliceToString(req []byte) (res string) {
	var reqStruct = (*reflect.SliceHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.StringHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len
	return
}
