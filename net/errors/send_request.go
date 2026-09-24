/* For license and copyright information please see the LEGAL file in the code repository */

package errs

var (
	SendRequest errSendRequest
)

type (
	errSendRequest struct{}
)

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (dt *errSendRequest) MediaType() string {
	return domainBaseMediatype + "send-request"
}
