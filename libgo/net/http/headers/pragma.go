/* For license and copyright information please see the LEGAL file in the code repository */

package http

// FixPragmaCacheControl do as RFC 7234, section 5.4: Treat [Pragma: no-cache] as [Cache-Control: no-cache]
func (h *Header) FixPragmaCacheControl() {
	if h.Header_Get(HeaderKey_Pragma) == "no-cache" {
		if h.Header_Get(HeaderKey_CacheControl) == "" {
			h.Header_Add(HeaderKey_CacheControl, "no-cache")
		}
	}
}
