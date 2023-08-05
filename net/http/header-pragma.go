/* For license and copyright information please see the LEGAL file in the code repository */

package http

// FixPragmaCacheControl do as RFC 7234, section 5.4: Treat [Pragma: no-cache] as [Cache-Control: no-cache]
func (h *header) FixPragmaCacheControl() {
	if h.Get(HeaderKeyPragma) == "no-cache" {
		if h.Gets(HeaderKeyCacheControl) == nil {
			h.Set(HeaderKeyCacheControl, "no-cache")
		}
	}
}
