/* For license and copyright information please see LEGAL file in repository */

package giti

// HTTPHandler is any object to be HTTP service handler.
type HTTPHandler interface {
	ServeHTTP(stream Stream, httpReq HTTPRequest, httpRes HTTPResponse) (err Error)
}

// HTTP Request Semantic
type HTTPRequest interface {
	// "message start-line" in HTTP/1.x or "pseudo-header fields" in HTTP/2.x||HTTP/3.x
	Method() string
	URI() HTTPURI
	Version() string
	SetMethod(method string)
	SetVersion(version string)

	Header() HTTPHeader
	HTTPBody

	Codec
}

// HTTP Response Semantic
type HTTPResponse interface {
	// "message start-line" in HTTP/1.x or "pseudo-header fields" in HTTP/2.x||HTTP/3.x
	Version() string
	StatusCode() string
	ReasonPhrase() string
	SetVersion(version string)
	SetStatus(statusCode, reasonPhrase string)

	Header() HTTPHeader
	HTTPBody

	Codec
}

// HTTP URI Semantic
type HTTPURI interface {
	Raw() string
	Scheme() string
	Authority() string
	Host() string
	Path() string
	Query() string
	Fragment() string
	Set(scheme, authority, path, query string)
}

// HTTP Header Semantic
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
