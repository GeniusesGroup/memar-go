/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"../compress/raw"
)

type RequestTest struct {
	name   string
	packet []byte
	req    Request // expected parse
	out    Request // parsed one
}

var requestTests = []RequestTest{
	{
		name:   "simple1",
		packet: []byte("GET / HTTP/1.1\r\n" + "\r\n"),
		req: Request{
			method: "GET",
			uri: URI{
				uri:       "/",
				uriAsByte: []byte("/"),
				scheme:    "",
				authority: "",
				path:      "/",
				query:     "",
				fragment:  "",
			},
			version: "HTTP/1.1",
			H:       header{},
			body: body{
				Codec: nil,
			},
		},
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
				uriAsByte: []byte("/m?2586547852"),
				scheme:    "",
				authority: "",
				path:      "/m",
				query:     "2586547852",
				fragment:  "",
			},
			version: "HTTP/1.1",
			H: header{
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
				Codec: &raw.ComDecom{
					Data: []byte(`{"Omid":"OMID"}`),
				},
			},
		},
	}, {
		name: "without-body",
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
			"\r\n"),
		req: Request{
			method: "POST",
			uri: URI{
				uri:       "/m?2586547852",
				uriAsByte: []byte("/m?2586547852"),
				scheme:    "",
				authority: "",
				path:      "/m",
				query:     "2586547852",
				fragment:  "",
			},
			version: "HTTP/1.1",
			H: header{
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
				Codec: nil,
			},
		},
	},
}

func TestRequest_Unmarshal(t *testing.T) {
	for _, tt := range requestTests {
		t.Run(tt.name, func(t *testing.T) {
			tt.out.Init()
			var err = tt.out.Unmarshal(tt.packet)
			if err != nil {
				t.Errorf("Request.Unmarshal() error = %v", err)
			}
			if tt.out.method != tt.req.method {
				t.Errorf("Request.Unmarshal(%q) - Method:\n\tgot  %v\n\twant %v\n", tt.packet, tt.out.method, tt.req.method)
			}
			if tt.out.uri.uri != tt.req.uri.uri {
				t.Errorf("Request.Unmarshal(%q) - URI:\n\tgot  %v\n\twant %v\n", tt.packet, tt.out.uri.uri, tt.req.uri.uri)
			}
			if tt.out.version != tt.req.version {
				t.Errorf("Request.Unmarshal(%q) - Version:\n\tgot  %v\n\twant %v\n", tt.packet, tt.out.version, tt.req.version)
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
			fmt.Println("cap--len of httpPacket:", cap(httpPacket), "--", len(httpPacket))
			fmt.Println("cap--len of tt.packet:", cap(tt.packet), "--", len(tt.packet))
			if httpPacket == nil {
				t.Errorf("Request.Marshal() return nil!")
			}
			if !bytes.Equal(tt.packet, httpPacket) {
				fmt.Println("encoded not same with original or just encode headers in not same order! ", tt.name)
				fmt.Println(string(httpPacket))
				// t.Errorf("Request.Marshal(%q):\n\tgot  %v\n\twant %v\n", tt.req, httpPacket, tt.packet)
			}
		})
	}
}

func BenchmarkRequest_Unmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var httpReq Request
		httpReq.Init()
		httpReq.Unmarshal(requestTests[1].packet)
	}
}

func BenchmarkRequest_Marshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		requestTests[1].req.Marshal()
	}
}
