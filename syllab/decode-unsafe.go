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
// This would blow up silently if return not limit by make(string, 0, UnsafeGetArrayLength(p))
func UnsafeGetString(p []byte, offset uint32) string {
	var slice = UnsafeGetByteArray(p, offset)
	return *(*string)(unsafe.Pointer(&slice))
}

// UnsafeGetByteArray decodes byte slice from the payload buffer in unsafe manner!
// This would blow up silently if return not limit by make([]byte, 0, UnsafeGetArrayLength(p))
func UnsafeGetByteArray(p []byte, offset uint32) []byte {
	var add uint32 = GetUInt32(p[offset:])
	var len uint32 = GetUInt32(p[offset+4:])
	return p[add : add+len]
}
