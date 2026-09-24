/* For license and copyright information please see the LEGAL file in the code repository */

package byteslice

type Safe struct {
	str []byte
	len int
}

func (self *Safe) Init(str []byte) {
	self.str = str
}

func (self *Safe) Append()  {}
func (self *Safe) Prepend() {}
func (self Safe) Equal(to Safe) {
	if self.len != to.len {
		// err =
		return
	}
}

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in self.
// If sep does not appear in s, cut returns s, "", false.
func (self *Safe) Cut(sep string) (before, after string, found bool) {
	// var i = strings.Index(s, sep)
	// if i >= 0 {
	// 	return s[:i], s[i+len(sep):], true
	// }
	// return s, "", false
	return
}

// CutByte slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in self.
// If sep does not appear in s, cut returns s, "", false.
func (self *Safe) CutByte(sep byte) (before, after string, found bool) {
	// var i = strings.IndexByte(s, sep)
	// if i >= 0 {
	// 	return s[:i], s[i+1:], true
	// }
	// return s, "", false
	return
}
