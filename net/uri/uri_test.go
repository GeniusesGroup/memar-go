/* For license and copyright information please see the LEGAL file in the code repository */

package uri

import (
	"testing"

	"memar/computer/buffer"
	error_p "memar/process/error/protocol"
	"memar/codec/string/ascii"
)

type uriTest struct {
	name       string
	encoded    string
	uri        URI[ascii.STR[buffer.String]] // expected parse
	out        URI[ascii.STR[buffer.String]] // parsed one
	wantURIEnd container_p.NumberOfElement
	wantError  error_p.Error
}

var uriTests = []uriTest{
	{
		name:    "asterisk-form",
		encoded: "*",
		uri: URI[ascii.STR[buffer.String]]{
			raw:    "*",
			scheme: "",
			AU: AU{
				authority: "",
			},
			path:     "",
			query:    "",
			fragment: "",
		},
		wantURIEnd: 1,
	}, {
		name:    "simple path",
		encoded: "/",
		uri: URI{
			raw:    "/",
			scheme: "",
			AU: AU{
				authority: "",
			},
			path:     "/",
			query:    "",
			fragment: "",
		},
		wantURIEnd: 1,
	}, {
		name:    "origin-form1",
		encoded: "/m?2586547852",
		uri: URI{
			raw:    "/m?2586547852#api",
			scheme: "",
			AU: AU{
				authority: "",
			},
			path:     "/m",
			query:    "2586547852",
			fragment: "api",
		},
		wantURIEnd: 17,
	}, {
		name:    "origin-form2",
		encoded: "/action/do/show/411?2586547852",
		uri: URI{
			raw:    "/action/do/show/411?2586547852#Test",
			scheme: "",
			AU: AU{
				authority: "",
			},
			path:     "/action/do/show/411",
			query:    "2586547852",
			fragment: "Test",
		},
		wantURIEnd: 35,
	}, {
		name:    "absolute-URI1",
		encoded: "https://tools.ietf.org/html/rfc2616",
		uri: URI{
			raw:    "https://tools.ietf.org/html/rfc2616#section-3.2",
			scheme: "https",
			AU: AU{
				authority: "tools.ietf.org",
			},
			path:     "/html/rfc2616",
			query:    "",
			fragment: "section-3.2",
		},
		wantURIEnd: 47,
	}, {
		name:    "absolute-URI2",
		encoded: "http://www.sabz.city/",
		uri: URI{
			raw:    "http://www.sabz.city/#file%20one%26two",
			scheme: "http",
			AU: AU{
				authority: "www.sabz.city",
			},
			path:     "/",
			query:    "",
			fragment: "file%20one%26two",
		},
		wantURIEnd: 38,
	}, {
		name:    "absolute-URI3",
		encoded: "https://www.sabz.city/pub/WWW/TheProject.html",
		uri: URI{
			raw:    "https://www.sabz.city/pub/WWW/TheProject.html",
			scheme: "https",
			AU: AU{
				authority: "www.sabz.city",
			},
			path:     "/pub/WWW/TheProject.html",
			query:    "",
			fragment: "",
		},
		wantURIEnd: 45,
	}, {
		name:    "absolute-URI4",
		encoded: "www.sabz.city/m?2586547852",
		uri: URI{
			raw:    "www.sabz.city/m?2586547852#api",
			scheme: "",
			AU: AU{
				authority: "www.sabz.city",
			},
			path:     "/m",
			query:    "2586547852",
			fragment: "api",
		},
		wantURIEnd: 30,
	}, {
		name:    "ftp1",
		encoded: "ftp://webmaster@www.sabz.city/",
		uri: URI{
			raw:    "ftp://webmaster@www.sabz.city/",
			scheme: "ftp",
			AU: AU{
				authority: "webmaster@www.sabz.city",
			},
			path:     "/",
			query:    "",
			fragment: "",
		},
		wantURIEnd: 30,
	}, {
		name:    "empty query",
		encoded: "http://www.sabz.city/",
		uri: URI{
			raw:    "http://www.sabz.city/?",
			scheme: "http",
			AU: AU{
				authority: "www.sabz.city",
			},
			path:     "/",
			query:    "",
			fragment: "",
		},
		wantURIEnd: 22,
	}, {
		name:    "embed uri in query",
		encoded: "http://www.sabz.city/repo?n=memar&m=modules/app/",
		uri: URI{
			raw:    "http://www.sabz.city/repo?n=memar&m=modules/app/",
			scheme: "http",
			AU: AU{
				authority: "www.sabz.city",
			},
			path:     "/repo",
			query:    "n=memar&m=modules/app/",
			fragment: "",
		},
		wantURIEnd: 48,
	},
}

func TestURI_Unmarshal(t *testing.T) {
	for _, tt := range uriTests {
		t.Run(tt.name, func(t *testing.T) {
			var gotURIEnd, err = tt.out.FromString(tt.uri.raw)
			if err != tt.wantError {
				t.Errorf("URI.Unmarshal(%q) = %v, want %v", tt.name, err, tt.wantError)
			}
			if gotURIEnd != tt.wantURIEnd {
				t.Errorf("URI.Unmarshal(%q) = %v, want %v", tt.name, gotURIEnd, tt.wantURIEnd)
			}
			if tt.out.scheme != tt.uri.scheme || tt.out.authority != tt.uri.authority || tt.out.path != tt.uri.path || tt.out.query != tt.uri.query || tt.out.fragment != tt.uri.fragment {
				t.Errorf("URI.Unmarshal(%q):\n\tgot  %v\n\twant %v\n", tt.name, tt.out, tt.uri)
			}
		})
	}
}

func TestURI_Marshal(t *testing.T) {
	for i := 1; i < len(uriTests); i++ { // start from 1 due to asterisk-form is not general form that we can use Set() method
		var uriTest = uriTests[i]
		uriTest.uri.Set(uriTest.uri.scheme, uriTest.uri.authority, uriTest.uri.path, uriTest.uri.query, uriTest.uri.fragment)
		t.Run(uriTest.name, func(t *testing.T) {
			var httpPacket = make([]byte, 0, uriTest.uri.SerializationLength())
			var n, err = uriTest.uri.Marshal(httpPacket)
			httpPacket = httpPacket[:n]
			if err != nil {
				t.Errorf("URI.Unmarshal():\n\tgot Error %v\n", err)
			}
			if uriTest.encoded != string(httpPacket) {
				t.Errorf("URI.Unmarshal():\n\tgot  %v\n\twant %v\n", string(httpPacket), uriTest.encoded)
			}
		})
	}
}
