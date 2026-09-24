/* For license and copyright information please see the LEGAL file in the code repository */

package errs

var (
	GuestNotAllow errGuestNotAllow
)

type (
	errGuestNotAllow struct{}
)

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (dt *errGuestNotAllow) MediaType() string { return domainBaseMediatype + "guest-not-allow" }

// memar error make guest-not-allowed ./errors -langs="eng,per"
