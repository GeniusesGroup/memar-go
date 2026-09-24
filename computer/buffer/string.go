/* For license and copyright information please see the LEGAL file in the code repository */

package buffer

// reflect.StringHeader, unsafe.String or unsafe.StringData
type String string

func (s *String) Append()  {}
func (s *String) Prepend() {}
func (s String) Equal(to String) {
	if s != to {
		// err =
		return
	}
}

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func (s *String) Cut(sep string) (before, after string, found bool) {
	// var i = strings.Index(s, sep)
	// if i >= 0 {
	// 	return s[:i], s[i+len(sep):], true
	// }
	// return s, "", false
	return
}

// CutByte slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func (s *String) CutByte(sep byte) (before, after string, found bool) {
	// var i = strings.IndexByte(s, sep)
	// if i >= 0 {
	// 	return s[:i], s[i+1:], true
	// }
	// return s, "", false
	return
}

func (s *String) FromString_Go(str string) {
	*s = String(str)
}
