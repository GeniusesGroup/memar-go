/* For license and copyright information please see the LEGAL file in the code repository */

package error

import (
	"libgo/protocol"
)

type ErrorType protocol.ErrorType

func (et *ErrorType) Set(errorType protocol.ErrorType)   { *et |= ErrorType(errorType) }
func (et *ErrorType) Unset(errorType protocol.ErrorType) { *et &= ^ErrorType(errorType) }
func (et ErrorType) Check(errorType protocol.ErrorType) bool {
	return errorType&protocol.ErrorType(et) == errorType
}

func (et Error) Internal() bool  { return et.Check(protocol.ErrorType_Internal) }
func (et Error) Temporary() bool { return et.Check(protocol.ErrorType_Temporary) }
func (et Error) Timeout() bool   { return et.Check(protocol.ErrorType_Timeout) }

func (et *Error) SetInternal()  { et.Set(protocol.ErrorType_Internal) }
func (et *Error) SetTemporary() { et.Set(protocol.ErrorType_Temporary) }
func (et *Error) SetTimeout()   { et.Set(protocol.ErrorType_Timeout) }
