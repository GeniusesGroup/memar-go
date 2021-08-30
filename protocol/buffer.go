/* For license and copyright information please see LEGAL file in repository */

package protocol

import "io"

type Buffer interface {
	Reset()

	Read(limit uint) []byte
	ReadAt(offset, limit uint) []byte
	Get() []byte
	GetUnread() []byte

	Write(s []byte)
	WriteAt(s []byte, offset uint)
	WriteString(s string)
	WriteStringAt(s string, offset uint)
	WriteByte(b byte) Error
	WriteByteAt(b byte, offset uint)
	Set(data []byte)

	Cap() int
	Len() int
	LenOfUnread() int

	Grow(n int)
	GrowLen(n int) (lenBeforeGrow int)

	// Split buffer to x part
	Split(size int)
	Parts() []Codec
	Part(id uint) Codec

	io.WriterTo
	io.ReaderFrom
}
