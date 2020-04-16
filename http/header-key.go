/* For license and copyright information please see LEGAL file in repository */

package http

// Common standard HTTP header keys
const (
	HeaderKeyHost          = "Host"
	HeaderKeyUpgrade       = "Upgrade"
	HeaderKeyUserAgent     = "User-Agent"
	HeaderKeyCookie        = "Cookie"
	HeaderKeySetCookie     = "Set-Cookie" // response header
	HeaderKeyDate          = "Date"
	HeaderKeyAuthorization = "Authorization"
	HeaderKeyCacheControl  = "Cache-Control"
	HeaderKeyOrigin        = "Origin"
	HeaderKeyVia           = "Via"
	HeaderKeyLastModified  = "Last-Modified" // response header
	HeaderKeyETag          = "ETag"          // response header
	HeaderKeyServer        = "Server"        // response header

	HeaderKeyAcceptContent  = "Accept"
	HeaderKeyAcceptCharset  = "Accept-Charset"
	HeaderKeyAcceptEncoding = "Accept-Encoding"
	HeaderKeyAcceptLanguage = "Accept-Language"
	HeaderKeyAcceptDatetime = "Accept-Datetime"
	HeaderKeyAcceptPatch    = "Accept-Patch"  // response header
	HeaderKeyAcceptRanges   = "Accept-Ranges" // response header

	HeaderKeyContentLength      = "Content-Length"
	HeaderKeyContentMD5         = "Content-MD5"
	HeaderKeyContentType        = "Content-Type"
	HeaderKeyContentDisposition = "Content-Disposition" // response header
	HeaderKeyContentEncoding    = "Content-Encoding"    // response header
	HeaderKeyContentLanguage    = "Content-Language"    // response header
	HeaderKeyContentLocation    = "Content-Location"    // response header
	HeaderKeyContentRange       = "Content-Range"       // response header

	HeaderKeyIfMatch           = "If-Match"
	HeaderKeyIfNoneMatch       = "If-None-Match"
	HeaderKeyIfModifiedSince   = "If-Modified-Since"
	HeaderKeyIfUnmodifiedSince = "If-Unmodified-Since"
	HeaderKeyIfRange           = "If-Range"

	HeaderKeyAccessControlAllowOrigin      = "Access-Control-Allow-Origin"      // response header
	HeaderKeyAccessControlAllowMethods     = "Access-Control-Allow-Methods"     // response header
	HeaderKeyAccessControlAllowCredentials = "Access-Control-Allow-Credentials" // response header
	HeaderKeyAccessControlAllowHeaders     = "Access-Control-Allow-Headers"     // response header
	HeaderKeyAccessControlExposeHeaders    = "Access-Control-Expose-Headers"    // response header
	HeaderKeyAccessControlMaxAge           = "Access-Control-Max-Age"           // response header
	HeaderKeyAccessControlRequestHeaders   = "Access-Control-Request-Headers"   // response header
	HeaderKeyAccessControlRequestMethod    = "Access-Control-Request-Method"    // response header
)
