/* For license and copyright information please see LEGAL file in repository */

package protocol

// SRPCHandler is any object to be sRPC service handler.
type SRPCHandler interface {
	// ServeSRPC method is sRPC handler of the service with Syllab codec data in the payload.
	ServeSRPC(st Stream) (err Error)
	// DoSRPC(req interface{}) (res Interface{}, err Error) Due to specific sign for each service, we can't standard it here.
}
