/* For license and copyright information please see LEGAL file in repository */

package benchmarks

import (
	"testing"
)

/*
	Assignment vs Copy performance

BenchmarkAssign-8   	  667114	      1696 ns/op	       0 B/op	       0 allocs/op
BenchmarkCopy-8     	 6520225	       188 ns/op	       0 B/op	       0 allocs/op
*/

const caSliceLen = 2400

var caTestBuf []byte

func init() {
	caTestBuf = make([]byte, caSliceLen)
	var j uint8
	for i := 0; i < caSliceLen; i++ {
		caTestBuf[i] = j
		j++
	}
}

func BenchmarkAssign(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t = make([]byte, caSliceLen)
		for i := 0; i < caSliceLen; i++ {
			t[i] = caTestBuf[i]
		}
	}
}

func BenchmarkCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t = make([]byte, caSliceLen)
		copy(t, caTestBuf)
	}
}
