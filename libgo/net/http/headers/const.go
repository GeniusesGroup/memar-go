/* For license and copyright information please see the LEGAL file in the code repository */

package headers

const (
	// TimeFormat is the time format to use when generating times in HTTP
	// headers. It is like time.RFC1123 but hard-codes GMT as the time
	// zone. The time being formatted must be in UTC for Format to
	// generate the correct format.
	TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
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

// Standard HTTP header keys
// https://www.iana.org/assignments/message-headers/message-headers.xhtml
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
const (
	// Request context
	Key_From           = "From"
	Key_Host           = "Host"
	Key_Referer        = "Referer"
	Key_ReferrerPolicy = "Referrer-Policy"
	Key_UserAgent      = "User-Agent"

	// Response context
	Key_Allow   = "Allow"
	Key_Server  = "Server"
	Key_ErrorID = "Error-ID"

	// Authentication
	Key_Authorization      = "Authorization"
	Key_ProxyAuthorization = "Proxy-Authorization"
	Key_ProxyAuthenticate  = "Proxy-Authenticate" // res
	Key_WWWAuthenticate    = "WWW-Authenticate"   // res

	// Caching
	Key_Age           = "Age"           // res
	Key_CacheControl  = "Cache-Control" // req & res
	Key_ClearSiteData = "Clear-Site-Data"
	Key_Expires       = "Expires" // res
	Key_Pragma        = "Pragma"  // req & res
	Key_Warning       = "Warning" // req & res
	Key_Vary          = "Vary"    // res

	// Conditionals
	Key_ETag              = "ETag" // res
	Key_IfMatch           = "If-Match"
	Key_IfNoneMatch       = "If-None-Match"
	Key_IfModifiedSince   = "If-Modified-Since"
	Key_IfUnmodifiedSince = "If-Unmodified-Since"
	Key_LastModified      = "Last-Modified" // res

	// Range requests
	Key_AcceptRanges = "Accept-Ranges" // res
	Key_ContentRange = "Content-Range" // res
	Key_IfRange      = "If-Range"
	Key_Range        = "Range"

	// Connection management
	Key_Connection = "Connection" // req & res
	Key_KeepAlive  = "Keep-Alive"
	Key_Upgrade    = "Upgrade"

	// CORS
	Key_AccessControlAllowOrigin      = "Access-Control-Allow-Origin"      // res
	Key_AccessControlAllowMethods     = "Access-Control-Allow-Methods"     // res
	Key_AccessControlAllowCredentials = "Access-Control-Allow-Credentials" // res
	Key_AccessControlAllowHeaders     = "Access-Control-Allow-Headers"     // res
	Key_AccessControlExposeHeaders    = "Access-Control-Expose-Headers"    // res
	Key_AccessControlMaxAge           = "Access-Control-Max-Age"           // res
	Key_AccessControlRequestHeaders   = "Access-Control-Request-Headers"   // res
	Key_AccessControlRequestMethod    = "Access-Control-Request-Method"    // res
	Key_Origin                        = "Origin"
	Key_TimingAllowOrigin             = "Timing-Allow-Origin"
	Key_XPermittedCrossDomainPolicies = "X-Permitted-Cross-Domain-Policies"

	// Content negotiation
	Key_AcceptContent  = "Accept"
	Key_AcceptCharset  = "Accept-Charset" // deprecated
	Key_AcceptEncoding = "Accept-Encoding"
	Key_AcceptLanguage = "Accept-Language"
	Key_AcceptDatetime = "Accept-Datetime"
	Key_AcceptPatch    = "Accept-Patch" // res

	// Message body information
	Key_ContentLength      = "Content-Length"      // req & res
	Key_ContentMD5         = "Content-MD5"         // req & res
	Key_ContentType        = "Content-Type"        // req & res
	Key_ContentDisposition = "Content-Disposition" // res
	Key_ContentEncoding    = "Content-Encoding"    // res
	Key_ContentLanguage    = "Content-Language"    // res
	Key_ContentLocation    = "Content-Location"    // res
	Key_TransferEncoding   = "Transfer-Encoding"   // res

	// Not ordered
	Key_Cookie                  = "Cookie"
	Key_SetCookie               = "Set-Cookie" // res
	Key_Date                    = "Date"       // req & res
	Key_Via                     = "Via"
	Key_Expect                  = "Expect"
	Key_Forwarded               = "Forwarded"
	Key_MaxForwards             = "Max-Forwards"
	Key_TE                      = "TE"
	Key_AltSvc                  = "Alt-Svc"                   // res
	Key_Link                    = "Link"                      // res
	Key_Location                = "Location"                  // res
	Key_P3P                     = "P3P"                       // res
	Key_PublicKeyPins           = "Public-Key-Pins"           // res
	Key_Refresh                 = "Refresh"                   // res
	Key_RetryAfter              = "Retry-After"               // res
	Key_StrictTransportSecurity = "Strict-Transport-Security" // res
	Key_Trailer                 = "Trailer"                   // res
	Key_Tk                      = "Tk"                        // res
	Key_XFrameOptions           = "X-Frame-Options"           // res
	Key_NonAuthoritativeReason  = "Non-Authoritative-Reason"  // res
)

// Standard HTTP header values
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
const (
	// Connection management
	Value_KeepAlive = "Keep-Alive"
	Value_Close     = "Close"

	// Message body information
	Value_Chunked  = "chunked"
	Value_Compress = "compress"
	Value_Deflate  = "deflate"
	Value_GZIP     = "gzip"
)
