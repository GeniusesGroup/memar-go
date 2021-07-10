/* For license and copyright information please see LEGAL file in repository */

package giti

import "io"

type Buffer interface {
	// Read(p []byte) (n int, err Error)  // same as io.Reader
	// Write(p []byte) (n int, err Error) // same as io.Writer
	
	Reset()

	Read(offset, limit uint) []byte
	WriteSlice(s []byte, offset, limit uint)
	WriteString(s string, offset, limit uint)
	WriteByte(b byte, offset, limit uint)
	Get() []byte
	GetUnread() []byte
	Set([]byte)
	AppendSlice(s []byte)
	AppendString(s string)
	AppendByte(b byte)

	Cap() int
	Len() int
	LenOfUnread() int

	Grow(n int)
	GrowLen(n int) (lenBeforeGrow int)

	// Split buffer to x part
	Split(size int)
	Parts() [][]byte
	Part(id int) []byte

	io.WriterTo
	io.ReaderFrom
}
