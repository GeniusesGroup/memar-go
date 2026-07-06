/* For license and copyright information please see LEGAL file in repository */

package syllab

import (
	"testing"
)

/*
go tool compile -S encode_test.go > encode_test_C.S
go tool objdump encode_test.o > encode_test_O.S
*/

var sliceLen = 4

func BenchmarkFunctionInlineByCompiler(b *testing.B) {
	var t = test{
		Num: 1587464654,
	}
	var buf = make([]byte, sliceLen)
	for n := 0; n < b.N; n++ {
		t.syllabEncoderCompiler(buf)
	}
}

func BenchmarkFunctionInlineByDev(b *testing.B) {
	var t = test{
		Num: 1587464654,
	}
	var buf = make([]byte, sliceLen)
	for n := 0; n < b.N; n++ {
		t.syllabEncoderDev(buf)
	}
}

func BenchmarkFunctionInlineByCompilerReturn(b *testing.B) {
	var t = test{
		Num: 1587464654,
	}
	for n := 0; n < b.N; n++ {
		t.syllabEncoderCompilerReturn()
	}
}

func BenchmarkFunctionInlineByDevReturn(b *testing.B) {
	var t = test{
		Num: 1587464654,
	}
	for n := 0; n < b.N; n++ {
		t.syllabEncoderReturn()
	}
}

type test struct {
	Num uint32
}

func (t *test) syllabEncoderCompiler(buf []byte) {
	setUInt32(buf, t.Num)
}

func setUInt32(p []byte, n uint32) {
	p[0] = byte(n)
	p[1] = byte(n >> 8)
	p[2] = byte(n >> 16)
	p[3] = byte(n >> 24)
}

func (t *test) syllabEncoderDev(buf []byte) {
	buf[0] = byte(t.Num)
	buf[1] = byte(t.Num >> 8)
	buf[2] = byte(t.Num >> 16)
	buf[3] = byte(t.Num >> 24)
}

func (t *test) syllabEncoderCompilerReturn() (buf []byte) {
	buf = make([]byte, sliceLen)
	setUInt32(buf, t.Num)
	return
}

func (t *test) syllabEncoderReturn() (buf []byte) {
	buf = make([]byte, sliceLen)

	buf[0] = byte(t.Num)
	buf[1] = byte(t.Num >> 8)
	buf[2] = byte(t.Num >> 16)
	buf[3] = byte(t.Num >> 24)
	return
}
