/* For license and copyright information please see LEGAL file in repository */

package protocol

// SRPCHandler is any object to be sRPC service handler.
type SRPCHandler interface {
	// ServeSRPC method is sRPC handler of the service with Syllab codec data in the payload.
	ServeSRPC(st Stream) (err Error)
	// DoSRPC(req interface{}) (res Interface{}, err Error) Due to specific sign for each service, we can't standard it here.

	// ServeSRPCDirect use to call a service without need to open any stream.
	// It can also use when service request data is smaller than network MTU.
	// Or use for time sensitive data like audio and video that streams shape in app layer
	ServeSRPCDirect(conn Connection, request []byte) (response []byte, err Error)
}
