/* For license and copyright information please see LEGAL file in repository */

package http

type requestStandards struct {
	Methods reqMethods
	Headers reqHeaders
}

type reqMethods struct {
	GET, POST, HEAD, PUT, DELETE, OPTIONS, CONNECT, TRACE, PATCH string
}

type reqHeaders struct {
	Accept             reqAccept
	Authorization      string
	CacheControl       string
	Connection         string
	Cookie             string
	Content            reqContent
	Date               string
	Expect             string
	Forwarded          string
	From               string
	Host               string
	If                 reqIf
	MaxForwards        string
	Origin             string
	Pragma             string
	ProxyAuthorization string
	Range              string
	Referer            string
	TE                 string
	UserAgent          string
	Upgrade            string
	Via                string
	Warning            string
}

type reqAccept struct {
	Content  string
	Charset  string
	Encoding string
	Language string
	Datetime string
}

type reqContent struct {
	Length string
	MD5    string
	Type   string
}

type reqIf struct {
	Match           string
	NoneMatch       string
	ModifiedSince   string
	UnmodifiedSince string
	Range           string
}

// RequestStandards : All text of Standard http request fileds
var RequestStandards = requestStandards{
	Methods: reqMethods{
		GET:     "GET",
		POST:    "POST",
		HEAD:    "HEAD",
		PUT:     "PUT",
		DELETE:  "DELETE",
		OPTIONS: "OPTIONS",
		CONNECT: "CONNECT",
		TRACE:   "TRACE",
		PATCH:   "PATCH"},
	Headers: reqHeaders{
		Accept: reqAccept{
			Content:  "Accept",
			Charset:  "Accept-Charset",
			Encoding: "Accept-Encoding",
			Language: "Accept-Language",
			Datetime: "Accept-Datetime"},
		Authorization: "Authorization",
		CacheControl:  "Cache-Control",
		Connection:    "Connection",
		Cookie:        "Cookie",
		Content: reqContent{
			Length: "Content-Length",
			MD5:    "Content-MD5",
			Type:   "Content-Type"},
		Date:      "Date",
		Expect:    "Expect",
		Forwarded: "Forwarded",
		From:      "From",
		Host:      "Host",
		If: reqIf{
			Match:           "If-Match",
			NoneMatch:       "If-None-Match",
			ModifiedSince:   "If-Modified-Since",
			UnmodifiedSince: "If-Unmodified-Since",
			Range:           "If-Range"},
		MaxForwards:        "Max-Forwards",
		Origin:             "Origin",
		Pragma:             "Pragma",
		ProxyAuthorization: "Proxy-Authorization",
		Range:              "Range",
		Referer:            "Referer",
		TE:                 "TE",
		UserAgent:          "User-Agent",
		Upgrade:            "Upgrade",
		Via:                "Via",
		Warning:            "Warning"}}
