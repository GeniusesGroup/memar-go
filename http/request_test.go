/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"runtime"
	"testing"

	"../compress"

	"github.com/valyala/fasthttp"
)

/*
note1: This benchmark is not apple to apple because libgo force RFCs rules in methods not in codec phase.
note2: This benchmark is not apple to apple because both net/http and fasthttp force to marshal||unmarshal by pass *bufio.Reader||*bufio.Writer that not forced by libgo

------------------------------------------requestTests[0].packet-------------------------------------------
BenchmarkNetHttpRequestDecode-8		387325	      2682 ns/op	    4721 B/op	       7 allocs/op
BenchmarkLibgoRequestUnmarshal-8	1000000	      1031 ns/op	    1712 B/op	       3 allocs/op
BenchmarkFastHTTPRequestDecode-8	631910	      1872 ns/op	    4272 B/op	       7 allocs/op

BenchmarkNetHttpRequestEncode-8		600348	      1979 ns/op	    1056 B/op			10 allocs/op
BenchmarkLibgoRequestMarshal-8		12630421	  100 ns/op			32 B/op				1 allocs/op
BenchmarkFastHTTPRequestEncode-8	822142	      1420 ns/op	    4208 B/op			3 allocs/op

note1: Libgo decoder performance better by: 2.6X from standard GO, 1.8X from FastHTTP
note2: Libgo encoder performance better by: 19.8X from standard GO, 14.2X from FastHTTP

------------------------------------------requestTests[1].packet-------------------------------------------
BenchmarkNetHttpRequestDecode-8		150092	      7604 ns/op	    5865 B/op	      22 allocs/op
BenchmarkLibgoRequestUnmarshal-8	551079	      2153 ns/op	    1712 B/op	       3 allocs/op
BenchmarkFastHTTPRequestDecode-8	193660	      6082 ns/op	    5713 B/op	      31 allocs/op

BenchmarkNetHttpRequestEncode-8		250158	      4937 ns/op	    2062 B/op	      15 allocs/op
BenchmarkLibgoRequestMarshal-8		1287526	       939 ns/op	     352 B/op	       1 allocs/op
BenchmarkLibgoRequestWriteTo-8		1272480	       953 ns/op	     352 B/op	       1 allocs/op
BenchmarkFastHTTPRequestEncode-8	693639	      1718 ns/op	    4208 B/op	       3 allocs/op

note1: Libgo decoder performance better by: 3.5X from standard GO, 2.8X from FastHTTP
note2: Libgo encoder performance better by: 5.2X from standard GO, 1.8X from FastHTTP
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

var test1Body = compress.Raw([]byte{})
var test2Body = compress.Raw([]byte(`{"Omid":"OMID"}`))

var requestTests = []RequestTest{
	{
		name:   "simple1",
		packet: []byte("GET /m HTTP/1.1\r\n" + "\r\n"),
		req: Request{
			method: "GET",
			uri: URI{
				uri:       "/m",
				scheme:    "",
				authority: "",
				path:      "/m",
				query:     "",
				fragment:  "",
			},
			version: "HTTP/1.1",
			header:  header{},
			body: body{
				Codec: &test1Body,
			},
		},
		out: NewRequest(),
	}, {
		name: "full1",
		packet: []byte("POST /m?2586547852 HTTP/1.1\r\n" +
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
			method: "POST",
			uri: URI{
				uri:       "/m?2586547852",
				scheme:    "",
				authority: "",
				path:      "/m",
				query:     "2586547852",
				fragment:  "",
			},
			version: "HTTP/1.1",
			header: header{
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
			body: body{
				Codec: &test2Body,
			},
		},
		out: NewRequest(),
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

func BenchmarkLibgoRequestUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var r = NewRequest()
		r.Unmarshal(requestTests[1].packet)
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
		b.Errorf("http.Request.Unmarshal() error = %v", err)
	}

	for n := 0; n < b.N; n++ {
		var b []byte
		var r = bytes.NewBuffer(b)
		unMarshaledRequestNetHTTPTest1.Write(r)
	}
}

func BenchmarkLibgoRequestMarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		requestTests[1].req.Marshal()
	}
}

func BenchmarkLibgoRequestWriteTo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		requestTests[1].req.WriteTo(io.Discard)
	}
}

func BenchmarkFastHTTPRequestEncode(b *testing.B) {
	var unMarshaledRequestFastHTTPTest1 fasthttp.Request
	var err error
	var br = bytes.NewBuffer(requestTests[1].packet)
	var fastBuf = bufio.NewReader(br)
	err = unMarshaledRequestFastHTTPTest1.Read(fastBuf)
	if err != nil {
		b.Errorf("fasthttp.Request.Unmarshal() error = %v", err)
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

func TestRequest_Unmarshal(t *testing.T) {
	for _, tt := range requestTests {
		t.Run(tt.name, func(t *testing.T) {
			var err = tt.out.Unmarshal(tt.packet)
			if err != nil {
				t.Errorf("Request.Unmarshal() error = %v", err)
			}
			if tt.out.method != tt.req.method {
				t.Errorf("Request.Unmarshal(%q):\n\tgot  %v\n\twant %v\n", tt.packet, tt.out.method, tt.req.method)
			}
			if !reflect.DeepEqual(tt.out.uri, tt.req.uri) {
				t.Errorf("Request.Unmarshal(%q):\n\tgot  %v\n\twant %v\n", tt.packet, tt.out.uri, tt.req.uri)
			}
			if tt.out.version != tt.req.version {
				t.Errorf("Request.Unmarshal(%q):\n\tgot  %v\n\twant %v\n", tt.packet, tt.out.version, tt.req.version)
			}
			// if !reflect.DeepEqual(tt.out.header.headers, tt.req.header.headers) {
			// 	t.Errorf("Request.Unmarshal(%q):\n\tgot  %v\n\twant %v\n", tt.packet, tt.out.header.headers, tt.req.header.headers)
			// }
			if !reflect.DeepEqual(tt.out.body, tt.req.body) {
				t.Errorf("Request.Unmarshal(%q):\n\tgot  %v\n\twant %v\n", tt.packet, tt.out.body, tt.req.body)
			}
		})
	}

	// s := req.header.GetSetCookies()[0]
	// err := s.CheckAndSanitize()
	// fmt.Fprintf(os.Stderr, "%v\n", err)
	// fmt.Fprintf(os.Stderr, "%v\n", s)
}

func TestRequest_Marshal(t *testing.T) {
	for _, tt := range requestTests {
		t.Run(tt.name, func(t *testing.T) {
			var httpPacket []byte = tt.req.Marshal()
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
		t.Errorf("fasthttp.Request.Unmarshal() error = %v", err)
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
