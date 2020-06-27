/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"net/textproto"
	"strings"
)

// header is represent HTTP header structure!
type header map[string][]string

// GetValues gets the values associated with the given key.
// Key must already be in CanonicalHeaderKey form.
func (h header) GetValues(key string) []string {
	if v := h[key]; len(v) > 0 {
		return v
	}
	return nil
}

// GetValue gets the first value associated with the given key.
// Key must already be in CanonicalHeaderKey form.
func (h header) GetValue(key string) string {
	if v := h[key]; len(v) > 0 {
		return v[0]
	}
	return ""
}

// Add adds the key, value pair to the header.
// It appends to any existing values associated with key.
// Key must already be in CanonicalHeaderKey form.
func (h header) Add(key, value string) {
	h[key] = append(h[key], value)
}

// SetValues sets the header entries associated with key to
// the single element value. It replaces any existing values associated with key.
// Key must already be in CanonicalHeaderKey form.
func (h header) SetValues(key string, values []string) {
	h[key] = values
}

func (h header) SetValue(key string, value string) {
	h[key] = []string{value}
}

// Del deletes the values associated with key.
// Key must already be in CanonicalHeaderKey form.
func (h header) Del(key string) {
	delete(h, key)
}

// GetCookies parses and returns the Cookie headers.
// By related RFC we just support one Cookie in header.
// https://tools.ietf.org/html/rfc6265#section-5.4
func (h header) GetCookies() (cookies []*Cookie) {
	var cookie = h.GetValue(HeaderKeyCookie)
	if len(cookie) == 0 {
		return
	}
	var index int
	cookies = make([]*Cookie, 0, 8)
	for {
		index = strings.IndexByte(cookie, ';')
		if index == -1 {
			return
		}
		var c Cookie
		c.UnMarshal(cookie[:index])
		cookies = append(cookies, &c)

		cookie = cookie[index+2:]
	}
}

// SetCookies parses and set them to Cookie header.
func (h header) SetCookies(cookies []*Cookie) {
	var b strings.Builder
	h.SetValue(HeaderKeyCookie, b.String())

	var ln = len(cookies)
	var cookie *Cookie
	var i int
	for ; ; i++ {
		cookie = cookies[i]
		b.WriteString(cookie.Name)
		b.WriteByte('=')
		b.WriteString(cookie.Value)
		if i < ln {
			b.WriteString(SemiColonSpace)
		} else {
			return
		}
	}
}

// GetSetCookies parses and returns the Set-Cookie headers.
// By related RFC must exist just one Set-Cookie in each line of header.
// https://tools.ietf.org/html/rfc6265#section-4.1.1
func (h header) GetSetCookies() (setCookies []*SetCookie) {
	var setCookie = h.GetValues(HeaderKeySetCookie)
	var setCookieCount = len(setCookie)
	if setCookieCount == 0 {
		return
	}
	setCookies = make([]*SetCookie, 0, setCookieCount)
	for i := 0; i < setCookieCount; i++ {
		var sc SetCookie
		sc.UnMarshal(setCookie[i])
		setCookies = append(setCookies, &sc)
	}
	return
}

// SetSetCookies parses and set given Set-Cookies to header.
// By related RFC must exist just one Set-Cookie in each line of header.
// https://tools.ietf.org/html/rfc6265#section-4.1.1
func (h header) SetSetCookies(setCookies []*SetCookie) {
	var ln = len(setCookies)
	for i := 0; i < ln; i++ {
		h.Add(HeaderKeySetCookie, setCookies[i].Marshal())
	}
}

// FixPragmaCacheControl do as RFC 7234, section 5.4: Treat [Pragma: no-cache] as [Cache-Control: no-cache]
func (h header) FixPragmaCacheControl() {
	if h.GetValue(HeaderKeyPragma) == "no-cache" {
		if h.GetValues(HeaderKeyCacheControl) == nil {
			h.SetValue(HeaderKeyCacheControl, "no-cache")
		}
	}
}

// Exclude eliminate headers by given keys!
func (h header) Exclude(exclude map[string]bool) {
	for key := range exclude {
		delete(h, key)
	}
}

// Marshal enecodes (h header) data to given httpPacket.
func (h header) Marshal(httpPacket *[]byte) {
	// TODO::: some header key must not inline by coma like set-cookie, ...
	for key, values := range h {
		*httpPacket = append(*httpPacket, key...)
		*httpPacket = append(*httpPacket, ColonSpace...)
		*httpPacket = append(*httpPacket, strings.Join(values, ",")...)
		*httpPacket = append(*httpPacket, CRLF...)
	}
}

// UnMarshal parses and decodes data of given httpPacket(without first line) to (h header).
// This method not respect to some RFCs like field-name in RFC7230, ... due to be more liberal in what it accept!
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (h header) UnMarshal(s string) (bodyStart int) {
	var valuesPool = make([]string, 16)
	var index, newLine int
	var key, value string
	var values []string
	for {
		newLine = strings.IndexByte(s, '\n')
		if newLine == -1 {
			// By https://tools.ietf.org/html/rfc2616#section-4 very simple http packet must end with CRLF even packet without header or body!
			// So this situation is not legal and means http packet is broken that don't even have new line. panic may occur!
			return bodyStart
		} else if newLine == 1 {
			// End of headers part of packet!
			return bodyStart + 2
		}

		index = strings.IndexByte(s, ':')
		if index == -1 || index > newLine {
			// Bad http packet! panic may occur!
			return bodyStart + 2 // Due to have one other "\r\n" after header end
		}
		key = s[:index]
		value = s[index+2 : newLine-1] // +2 due to have a space after colon force by RFC && -1 due to have "\r\n"
		newLine++
		s = s[newLine:]
		bodyStart += newLine

		values = h[key]
		if values == nil && len(valuesPool) > 0 {
			// More than likely this will be a single-element key. Most headers aren't multi-valued.
			// Set the capacity on strs[0] to 1, so any future append won't extend the slice into the other strings.
			values = valuesPool[:1:1]
			valuesPool = valuesPool[1:]
			values[0] = value
			h[key] = values
		} else {
			h[key] = append(values, value)
		}
	}
}

// CanonicalHeaderKey returns the canonical format of the header key s.
// The canonicalization converts the first letter and any letter following a hyphen to upper case;
// the rest are converted to lowercase. For example, the canonical key for "accept-encoding" is "Accept-Encoding".
// If s contains a space or invalid header field bytes, it is returned without modifications.
func CanonicalHeaderKey(s string) string { return textproto.CanonicalMIMEHeaderKey(s) }
