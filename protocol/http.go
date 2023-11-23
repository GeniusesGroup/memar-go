/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// HTTPHandler is any object to be HTTP service handler.
// Some other protocol like gRPC, SOAP, ... must implement inside HTTP, If they are use HTTP as a transfer protocol.
type HTTPHandler interface {
	// Fill just if any http handler needed. Suggest use simple immutable path,
	// not variable data included in path describe here:https://www.rfc-editor.org/rfc/rfc6570 e.g. "/product?id=1" instead of "/product/1/"
	// API services can set like "/s?{{.ServiceID}}" but it is not efficient, instead find services by ID as base64
	URI() string // suggest use just URI.Path

	// ** NOTE: Due to reuse underling buffer If caller need to keep any data from httpReq or httpRes it must make a copy and
	// ** prevent from keep a reference to any data get from these two interface after return.
	// **
	ServeHTTP(sk Socket, httpReq HTTPRequest, httpRes HTTPResponse) (err Error)
	// ServeGoHTTP(httpRes http.ResponseWriter, httpReq *http.Request) Due to bad sign, we can't standard it here.

	// Call service remotely by HTTP protocol
	// doHTTP(req any) (res any, err Error) Due to specific sign for each service, we can't standard it here.
}

// HTTPRequest indicate HTTP request semantic.
type HTTPRequest interface {
	HTTP_PseudoHeader_Request
	Header() HTTPHeader
	HTTPBody
}

// HTTPResponse indicate HTTP response semantic.
type HTTPResponse interface {
	HTTP_PseudoHeader_Response
	Header() HTTPHeader
	HTTPBody
}

// HTTP_PseudoHeader_Request indicate request pseudo header.
// "message start-line" in HTTP/1.x or "pseudo-header fields" in HTTP/2.x||HTTP/3.x
type HTTP_PseudoHeader_Request interface {
	Method() string
	// https://datatracker.ietf.org/doc/html/rfc2616#section-3.2
	// http_URL = "http:" "//" host [ ":" port ] [ abs_path [ "?" query ]]
	// .Scheme() string // always return "http"
	URI() URI
	Version() string

	SetMethod(method string)
	SetVersion(version string)
}

// HTTP_PseudoHeader_Response indicate response pseudo header.
// "message start-line" in HTTP/1.x or "pseudo-header fields" in HTTP/2.x||HTTP/3.x
type HTTP_PseudoHeader_Response interface {
	Version() string
	StatusCode() string
	ReasonPhrase() string

	SetVersion(version string)
	SetStatus(statusCode, reasonPhrase string)
}

// HTTPHeader indicate HTTP header semantic.
type HTTPHeader interface {
	Get(key string) (value string)
	Gets(key string) (values []string)
	Add(key, value string)
	Adds(key string, values []string)
	Set(key string, value string)
	Sets(key string, values []string)
	Del(key string)
}

// HTTP Body Semantic
type HTTPBody interface {
	Body() Codec
	SetBody(codec Codec)
}
