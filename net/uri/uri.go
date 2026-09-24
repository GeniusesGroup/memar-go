/* For license and copyright information please see the LEGAL file in the code repository */

package uri

import (
	container_p "memar/computer/adt/container/protocol"
	buffer_p "memar/computer/buffer/protocol"
	"memar/computer/buffer/byteslice/convert"
	"memar/computer/datatype"
	error_p "memar/process/error/protocol"
	string_p "memar/codec/string/protocol"
)

// Parsed use to FIX `URI` name with its method as `URI()` when embed to other capsule.
type Parsed = URI

// URI store http URI parts.
// https://tools.ietf.org/html/rfc3986
// https://tools.ietf.org/html/rfc2616#section-3.2
// https://tools.ietf.org/html/rfc2616#section-5.1.2
// Request-URI = "*" | absoluteURI | abs_path | authority
// http_URL = "http:" "//" host [ ":" port ] [ abs_path [ "?" query ]]
type URI struct {
	datatype.DataType

	raw string

	scheme string // = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
	AU
	path     string //
	query    string // encoded query values, without '?'
	fragment string // fragment for references, without '#'
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (u *URI) Init(uri string) (err error_p.Error) {
	_, err = u.FromString(uri)
	return
}
func (u *URI) Reinit(uri string) (err error_p.Error) {
	_, err = u.FromString(uri)
	return
}
func (u *URI) Deinit() (err error_p.Error) {
	return
}

func (u *URI) Set(scheme, authority, path, query, fragment string) {
	u.discardRaw()
	u.scheme, u.authority, u.path, u.query, u.fragment = scheme, authority, path, query, fragment
}

func (u *URI) URI() string      { return u.raw }
func (u *URI) Scheme() string   { return u.scheme }
func (u *URI) Path() string     { return u.path }
func (u *URI) Query() string    { return u.query }
func (u *URI) Fragment() string { return u.fragment }

func (u *URI) SetURI(uri string)    { u.discardRaw(); u.raw = uri }
func (u *URI) SetScheme(s string)   { u.discardRaw(); u.scheme = s }
func (u *URI) SetPath(p string)     { u.discardRaw(); u.path = p }
func (u *URI) SetQuery(q string)    { u.discardRaw(); u.query = q }
func (u *URI) SetFragment(f string) { u.discardRaw(); u.fragment = f }

func (u *URI) ParseQuery() (q Query) { q.Init(u.query); return }

// IsAbs reports whether the URL is absolute.
// Absolute means that it has a non-empty scheme.
func (u *URI) IsAbs() bool { return u.scheme.OccupiedLength() > 0 }

//memar:impl memar/codec/protocol.Field_Length
func (u *URI) SerializationLength() (ln container_p.NumberOfElement) {
	ln = u.raw.OccupiedLength()
	if ln == 0 {
		ln = u.serializationLength()
	}
	return
}

//memar:impl memar/protocol.Decoder
func (u *URI) Decode(source buffer_p.Buffer) (err error_p.Error) {
	var char, _ = source.Peek()
	if char == sign_Asterisk {
		u.raw = sign_Asterisk_String
		source.Pop()
		return
	}

	var originForm bool
	char, _ = source.Peek()
	if char == '/' {
		originForm = true
	}
	var authorityStartIndex, pathStartIndex, questionIndex, numberSignIndex int
	// Don't need to continue loop anymore if we see Space character.
	for char != sign_SP {
		switch char {
		case sign_Colon:
			// Check : mark is first appear before any start||end sign or it is part of others!
			if authorityStartIndex == 0 {
				u.scheme = s[:i]
				i += 2                      // next loop will i+=1 so we just add i+=2
				authorityStartIndex = i + 1 // +3 due to have ://
			}
		case sign_Slash:
			// Just check slash in middle of URI! If URI in origin form pathStartIndex always be 0!
			if authorityStartIndex != 0 && pathStartIndex == 0 {
				pathStartIndex = i
				u.authority = s[authorityStartIndex:pathStartIndex]
			} else if !originForm && pathStartIndex == 0 && i != 0 {
				pathStartIndex = i
				u.authority = s[:i]
			}
		case sign_Question:
			// Check ? mark is first appear or it is part of some query key||value!
			if questionIndex == 0 {
				questionIndex = i
				u.path = s[pathStartIndex:questionIndex]
			}
		case sign_NumberSign:
			if numberSignIndex == 0 {
				numberSignIndex = i
				if questionIndex == 0 {
					u.path = s[pathStartIndex:numberSignIndex]
				} else {
					u.query = s[questionIndex+1 : numberSignIndex] // +1 due to we don't need '?'
				}
			}
		}

		char, _ = source.Peek()
	}

	uriEnd = container_p.NumberOfElement(i)
	if questionIndex == 0 && numberSignIndex == 0 {
		u.path = s[pathStartIndex:uriEnd]
	}
	if numberSignIndex != 0 {
		u.fragment = s[numberSignIndex+1 : uriEnd] // +1 due to we don't need '#'
	}
	if questionIndex != 0 && numberSignIndex == 0 {
		u.query = s[questionIndex+1 : uriEnd] // +1 due to we don't need '?'
	}

	// u.raw = source.
	return
}

//memar:impl memar/protocol.Encoder
func (u *URI) Encode(destination buffer_p.Buffer) (err error_p.Error) {
	if !u.raw.IsEmpty() {
		_, err = destination.Concat(u.raw)
	} else {
		if !u.scheme.IsEmpty() {
			_, err = destination.Append(convert.UnsafeStringToByteSlice(u.scheme)...)
			if err != nil {
				return
			}
			_, err = destination.Append(convert.UnsafeStringToByteSlice("://")...)
			if err != nil {
				return
			}
		}
		_, err = destination.Append(convert.UnsafeStringToByteSlice(u.authority)...)
		if err != nil {
			return
		}
		if u.path == "" {
			_, err = destination.Append(sign_Slash)
			if err != nil {
				return
			}
		} else {
			_, err = destination.Append(convert.UnsafeStringToByteSlice(u.path)...)
			if err != nil {
				return
			}
		}
		if u.query != "" {
			_, err = destination.Append(sign_Question)
			if err != nil {
				return
			}
			_, err = destination.Append(convert.UnsafeStringToByteSlice(u.query)...)
			if err != nil {
				return
			}
		}
		if u.fragment != "" {
			_, err = destination.Append(sign_NumberSign)
			if err != nil {
				return
			}
			_, err = destination.Append(convert.UnsafeStringToByteSlice(u.fragment)...)
			if err != nil {
				return
			}
		}
	}
	return
}

// Marshal encode URI data and return it.
func (u *URI) Marshal(destination []byte) (n container_p.NumberOfElement, err error_p.Error) {
	n = u.SerializationLength()
	var desCap = cap(destination) - len(destination)
	if n > container_p.NumberOfElement(desCap) {
		// TODO::: return proper error
		// err =
		return
	}
	if u.raw != "" {
		copy(destination[len(destination):], u.raw)
	} else {
		_ = u.marshalTo(destination)
	}
	return
}

// Unmarshal use to parse and decode given URI to u
func (u *URI) Unmarshal(source []byte) (n container_p.NumberOfElement, err error_p.Error) {
	n, err = u.FromString(convert.UnsafeByteSliceToString(source))
	return
}

// FromString use to parse and decode given URI to u
func (u *URI) FromString(s string) (uriEnd container_p.NumberOfElement, err error_p.Error) {
	if s[0] == sign_Asterisk {
		uriEnd = 1
	} else {
		var originForm bool
		if s[0] == '/' {
			originForm = true
		}

		var authorityStartIndex, pathStartIndex, questionIndex, numberSignIndex int
		var ln = len(s)
		var i int
	Loop:
		for i = 0; i < ln; i++ {
			switch s[i] {
			case sign_Colon:
				// Check : mark is first appear before any start||end sign or it is part of others!
				if authorityStartIndex == 0 {
					u.scheme = s[:i]
					i += 2                      // next loop will i+=1 so we just add i+=2
					authorityStartIndex = i + 1 // +3 due to have ://
				}
			case sign_Slash:
				// Just check slash in middle of URI! If URI in origin form pathStartIndex always be 0!
				if authorityStartIndex != 0 && pathStartIndex == 0 {
					pathStartIndex = i
					u.authority = s[authorityStartIndex:pathStartIndex]
				} else if !originForm && pathStartIndex == 0 && i != 0 {
					pathStartIndex = i
					u.authority = s[:i]
				}
			case sign_Question:
				// Check ? mark is first appear or it is part of some query key||value!
				if questionIndex == 0 {
					questionIndex = i
					u.path = s[pathStartIndex:questionIndex]
				}
			case sign_NumberSign:
				if numberSignIndex == 0 {
					numberSignIndex = i
					if questionIndex == 0 {
						u.path = s[pathStartIndex:numberSignIndex]
					} else {
						u.query = s[questionIndex+1 : numberSignIndex] // +1 due to we don't need '?'
					}
				}
			case sign_SP:
				// Don't need to continue loop anymore
				break Loop
			}
		}

		uriEnd = container_p.NumberOfElement(i)
		if questionIndex == 0 && numberSignIndex == 0 {
			u.path = s[pathStartIndex:uriEnd]
		}
		if numberSignIndex != 0 {
			u.fragment = s[numberSignIndex+1 : uriEnd] // +1 due to we don't need '#'
		}
		if questionIndex != 0 && numberSignIndex == 0 {
			u.query = s[questionIndex+1 : uriEnd] // +1 due to we don't need '?'
		}
	}

	u.raw = s[:uriEnd]
	return
}

func (u *URI) marshalTo(httpPacket []byte) []byte {
	var uriStart = len(httpPacket)
	if u.scheme != "" {
		httpPacket = append(httpPacket, u.scheme...)
		httpPacket = append(httpPacket, "://"...)
	}
	httpPacket = append(httpPacket, u.authority...)
	if u.path == "" {
		httpPacket = append(httpPacket, sign_Slash)
	} else {
		httpPacket = append(httpPacket, u.path...)
	}
	if u.query != "" {
		httpPacket = append(httpPacket, sign_Question)
		httpPacket = append(httpPacket, u.query...)
	}

	// TODO::: below code cause memory leak if dev use u.raw in other places due to GC can't free whole http packet
	u.raw = convert.UnsafeByteSliceToString(httpPacket[uriStart:])
	return httpPacket
}

func (u *URI) serializationLength() (ln container_p.NumberOfElement) {
	ln = 4 // 4 == len("://")+len("?")
	ln += container_p.NumberOfElement(len(u.scheme) + len(u.authority) + len(u.path) + len(u.query) + len(u.fragment))
	return
}

// Discard any exciting data in raw uri.
// It is use when any part of uri want to change after Init()
func (u *URI) discardRaw() {
	u.raw = ""
}
