/* For license and copyright information please see LEGAL file in repository */

package http

import "../convert"

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

// Marshal encode URI data to given httpPacket and update u.Raw
func (u *URI) Marshal(httpPacket []byte) []byte {
	var startLen int = len(httpPacket)

	if u.Scheme != "" {
		httpPacket = append(httpPacket, u.Scheme...)
		httpPacket = append(httpPacket, "://"...)
	}
	httpPacket = append(httpPacket, u.Authority...)
	httpPacket = append(httpPacket, u.Path...)
	if u.Path == "" {
		httpPacket = append(httpPacket, Slash)
	}
	if u.Query != "" {
		httpPacket = append(httpPacket, Question)
		httpPacket = append(httpPacket, u.Query...)
	}

	var raw []byte = httpPacket[startLen:]
	u.Raw = convert.UnsafeByteSliceToString(raw)

	return httpPacket
}

// UnMarshal use to parse and decode given raw URI to u
func (u *URI) UnMarshal(s string) (uriEnd int) {
	if s[0] == Asterisk {
		u.Path = s[:1]
		u.Raw = s[:1]
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
				u.Scheme = s[:i]
				i += 2                      // next loop will i+=1 so we just add i+=2
				authorityStartIndex = i + 1 // +3 due to have ://
			}
		case Slash:
			// Just check slash in middle of URI! If URI in origin form pathStartIndex always be 0!
			if authorityStartIndex != 0 && pathStartIndex == 0 {
				pathStartIndex = i
				u.Authority = s[authorityStartIndex:pathStartIndex]
			} else if !originForm && pathStartIndex == 0 && i != 0 {
				pathStartIndex = i
				u.Authority = s[:i]
			}
		case Question:
			// Check ? mark is first appear or it is part of some query key||value!
			if questionIndex == 0 {
				questionIndex = i
				u.Path = s[pathStartIndex:questionIndex]
			}
		case NumberSign:
			if numberSignIndex == 0 {
				numberSignIndex = i
				if questionIndex == 0 {
					u.Path = s[pathStartIndex:numberSignIndex]
				} else {
					u.Query = s[questionIndex+1 : numberSignIndex] // +1 due to we don't need '?'
				}
			}
		case SP:
			uriEnd = i
			if questionIndex == 0 && numberSignIndex == 0 {
				u.Path = s[pathStartIndex:uriEnd]
			}
			if numberSignIndex != 0 {
				u.Fragment = s[numberSignIndex+1 : uriEnd] // +1 due to we don't need '#'
			}
			if questionIndex != 0 && numberSignIndex == 0 {
				u.Query = s[questionIndex+1 : uriEnd] // +1 due to we don't need '?'
			}
			u.Raw = s[:uriEnd]
			// Don't need to continue loop!
			return
		}
	}
	return
}

// Len return length of Marshal()
func (u *URI) Len() int {
	return len(u.Scheme) + len(u.Authority) + len(u.Path) + len(u.Query) + 4 // 4 = len("://")+len("?")
}
