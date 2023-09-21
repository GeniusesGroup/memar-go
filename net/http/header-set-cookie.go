/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"net"
	"strconv"
	"strings"
	"time"

	"memar/convert"
	errs "memar/net/http/errors"
	"memar/protocol"
)

// GetSetCookies parses and returns the Set-Cookie headers.
// By related RFC must exist just one Set-Cookie in each line of header.
// https://tools.ietf.org/html/rfc6265#section-4.1.1
func (h *header) SetCookies() (setCookies []SetCookie) {
	var scs = h.Gets(HeaderKeySetCookie)
	var setCookieCount = len(scs)
	if setCookieCount == 0 {
		return
	}
	setCookies = make([]SetCookie, setCookieCount)
	for i := 0; i < setCookieCount; i++ {
		setCookies[i].Unmarshal(scs[i])
	}
	return
}

// MarshalSetCookies parses and set given Set-Cookies to the header.
// By related RFC must exist just one Set-Cookie in each line of header.
// https://tools.ietf.org/html/rfc6265#section-4.1.1
func (h *header) MarshalSetCookies(setCookies []SetCookie) {
	var ln = len(setCookies)
	for i := 0; i < ln; i++ {
		h.Add(HeaderKeySetCookie, setCookies[i].Marshal())
	}
}

/*
SetCookie structure and methods implement by https://tools.ietf.org/html/rfc6265#section-4.1

MaxAge:
	MaxAge=0 means no 'Max-Age' attribute specified.
	MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	MaxAge>0 means Max-Age attribute present and given in seconds
*/

// A SetCookie represents an HTTP cookie as sent in the Set-Cookie header of an
// HTTP response or the Cookie header of an HTTP request.
type SetCookie struct {
	Name     string
	Value    string
	Path     string // optional
	Domain   string // optional
	Expires  string // optional
	MaxAge   string // optional
	Secure   bool   // optional
	HTTPOnly bool   // optional
	SameSite string // optional
}

// CheckAndSanitize use to check if the set-cookie is in standard by related RFCs.
func (sc *SetCookie) CheckAndSanitize() (err protocol.Error) {
	sc.Name, err = sanitizeCookieName(sc.Name)
	if err != nil {
		return
	}
	sc.Value, err = sanitizeCookieValue(sc.Value)
	if err != nil {
		return
	}
	sc.Path, err = sanitizeCookiePath(sc.Path)
	if err != nil {
		return
	}
	sc.Domain, err = sanitizeCookieDomain(sc.Domain)
	return
}

// GetExpire return the set-cookie expire in time.Time structure.
func (sc *SetCookie) GetExpire() (expTime time.Time) {
	var err error
	expTime, err = time.Parse(time.RFC1123, sc.Expires)
	if err != nil {
		expTime, err = time.Parse("Mon, 02-Jan-2006 15:04:05 MST", sc.Expires)
		if err != nil {
			expTime = time.Time{}
		}
	}
	return
}

// SetExpire use to set expire time by time.Time instead of raw string!
// IETF RFC 6265 Section 5.1.1.5, the year must not be less than 1601 but don't force or check here!
func (sc *SetCookie) SetExpire(expTime time.Time) {
	sc.Expires = expTime.UTC().Format(TimeFormat)
}

// GetMaxAge returns Max-Age value in Int instead of raw string!
func (sc *SetCookie) GetMaxAge() (maxAge int) {
	var err error
	maxAge, err = strconv.Atoi(sc.MaxAge)
	if err != nil {
		maxAge = 0
	}
	return
}

// SetMaxAge use to set Max-Age value by Int instead of raw string!
func (sc *SetCookie) SetMaxAge(maxAge int) {
	if maxAge > 0 {
		sc.MaxAge = strconv.FormatUint(uint64(maxAge), 10)
	} else if maxAge <= 0 {
		sc.MaxAge = "0"
	}
}

// SetCookieSameSite allows a server define a cookie attribute making it impossible to
// the browser send this cookie along with cross-site requests. The main goal
// is mitigate the risk of cross-origin information leakage, and provides some
// protection against cross-site request forgery attacks.
//
// See https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00 for details.
type SetCookieSameSite int

const (
	SetCookieSameSiteDefaultMode SetCookieSameSite = iota + 1
	// Cookies are allowed to be sent with top-level navigations and will be sent along with GET request
	// initiated by third party website. This is the default value in modern browsers.
	SetCookieSameSiteLaxMode
	// Cookies will only be sent in a first-party context and not be sent along with requests initiated by third party websites.
	SetCookieSameSiteStrictMode
	// Cookies will be sent in all contexts, i.e sending cross-origin is allowed.
	SetCookieSameSiteNoneMode
)

// GetSameSite returns Same-Site value in SetCookieSameSite type instead of raw string!
func (sc *SetCookie) GetSameSite() (sameSite SetCookieSameSite) {
	var lowerVal = strings.ToLower(sc.SameSite)
	switch lowerVal {
	case "lax":
		sameSite = SetCookieSameSiteLaxMode
	case "strict":
		sameSite = SetCookieSameSiteStrictMode
	case "none":
		sameSite = SetCookieSameSiteNoneMode
	default:
		sameSite = SetCookieSameSiteDefaultMode
	}
	return
}

// SetSameSite use to set Same-Site value by SetCookieSameSite type instead of raw string!
func (sc *SetCookie) SetSameSite(sameSite SetCookieSameSite) {
	switch sameSite {
	case SetCookieSameSiteDefaultMode:
		// TODO::: Why?? Really why standard mix boolean and value attribute!! we use space for boolean!
		sc.SameSite = " "
	case SetCookieSameSiteLaxMode:
		sc.SameSite = "Lax"
	case SetCookieSameSiteStrictMode:
		sc.SameSite = "Strict"
	case SetCookieSameSiteNoneMode:
		sc.SameSite = "none"
	}
}

// Marshal returns the serialization of the set-cookie.
func (sc *SetCookie) Marshal() string {
	// TODO::: make buffer by needed size.
	var b strings.Builder

	b.WriteString(sc.Name)
	b.WriteByte('=')
	b.WriteString(sc.Value)

	if len(sc.Path) > 0 {
		b.WriteString("; Path=")
		b.WriteString(sc.Path)
	}
	if len(sc.Domain) > 0 {
		b.WriteString("; Domain=")
		b.WriteString(sc.Domain)
	}
	if len(sc.Expires) > 0 {
		b.WriteString("; Expires=")
		b.WriteString(sc.Expires)
	}
	if len(sc.MaxAge) > 0 {
		b.WriteString("; Max-Age=")
		b.WriteString(sc.MaxAge)
	}
	if sc.HTTPOnly {
		b.WriteString("; HttpOnly")
	}
	if sc.Secure {
		b.WriteString("; Secure")
	}
	if len(sc.SameSite) > 0 {
		b.WriteString("; SameSite=")
		b.WriteString(sc.SameSite)
	}
	return b.String()
}

// Unmarshal parse given set-cookie value to sc and return.
// set-cookie value must be in standard or use CheckAndSanitize() if you desire after Unmarshaling!
// In some bad packet may occur panic, handle panic by recover otherwise app will crash and exit!
func (sc *SetCookie) Unmarshal(setCookie string) {
	var index = strings.IndexByte(setCookie, '=')
	// First check no equal(=) sign or empty name or value
	if index < 1 {
		return
	}
	sc.Name = setCookie[:index]
	setCookie = setCookie[index+1:]

	index = strings.IndexByte(setCookie, ';')
	if index == -1 {
		sc.Value = setCookie
		return
	}
	sc.Value = setCookie[:index]

	var attr, val string
	var nextSemiColonIndex int = index
	var end bool
	for !end {
		setCookie = setCookie[nextSemiColonIndex+2:] // +2 due to also have a space after semicolon

		nextSemiColonIndex = strings.IndexByte(setCookie, ';')
		if nextSemiColonIndex == -1 {
			nextSemiColonIndex = len(setCookie)
			end = true
		}

		index = strings.IndexByte(setCookie, '=')
		if index == -1 {
			// Boolean attribute
			attr = strings.ToLower(setCookie[:nextSemiColonIndex])
		} else {
			// Value attribute
			attr = strings.ToLower(setCookie[:index])
			val = setCookie[index+1 : nextSemiColonIndex]
		}

		switch attr {
		case "samesite":
			sc.SameSite = val
			continue
		case "secure":
			sc.Secure = true
			continue
		case "httponly":
			sc.HTTPOnly = true
			continue
		case "domain":
			sc.Domain = val
			continue
		case "max-age":
			sc.MaxAge = val
			continue
		case "expires":
			sc.Expires = val
			continue
		case "path":
			sc.Path = val
			continue
		}
	}
}

// path-av           = "Path=" path-value
// path-value        = <any CHAR except CTLs or ";">
// Don't check for ; due to Unmarshal will panic for bad cookie!!
func sanitizeCookiePath(v string) (path string, err protocol.Error) {
	var ln = len(v)
	var buf = make([]byte, 0, ln)
	var b byte
	for i := 0; i < ln; i++ {
		b = v[i]
		if 0x20 <= b && b < 0x7f {
			buf = append(buf, b)
		} else {
			err = &errs.ErrCookieBadPath
		}
	}
	path = convert.UnsafeByteSliceToString(buf)
	return
}

// A sc.Domain containing illegal characters is not sanitized but simply dropped which turns the cookie
// into a host-only cookie. A leading dot is okay but won't be sent.
func sanitizeCookieDomain(d string) (domain string, err protocol.Error) {
	var ln = len(d)
	if ln == 0 || ln > 255 {
		return domain, &errs.ErrCookieBadDomain
	}

	// A cookie a domain attribute may start with a leading dot.
	if d[0] == '.' {
		d = d[1:]
		ln--
	}

	var last byte = '.'
	var partlen int
	var b byte
	for i := 0; i < ln; i++ {
		b = d[i]
		switch {
		default:
			return domain, &errs.ErrCookieBadDomain
		case 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z':
			// No '_' allowed here (in contrast to package net).
			partlen++
		case '0' <= b && b <= '9':
			// fine
			partlen++
		case b == '-':
			// Byte before dash cannot be dot.
			if last == '.' {
				return domain, &errs.ErrCookieBadDomain
			}
			partlen++
		case b == '.':
			// Byte before dot cannot be dot or dash.
			if last == '.' || last == '-' {
				return domain, &errs.ErrCookieBadDomain
			}
			if partlen > 63 || partlen == 0 {
				return domain, &errs.ErrCookieBadDomain
			}
			partlen = 0
		}
		last = b
	}
	// TODO::: is end . legal??
	if last != '-' && last != '.' && partlen < 64 {
		return d, err
	}

	if net.ParseIP(d) != nil {
		return d, err
	}

	return domain, &errs.ErrCookieBadDomain
}
