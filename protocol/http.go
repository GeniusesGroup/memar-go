/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// HTTPHandler is any object to be HTTP service handler.
type HTTPHandler interface {
	ServeHTTP(stream Stream, httpReq HTTPRequest, httpRes HTTPResponse) (err Error)
	// ServeGoHTTP(httpRes http.ResponseWriter, httpReq *http.Request) Due to bad sign, we can't standard it here.

	// Call service remotely by HTTP protocol
	// doHTTP(req any) (res any, err Error) Due to specific sign for each service, we can't standard it here.
}

// HTTPRequest indicate request semantic.
type HTTPRequest interface {
	HTTP_PseudoHeader_Request
	Header() HTTPHeader
	HTTPBody

	Codec
}

// HTTPResponse indicate response semantic.
type HTTPResponse interface {
	HTTP_PseudoHeader_Response
	Header() HTTPHeader
	HTTPBody

	Codec
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
