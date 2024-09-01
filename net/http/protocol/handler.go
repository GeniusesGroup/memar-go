/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

import (
	net_p "memar/net/protocol"
	uri_p "memar/net/uri/protocol"
	"memar/protocol"
	string_p "memar/string/protocol"
)

// Handler is any object to be HTTP service handler.
// Some other protocol like gRPC, SOAP, ... must implement inside HTTP, If they are use HTTP as a transfer protocol.
type Handler[HTTP_REQ Request /*[STR]*/, HTTP_RES Response /*[STR]*/, STR string_p.String, REQ, RES any] interface {
	// Fill just if any http handler need route by path instead of ServiceID().
	//
	// **ATTENTION** - As describe here https://www.rfc-editor.org/rfc/rfc6570
	// JUST use simple immutable path and DO NOT variable data included in path.
	// e.g. "/product?id=1" instead of "/product/1/"
	uri_p.Path[STR]

	// ** NOTE: Due to reuse underling buffer If caller need to keep any data from httpReq or httpRes it must make a copy and
	// ** prevent from keep a reference to any data get from these two interface after return.
	// **
	ServeHTTP(sk net_p.Socket, req Request, res Response) (err protocol.Error)

	// Call service remotely by HTTP protocol
	// doHTTP(req REQ) (res RES, err Error)
}
