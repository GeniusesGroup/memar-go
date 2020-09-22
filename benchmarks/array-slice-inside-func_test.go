/* For license and copyright information please see LEGAL file in repository */

package benchmarks

import (
	"crypto/sha512"
	"testing"
)

/*
	Array vs Slice initialize inside a func

BenchmarkArrayInsideFunc-8   	 1830218	       654 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceInsideFunc-8   	 1829384	       658 ns/op	       0 B/op	       0 allocs/op
*/

func BenchmarkArrayInsideFunc(b *testing.B) {
	var id uint64 = 4150904594571984896
	for n := 0; n < b.N; n++ {
		var buf [8]byte // 8

		buf[0] = byte(id)
		buf[1] = byte(id >> 8)
		buf[2] = byte(id >> 16)
		buf[3] = byte(id >> 24)
		buf[4] = byte(id >> 32)
		buf[5] = byte(id >> 40)
		buf[6] = byte(id >> 48)
		buf[7] = byte(id >> 56)

		sha512.Sum512_256(buf[:])
	}
}

func BenchmarkSliceInsideFunc(b *testing.B) {
	var id uint64 = 4150904594571984896
	for n := 0; n < b.N; n++ {
		var buf = make([]byte, 8)

		buf[0] = byte(id)
		buf[1] = byte(id >> 8)
		buf[2] = byte(id >> 16)
		buf[3] = byte(id >> 24)
		buf[4] = byte(id >> 32)
		buf[5] = byte(id >> 40)
		buf[6] = byte(id >> 48)
		buf[7] = byte(id >> 56)

		sha512.Sum512_256(buf)
	}
}
