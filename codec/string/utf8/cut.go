/* For license and copyright information please see the LEGAL file in the code repository */

package utf8

import (
	"strings"
)

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func Cut(s, sep string) (before, after string, found bool) {
	var i = strings.Index(s, sep)
	if i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

// CutByte slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func CutByte(s string, sep byte) (before, after string, found bool) {
	var i = strings.IndexByte(s, sep)
	if i >= 0 {
		return s[:i], s[i+1:], true
	}
	return s, "", false
}
