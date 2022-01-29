/* For license and copyright information please see LEGAL file in repository */

package time

import (
	"testing"
	"time"
)

/*
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i7-2670QM CPU @ 2.20GHz
Benchmark_goUnixMilli-8   		100000000	        10.26 ns/op	       0 B/op	       0 allocs/op
Benchmark_UnixNowMilli-8    	200743888	        5.936 ns/op	       0 B/op	       0 allocs/op
Benchmark_RuntimeMonotonic-8   	278594536	        4.172 ns/op	       0 B/op	       0 allocs/op
*/

func Benchmark_goUnixMilli(b *testing.B) {
	for n := 0; n < b.N; n++ {
		time.Now().UnixMilli()
	}
}

func Benchmark_UnixNowMilli(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UnixNowMilli()
	}
}

func Benchmark_RuntimeMonotonic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RuntimeMonotonic()
	}
}
