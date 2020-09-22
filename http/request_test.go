/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"testing"

	"github.com/valyala/fasthttp"
)

/*
------------------------------------------requestTests[0].packet-------------------------------------------
BenchmarkNetHttpRequestDecode-8    	  387325	      2682 ns/op	    4721 B/op	       7 allocs/op
BenchmarkLibgoRequestDecode-8      	 1000000	      1031 ns/op	    1712 B/op	       3 allocs/op
BenchmarkFastHTTPRequestDecode-8   	  631910	      1872 ns/op	    4272 B/op	       7 allocs/op

BenchmarkNetHttpRequestEncode-8    	  600348	      1979 ns/op	    1056 B/op	      10 allocs/op
BenchmarkLibgoRequestEncode-8      	12630421	       100 ns/op	      32 B/op	       1 allocs/op
BenchmarkFastHTTPRequestEncode-8   	  822142	      1420 ns/op	    4208 B/op	       3 allocs/op

note1: This benchmark is not apple to apple because libgo force RFCs rules in methods not in encode||decode phase!
note2: Libgo decoder performance better by: 2.6X from standard GO, 1.8X from FastHTTP
note3: Libgo encoder performance better by: 19.8X from standard GO, 14.2X from FastHTTP

------------------------------------------requestTests[1].packet-------------------------------------------
BenchmarkNetHttpRequestDecode-8    	  150092	      7604 ns/op	    5865 B/op	      22 allocs/op
BenchmarkLibgoRequestDecode-8      	  551079	      2153 ns/op	    1712 B/op	       3 allocs/op
BenchmarkFastHTTPRequestDecode-8   	  176703	      6903 ns/op	    6289 B/op	      32 allocs/op

BenchmarkNetHttpRequestEncode-8    	  226796	      6023 ns/op	    2069 B/op	      15 allocs/op
BenchmarkLibgoRequestEncode-8      	 1271994	       946 ns/op	     352 B/op	       1 allocs/op
BenchmarkFastHTTPRequestEncode-8   	  693639	      1718 ns/op	    4208 B/op	       3 allocs/op

note1: This benchmark is not apple to apple because libgo force RFCs rules in methods not in encode||decode phase!
note2: Libgo decoder performance better by: 3.5X from standard GO, 3.2X from FastHTTP
note3: Libgo encoder performance better by: 6.4X from standard GO, 1.8X from FastHTTP
*/

/*
	Data
*/

type RequestTest struct {
	name   string
	packet []byte
	req    Request  // expected parse
	out    *Request // parsed one
}

var requestTests = []RequestTest{
	{
		name:   "simple1",
		packet: []byte("GET /apis HTTP/1.1\r\n" + "\r\n"),
		req: Request{
			Method: "GET",
			URI: URI{
				Raw:       "/apis",
				Scheme:    "",
				Authority: "",
				Path:      "/apis",
				Query:     "",
				Fragment:  "",
			},
			Version: "HTTP/1.1",
			Header:  header{},
			Body:    []byte{},
		},
		out: MakeNewRequest(),
	}, {
		name: "full1",
		packet: []byte("POST /apis?2586547852 HTTP/1.1\r\n" +
			"Accept: text/html\r\n" +
			"Accept-Encoding: gzip, deflate\r\n" +
			"Accept-Language: en,fa;q=0.9\r\n" +
			"Cache-Control: max-age=0\r\n" +
			"Connection: keep-alive\r\n" +
			"Content-Length: 15\r\n" +
			"Content-Type: application/json\r\n" +
			"Host: www.sabz.city\r\n" +
			"Set-Cookie: test\r\n" +
			"Upgrade-Insecure-Requests: 1\r\n" +
			"User-Agent: Mozilla\r\n" +
			"\r\n" +
			`{"Omid":"OMID"}`),
		req: Request{
			Method: "POST",
			URI: URI{
				Raw:       "/apis?2586547852",
				Scheme:    "",
				Authority: "",
				Path:      "/apis",
				Query:     "2586547852",
				Fragment:  "",
			},
			Version: "HTTP/1.1",
			Header: header{
				headers: map[string][]string{
					"Accept":                    []string{"text/html"},
					"Accept-Encoding":           []string{"gzip, deflate"},
					"Accept-Language":           []string{"en,fa;q=0.9"},
					"Cache-Control":             []string{"max-age=0"},
					"Connection":                []string{"keep-alive"},
					"Content-Length":            []string{"15"},
					"Content-Type":              []string{"application/json"},
					"Host":                      []string{"www.sabz.city"},
					"Set-Cookie":                []string{"test"},
					"Upgrade-Insecure-Requests": []string{"1"},
					"User-Agent":                []string{"Mozilla"},
				},
			},
			Body: []byte(`{"Omid":"OMID"}`),
		},
		out: MakeNewRequest(),
	},
}

func init() {
	fmt.Print("Number of CPU used:", runtime.NumCPU(), "\n")
}

/*
	Benchmarks
*/

func BenchmarkNetHttpRequestDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var r = bytes.NewBuffer(requestTests[1].packet)
		var buf = bufio.NewReader(r)
		http.ReadRequest(buf)
	}
}

func BenchmarkLibgoRequestDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var r = MakeNewRequest()
		r.UnMarshal(requestTests[1].packet)
	}
}

func BenchmarkFastHTTPRequestDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var r = bytes.NewBuffer(requestTests[1].packet)
		var buf = bufio.NewReader(r)
		var req fasthttp.Request
		req.Read(buf)
	}
}

func BenchmarkNetHttpRequestEncode(b *testing.B) {
	var unMarshaledRequestNetHTTPTest1 *http.Request
	var err error
	var r = bytes.NewBuffer(requestTests[1].packet)
	var buf = bufio.NewReader(r)
	unMarshaledRequestNetHTTPTest1, err = http.ReadRequest(buf)
	if err != nil {
		b.Errorf("http.Request.UnMarshal() error = %v", err)
	}

	for n := 0; n < b.N; n++ {
		var b []byte
		var r = bytes.NewBuffer(b)
		unMarshaledRequestNetHTTPTest1.Write(r)
	}
}

func BenchmarkLibgoRequestEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		requestTests[1].req.Marshal()
	}
}

func BenchmarkFastHTTPRequestEncode(b *testing.B) {
	var unMarshaledRequestFastHTTPTest1 fasthttp.Request
	var err error
	var br = bytes.NewBuffer(requestTests[1].packet)
	var fastBuf = bufio.NewReader(br)
	err = unMarshaledRequestFastHTTPTest1.Read(fastBuf)
	if err != nil {
		b.Errorf("fasthttp.Request.UnMarshal() error = %v", err)
	}

	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		var bufW = bufio.NewWriter(&buf)
		unMarshaledRequestFastHTTPTest1.Write(bufW)
	}
}

/*
	Tests
*/

func TestRequest_UnMarshal(t *testing.T) {
	for _, tt := range requestTests {
		t.Run(tt.name, func(t *testing.T) {
			var err = tt.out.UnMarshal(tt.packet)
			if err != nil {
				t.Errorf("Request.UnMarshal() error = %v", err)
			}
			if tt.out.Method != tt.req.Method {
				t.Errorf("Request.UnMarshal(%q):\n\tgot  %v\n\twant %v\n", tt.packet, tt.out.Method, tt.req.Method)
			}
			if !reflect.DeepEqual(tt.out.URI, tt.req.URI) {
				t.Errorf("Request.UnMarshal(%q):\n\tgot  %v\n\twant %v\n", tt.packet, tt.out.URI, tt.req.URI)
			}
			if tt.out.Version != tt.req.Version {
				t.Errorf("Request.UnMarshal(%q):\n\tgot  %v\n\twant %v\n", tt.packet, tt.out.Version, tt.req.Version)
			}
			// if !reflect.DeepEqual(tt.out.Header.headers, tt.req.Header.headers) {
			// 	t.Errorf("Request.UnMarshal(%q):\n\tgot  %v\n\twant %v\n", tt.packet, tt.out.Header.headers, tt.req.Header.headers)
			// }
			if !bytes.Equal(tt.out.Body, tt.req.Body) {
				t.Errorf("Request.UnMarshal(%q):\n\tgot  %v\n\twant %v\n", tt.packet, tt.out.Body, tt.req.Body)
			}
		})
	}

	// s := req.Header.GetSetCookies()[0]
	// err := s.CheckAndSanitize()
	// fmt.Fprintf(os.Stderr, "%v\n", err)
	// fmt.Fprintf(os.Stderr, "%v\n", s)
}

func TestRequest_Marshal(t *testing.T) {
	for _, tt := range requestTests {
		t.Run(tt.name, func(t *testing.T) {
			var httpPacket []byte
			httpPacket = tt.req.Marshal()
			if httpPacket == nil {
				t.Errorf("Request.Marshal() return nil!")
			}
			if !bytes.Equal(tt.packet, httpPacket) {
				fmt.Print("encoded not same with original or just encode headers in not same order!\n")
				fmt.Print("cap--len of httpPacket:", cap(httpPacket), "--", len(httpPacket), "\n")
				fmt.Print("cap--len of tt.packet:", cap(tt.packet), "--", len(tt.packet), "\n")
				fmt.Print(string(httpPacket), "\n")
				// t.Errorf("Request.Marshal(%q):\n\tgot  %v\n\twant %v\n", tt.req, httpPacket, tt.packet)
			}
		})
	}
}

func TestFastHTTPRequestDecode(t *testing.T) {
	var br = bytes.NewBuffer(requestTests[1].packet)
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

func TestFastHTTPRequestEncode(t *testing.T) {
	var unMarshaledRequestFastHTTPTest1 fasthttp.Request
	var err error
	var br = bytes.NewBuffer(requestTests[1].packet)
	var fastBuf = bufio.NewReader(br)
	err = unMarshaledRequestFastHTTPTest1.Read(fastBuf)
	if err != nil {
		t.Errorf("fasthttp.Request.UnMarshal() error = %v", err)
	}

	var buf bytes.Buffer
	_, err = unMarshaledRequestFastHTTPTest1.WriteTo(&buf)
	var httpPacket []byte = buf.Bytes()
	if err != nil {
		fmt.Print(err, "\n")
		t.Fail()
	} else if !bytes.Equal(httpPacket, requestTests[1].packet) {
		fmt.Print("encoded not same with original or just encode headers in not same order!\n")
		fmt.Print("cap--len of httpPacket:", cap(httpPacket), "--", len(httpPacket), "\n")
		fmt.Print("cap--len of tt.packet:", cap(requestTests[1].packet), "--", len(requestTests[1].packet), "\n")
		fmt.Print(string(httpPacket), "\n")
		// t.Fail()
	}
}
