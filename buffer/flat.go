/* For license and copyright information please see the LEGAL file in the code repository */

package buffer

import (
	"fmt"

	"memar/protocol"
)

// smallBufferSize is an initial allocation minimal capacity.
const smallBufferSize = 64

type Flat struct {
	buf []byte
	off int // read at &buf[off], write at &buf[len(buf)]
}

func (b *Flat) Init(buf []byte) {
	b.buf = buf
}

func (b *Flat) Reset() {
	b.buf = b.buf[:0]
	b.off = 0
}

func (b *Flat) Bytes() []byte  { return b.buf }
func (b *Flat) String() string { return string(b.buf) }

func (b *Flat) Read(p []byte) (n int, err error) {
	var unReadLen = len(b.buf) - b.off
	if n < 0 || n > unReadLen {
		n = unReadLen
	}
	copy(p, b.buf[b.off:b.off+n])
	b.off += n
	return
}

func (b *Flat) ReadByte() (bt byte, err error) {
	var unReadLen = len(b.buf) - b.off
	if unReadLen > 1 {
		bt = b.buf[b.off]
		b.off++
	}
	return
}

func (b *Flat) ReadAt(offset, limit uint) []byte {
	b.off = len(b.buf)
	return b.buf[b.off:]
}

func (b *Flat) Set(p []byte) { b.buf = p }

func (b *Flat) Write(p []byte) (n int, err error) {
	b.buf = append(b.buf, p...)
	n = len(p)
	return
}

func (b *Flat) WriteString(s string) (n int, err error) {
	b.buf = append(b.buf, s...)
	n = len(s)
	return
}
func (b *Flat) WriteByte(bt byte) (err error) { b.buf = append(b.buf, bt); return }

func (b *Flat) Cap() int         { return cap(b.buf) }
func (b *Flat) Len() int         { return len(b.buf) }
func (b *Flat) LenOfUnread() int { return len(b.buf) - b.off }

func (b *Flat) Grow(n int) {
	if b.buf == nil {
		if n <= smallBufferSize {
			n = smallBufferSize
		}
		b.buf = make([]byte, 0, n)
		return
	}

	var buffLen = len(b.buf)
	var buffCap = cap(b.buf)
	if n <= buffCap-buffLen {
		return
	}

	var buf = makeSlice(2*buffCap + n)
	copy(buf, b.buf)
	b.buf = buf[:buffLen]
}

func (b *Flat) GrowLen(n int) int {
	if b.buf == nil {
		if n <= smallBufferSize {
			n = smallBufferSize
		}
		b.buf = make([]byte, 0, n)
		return 0
	}

	var buffLen = len(b.buf)
	var buffCap = cap(b.buf)
	if n <= buffCap-buffLen {
		b.buf = b.buf[:buffLen+n]
		return buffLen
	}

	var buf = makeSlice(2*buffCap + n)
	copy(buf, b.buf)
	b.buf = buf[:buffLen+n]
	return buffLen
}

func (b *Flat) WriteTo(w protocol.Writer) (n int, err protocol.Error) {
	// io.Copy(w, b)
	return
}
func (b *Flat) ReadFrom(r protocol.Reader) (n int, err protocol.Error) {
	// io.Copy(b, r)
	return
}
func (b *Flat) Close() (err error) { return }

// makeSlice allocates a slice of size n. If the allocation fails, it panics
// with ErrTooLarge.
func makeSlice(n int) []byte {
	// If the make fails, give a known error.
	defer func() {
		var r = recover()
		if r != nil {
			panic(fmt.Sprintln("buffer - makeSlice: ", r))
		}
	}()
	return make([]byte, n)
}
