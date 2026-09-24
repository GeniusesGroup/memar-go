/* For license and copyright information please see the LEGAL file in the code repository */

package ascii

import (
	buffer_p "memar/computer/buffer/protocol"
	error_p "memar/process/error/protocol"
)

type STR struct {
	buf buffer_p.Buffer
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (s *STR) Init(buf buffer_p.Buffer) (err error_p.Error) {
	s.buf = buf
	return
}
func (s *STR) Reinit(buf buffer_p.Buffer) (err error_p.Error) {
	s.buf = buf
	return
}
func (s *STR) Deinit(buf buffer_p.Buffer) (err error_p.Error) {
	// s.buf = nil
	return
}

func (s *STR) Buffer() buffer_p.Buffer { return s.buf }

func (s *STR) Append()  {}
func (s *STR) Prepend() {}
func (s STR) Equal(to STR) {
	// if s.buf.Equal(to) {
	// err =
	// return
	// }
}

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func (s *STR) Cut(sep string) (before, after string, found bool) {
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
func (s *STR) CutByte(sep byte) (before, after string, found bool) {
	// var i = strings.IndexByte(s, sep)
	// if i >= 0 {
	// 	return s[:i], s[i+1:], true
	// }
	// return s, "", false
	return
}
