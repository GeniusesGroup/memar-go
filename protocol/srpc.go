/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// SRPCHandler is any object to be sRPC service handler.
type SRPCHandler interface {
	// ServeSRPC method is sRPC handler of the service with Syllab codec data in the payload.
	ServeSRPC(st Stream, rawReq Codec) (rawRes Codec, err Error)

	// Call service remotely by sRPC protocol
	// doSRPC(req any) (res any, err Error) Due to specific sign for each service, we can't standard it here.

	// ServeSRPCDirect use to call a service without need to open any stream.
	// It can also use when service request data is smaller than network MTU.
	// Or use for time sensitive data like audio and video that streams shape in app layer
	ServeSRPCDirect(conn Connection, request []byte) (response []byte, err Error)
}
