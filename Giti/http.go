/* For license and copyright information please see LEGAL file in repository */

package giti

type HTTPRequest interface {
	Method() string
	SetMethod(method string)
	URI() HTTPURI
	Version() string
	SetVersion(version string)

	Header() HTTPHeader
	Body() HTTPBody

	// Helpers methods

	// GetHost returns host of request by RFC 7230, section 5.3 rules: Must treat
	//		GET / HTTP/1.1
	//		Host: www.sabz.city
	// and
	//		GET https://www.sabz.city/ HTTP/1.1
	//		Host: apis.sabz.city
	// the same. In the second case, any Host line is ignored.
	GetHost() (host string)

	Codec
}

type HTTPResponse interface {
	Version() string
	SetVersion(version string)
	StatusCode() string
	ReasonPhrase() string
	SetStatus(statusCode, reasonPhrase string)

	Header() HTTPHeader
	Body() HTTPBody

	Codec
}

type HTTPURI interface {
	Raw() string
	Scheme() string
	Authority() string
	Path() string
	Query() string
	Fragment() string
	Set(scheme, authority, path, query string)
}

type HTTPHeader interface {
	Get(key string) (value string)
	Gets(key string) (values []string)
	Add(key, value string)
	Adds(key string, values []string)
	Set(key string, value string)
	Sets(key string, values []string)
	Del(key string)
}

type HTTPBody interface {
	Raw() []byte // only for read from peer!
	BodyCodec() Codec
	SetBodyCodec(codec Codec)
}
