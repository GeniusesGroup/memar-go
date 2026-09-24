/* For license and copyright information please see the LEGAL file in the code repository */

package http_headers

import (
	"strings"

	container_p "memar/computer/adt/container/protocol"
	"memar/computer/buffer/byteslice/convert"
	http_p "memar/libgo/net/http/protocol"
	error_p "memar/process/error/protocol"
)

func AddCookies(h http_p.Header, cookies string) {
	h.Header_Add(Key_Cookie, cookies)
}

// AddCookie will marshal and add given cookies to the Cookie header.
func AddCookie(h http_p.Header, cookies ...Cookie) {
	// TODO::: make buffer by needed size.
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
	h.Header_Add(Key_Cookie, b.String())
}

// Cookies store all the Cookie headers.
// Cookies parses and returns the Cookie headers.
// By related RFC we just support one Cookie in header.
// https://tools.ietf.org/html/rfc6265#section-5.4
type Cookies string

func (self *Cookies) Init(h http_p.Header) {
	var cookies = h.Header_Get(Key_Cookie)
	self = Cookies(cookies)
}

func (c *Cookies) Next() (cookie Cookie, exist bool) {
	var cs = *c

	if len(cs) == 0 {
		exist = false
		return
	}

	var dl, _ = cookie.FromString(string(cs))
	var remain = cs[dl:]
	if len(remain) > 0 {
		exist = true
		*c = Cookies(remain)
	}
	return
}

// Cookie represents an HTTP cookie as sent in the Cookie header of an HTTP request.
// implement by https://tools.ietf.org/html/rfc6265#section-4.2
type Cookie struct {
	Name  string
	Value string
}

// CheckAndSanitize check if the cookie is in standard by RFC and try to fix them. It returns last error!
func (c *Cookie) CheckAndSanitize() (err error_p.Error) {
	c.Name, err = sanitizeCookieName(c.Name)
	if err != nil {
		return
	}
	c.Value, err = sanitizeCookieValue(c.Value)
	return
}

// Marshal returns the serialization of the cookie.
func (c *Cookie) Marshal() string {
	return c.Name + "=" + c.Value
}

// Unmarshal parse given cookie value to c and return
func (c *Cookie) Unmarshal(cookie string) {
	var equalIndex = strings.IndexByte(cookie, '=')
	// First check no equal(=) sign or empty name or value
	if equalIndex < 1 || equalIndex == len(cookie)-1 {
		return
	}
	c.Name = cookie[:equalIndex]
	c.Value = cookie[equalIndex+1:]
}

// FromString parse given cookies value to c and return
func (c *Cookie) FromString(cookies string) (dl container_p.NumberOfElement, err error_p.Error) {
	var equalIndex = strings.IndexByte(cookies, '=')
	// First check no equal(=) sign or empty name or value
	if equalIndex < 1 || equalIndex == len(cookies)-1 {
		// err =
		return
	}
	c.Name = cookies[:equalIndex]

	var SemiColonIndex = strings.IndexByte(cookies, ';')
	if SemiColonIndex == -1 {
		c.Value = cookies[equalIndex+1:]
		dl = container_p.NumberOfElement(len(cookies) - 1)
	} else {
		c.Value = cookies[equalIndex+1 : SemiColonIndex]
		dl = container_p.NumberOfElement(SemiColonIndex + 2) // Due to have space after semi colon do +2
	}
	return
}

func sanitizeCookieName(n string) (name string, err error_p.Error) {
	var ln = len(n)
	var buf = make([]byte, 0, ln)
	var b byte
	for i := 0; i < ln; i++ {
		b = n[i]
		if b == '\n' || b == '\r' {
			buf = append(buf, '-')
			// err = &errs.ErrCookieBadName
		} else {
			buf = append(buf, b)
		}
	}
	name = convert.UnsafeByteSliceToString(buf)
	return
}

// https://tools.ietf.org/html/rfc6265#section-4.1.1
// cookie-value      = *cookie-octet / ( DQUOTE *cookie-octet DQUOTE )
// cookie-octet      = %x21 / %x23-2B / %x2D-3A / %x3C-5B / %x5D-7E
//
//	; US-ASCII characters excluding CTLs,
//	; whitespace, DQUOTE, comma, semicolon,
//	; and backslash
//
// Don't check for ; due to Unmarshal will panic for bad cookie!!
func sanitizeCookieValue(v string) (value string, err error_p.Error) {
	var ln = len(v)
	var buf = make([]byte, 0, ln)
	var b byte
	for i := 0; i < ln; i++ {
		b = v[i]
		if 0x20 <= b && b < 0x7f && b != ' ' && b != '"' && b != ',' && b != '\\' {
			buf = append(buf, b)
		} else {
			// err = &errs.ErrCookieBadValue
		}
	}
	value = convert.UnsafeByteSliceToString(buf)
	return
}
