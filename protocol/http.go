/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Other languages:
// - https://www.php-fig.org/psr/psr-7/
// - https://nodejs.org/api/http.html#requestgetheaders

// HTTPHandler is any object to be HTTP service handler.
// Some other protocol like gRPC, SOAP, ... must implement inside HTTP, If they are use HTTP as a transfer protocol.
type HTTPHandler /*[REQ, RES any]*/ interface {
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
	// DoHTTP(req REQ) (res RES, err Error) Due to specific sign for each service, we can't standard it here.
}

// HTTPRequest indicate HTTP request semantic.
type HTTPRequest interface {
	HTTP_PseudoHeader_Request
	HTTP_Header
	HTTP_Body
}

// HTTPResponse indicate HTTP response semantic.
type HTTPResponse interface {
	HTTP_PseudoHeader_Response
	HTTP_Header
	HTTP_Body
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
type HTTP_Header interface {
	Header_Get(key string) (value string)
	Header_Add(key, value string)
	// Header_Set is same as Header_Del() >> Header_Add()
	Header_Set(key, value string)
	Header_Del(key string)
}

// some header fields such as "Set-Cookie", "WWW-Authenticate", "Proxy-Authenticate" break multiple values
// separate by comma and use multi line same key! implementations MUST provide iteration mechanism over all header fields.
type HTTP_Header_Iteration Iteration_KV[string, string]

// HTTP Body Semantic
type HTTP_Body interface {
	Body() Codec
	SetBody(codec Codec)
}
