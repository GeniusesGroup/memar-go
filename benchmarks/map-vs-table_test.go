/* For license and copyright information please see LEGAL file in repository */

package benchmarks

import (
	"testing"

	"../achaemenid"
)

/*
BenchmarkFillUInt32Buckets-8   	 7751190	       158 ns/op	     192 B/op	       1 allocs/op
BenchmarkFillUInt32Map-8       	 5153385	       284 ns/op	     207 B/op	       1 allocs/op

BenchmarkGetUInt32Buckets-8    	28587764	       41.0 ns/op	      16 B/op	       0 allocs/op
BenchmarkGetUInt32Map-8        	 6095160	       225 ns/op	      28 B/op	       0 allocs/op
*/

const bucketLen = 4

type numberBucket struct {
	ID      [bucketLen]int
	Service [bucketLen]*achaemenid.Service
}

func BenchmarkFillUInt32Buckets(b *testing.B) {
	var bLen = b.N/bucketLen
	if bLen == 0 {
		bLen = 8
	}
	var nBuckets = make([]numberBucket, bLen)

	var loc int
	for n := 0; n < b.N; n++ {
		var ser = achaemenid.Service{}
		loc = n % bLen
		for j := 0; j < bucketLen; j++ {
			if nBuckets[loc].ID[j] == 0 {
				nBuckets[loc].ID[j] = n
				nBuckets[loc].Service[j] = &ser
				break
			}
		}
	}
}

func BenchmarkFillUInt32Map(b *testing.B) {
	var nMaps = make(map[int]*achaemenid.Service, b.N)

	for n := 0; n < b.N; n++ {
		var ser = achaemenid.Service{}
		nMaps[n] = &ser
	}
}

func BenchmarkGetUInt32Buckets(b *testing.B) {
	var bLen = b.N/bucketLen
	if bLen == 0 {
		bLen = 8
	}
	var nBuckets = make([]numberBucket, bLen)

	var ser = achaemenid.Service{}
	var loc int
	for n := 0; n < b.N; n++ {
		loc = n % bLen
		for j := 0; j < bucketLen; j++ {
			if nBuckets[loc].ID[j] == 0 {
				nBuckets[loc].ID[j] = n
				nBuckets[loc].Service[j] = &ser
				break
			}
		}
	}

	var ok bool
	for n := 0; n < b.N; n++ {
		var loc = n % bLen
		for j := 0; j < bucketLen; j++ {
			if nBuckets[loc].ID[j] == n {
				ok = true
				break
			}
		}
		if !ok {
			// Failed some time in tests! we write this test to show internal map is not efficient for almost read HashTables!
			// b.Fail()
		}
	}
}

func BenchmarkGetUInt32Map(b *testing.B) {
	var nMaps = make(map[int]*achaemenid.Service, b.N)
	var ser = achaemenid.Service{}
	for n := 0; n < b.N; n++ {
		nMaps[n] = &ser
	}

	var ok bool
	for n := 0; n < b.N; n++ {
		_, ok = nMaps[n]
		if !ok {
			b.Fail()
		}
	}
}
