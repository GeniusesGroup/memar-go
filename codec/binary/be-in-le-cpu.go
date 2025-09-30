// For license and copyright information please see the LEGAL file in the code repository

//go:build 386 || amd64 || amd64p32 || arm || arm64 || loong64 || mips64le || mips64p32le || mipsle || ppc64le || riscv || riscv64 || wasm

package binary

import "unsafe"

// BigEndian is the big-endian implementation to get|set from be binary to le cpu
type BigEndian []byte

func (b BigEndian) Uint16() uint16 {
	_ = b[1] // bounds check hint to compiler; see go.dev/issue/14808
	return uint16(b[1]) | uint16(b[0])<<8
}
func (b BigEndian) Uint32() uint32 {
	_ = b[3] // bounds check hint to compiler; see go.dev/issue/14808
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func (b BigEndian) Uint64() uint64 {
	_ = b[7] // bounds check hint to compiler; see go.dev/issue/14808
	return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
}

func (b BigEndian) PutUint16(v uint16) {
	_ = b[1] // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 8)
	b[1] = byte(v)
}
func (b BigEndian) PutUint32(v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
}
func (b BigEndian) PutUint64(v uint64) {
	_ = b[7] // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
}

func init() {
	i := uint32(1)
	b := (*[4]byte)(unsafe.Pointer(&i))
	if b[0] != 1 {
		panic("Expect LittleEndian CPU but have BigEndian CPU that cause many problem in other packages")
	}
}
