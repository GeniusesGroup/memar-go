/* For license and copyright information please see LEGAL file in repository */

package unix

import (
	"testing"
	"time"
	_ "unsafe" // for go:linkname
)

/*
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i7-2670QM CPU @ 2.20GHz
Benchmark_goUnixMilli-8   	100000000	        10.26 ns/op	       0 B/op	       0 allocs/op
Benchmark_NowMilli-8    	200743888	        5.936 ns/op	       0 B/op	       0 allocs/op
Benchmark_RuntimeNano-8   	278594536	        4.172 ns/op	       0 B/op	       0 allocs/op
*/

func Benchmark_goUnixMilli(b *testing.B) {
	for n := 0; n < b.N; n++ {
		time.Now().UnixMilli()
	}
}

func Benchmark_NowMilli(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Now().MilliElapsed()
	}
}

func Benchmark_RuntimeMonotonic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RuntimeNano()
	}
}

// RuntimeNano returns the current value of the runtime monotonic clock in nanoseconds.
// It isn't not wall clock, Use in tasks like timeout, ...
//go:linkname RuntimeNano runtime.nanotime
func RuntimeNano() int64
