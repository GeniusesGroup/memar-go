/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"strconv"
	"testing"
)

/*
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i7-2670QM CPU @ 2.20GHz
Benchmark_StringToUint8Base10-8    	141969729	        8.625 ns/op	       0 B/op	       0 allocs/op
Benchmark_strconvParseUint8-8      	62640288	        18.21 ns/op	       0 B/op	       0 allocs/op
Benchmark_StringToUint32Base10-8   	78345345	        14.67 ns/op	       0 B/op	       0 allocs/op
Benchmark_strconvParseUint32-8      29286721	        40.92 ns/op	       0 B/op	       0 allocs/op
Benchmark_StringToUint64Base10-8   	31770237	        33.70 ns/op	       0 B/op	       0 allocs/op
Benchmark_strconvParseUint64-8   	15955197	        76.97 ns/op	       0 B/op	       0 allocs/op
*/

func Benchmark_StringToUint8Base10(b *testing.B) {
	var num = "255"
	for n := 0; n < b.N; n++ {
		StringToUint8Base10(num)
	}
}

func Benchmark_strconvParseUint8(b *testing.B) {
	var num = "255"
	for n := 0; n < b.N; n++ {
		strconv.ParseUint(num, 10, 8)
	}
}

func Benchmark_StringToUint32Base10(b *testing.B) {
	var num = "4294967295"
	for n := 0; n < b.N; n++ {
		StringToUint32Base10(num)
	}
}

func Benchmark_strconvParseUint32(b *testing.B) {
	var num = "4294967295"
	for n := 0; n < b.N; n++ {
		strconv.ParseUint(num, 10, 32)
	}
}

func Benchmark_StringToUint64Base10(b *testing.B) {
	var num = "18446744073709551615"
	for n := 0; n < b.N; n++ {
		StringToUint64Base10(num)
	}
}

func Benchmark_strconvParseUint64(b *testing.B) {
	var num = "18446744073709551615"
	for n := 0; n < b.N; n++ {
		strconv.ParseUint(num, 10, 64)
	}
}
