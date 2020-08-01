/* For license and copyright information please see LEGAL file in repository */

package http

import "unsafe"

// URI store raw URI and all parts of it.
// https://tools.ietf.org/html/rfc2616#section-3.2
// https://tools.ietf.org/html/rfc2616#section-5.1.2
// https://tools.ietf.org/html/rfc3986
// Request-URI = "*" | absoluteURI | abs_path | authority
// http_URL = "http:" "//" host [ ":" port ] [ abs_path [ "?" query ]]
type URI struct {
	Raw       string
	Scheme    string // = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
	Authority string // = [ userinfo "@" ] host [ ":" port ]
	Path      string //
	Query     string // encoded query values, without '?'
	Fragment  string // fragment for references, without '#'
}

// Marshal encode URI data to u.RawURI
func (u *URI) Marshal() {
	var buf = make([]byte, len(u.Scheme)+len(u.Authority)+len(u.Path)+len(u.Query))
	if u.Scheme != "" {
		buf = append(buf, u.Scheme...)
		buf = append(buf, "://"...)
	}
	buf = append(buf, u.Authority...)
	buf = append(buf, u.Path...)
	if u.Query != "" {
		buf = append(buf, '?')
		buf = append(buf, u.Path...)
	}
	u.Raw = *(*string)(unsafe.Pointer(&buf))
}

// MarshalRequestURI returns the encoded path?query
// string that would be used in an HTTP request for u.
func (u *URI) MarshalRequestURI() string {
	if len(u.Path) > 0 && len(u.Query) > 0 {
		return u.Path + "?" + u.Query
	} else if len(u.Path) > 0 {
		return u.Path
	}
	return "/"
}

// UnMarshal use to parse and decode given raw URI to u
func (u *URI) UnMarshal(raw string) {
	u.Raw = raw

	if raw == "*" {
		u.Path = "*"
		return
	}

	var pathStartIndex, pathEndIndex int
	var ln = len(raw)
	for i := 0; i < ln; i++ {
		switch raw[i] {
		case ':':
			pathStartIndex = i + 3 // +3 due to have ://
			u.Scheme = raw[:i]
		case '?':
			pathEndIndex = i
			u.Query = raw[i+1:] // +1 due to we don't need '?'
		case '#':
			if pathEndIndex == 0 {
				pathEndIndex = i
			}
			u.Fragment = raw[i+1:] // +1 due to we don't need '#'
			// Don't need to continue loop!
			break
		}
	}
	if pathEndIndex == 0 {
		pathEndIndex = ln
	}
	u.Path = raw[pathStartIndex:pathEndIndex]

	ln = len(u.Fragment)
	if ln > 0 && u.Query != "" {
		u.Query = u.Query[:len(u.Query)-ln-1] // -1 due to we don't need '#'
	}
}
