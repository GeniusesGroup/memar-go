/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"testing"
)

type uriTest struct {
	name       string
	raw        string
	encoded    string
	uri        URI // expected parse
	out        URI // parsed one
	wantURIEnd int
}

var uriTests = []uriTest{
	{
		name:    "asterisk-form",
		raw:     "* ",
		encoded: "*",
		uri: URI{
			uri:       "*",
			scheme:    "",
			authority: "",
			path:      "",
			query:     "",
			fragment:  "",
		},
		wantURIEnd: 1,
	}, {
		name:    "simple path",
		raw:     "/ ",
		encoded: "/",
		uri: URI{
			uri:       "/",
			scheme:    "",
			authority: "",
			path:      "/",
			query:     "",
			fragment:  "",
		},
		wantURIEnd: 1,
	}, {
		name:    "origin-form1",
		raw:     "/m?2586547852#api ",
		encoded: "/m?2586547852",
		uri: URI{
			uri:       "/m?2586547852#api",
			scheme:    "",
			authority: "",
			path:      "/m",
			query:     "2586547852",
			fragment:  "api",
		},
		wantURIEnd: 17,
	}, {
		name:    "origin-form2",
		raw:     "/action/do/show/411?2586547852#Test ",
		encoded: "/action/do/show/411?2586547852",
		uri: URI{
			uri:       "/action/do/show/411?2586547852#Test",
			scheme:    "",
			authority: "",
			path:      "/action/do/show/411",
			query:     "2586547852",
			fragment:  "Test",
		},
		wantURIEnd: 35,
	}, {
		name:    "absolute-URI1",
		raw:     "https://tools.ietf.org/html/rfc2616#section-3.2 ",
		encoded: "https://tools.ietf.org/html/rfc2616",
		uri: URI{
			uri:       "https://tools.ietf.org/html/rfc2616#section-3.2",
			scheme:    "https",
			authority: "tools.ietf.org",
			path:      "/html/rfc2616",
			query:     "",
			fragment:  "section-3.2",
		},
		wantURIEnd: 47,
	}, {
		name:    "absolute-URI2",
		raw:     "http://www.sabz.city/#file%20one%26two ",
		encoded: "http://www.sabz.city/",
		uri: URI{
			uri:       "http://www.sabz.city/#file%20one%26two",
			scheme:    "http",
			authority: "www.sabz.city",
			path:      "/",
			query:     "",
			fragment:  "file%20one%26two",
		},
		wantURIEnd: 38,
	}, {
		name:    "absolute-URI3",
		raw:     "https://www.sabz.city/pub/WWW/TheProject.html ",
		encoded: "https://www.sabz.city/pub/WWW/TheProject.html",
		uri: URI{
			uri:       "https://www.sabz.city/pub/WWW/TheProject.html",
			scheme:    "https",
			authority: "www.sabz.city",
			path:      "/pub/WWW/TheProject.html",
			query:     "",
			fragment:  "",
		},
		wantURIEnd: 45,
	}, {
		name:    "absolute-URI4",
		raw:     "www.sabz.city/m?2586547852#api ",
		encoded: "www.sabz.city/m?2586547852",
		uri: URI{
			uri:       "www.sabz.city/m?2586547852#api",
			scheme:    "",
			authority: "www.sabz.city",
			path:      "/m",
			query:     "2586547852",
			fragment:  "api",
		},
		wantURIEnd: 30,
	}, {
		name:    "ftp1",
		raw:     "ftp://webmaster@www.sabz.city/ ",
		encoded: "ftp://webmaster@www.sabz.city/",
		uri: URI{
			uri:       "ftp://webmaster@www.sabz.city/",
			scheme:    "ftp",
			authority: "webmaster@www.sabz.city",
			path:      "/",
			query:     "",
			fragment:  "",
		},
		wantURIEnd: 30,
	}, {
		name:    "empty query",
		raw:     "http://www.sabz.city/? ",
		encoded: "http://www.sabz.city/",
		uri: URI{
			uri:       "http://www.sabz.city/?",
			scheme:    "http",
			authority: "www.sabz.city",
			path:      "/",
			query:     "",
			fragment:  "",
		},
		wantURIEnd: 22,
	},
}

func TestURI_Unmarshal(t *testing.T) {
	for _, tt := range uriTests {
		t.Run(tt.name, func(t *testing.T) {
			var gotURIEnd = tt.out.unmarshalFrom(tt.raw)
			if gotURIEnd != tt.wantURIEnd {
				t.Errorf("URI.Unmarshal(%q) = %v, want %v",tt.name, gotURIEnd, tt.wantURIEnd)
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
		uriTest.uri.Set(uriTest.uri.scheme, uriTest.uri.authority, uriTest.uri.path, uriTest.uri.query)
		t.Run(uriTest.name, func(t *testing.T) {
			var httpPacket = uriTest.uri.Marshal()
			if uriTest.encoded != string(httpPacket) {
				t.Errorf("URI.Unmarshal():\n\tgot  %v\n\twant %v\n", string(httpPacket), uriTest.encoded)
			}
		})
	}
}
