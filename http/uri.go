/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"io"

	"../convert"
	"../protocol"
)

// URI store http URI parts.
// https://tools.ietf.org/html/rfc2616#section-3.2
// https://tools.ietf.org/html/rfc2616#section-5.1.2
// https://tools.ietf.org/html/rfc3986
// Request-URI = "*" | absoluteURI | abs_path | authority
// http_URL = "http:" "//" host [ ":" port ] [ abs_path [ "?" query ]]
type URI struct {
	uri       string
	uriAsByte []byte
	scheme    string // = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
	authority string // host [ ":" port ]
	host      string // host without port if any exist in authority
	path      string //
	query     string // encoded query values, without '?'
	fragment  string // fragment for references, without '#'
}

func (u *URI) Init(uri string) {
	u.uri = uri
	u.uriAsByte = convert.UnsafeStringToByteSlice(uri)
	// TODO::: anything else?
}

func (u *URI) Set(scheme, authority, path, query string) {
	u.scheme, u.authority, u.path, u.query = scheme, authority, path, query

	u.uriAsByte = make([]byte, 0, u.len())
	if u.scheme != "" {
		u.uriAsByte = append(u.uriAsByte, u.scheme...)
		u.uriAsByte = append(u.uriAsByte, "://"...)
	}
	u.uriAsByte = append(u.uriAsByte, u.authority...)
	if u.path == "" {
		u.uriAsByte = append(u.uriAsByte, Slash)
	} else {
		u.uriAsByte = append(u.uriAsByte, u.path...)
	}
	if u.query != "" {
		u.uriAsByte = append(u.uriAsByte, Question)
		u.uriAsByte = append(u.uriAsByte, u.query...)
	}

	u.uri = convert.UnsafeByteSliceToString(u.uriAsByte)
}

func (u *URI) URI() string       { return u.uri }
func (u *URI) Scheme() string    { return u.scheme }
func (u *URI) Authority() string { return u.authority }
func (u *URI) Host() string      { return u.host }
func (u *URI) Path() string      { return u.path }
func (u *URI) Query() string     { return u.query }
func (u *URI) Fragment() string  { return u.fragment }

// checkHost check host of request by RFC 7230, section 5.3 rules: Must treat
//		GET / HTTP/1.1
//		Host: www.sabz.city
// and
//		GET https://www.sabz.city/ HTTP/1.1
//		Host: apis.sabz.city
// the same. In the second case, any Host line is ignored.
func (u *URI) checkHost(h *header) {
	if u.authority == "" {
		u.host = h.Get(HeaderKeyHost)
	}
	// TODO::: decode host (remove port if exist) from authority
	u.host = u.authority
}

/*
********** protocol.Codec interface **********
 */

func (u *URI) MediaType() string    { return "application/uri" } // application/x-www-form-urlencoded
func (u *URI) CompressType() string { return "" }

func (u *URI) Decode(reader io.Reader) (err protocol.Error) {
	// TODO:::
	return
}

func (u *URI) Encode(writer io.Writer) (err error) {
	var encodedURI = u.Marshal()
	_, err = writer.Write(encodedURI)
	return
}

// Marshal encode URI data and return it.
func (u *URI) Marshal() (encodedURI []byte) {
	return u.uriAsByte
}

// MarshalTo encode URI data to given httpPacket and update u.uri and return httpPacket with new len.
func (u *URI) MarshalTo(httpPacket []byte) []byte {
	httpPacket = append(httpPacket, u.uriAsByte...)
	return httpPacket
}

// Unmarshal use to parse and decode given URI to u
func (u *URI) Unmarshal(s string) (uriEnd int) {
	if s[0] == Asterisk {
		u.uri = s[:1]
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
			u.Init(s[:uriEnd])
			// Don't need to continue loop!
			return
		}
	}
	return
}

// Len return length of Marshal()
func (u *URI) Len() (ln int) { return len(u.uriAsByte) }

func (u *URI) len() (ln int) {
	ln = 4 // 4 == len("://")+len("?")
	ln += len(u.scheme) + len(u.authority) + len(u.path) + len(u.query)
	return
}

/*
********** io package interfaces **********
 */

func (u *URI) WriteTo(writer io.Writer) (n int64, err error) {
	var encodedURI = u.Marshal()
	var writeLength int
	writeLength, err = writer.Write(encodedURI)
	n = int64(writeLength)
	return
}
