/* For license and copyright information please see LEGAL file in repository */

package syllab

import (
	"unsafe"
)

/*
	************************************************************
	**********************Fixed Size ARRAY**********************
	************************************************************
	***********************PAY ATTENTION************************
	By use below helper functions you can't achieve max performance!
	Use code generation to prevent unneeded memory alloc by CompleteMethods()!
*/

// Unsafe don't give any efficient in fixed size arrays! Use decode safe functions!

/*
************************************************************
*******************Dynamically size ARRAY*******************
************************************************************
 */

// UnsafeGetString decodes string from the payload buffer in unsafe manner!
func UnsafeGetString(p []byte, stackIndex uint32) string {
	var add uint32 = GetUInt32(p, stackIndex)
	var len uint32 = GetUInt32(p, stackIndex+4)
	var slice = p[add : add+len]
	return *(*string)(unsafe.Pointer(&slice))
}

// UnsafeGetByteArray decodes byte slice from the payload buffer in unsafe manner!
func UnsafeGetByteArray(p []byte, stackIndex uint32) []byte {
	var add uint32 = GetUInt32(p, stackIndex)
	var len uint32 = GetUInt32(p, stackIndex+4)
	return p[add : add+len]
}
