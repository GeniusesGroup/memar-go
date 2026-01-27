/* For license and copyright information please see LEGAL file in repository */

package pool

import (
	"testing"
)

/*
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i7-2670QM CPU @ 2.20GHz

poolLen:10000
BenchmarkRealloc-8   	   55225	     23272 ns/op	   81922 B/op	       1 allocs/op
BenchmarkEmpty-8     	   31210	     51066 ns/op	      10 B/op	       0 allocs/op

poolLen:100000
BenchmarkRealloc-8   	    5335	    212914 ns/op	  802969 B/op	       1 allocs/op
BenchmarkEmpty-8     	    2919	    412944 ns/op	    1100 B/op	       0 allocs/op

poolLen:300000
BenchmarkRealloc-8   	    2193	    590969 ns/op	 2401359 B/op	       1 allocs/op
BenchmarkEmpty-8     	     508	   2448291 ns/op	   18899 B/op	       0 allocs/op
*/

const poolLen = 300000

func BenchmarkRealloc(b *testing.B) {
	var pool1 = make([]interface{}, poolLen/2)
	for n := 0; n < b.N; n++ {
		var pool2 = pool1
		pool1 = make([]interface{}, poolLen/2)
		_ = pool1[0]
		_ = pool2[0]
	}
}

func BenchmarkEmpty(b *testing.B) {
	var pool1 = make([]interface{}, poolLen)
	var pool2 = make([]interface{}, poolLen)
	for n := 0; n < b.N; n++ {
		copy(pool2, pool1)
		for i := 0; i < poolLen/2; i++ {
			pool2[i] = nil
		}
		_ = pool1[0]
		_ = pool2[0]
		for i := 0; i < poolLen/2; i++ {
			pool1[i] = nil
		}
	}
}
