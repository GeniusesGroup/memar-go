/* For license and copyright information please see the LEGAL file in the code repository */

package http_headers

import (
	http_p "memar/libgo/net/http/protocol"
)

// FixPragmaCacheControl do as RFC 7234, section 5.4: Treat [Pragma: no-cache] as [Cache-Control: no-cache]
func FixPragmaCacheControl(h http_p.Header) {
	if h.Header_Get(Key_Pragma) == "no-cache" {
		if h.Header_Get(Key_CacheControl) == "" {
			h.Header_Add(Key_CacheControl, "no-cache")
		}
	}
}
