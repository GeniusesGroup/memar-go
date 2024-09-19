/* For license and copyright information please see the LEGAL file in the code repository */

package service_p

import (
	datatype_p "memar/datatype/protocol"
	error_p "memar/error/protocol"
)

// Handlers is just test (approver) interface and MUST NOT use directly in any signature.
// Due to Golang import cycle problem we can't use `net_p.Socket`
type Handlers[SK any /*net_p.Socket*/, ReqT, ResT datatype_p.DataType] interface {
	// Call service locally by import service package to other one
	Process(sk SK, req ReqT) (res ResT, err error_p.Error)
	//
	// Call service remotely by preferred(SDK generator choose) protocol.
	Do(req ReqT) (res ResT, err error_p.Error)
}
