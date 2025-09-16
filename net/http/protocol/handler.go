/* For license and copyright information please see the LEGAL file in the code repository */

package http_p

import (
	error_p "memar/error/protocol"
	net_p "memar/net/protocol"
	uri_p "memar/net/uri/protocol"
)

// Handler is any object to be HTTP service handler.
// Some other protocol like gRPC, SOAP, ... must implement inside HTTP, If they are use HTTP as a transfer protocol.
type Handler /*[HTTP_REQ Request, HTTP_RES Response, STR string_p.String]*/ interface {
	// Fill just if any http handler need route by path instead of ServiceID().
	//
	// **ATTENTION**
	// - Just use `URI.Path` since other part of URI not fit in `memar` rules. e.g.
	// 		- If add `URI.Host`, means application will serve for multiple domains. In TLS world, it MUST handle in upper layers.
	// 		- If add `URI.Port`, means we lock to old protocols like `tcp`.
	// - As describe here https://www.rfc-editor.org/rfc/rfc6570
	// 	JUST use simple immutable path and DO NOT variable data included in path.
	// 	e.g. "/product?id=1" instead of "/product/1/"
	uri_p.Path

	// ** NOTE: Due to reuse underling buffer If caller need to keep any data from httpReq or httpRes it must make a copy and
	// ** prevent from keep a reference to any data get from these two interface after return.
	// **
	ServeHTTP(sk net_p.Socket, req Request, res Response) (err error_p.Error)

	// Call service remotely by HTTP protocol
	// doHTTP[REQ operation_p.Request, RES operation_p.Response](req REQ) (res RES, err error_p.Error)
}
