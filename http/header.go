/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"net/textproto"
)

// Header is represent HTTP header structure!
type Header map[string][]string

// Get gets the value associated with the given key.
// Key must already be in CanonicalHeaderKey form.
func (h Header) Get(key string) []string {
	if v := h[key]; len(v) > 0 {
		return v
	}
	return nil
}

// Add adds the key, value pair to the header.
// It appends to any existing values associated with key.
// Key must already be in CanonicalHeaderKey form.
func (h Header) Add(key, value string) {
	h[key] = append(h[key], value)
}

// Set sets the header entries associated with key to
// the single element value. It replaces any existing
// values associated with key.
// Key must already be in CanonicalHeaderKey form.
func (h Header) Set(key string, value []string) {
	h[key] = value
}

// Del deletes the values associated with key.
// Key must already be in CanonicalHeaderKey form.
func (h Header) Del(key string) {
	delete(h, key)
}

// CanonicalHeaderKey returns the canonical format of the
// header key s. The canonicalization converts the first
// letter and any letter following a hyphen to upper case;
// the rest are converted to lowercase. For example, the
// canonical key for "accept-encoding" is "Accept-Encoding".
// If s contains a space or invalid header field bytes, it is
// returned without modifications.
func CanonicalHeaderKey(s string) string { return textproto.CanonicalMIMEHeaderKey(s) }

/*
--- Some useful logic http header code ---
*/

// get the http cookie from request header
// var cookie *http.Cookie
// cookie.ParseCookie(Header.Get(http.HeaderKeyCookie)[0])

// get the http set-cookie from request header
// var cookie *http.Cookie
// cookie.ParseCookie(Header.Get(http.HeaderKeySetCookie)[0])

// set desire cookie in response header
// Header.Add(http.HeaderKeySetCookie, cookie.String())
