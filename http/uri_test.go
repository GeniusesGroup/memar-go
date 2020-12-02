/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"reflect"
	"testing"
)

type URITest struct {
	name       string
	raw        string
	encoded    string
	uri        URI // expected parse
	out        URI // parsed one
	wantURIEnd int
}

var urlTests = []URITest{
	{
		name:    "asterisk-form",
		raw:     "* ",
		encoded: "*",
		uri: URI{
			Raw:       "*",
			Scheme:    "",
			Authority: "",
			Path:      "*",
			Query:     "",
			Fragment:  "",
		},
		wantURIEnd: 1,
	}, {
		name:    "origin-form1",
		raw:     "/apis?2586547852#api ",
		encoded: "/apis?2586547852",
		uri: URI{
			Raw:       "/apis?2586547852#api",
			Scheme:    "",
			Authority: "",
			Path:      "/apis",
			Query:     "2586547852",
			Fragment:  "api",
		},
		wantURIEnd: 20,
	}, {
		name:    "origin-form2",
		raw:     "/action/do/show/411?2586547852#Test ",
		encoded: "/action/do/show/411?2586547852",
		uri: URI{
			Raw:       "/action/do/show/411?2586547852#Test",
			Scheme:    "",
			Authority: "",
			Path:      "/action/do/show/411",
			Query:     "2586547852",
			Fragment:  "Test",
		},
		wantURIEnd: 35,
	}, {
		name:    "absolute-URI1",
		raw:     "https://tools.ietf.org/html/rfc2616#section-3.2 ",
		encoded: "https://tools.ietf.org/html/rfc2616",
		uri: URI{
			Raw:       "https://tools.ietf.org/html/rfc2616#section-3.2",
			Scheme:    "https",
			Authority: "tools.ietf.org",
			Path:      "/html/rfc2616",
			Query:     "",
			Fragment:  "section-3.2",
		},
		wantURIEnd: 47,
	}, {
		name:    "absolute-URI2",
		raw:     "http://www.sabz.city/#file%20one%26two ",
		encoded: "http://www.sabz.city/",
		uri: URI{
			Raw:       "http://www.sabz.city/#file%20one%26two",
			Scheme:    "http",
			Authority: "www.sabz.city",
			Path:      "/",
			Query:     "",
			Fragment:  "file%20one%26two",
		},
		wantURIEnd: 38,
	}, {
		name:    "absolute-URI3",
		raw:     "https://www.sabz.city/pub/WWW/TheProject.html ",
		encoded: "https://www.sabz.city/pub/WWW/TheProject.html",
		uri: URI{
			Raw:       "https://www.sabz.city/pub/WWW/TheProject.html",
			Scheme:    "https",
			Authority: "www.sabz.city",
			Path:      "/pub/WWW/TheProject.html",
			Query:     "",
			Fragment:  "",
		},
		wantURIEnd: 45,
	}, {
		name:    "absolute-URI4",
		raw:     "www.sabz.city/apis?2586547852#api ",
		encoded: "www.sabz.city/apis?2586547852",
		uri: URI{
			Raw:       "www.sabz.city/apis?2586547852#api",
			Scheme:    "",
			Authority: "www.sabz.city",
			Path:      "/apis",
			Query:     "2586547852",
			Fragment:  "api",
		},
		wantURIEnd: 33,
	}, {
		name:    "ftp1",
		raw:     "ftp://webmaster@www.sabz.city/ ",
		encoded: "ftp://webmaster@www.sabz.city/",
		uri: URI{
			Raw:       "ftp://webmaster@www.sabz.city/",
			Scheme:    "ftp",
			Authority: "webmaster@www.sabz.city",
			Path:      "/",
			Query:     "",
			Fragment:  "",
		},
		wantURIEnd: 30,
	}, {
		name:    "empty query",
		raw:     "http://www.sabz.city/? ",
		encoded: "http://www.sabz.city/",
		uri: URI{
			Raw:       "http://www.sabz.city/?",
			Scheme:    "http",
			Authority: "www.sabz.city",
			Path:      "/",
			Query:     "",
			Fragment:  "",
		},
		wantURIEnd: 22,
	},
}

func TestURI_UnMarshal(t *testing.T) {
	for _, tt := range urlTests {
		t.Run(tt.name, func(t *testing.T) {
			var gotURIEnd = tt.out.UnMarshal(tt.raw)
			if gotURIEnd != tt.wantURIEnd {
				t.Errorf("URI.UnMarshal() = %v, want %v", gotURIEnd, tt.wantURIEnd)
			}
			if !reflect.DeepEqual(tt.out, tt.uri) {
				t.Errorf("URI.UnMarshal(%q):\n\tgot  %v\n\twant %v\n", tt.raw, tt.out, tt.uri)
			}
		})
	}
}

func TestURI_Marshal(t *testing.T) {
	for _, tt := range urlTests {
		t.Run(tt.name, func(t *testing.T) {
			var httpPacket []byte
			httpPacket = tt.uri.Marshal(httpPacket)
			if tt.encoded != string(httpPacket) {
				t.Errorf("URI.UnMarshal():\n\tgot  %v\n\twant %v\n", string(httpPacket), tt.encoded)
			}
		})
	}
}
