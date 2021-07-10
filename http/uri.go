/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"../convert"
	"../giti"
)

// URI store raw URI and all parts of it.
// https://tools.ietf.org/html/rfc2616#section-3.2
// https://tools.ietf.org/html/rfc2616#section-5.1.2
// https://tools.ietf.org/html/rfc3986
// Request-URI = "*" | absoluteURI | abs_path | authority
// http_URL = "http:" "//" host [ ":" port ] [ abs_path [ "?" query ]]
type URI struct {
	raw       string
	scheme    string // = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
	authority string // = [ userinfo "@" ] host [ ":" port ]
	path      string //
	query     string // encoded query values, without '?'
	fragment  string // fragment for references, without '#'
}

func (u *URI) Raw() string       { return u.raw }
func (u *URI) Scheme() string    { return u.scheme }
func (u *URI) Authority() string { return u.authority }
func (u *URI) Path() string      { return u.path }
func (u *URI) Query() string     { return u.query }
func (u *URI) Fragment() string  { return u.fragment }
func (u *URI) Set(scheme, authority, path, query string) {
	u.scheme, u.authority, u.path, u.query = scheme, authority, path, query
}

func (u *URI) Decode(buf giti.Buffer) (err giti.Error) {
	// TODO:::
	return
}

func (u *URI) Encode(buf giti.Buffer) {
	buf.Set(u.Marshal(buf.Get()))
}

// Marshal encode URI data to given httpPacket and update u.raw
func (u *URI) Marshal(httpPacket []byte) []byte {
	var startLen int = len(httpPacket)

	if u.scheme != "" {
		httpPacket = append(httpPacket, u.scheme...)
		httpPacket = append(httpPacket, "://"...)
	}
	httpPacket = append(httpPacket, u.authority...)
	httpPacket = append(httpPacket, u.path...)
	if u.path == "" {
		httpPacket = append(httpPacket, Slash)
	}
	if u.query != "" {
		httpPacket = append(httpPacket, Question)
		httpPacket = append(httpPacket, u.query...)
	}

	var raw []byte = httpPacket[startLen:]
	u.raw = convert.UnsafeByteSliceToString(raw)

	return httpPacket
}

// UnMarshal use to parse and decode given raw URI to u
func (u *URI) UnMarshal(s string) (uriEnd int) {
	if s[0] == Asterisk {
		u.path = s[:1]
		u.raw = s[:1]
		return 1
	}

	var originForm bool
	if s[0] == '/' {
		originForm = true
	}

	var authorityStartIndex, pathStartIndex, questionIndex, numberSignIndex int
	var ln = len(s)
	for i := 0; i < ln; i++ {
		switch s[i] {
		case Colon:
			// Check : mark is first appear before any start||end sign or it is part of others!
			if authorityStartIndex == 0 {
				u.scheme = s[:i]
				i += 2                      // next loop will i+=1 so we just add i+=2
				authorityStartIndex = i + 1 // +3 due to have ://
			}
		case Slash:
			// Just check slash in middle of URI! If URI in origin form pathStartIndex always be 0!
			if authorityStartIndex != 0 && pathStartIndex == 0 {
				pathStartIndex = i
				u.authority = s[authorityStartIndex:pathStartIndex]
			} else if !originForm && pathStartIndex == 0 && i != 0 {
				pathStartIndex = i
				u.authority = s[:i]
			}
		case Question:
			// Check ? mark is first appear or it is part of some query key||value!
			if questionIndex == 0 {
				questionIndex = i
				u.path = s[pathStartIndex:questionIndex]
			}
		case NumberSign:
			if numberSignIndex == 0 {
				numberSignIndex = i
				if questionIndex == 0 {
					u.path = s[pathStartIndex:numberSignIndex]
				} else {
					u.query = s[questionIndex+1 : numberSignIndex] // +1 due to we don't need '?'
				}
			}
		case SP:
			uriEnd = i
			if questionIndex == 0 && numberSignIndex == 0 {
				u.path = s[pathStartIndex:uriEnd]
			}
			if numberSignIndex != 0 {
				u.fragment = s[numberSignIndex+1 : uriEnd] // +1 due to we don't need '#'
			}
			if questionIndex != 0 && numberSignIndex == 0 {
				u.query = s[questionIndex+1 : uriEnd] // +1 due to we don't need '?'
			}
			u.raw = s[:uriEnd]
			// Don't need to continue loop!
			return
		}
	}
	return
}

// Len return length of Marshal()
func (u *URI) Len() (ln int) {
	ln = 4 // 4 == len("://")+len("?")
	ln += len(u.scheme) + len(u.authority) + len(u.path) + len(u.query)
	return
}
