// For license and copyright information please see the LEGAL file in the code repository

//go:build 386 || amd64 || amd64p32 || arm || arm64 || loong64 || mips64le || mips64p32le || mipsle || ppc64le || riscv || riscv64 || wasm

package binary

import (
	"unsafe"
)

// LittleEndian is the little-endian implementation to get|set from le binary to le cpu
type LittleEndian []byte

func (b LittleEndian) Uint16() uint16 {
	_ = b[1] // bounds check hint to compiler; see go.dev/issue/14808
	return uint16(b[0]) | uint16(b[1])<<8
}
func (b LittleEndian) Uint32() uint32 {
	_ = b[3] // bounds check hint to compiler; see go.dev/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}
func (b LittleEndian) Uint64() uint64 {
	_ = b[7] // bounds check hint to compiler; see go.dev/issue/14808
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

func (b LittleEndian) PutUint16(v uint16) {
	_ = b[1] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
}
func (b LittleEndian) PutUint32(v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}
func (b LittleEndian) PutUint64(v uint64) {
	_ = b[7] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
}

// *data = math.Float32frombits(order.Uint32(bs))
// *data = math.Float64frombits(order.Uint64(bs))

func init() {
	i := uint32(1)
	b := (*[4]byte)(unsafe.Pointer(&i))
	if b[0] != 1 {
		panic("Expect LittleEndian CPU but have BigEndian CPU that cause many problem in other packages")
	}
}
