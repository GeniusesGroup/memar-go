/* For license and copyright information please see LEGAL file in repository */

package http

// https://en.wikipedia.org/wiki/List_of_HTTP_header_fields
type responseStandards struct {
	Headers resHeaders
	Status  resStatus
}

type resHeaders struct {
	AccessControl           resAccessControl
	Accept                  resAccept
	Age                     string
	Allow                   string
	AltSvc                  string
	CacheControl            string
	Connection              string
	Content                 resContent
	Date                    string
	ETag                    string
	Expires                 string
	LastModified            string
	Link                    string
	Location                string
	P3P                     string
	Pragma                  string
	ProxyAuthenticate       string
	PublicKeyPins           string
	Refresh                 string
	RetryAfter              string
	Server                  string
	SetCookie               string
	StrictTransportSecurity string
	Trailer                 string
	TransferEncoding        string
	Tk                      string
	Upgrade                 string
	Vary                    string
	Via                     string
	Warning                 string
	WWWAuthenticate         string
	XFrameOptions           string
}

type resAccessControl struct {
	AllowOrigin      string
	AllowMethods     string
	AllowCredentials string
	AllowHeaders     string
	ExposeHeaders    string
	MaxAge           string
	RequestHeaders   string
	RequestMethod    string
}

type resAccept struct {
	Patch  string
	Ranges string
}

type resContent struct {
	Disposition string
	Encoding    string
	Language    string
	Length      string
	Location    string
	MD5         string
	Range       string
	Type        string
}

type resStatus struct {
	Continue           statusStruct // RFC 7231, 6.2.1
	SwitchingProtocols statusStruct // RFC 7231, 6.2.2
	Processing         statusStruct // RFC 2518, 10.1

	OK                   statusStruct // RFC 7231, 6.3.1
	Created              statusStruct // RFC 7231, 6.3.2
	Accepted             statusStruct // RFC 7231, 6.3.3
	NonAuthoritativeInfo statusStruct // RFC 7231, 6.3.4
	NoContent            statusStruct // RFC 7231, 6.3.5
	ResetContent         statusStruct // RFC 7231, 6.3.6
	PartialContent       statusStruct // RFC 7233, 4.1
	MultiStatus          statusStruct // RFC 4918, 11.1
	AlreadyReported      statusStruct // RFC 5842, 7.1
	IMUsed               statusStruct // RFC 3229, 10.4.1

	MultipleChoices   statusStruct // RFC 7231, 6.4.1
	MovedPermanently  statusStruct // RFC 7231, 6.4.2
	Found             statusStruct // RFC 7231, 6.4.3
	SeeOther          statusStruct // RFC 7231, 6.4.4
	NotModified       statusStruct // RFC 7232, 4.1
	UseProxy          statusStruct // RFC 7231, 6.4.5
	SwitchProxy       statusStruct // RFC 7231, 6.4.6 (Unused)
	TemporaryRedirect statusStruct // RFC 7231, 6.4.7
	PermanentRedirect statusStruct // RFC 7538, 3

	BadRequest                  statusStruct // RFC 7231, 6.5.1
	Unauthorized                statusStruct // RFC 7235, 3.1
	PaymentRequired             statusStruct // RFC 7231, 6.5.2
	Forbidden                   statusStruct // RFC 7231, 6.5.3
	NotFound                    statusStruct // RFC 7231, 6.5.4
	MethodNotAllowed            statusStruct // RFC 7231, 6.5.5
	NotAcceptable               statusStruct // RFC 7231, 6.5.6
	ProxyAuthRequired           statusStruct // RFC 7235, 3.2
	RequestTimeout              statusStruct // RFC 7231, 6.5.7
	Conflict                    statusStruct // RFC 7231, 6.5.8
	Gone                        statusStruct // RFC 7231, 6.5.9
	LengthRequired              statusStruct // RFC 7231, 6.5.10
	PreconditionFailed          statusStruct // RFC 7232, 4.2
	PayloadTooLarge             statusStruct // RFC 7231, 6.5.11
	URITooLong                  statusStruct // RFC 7231, 6.5.12
	UnsupportedMediaType        statusStruct // RFC 7231, 6.5.13
	RangeNotSatisfiable         statusStruct // RFC 7233, 4.4
	ExpectationFailed           statusStruct // RFC 7231, 6.5.14
	Teapot                      statusStruct // RFC 7168, 2.3.3
	UnprocessableEntity         statusStruct // RFC 4918, 11.2
	Locked                      statusStruct // RFC 4918, 11.3
	FailedDependency            statusStruct // RFC 4918, 11.4
	UpgradeRequired             statusStruct // RFC 7231, 6.5.15
	PreconditionRequired        statusStruct // RFC 6585, 3
	TooManyRequests             statusStruct // RFC 6585, 4
	RequestHeaderFieldsTooLarge statusStruct // RFC 6585, 5
	UnavailableForLegalReasons  statusStruct // RFC 7725, 3

	InternalServerError           statusStruct // RFC 7231, 6.6.1
	NotImplemented                statusStruct // RFC 7231, 6.6.2
	BadGateway                    statusStruct // RFC 7231, 6.6.3
	ServiceUnavailable            statusStruct // RFC 7231, 6.6.4
	GatewayTimeout                statusStruct // RFC 7231, 6.6.5
	HTTPVersionNotSupported       statusStruct // RFC 7231, 6.6.6
	VariantAlsoNegotiates         statusStruct // RFC 2295, 8.1
	InsufficientStorage           statusStruct // RFC 4918, 11.5
	LoopDetected                  statusStruct // RFC 5842, 7.2
	NotExtended                   statusStruct // RFC 2774, 7
	NetworkAuthenticationRequired statusStruct // RFC 6585, 6
}

type statusStruct struct {
	Code int
	Text string
}

// ResponseStandards : All text of Standard http response fileds.
var ResponseStandards = responseStandards{
	Headers: resHeaders{
		AccessControl: resAccessControl{
			AllowOrigin:      "Access-Control-Allow-Origin",
			AllowMethods:     "Access-Control-Allow-Methods",
			AllowCredentials: "Access-Control-Allow-Credentials",
			AllowHeaders:     "Access-Control-Allow-Headers",
			ExposeHeaders:    "Access-Control-Expose-Headers",
			MaxAge:           "Access-Control-Max-Age",
			RequestHeaders:   "Access-Control-Request-Headers",
			RequestMethod:    "Access-Control-Request-Method"},
		Accept: resAccept{
			Patch:  "Accept-Patch",
			Ranges: "Accept-Ranges"},
		Age:          "Age",
		Allow:        "Allow",
		AltSvc:       "Alt-Svc",
		CacheControl: "Cache-Control",
		Connection:   "Connection",
		Content: resContent{
			Disposition: "Content-Disposition",
			Encoding:    "Content-Encoding",
			Language:    "Content-Language",
			Length:      "Content-Length",
			Location:    "Content-Location",
			MD5:         "Content-MD5",
			Range:       "Content-Range",
			Type:        "Content-Type"},
		Date:                    "Date",
		ETag:                    "ETag",
		Expires:                 "Expires",
		LastModified:            "Last-Modified",
		Link:                    "Link",
		Location:                "Location",
		P3P:                     "P3P",
		Pragma:                  "Pragma",
		ProxyAuthenticate:       "Proxy-Authenticate",
		PublicKeyPins:           "Public-Key-Pins",
		Refresh:                 "Refresh",
		RetryAfter:              "Retry-After",
		Server:                  "Server",
		SetCookie:               "Set-Cookie",
		StrictTransportSecurity: "Strict-Transport-Security",
		Trailer:                 "Trailer",
		TransferEncoding:        "Transfer-Encoding",
		Tk:                      "Tk",
		Upgrade:                 "Upgrade",
		Vary:                    "Vary",
		Via:                     "Via",
		Warning:                 "Warning",
		WWWAuthenticate:         "WWW-Authenticate",
		XFrameOptions:           "X-Frame-Options"},
	Status: resStatus{
		Continue:                      statusStruct{Code: 100, Text: "Continue"},
		SwitchingProtocols:            statusStruct{Code: 101, Text: "Switching Protocols"},
		Processing:                    statusStruct{Code: 102, Text: "Processing"},
		OK:                            statusStruct{Code: 200, Text: "OK"},
		Created:                       statusStruct{Code: 201, Text: "Created"},
		Accepted:                      statusStruct{Code: 202, Text: "Accepted"},
		NonAuthoritativeInfo:          statusStruct{Code: 203, Text: "Non-Authoritative Info"},
		NoContent:                     statusStruct{Code: 204, Text: "No Content"},
		ResetContent:                  statusStruct{Code: 205, Text: "Reset Content"},
		PartialContent:                statusStruct{Code: 206, Text: "Partial Content"},
		MultiStatus:                   statusStruct{Code: 207, Text: "Multi-Status"},
		AlreadyReported:               statusStruct{Code: 208, Text: "Already Reported"},
		IMUsed:                        statusStruct{Code: 226, Text: "IM Used"},
		MultipleChoices:               statusStruct{Code: 300, Text: "Multiple Choices"},
		MovedPermanently:              statusStruct{Code: 301, Text: "Moved Permanently"},
		Found:                         statusStruct{Code: 302, Text: "Found"},
		SeeOther:                      statusStruct{Code: 303, Text: "See Other"},
		NotModified:                   statusStruct{Code: 304, Text: "Not Modified"},
		UseProxy:                      statusStruct{Code: 305, Text: "Use Proxy"},
		SwitchProxy:                   statusStruct{Code: 306, Text: "Switch Proxy"},
		TemporaryRedirect:             statusStruct{Code: 307, Text: "Temporary Redirect"},
		PermanentRedirect:             statusStruct{Code: 308, Text: "Permanent Redirect"},
		BadRequest:                    statusStruct{Code: 400, Text: "Bad Request"},
		Unauthorized:                  statusStruct{Code: 401, Text: "Unauthorized"},
		PaymentRequired:               statusStruct{Code: 402, Text: "Payment Required"},
		Forbidden:                     statusStruct{Code: 403, Text: "Forbidden"},
		NotFound:                      statusStruct{Code: 404, Text: "Not Found"},
		MethodNotAllowed:              statusStruct{Code: 405, Text: "Method Not Allowed"},
		NotAcceptable:                 statusStruct{Code: 406, Text: "Not Acceptable"},
		ProxyAuthRequired:             statusStruct{Code: 407, Text: "Proxy Authentication Required"},
		RequestTimeout:                statusStruct{Code: 408, Text: "Request Timeout"},
		Conflict:                      statusStruct{Code: 409, Text: "Conflict"},
		Gone:                          statusStruct{Code: 410, Text: "Gone"},
		LengthRequired:                statusStruct{Code: 411, Text: "Length Required"},
		PreconditionFailed:            statusStruct{Code: 412, Text: "Precondition Failed"},
		PayloadTooLarge:               statusStruct{Code: 413, Text: "Payload Too Large"},
		URITooLong:                    statusStruct{Code: 414, Text: "URI Too Long"},
		UnsupportedMediaType:          statusStruct{Code: 415, Text: "Unsupported Media Type"},
		RangeNotSatisfiable:           statusStruct{Code: 416, Text: "Range Not Satisfiable"},
		ExpectationFailed:             statusStruct{Code: 417, Text: "Expectation Failed"},
		Teapot:                        statusStruct{Code: 418, Text: "I'm a teapot"},
		UnprocessableEntity:           statusStruct{Code: 422, Text: "Unprocessable Entity"},
		Locked:                        statusStruct{Code: 423, Text: "Locked"},
		FailedDependency:              statusStruct{Code: 424, Text: "Failed Dependency"},
		UpgradeRequired:               statusStruct{Code: 426, Text: "Upgrade Required"},
		PreconditionRequired:          statusStruct{Code: 428, Text: "Precondition Required"},
		TooManyRequests:               statusStruct{Code: 429, Text: "Too Many Requests"},
		RequestHeaderFieldsTooLarge:   statusStruct{Code: 431, Text: "Request Header Fields Too Large"},
		UnavailableForLegalReasons:    statusStruct{Code: 451, Text: "Unavailable For Legal Reasons"},
		InternalServerError:           statusStruct{Code: 500, Text: "Internal Server Error"},
		NotImplemented:                statusStruct{Code: 501, Text: "Not Implemented"},
		BadGateway:                    statusStruct{Code: 502, Text: "Bad Gateway"},
		ServiceUnavailable:            statusStruct{Code: 503, Text: "Service Unavailable"},
		GatewayTimeout:                statusStruct{Code: 504, Text: "Gateway Timeout"},
		HTTPVersionNotSupported:       statusStruct{Code: 505, Text: "HTTP Version Not Supported"},
		VariantAlsoNegotiates:         statusStruct{Code: 506, Text: "Variant Also Negotiates"},
		InsufficientStorage:           statusStruct{Code: 507, Text: "Insufficient Storage"},
		LoopDetected:                  statusStruct{Code: 508, Text: "Loop Detected"},
		NotExtended:                   statusStruct{Code: 510, Text: "Not Extended"},
		NetworkAuthenticationRequired: statusStruct{Code: 511, Text: "Network Authentication Required"}}}
