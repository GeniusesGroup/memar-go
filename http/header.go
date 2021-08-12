/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"strings"

	"../giti"
)

// header is represent HTTP header structure!
type header struct {
	headers    map[string][]string
	valuesPool []string
}

func (h *header) init() {
	h.headers = make(map[string][]string, 16)
	h.valuesPool = make([]string, 16)
}

// Get returns the first value associated with the given key.
// Key must already be in CanonicalHeaderKey form.
func (h *header) Get(key string) string {
	if v := h.headers[key]; len(v) > 0 {
		return v[0]
	}
	return ""
}

// Gets returns all values associated with the given key.
// Key must already be in CanonicalHeaderKey form.
func (h *header) Gets(key string) []string {
	if v := h.headers[key]; len(v) > 0 {
		return v
	}
	return nil
}

// Add adds the key, value pair to the header.
// It appends to any existing values associated with key.
// Key must already be in CanonicalHeaderKey form.
func (h *header) Add(key, value string) {
	var values []string
	values = h.headers[key]
	if values == nil {
		h.Set(key, value)
	} else {
		h.headers[key] = append(values, value)
	}
}

// Adds append given values to end of given key exiting values!
// Key must already be in CanonicalHeaderKey form.
func (h *header) Adds(key string, values []string) {
	h.headers[key] = append(h.headers[key], values...)
}

// Set replace given value in given key.
// Key must already be in CanonicalHeaderKey form.
func (h *header) Set(key string, value string) {
	if len(h.valuesPool) == 0 {
		h.valuesPool = make([]string, 16)
	}
	// More than likely this will be a single-element key. Most headers aren't multi-valued.
	// Set the capacity on valuesPool[0] to 1, so any future append won't extend the slice into the other strings.
	var values []string = h.valuesPool[:1:1]
	h.valuesPool = h.valuesPool[1:]
	values[0] = value
	h.headers[key] = values
}

// Sets sets the header entries associated with key to
// the single element value. It replaces any existing values associated with key.
// Key must already be in CanonicalHeaderKey form.
func (h *header) Sets(key string, values []string) {
	h.headers[key] = values
}

// Del deletes the values associated with key.
// Key must already be in CanonicalHeaderKey form.
func (h *header) Del(key string) {
	delete(h.headers, key)
}

// GetCookies parses and returns the Cookie headers.
// By related RFC we just support one Cookie in header.
// https://tools.ietf.org/html/rfc6265#section-5.4
func (h *header) GetCookies() (cookies []Cookie) {
	var cookie = h.Get(HeaderKeyCookie)
	if len(cookie) == 0 {
		return
	}
	var index int
	cookies = make([]Cookie, 0, 8)
	var c Cookie
	for {
		index = strings.IndexByte(cookie, ';')
		if index == -1 {
			c.UnMarshal(cookie)
			cookies = append(cookies, c)
			return
		}
		c.UnMarshal(cookie[:index])
		cookies = append(cookies, c)

		cookie = cookie[index+2:]
	}
}

// SetCookies parses and set them to Cookie header.
func (h *header) SetCookies(cookies []Cookie) {
	var b strings.Builder
	var ln = len(cookies)
	var i int
	for ; ; i++ {
		b.WriteString(cookies[i].Name)
		b.WriteByte('=')
		b.WriteString(cookies[i].Value)
		if i < ln {
			b.WriteString(SemiColonSpace)
		} else {
			break
		}
	}
	h.Set(HeaderKeyCookie, b.String())
}

// GetSetCookies parses and returns the Set-Cookie headers.
// By related RFC must exist just one Set-Cookie in each line of header.
// https://tools.ietf.org/html/rfc6265#section-4.1.1
func (h *header) GetSetCookies() (setCookies []SetCookie) {
	var scs = h.Gets(HeaderKeySetCookie)
	var setCookieCount = len(scs)
	if setCookieCount == 0 {
		return
	}
	setCookies = make([]SetCookie, setCookieCount)
	for i := 0; i < setCookieCount; i++ {
		setCookies[i].UnMarshal(scs[i])
	}
	return
}

// SetSetCookies parses and set given Set-Cookies to header.
// By related RFC must exist just one Set-Cookie in each line of header.
// https://tools.ietf.org/html/rfc6265#section-4.1.1
func (h *header) SetSetCookies(setCookies []SetCookie) {
	var ln = len(setCookies)
	for i := 0; i < ln; i++ {
		h.Add(HeaderKeySetCookie, setCookies[i].Marshal())
	}
}

// FixPragmaCacheControl do as RFC 7234, section 5.4: Treat [Pragma: no-cache] as [Cache-Control: no-cache]
func (h *header) FixPragmaCacheControl() {
	if h.Get(HeaderKeyPragma) == "no-cache" {
		if h.Gets(HeaderKeyCacheControl) == nil {
			h.Set(HeaderKeyCacheControl, "no-cache")
		}
	}
}

// Exclude eliminate headers by given keys!
func (h *header) Exclude(exclude map[string]bool) {
	for key := range exclude {
		delete(h.headers, key)
	}
}

/*
********** giti.Codec interface **********
 */

func (h *header) Decode(buf giti.Buffer) (err giti.Error) {
	// TODO:::
	return
}

func (h *header) Encode(buf giti.Buffer) {
	buf.Set(h.MarshalTo(buf.Get()))
}

// Marshal enecodes whole h *header data and return httpHeader!
func (h *header) Marshal() (httpHeader []byte) {
	httpHeader = make([]byte, 0, h.Len())
	httpHeader = h.MarshalTo(httpHeader)
	return
}

// MarshalTo enecodes (h *header) data to given httpPacket.
func (h *header) MarshalTo(httpPacket []byte) []byte {
	// TODO::: some header key must not inline by coma like set-cookie, ...
	for key, values := range h.headers {
		if key == HeaderKeySetCookie {
			for _, value := range values {
				httpPacket = append(httpPacket, key...)
				httpPacket = append(httpPacket, ColonSpace...)
				httpPacket = append(httpPacket, value...)
				httpPacket = append(httpPacket, CRLF...)
			}
		} else {
			httpPacket = append(httpPacket, key...)
			httpPacket = append(httpPacket, ColonSpace...)
			for _, value := range values {
				httpPacket = append(httpPacket, value...)
				httpPacket = append(httpPacket, Coma)
			}
			httpPacket = httpPacket[:len(httpPacket)-1] // Remove trailing comma
			httpPacket = append(httpPacket, CRLF...)
		}
	}
	return httpPacket
}

// UnMarshal parses and decodes data of given httpPacket(without first line) to (h *header).
// This method not respect to some RFCs like field-name in RFC7230, ... due to be more liberal in what it accept!
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (h *header) UnMarshal(s string) (headerEnd int) {
	var colonIndex, newLine int
	var key, value string
	for {
		newLine = strings.IndexByte(s, '\r')
		if newLine < 3 {
			// newLine == -1 >> broken or malformed packet, panic may occur!
			// newLine == 0 >> End of headers part of packet, no panic
			// 1 < newLine > 3 >> bad header || broken || malformed packet, panic may occur!
			return headerEnd
		}

		colonIndex = strings.IndexByte(s[:newLine], ':')
		if colonIndex == -1 {
			// Header key without value!?? Bad http packet!??
			newLine += 2 // +2 due to have "\r\n" at end of each *header line
			s = s[newLine:]
			headerEnd += newLine
			continue
		}

		key = s[:colonIndex]
		value = s[colonIndex+2 : newLine] // +2 due to have a space after colon force by RFC &&
		h.Add(key, value)                 // TODO::: is legal to have multiple key in request header or use h.Set()??

		newLine += 2 // +2 due to have "\r\n" at end of each *header line
		s = s[newLine:]
		headerEnd += newLine
	}
}

// Len returns length of encoded header!
func (h *header) Len() (ln int) {
	for key, values := range h.headers {
		ln += len(key)
		ln += 4 // 4=len(ColonSpace)+len(CRLF)
		for _, value := range values {
			ln += len(value)
			ln++ // 1=len(Coma)
		}
	}
	return
}
