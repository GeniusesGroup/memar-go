/* For license and copyright information please see the LEGAL file in the code repository */

package http

const (
	Length_Min_Packet     = 64
	Length_Max_Header     = 8192
	Length_Max_StatusCode = 4 // plus one for ' ' character after method name.
	// Due to RFC(https://datatracker.ietf.org/doc/html/rfc2616#section-6.1.1) can't indicate phrase max length
	Length_Max_VersionMaxLength = 9 // plus one for ' ' or '\r' character after method name.
	Length_Max_MethodMaxLength  = 8 // plus one for ' ' character after method name.

	// TimeFormat is the time format to use when generating times in HTTP
	// headers. It is like time.RFC1123 but hard-codes GMT as the time
	// zone. The time being formatted must be in UTC for Format to
	// generate the correct format.
	TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

	headerInitLen = 16
)

// Some default values
const (
	DefaultUserAgent = "Memar-Client"
	DefaultServer    = "Memar"

	SP             byte   = ' '  // <US-ASCII SP, space (32)>
	HT             byte   = '	'  // <US-ASCII HT, horizontal-tab (9)>
	CR             byte   = '\r' // <US-ASCII CR, carriage return (13)>
	LF             byte   = '\n' // <US-ASCII LF, linefeed (10)>
	Colon          byte   = ':'
	NumberSign     byte   = '#'
	Comma          byte   = ','
	Question       byte   = '?'
	Slash          byte   = '/'
	Asterisk       byte   = '*'
	CRLF           string = "\r\n"
	ColonSpace     string = ": "
	SemiColonSpace string = "; "
)

// Standard HTTP versions
const (
	Version_HTTP1  = "HTTP/1.0"
	Version_HTTP11 = "HTTP/1.1"
	Version_HTTP2  = "HTTP/2.0"
	Version_HTTP3  = "HTTP/3.0"
)

// Standard HTTP methods
// https://tools.ietf.org/html/rfc7231#section-4
const (
	Method_GET     = "GET"
	Method_POST    = "POST"
	Method_HEAD    = "HEAD"
	Method_PUT     = "PUT"
	Method_DELETE  = "DELETE"
	Method_OPTIONS = "OPTIONS"
	Method_CONNECT = "CONNECT"
	Method_TRACE   = "TRACE"
	Method_PATCH   = "PATCH" // https://tools.ietf.org/html/rfc5789#section-2
)

// Standard HTTP header keys
// https://www.iana.org/assignments/message-headers/message-headers.xhtml
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
const (
	// Request context
	HeaderKey_From           = "From"
	HeaderKey_Host           = "Host"
	HeaderKey_Referer        = "Referer"
	HeaderKey_ReferrerPolicy = "Referrer-Policy"
	HeaderKey_UserAgent      = "User-Agent"

	// Response context
	HeaderKey_Allow   = "Allow"
	HeaderKey_Server  = "Server"
	HeaderKey_ErrorID = "Error-ID"

	// Authentication
	HeaderKey_Authorization      = "Authorization"
	HeaderKey_ProxyAuthorization = "Proxy-Authorization"
	HeaderKey_ProxyAuthenticate  = "Proxy-Authenticate" // res
	HeaderKey_WWWAuthenticate    = "WWW-Authenticate"   // res

	// Caching
	HeaderKey_Age           = "Age"           // res
	HeaderKey_CacheControl  = "Cache-Control" // req & res
	HeaderKey_ClearSiteData = "Clear-Site-Data"
	HeaderKey_Expires       = "Expires" // res
	HeaderKey_Pragma        = "Pragma"  // req & res
	HeaderKey_Warning       = "Warning" // req & res
	HeaderKey_Vary          = "Vary"    // res

	// Conditionals
	HeaderKey_ETag              = "ETag" // res
	HeaderKey_IfMatch           = "If-Match"
	HeaderKey_IfNoneMatch       = "If-None-Match"
	HeaderKey_IfModifiedSince   = "If-Modified-Since"
	HeaderKey_IfUnmodifiedSince = "If-Unmodified-Since"
	HeaderKey_LastModified      = "Last-Modified" // res

	// Range requests
	HeaderKey_AcceptRanges = "Accept-Ranges" // res
	HeaderKey_ContentRange = "Content-Range" // res
	HeaderKey_IfRange      = "If-Range"
	HeaderKey_Range        = "Range"

	// Connection management
	HeaderKey_Connection = "Connection" // req & res
	HeaderKey_KeepAlive  = "Keep-Alive"
	HeaderKey_Upgrade    = "Upgrade"

	// CORS
	HeaderKey_AccessControlAllowOrigin      = "Access-Control-Allow-Origin"      // res
	HeaderKey_AccessControlAllowMethods     = "Access-Control-Allow-Methods"     // res
	HeaderKey_AccessControlAllowCredentials = "Access-Control-Allow-Credentials" // res
	HeaderKey_AccessControlAllowHeaders     = "Access-Control-Allow-Headers"     // res
	HeaderKey_AccessControlExposeHeaders    = "Access-Control-Expose-Headers"    // res
	HeaderKey_AccessControlMaxAge           = "Access-Control-Max-Age"           // res
	HeaderKey_AccessControlRequestHeaders   = "Access-Control-Request-Headers"   // res
	HeaderKey_AccessControlRequestMethod    = "Access-Control-Request-Method"    // res
	HeaderKey_Origin                        = "Origin"
	HeaderKey_TimingAllowOrigin             = "Timing-Allow-Origin"
	HeaderKey_XPermittedCrossDomainPolicies = "X-Permitted-Cross-Domain-Policies"

	// Content negotiation
	HeaderKey_AcceptContent  = "Accept"
	HeaderKey_AcceptCharset  = "Accept-Charset" // deprecated
	HeaderKey_AcceptEncoding = "Accept-Encoding"
	HeaderKey_AcceptLanguage = "Accept-Language"
	HeaderKey_AcceptDatetime = "Accept-Datetime"
	HeaderKey_AcceptPatch    = "Accept-Patch" // res

	// Message body information
	HeaderKey_ContentLength      = "Content-Length"      // req & res
	HeaderKey_ContentMD5         = "Content-MD5"         // req & res
	HeaderKey_ContentType        = "Content-Type"        // req & res
	HeaderKey_ContentDisposition = "Content-Disposition" // res
	HeaderKey_ContentEncoding    = "Content-Encoding"    // res
	HeaderKey_ContentLanguage    = "Content-Language"    // res
	HeaderKey_ContentLocation    = "Content-Location"    // res
	HeaderKey_TransferEncoding   = "Transfer-Encoding"   // res

	// Not ordered
	HeaderKey_Cookie                  = "Cookie"
	HeaderKey_SetCookie               = "Set-Cookie" // res
	HeaderKey_Date                    = "Date"       // req & res
	HeaderKey_Via                     = "Via"
	HeaderKey_Expect                  = "Expect"
	HeaderKey_Forwarded               = "Forwarded"
	HeaderKey_MaxForwards             = "Max-Forwards"
	HeaderKey_TE                      = "TE"
	HeaderKey_AltSvc                  = "Alt-Svc"                   // res
	HeaderKey_Link                    = "Link"                      // res
	HeaderKey_Location                = "Location"                  // res
	HeaderKey_P3P                     = "P3P"                       // res
	HeaderKey_PublicKeyPins           = "Public-Key-Pins"           // res
	HeaderKey_Refresh                 = "Refresh"                   // res
	HeaderKey_RetryAfter              = "Retry-After"               // res
	HeaderKey_StrictTransportSecurity = "Strict-Transport-Security" // res
	HeaderKey_Trailer                 = "Trailer"                   // res
	HeaderKey_Tk                      = "Tk"                        // res
	HeaderKey_XFrameOptions           = "X-Frame-Options"           // res
	HeaderKey_NonAuthoritativeReason  = "Non-Authoritative-Reason"  // res
)

// Standard HTTP header values
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
const (
	// Connection management
	HeaderValue_KeepAlive = "Keep-Alive"
	HeaderValue_Close     = "Close"

	// Message body information
	HeaderValue_Chunked  = "chunked"
	HeaderValue_Compress = "compress"
	HeaderValue_Deflate  = "deflate"
	HeaderValue_GZIP     = "gzip"
)

// HTTP Status codes
// https://tools.ietf.org/html/rfc7231#section-6
const (
	StatusContinue                 uint16 = 100 // RFC 7231, 6.2.1
	StatusContinueCode                    = "100"
	StatusContinuePhrase                  = "Continue"
	StatusSwitchingProtocols       uint16 = 101 // RFC 7231, 6.2.2
	StatusSwitchingProtocolsCode          = "101"
	StatusSwitchingProtocolsPhrase        = "Switching Protocols"
	StatusProcessing               uint16 = 102 // RFC 2518, 10.1
	StatusProcessingCode                  = "102"
	StatusProcessingPhrase                = "Processing"

	StatusOK                         uint16 = 200 // RFC 7231, 6.3.1
	StatusOKCode                            = "200"
	StatusOKPhrase                          = "OK"
	StatusCreated                    uint16 = 201 // RFC 7231, 6.3.2
	StatusCreatedCode                       = "201"
	StatusCreatedPhrase                     = "Created"
	StatusAccepted                   uint16 = 202 // RFC 7231, 6.3.3
	StatusAcceptedCode                      = "202"
	StatusAcceptedPhrase                    = "Accepted"
	StatusNonAuthoritativeInfo       uint16 = 203 // RFC 7231, 6.3.4
	StatusNonAuthoritativeInfoCode          = "203"
	StatusNonAuthoritativeInfoPhrase        = "Non-Authoritative Information"
	StatusNoContent                  uint16 = 204 // RFC 7231, 6.3.5
	StatusNoContentCode                     = "204"
	StatusNoContentPhrase                   = "No Content"
	StatusResetContent               uint16 = 205 // RFC 7231, 6.3.6
	StatusResetContentCode                  = "205"
	StatusResetContentPhrase                = "Reset Content"
	StatusPartialContent             uint16 = 206 // RFC 7233, 4.1
	StatusPartialContentCode                = "206"
	StatusPartialContentPhrase              = "Partial Content"
	StatusMultiStatus                uint16 = 207 // RFC 4918, 11.1
	StatusMultiStatusCode                   = "207"
	StatusMultiStatusPhrase                 = "Multi-Status"
	StatusAlreadyReported            uint16 = 208 // RFC 5842, 7.1
	StatusAlreadyReportedCode               = "208"
	StatusAlreadyReportedPhrase             = "Already Reported"
	StatusIMUsed                     uint16 = 226 // RFC 3229, 10.4.1
	StatusIMUsedCode                        = "226"
	StatusIMUsedPhrase                      = "IM Used"

	StatusMultipleChoices         uint16 = 300 // RFC 7231, 6.4.1
	StatusMultipleChoicesCode            = "300"
	StatusMultipleChoicesPhrase          = "Multiple Choices"
	StatusMovedPermanently        uint16 = 301 // RFC 7231, 6.4.2
	StatusMovedPermanentlyCode           = "301"
	StatusMovedPermanentlyPhrase         = "Moved Permanently"
	StatusFound                   uint16 = 302 // RFC 7231, 6.4.3
	StatusFoundCode                      = "302"
	StatusFoundPhrase                    = "Found"
	StatusSeeOther                uint16 = 303 // RFC 7231, 6.4.4
	StatusSeeOtherCode                   = "303"
	StatusSeeOtherPhrase                 = "See Other"
	StatusNotModified             uint16 = 304 // RFC 7232, 4.1
	StatusNotModifiedCode                = "304"
	StatusNotModifiedPhrase              = "Not Modified"
	StatusUseProxy                uint16 = 305 // RFC 7231, 6.4.5
	StatusUseProxyCode                   = "305"
	StatusUseProxyPhrase                 = "Use Proxy"
	StatusSwitchProxy             uint16 = 306 // RFC 7231, 6.4.6 (Unused)
	StatusSwitchProxyCode                = "306"
	StatusSwitchProxyPhrase              = "Switch Proxy"
	StatusTemporaryRedirect       uint16 = 307 // RFC 7231, 6.4.7
	StatusTemporaryRedirectCode          = "307"
	StatusTemporaryRedirectPhrase        = "Temporary Redirect"
	StatusPermanentRedirect       uint16 = 308 // RFC 7538, 3
	StatusPermanentRedirectCode          = "308"
	StatusPermanentRedirectPhrase        = "Permanent Redirect"

	StatusBadRequest                       uint16 = 400 // RFC 7231, 6.5.1
	StatusBadRequestCode                          = "400"
	StatusBadRequestPhrase                        = "Bad Request"
	StatusUnauthorized                     uint16 = 401 // RFC 7235, 3.1
	StatusUnauthorizedCode                        = "401"
	StatusUnauthorizedPhrase                      = "Unauthorized"
	StatusPaymentRequired                  uint16 = 402 // RFC 7231, 6.5.2
	StatusPaymentRequiredCode                     = "402"
	StatusPaymentRequiredPhrase                   = "Payment Required"
	StatusForbidden                        uint16 = 403 // RFC 7231, 6.5.3
	StatusForbiddenCode                           = "403"
	StatusForbiddenPhrase                         = "Forbidden"
	StatusNotFound                         uint16 = 404 // RFC 7231, 6.5.4
	StatusNotFoundCode                            = "404"
	StatusNotFoundPhrase                          = "Not Found"
	StatusMethodNotAllowed                 uint16 = 405 // RFC 7231, 6.5.5
	StatusMethodNotAllowedCode                    = "405"
	StatusMethodNotAllowedPhrase                  = "Method Not Allowed"
	StatusNotAcceptable                    uint16 = 406 // RFC 7231, 6.5.6
	StatusNotAcceptableCode                       = "406"
	StatusNotAcceptablePhrase                     = "Not Acceptable"
	StatusProxyAuthRequired                uint16 = 407 // RFC 7235, 3.2
	StatusProxyAuthRequiredCode                   = "407"
	StatusProxyAuthRequiredPhrase                 = "Proxy Authentication Required"
	StatusRequestTimeout                   uint16 = 408 // RFC 7231, 6.5.7
	StatusRequestTimeoutCode                      = "408"
	StatusRequestTimeoutPhrase                    = "Request Timeout"
	StatusConflict                         uint16 = 409 // RFC 7231, 6.5.8
	StatusConflictCode                            = "409"
	StatusConflictPhrase                          = "Conflict"
	StatusGone                             uint16 = 410 // RFC 7231, 6.5.9
	StatusGoneCode                                = "410"
	StatusGonePhrase                              = "Gone"
	StatusLengthRequired                   uint16 = 411 // RFC 7231, 6.5.10
	StatusLengthRequiredCode                      = "411"
	StatusLengthRequiredPhrase                    = "Length Required"
	StatusPreconditionFailed               uint16 = 412 // RFC 7232, 4.2
	StatusPreconditionFailedCode                  = "412"
	StatusPreconditionFailedPhrase                = "Precondition Failed"
	StatusPayloadTooLarge                  uint16 = 413 // RFC 7231, 6.5.11
	StatusPayloadTooLargeCode                     = "413"
	StatusPayloadTooLargePhrase                   = "Payload Too Large"
	StatusURITooLong                       uint16 = 414 // RFC 7231, 6.5.12
	StatusURITooLongCode                          = "414"
	StatusURITooLongPhrase                        = "URI Too Long"
	StatusUnsupportedMediaType             uint16 = 415 // RFC 7231, 6.5.13
	StatusUnsupportedMediaTypeCode                = "415"
	StatusUnsupportedMediaTypePhrase              = "Unsupported Media Type"
	StatusRangeNotSatisfiable              uint16 = 416 // RFC 7233, 4.4
	StatusRangeNotSatisfiableCode                 = "416"
	StatusRangeNotSatisfiablePhrase               = "Requested Range Not Satisfiable"
	StatusExpectationFailed                uint16 = 417 // RFC 7231, 6.5.14
	StatusExpectationFailedCode                   = "417"
	StatusExpectationFailedPhrase                 = "Expectation Failed"
	StatusTeapot                           uint16 = 418 // RFC 7168, 2.3.3
	StatusTeapotCode                              = "418"
	StatusTeapotPhrase                            = "I'm a teapot"
	StatusUnprocessableEntity              uint16 = 422 // RFC 4918, 11.2
	StatusUnprocessableEntityCode                 = "422"
	StatusUnprocessableEntityPhrase               = "Unprocessable Entity"
	StatusLocked                           uint16 = 423 // RFC 4918, 11.3
	StatusLockedCode                              = "423"
	StatusLockedPhrase                            = "Locked"
	StatusFailedDependency                 uint16 = 424 // RFC 4918, 11.4
	StatusFailedDependencyCode                    = "424"
	StatusFailedDependencyPhrase                  = "Failed Dependency"
	StatusUpgradeRequired                  uint16 = 426 // RFC 7231, 6.5.15
	StatusUpgradeRequiredCode                     = "426"
	StatusUpgradeRequiredPhrase                   = "Upgrade Required"
	StatusPreconditionRequired             uint16 = 428 // RFC 6585, 3
	StatusPreconditionRequiredCode                = "428"
	StatusPreconditionRequiredPhrase              = "Precondition Required"
	StatusTooManyRequests                  uint16 = 429 // RFC 6585, 4
	StatusTooManyRequestsCode                     = "429"
	StatusTooManyRequestsPhrase                   = "Too Many Requests"
	StatusHeaderFieldsTooLarge             uint16 = 431 // RFC 6585, 5
	StatusHeaderFieldsTooLargeCode                = "431"
	StatusHeaderFieldsTooLargePhrase              = "Header Fields Too Large"
	StatusUnavailableForLegalReasons       uint16 = 451 // RFC 7725, 3
	StatusUnavailableForLegalReasonsCode          = "451"
	StatusUnavailableForLegalReasonsPhrase        = "Unavailable For Legal Reasons"

	StatusInternalServerError                 uint16 = 500 // RFC 7231, 6.6.1
	StatusInternalServerErrorCode                    = "500"
	StatusInternalServerErrorPhrase                  = "Internal Server Error"
	StatusNotImplemented                      uint16 = 501 // RFC 7231, 6.6.2
	StatusNotImplementedCode                         = "501"
	StatusNotImplementedPhrase                       = "Not Implemented"
	StatusBadGateway                          uint16 = 502 // RFC 7231, 6.6.3
	StatusBadGatewayCode                             = "502"
	StatusBadGatewayPhrase                           = "Bad Gateway"
	StatusServiceUnavailable                  uint16 = 503 // RFC 7231, 6.6.4
	StatusServiceUnavailableCode                     = "503"
	StatusServiceUnavailablePhrase                   = "Service Unavailable"
	StatusGatewayTimeout                      uint16 = 504 // RFC 7231, 6.6.5
	StatusGatewayTimeoutCode                         = "504"
	StatusGatewayTimeoutPhrase                       = "Gateway Timeout"
	StatusHTTPVersionNotSupported             uint16 = 505 // RFC 7231, 6.6.6
	StatusHTTPVersionNotSupportedCode                = "505"
	StatusHTTPVersionNotSupportedPhrase              = "HTTP Version Not Supported"
	StatusVariantAlsoNegotiates               uint16 = 506 // RFC 2295, 8.1
	StatusVariantAlsoNegotiatesCode                  = "506"
	StatusVariantAlsoNegotiatesPhrase                = "Variant Also Negotiates"
	StatusInsufficientStorage                 uint16 = 507 // RFC 4918, 11.5
	StatusInsufficientStorageCode                    = "507"
	StatusInsufficientStoragePhrase                  = "Insufficient Storage"
	StatusLoopDetected                        uint16 = 508 // RFC 5842, 7.2
	StatusLoopDetectedCode                           = "508"
	StatusLoopDetectedPhrase                         = "Loop Detected"
	StatusNotExtended                         uint16 = 510 // RFC 2774, 7
	StatusNotExtendedCode                            = "510"
	StatusNotExtendedPhrase                          = "Not Extended"
	StatusNetworkAuthenticationRequired       uint16 = 511 // RFC 6585, 6
	StatusNetworkAuthenticationRequiredCode          = "511"
	StatusNetworkAuthenticationRequiredPhrase        = "Network Authentication Required"
)

var statusText = map[uint16]string{
	StatusContinue:           StatusContinuePhrase,
	StatusSwitchingProtocols: StatusSwitchingProtocolsPhrase,
	StatusProcessing:         StatusProcessingPhrase,

	StatusOK:                   StatusOKPhrase,
	StatusCreated:              StatusCreatedPhrase,
	StatusAccepted:             StatusAcceptedPhrase,
	StatusNonAuthoritativeInfo: StatusNonAuthoritativeInfoPhrase,
	StatusNoContent:            StatusNoContentPhrase,
	StatusResetContent:         StatusResetContentPhrase,
	StatusPartialContent:       StatusPartialContentPhrase,
	StatusMultiStatus:          StatusMultiStatusPhrase,
	StatusAlreadyReported:      StatusAlreadyReportedPhrase,
	StatusIMUsed:               StatusIMUsedPhrase,

	StatusMultipleChoices:   StatusMultipleChoicesPhrase,
	StatusMovedPermanently:  StatusMovedPermanentlyPhrase,
	StatusFound:             StatusFoundPhrase,
	StatusSeeOther:          StatusSeeOtherPhrase,
	StatusNotModified:       StatusNotModifiedPhrase,
	StatusUseProxy:          StatusUseProxyPhrase,
	StatusSwitchProxy:       StatusSwitchProxyPhrase,
	StatusTemporaryRedirect: StatusTemporaryRedirectPhrase,
	StatusPermanentRedirect: StatusPermanentRedirectPhrase,

	StatusBadRequest:                 StatusBadRequestPhrase,
	StatusUnauthorized:               StatusUnauthorizedPhrase,
	StatusPaymentRequired:            StatusPaymentRequiredPhrase,
	StatusForbidden:                  StatusForbiddenPhrase,
	StatusNotFound:                   StatusNotFoundPhrase,
	StatusMethodNotAllowed:           StatusMethodNotAllowedPhrase,
	StatusNotAcceptable:              StatusNotAcceptablePhrase,
	StatusProxyAuthRequired:          StatusProxyAuthRequiredPhrase,
	StatusRequestTimeout:             StatusRequestTimeoutPhrase,
	StatusConflict:                   StatusConflictPhrase,
	StatusGone:                       StatusGonePhrase,
	StatusLengthRequired:             StatusLengthRequiredPhrase,
	StatusPreconditionFailed:         StatusPreconditionFailedPhrase,
	StatusPayloadTooLarge:            StatusPayloadTooLargePhrase,
	StatusURITooLong:                 StatusURITooLongPhrase,
	StatusUnsupportedMediaType:       StatusUnsupportedMediaTypePhrase,
	StatusRangeNotSatisfiable:        StatusRangeNotSatisfiablePhrase,
	StatusExpectationFailed:          StatusExpectationFailedPhrase,
	StatusTeapot:                     StatusTeapotPhrase,
	StatusUnprocessableEntity:        StatusUnprocessableEntityPhrase,
	StatusLocked:                     StatusLockedPhrase,
	StatusFailedDependency:           StatusFailedDependencyPhrase,
	StatusUpgradeRequired:            StatusUpgradeRequiredPhrase,
	StatusPreconditionRequired:       StatusPreconditionRequiredPhrase,
	StatusTooManyRequests:            StatusTooManyRequestsPhrase,
	StatusHeaderFieldsTooLarge:       StatusHeaderFieldsTooLargePhrase,
	StatusUnavailableForLegalReasons: StatusUnavailableForLegalReasonsPhrase,

	StatusInternalServerError:           StatusInternalServerErrorPhrase,
	StatusNotImplemented:                StatusNotImplementedPhrase,
	StatusBadGateway:                    StatusBadGatewayPhrase,
	StatusServiceUnavailable:            StatusServiceUnavailablePhrase,
	StatusGatewayTimeout:                StatusGatewayTimeoutPhrase,
	StatusHTTPVersionNotSupported:       StatusHTTPVersionNotSupportedPhrase,
	StatusVariantAlsoNegotiates:         StatusVariantAlsoNegotiatesPhrase,
	StatusInsufficientStorage:           StatusInsufficientStoragePhrase,
	StatusLoopDetected:                  StatusLoopDetectedPhrase,
	StatusNotExtended:                   StatusNotExtendedPhrase,
	StatusNetworkAuthenticationRequired: StatusNetworkAuthenticationRequiredPhrase,
}

// GetStatusText returns a text for the HTTP Status code. It returns the empty
// string if the code is unknown.
func GetStatusText(code uint16) string {
	return statusText[code]
}
