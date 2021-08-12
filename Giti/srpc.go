/* For license and copyright information please see LEGAL file in repository */

package giti

// SRPCHandler is any object to be sRPC service handler.
type SRPCHandler interface {
	// ServeSRPCSyllab method is sRPC handler of the service with Syllab codec data in the stream payload.
	ServeSRPCSyllab(st Stream) (err Error)
}
