/* For license and copyright information please see the LEGAL file in the code repository */

package srpc_p

import (
	error_p "memar/error/protocol"
	net_p "memar/net/protocol"
)

// Handler is any object to be sRPC service handler.
type Handler interface {
	// ServeSRPC method is sRPC handler of the service with Syllab codec data in the payload.
	ServeSRPC(sk net_p.Socket) (err error_p.Error)

	// Call service remotely by sRPC protocol
	// doSRPC(req any) (res any, err error_p.Error) Due to specific sign for each service, we can't have it here.
}
