/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/valyala/fasthttp"
)

/*
Benchmark1NetHttpRequestDecode-8    	  155938	      7578 ns/op	    5865 B/op	      22 allocs/op
Benchmark1LibgoRequestDecode-8      	  532573	      2189 ns/op	    1712 B/op	       3 allocs/op
Benchmark1FastHTTPRequestDecode-8   	  184724	      6464 ns/op	    6289 B/op	      32 allocs/op

Benchmark1NetHttpRequestEncode-8    	  236688	      5159 ns/op	    2069 B/op	      15 allocs/op
Benchmark1LibgoRequestEncode-8      	 1202239	       968 ns/op	     336 B/op	       1 allocs/op
Benchmark1FastHTTPRequestEncode-8   	  631375	      1719 ns/op	    4208 B/op	       3 allocs/op

note1: This benchmark is not apple to apple because libgo force rules in methods not in encode||decode phase!
note2: Libgo decoder performance better by: 3.5X from standard GO, 3X from FastHTTP
note3: Libgo encoder performance better by: 5.3X from standard GO, 1.8X from FastHTTP
*/

/*
	Test data
*/

var marshaledRequestTest1 = []byte("POST /apis?2586547852 HTTP/1.1\r\n" +
	"Host: www.sabz.city\r\n" +
	"Connection: keep-alive\r\n" +
	"Set-Cookie: test\r\n" +
	"Cache-Control: max-age=0\r\n" +
	"Upgrade-Insecure-Requests: 1\r\n" +
	"User-Agent: Mozilla\r\n" +
	"Accept: text/html\r\n" +
	"Accept-Encoding: gzip, deflate\r\n" +
	"Accept-Language: en,fa;q=0.9\r\n" +
	"Content-Type: application/json\r\n" +
	"Content-Length: 15\r\n" +
	"\r\n" +
	`{"Omid":"OMID"}`)
var unMarshaledRequestNetHTTPTest1 *http.Request
var unMarshaledRequestLibgoTest1 *Request
var unMarshaledRequestFastHTTPTest1 fasthttp.Request

func init() {
	var err error

	var r = bytes.NewBuffer(marshaledRequestTest1)
	var buf = bufio.NewReader(r)
	unMarshaledRequestNetHTTPTest1, err = http.ReadRequest(buf)
	if err != nil {
		fmt.Print(err, "\n")
	}

	unMarshaledRequestLibgoTest1 = MakeNewRequest()
	err = unMarshaledRequestLibgoTest1.UnMarshal(marshaledRequestTest1)
	if err != nil {
		fmt.Print(err, "\n")
	}

	var br = bytes.NewBuffer(marshaledRequestTest1)
	var fastBuf = bufio.NewReader(br)
	err = unMarshaledRequestFastHTTPTest1.Read(fastBuf)
	if err != nil {
		fmt.Print(err, "\n")
	}
}

/*
	Decode && Encode Benchmarks
*/

func Benchmark1NetHttpRequestDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var r = bytes.NewBuffer(marshaledRequestTest1)
		var buf = bufio.NewReader(r)
		http.ReadRequest(buf)
	}
}

func Benchmark1LibgoRequestDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var r = MakeNewRequest()
		r.UnMarshal(marshaledRequestTest1)
	}
}

func Benchmark1FastHTTPRequestDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var r = bytes.NewBuffer(marshaledRequestTest1)
		var buf = bufio.NewReader(r)
		var req fasthttp.Request
		req.Read(buf)
	}
}

func Benchmark1NetHttpRequestEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var b []byte
		var r = bytes.NewBuffer(b)
		unMarshaledRequestNetHTTPTest1.Write(r)
	}
}

func Benchmark1LibgoRequestEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unMarshaledRequestLibgoTest1.Marshal()
	}
}

func Benchmark1FastHTTPRequestEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		var bufW = bufio.NewWriter(&buf)
		unMarshaledRequestFastHTTPTest1.Write(bufW)
	}
}

/*
	Encode Tests
*/

func Test1LibgoRequestDecode(t *testing.T) {
	var req = MakeNewRequest()
	var err = req.UnMarshal(marshaledRequestTest1)
	if err != nil {
		fmt.Print(err, "\n")
		t.Fail()
	} else if string(req.Body) != `{"Omid":"OMID"}` {
		fmt.Print("Decoded Body not same\n")
		t.Fail()
	}
	fmt.Print(len(req.Header.valuesPool), "\n")
	fmt.Print(req, "\n")

	// fmt.Fprintf(os.Stderr, "%v\n", req.Method)
	// fmt.Fprintf(os.Stderr, "%v\n", req.URI)
	// fmt.Fprintf(os.Stderr, "%v\n", req.Version)
	// fmt.Fprintf(os.Stderr, "%v\n", req.Header)
	// marshaledRequestTest1[30] = '-'
	// fmt.Fprintf(os.Stderr, "%v\n", req.Header)

	// s := req.Header.GetSetCookies()[0]
	// err := s.CheckAndSanitize()
	// fmt.Fprintf(os.Stderr, "%v\n", err)
	// fmt.Fprintf(os.Stderr, "%v\n", s)

	// fmt.Fprintf(os.Stderr, "%v\n", string(req.Body))
}

func Test1LibgoRequestEncode(t *testing.T) {
	var httpPacket = unMarshaledRequestLibgoTest1.Marshal()
	if httpPacket == nil {
		fmt.Print("nil encode packet", "\n")
		t.Fail()
	} else if !bytes.Equal(httpPacket, marshaledRequestTest1) {
		fmt.Print("encoded not same with original\n")
		fmt.Print("cap--len of httpPacket:", cap(httpPacket), "--", len(httpPacket), "\n")
		fmt.Print("cap--len of httpPacket:", cap(marshaledRequestTest1), "--", len(marshaledRequestTest1), "\n")
		fmt.Print(string(httpPacket), "\n")
		t.Fail()
	}
}

func Test1FastHTTPRequestDecode(t *testing.T) {
	var br = bytes.NewBuffer(marshaledRequestTest1)
	var fastBuf = bufio.NewReader(br)
	var req fasthttp.Request
	var err = req.Read(fastBuf)
	if err != nil {
		fmt.Print(err, "\n")
		t.Fail()
	}
	fmt.Print(req.URI().String(), "\n")
	fmt.Print(req.Header.String(), "\n")
	fmt.Print(string(req.Body()), "\n")
}

func Test1FastHTTPRequestEncode(t *testing.T) {
	var buf bytes.Buffer
	var _, err = unMarshaledRequestFastHTTPTest1.WriteTo(&buf)
	var httpPacket []byte = buf.Bytes()
	if err != nil {
		fmt.Print(err, "\n")
		t.Fail()
	} else if !bytes.Equal(httpPacket, marshaledRequestTest1) {
		fmt.Print("cap--len of httpPacket:", cap(httpPacket), "--", len(httpPacket), "\n")
		fmt.Print("cap--len of httpPacket:", cap(marshaledRequestTest1), "--", len(marshaledRequestTest1), "\n")
		fmt.Print("encoded not same with original\n")
		t.Fail()
	}
	fmt.Print(string(httpPacket), "\n")
}
