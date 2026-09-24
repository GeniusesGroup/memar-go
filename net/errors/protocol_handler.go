/* For license and copyright information please see the LEGAL file in the code repository */

package errs

var (
	ProtocolHandler errProtocolHandler
)

type (
	errProtocolHandler struct{}
)

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (dt *errProtocolHandler) MediaType() string {
	return domainBaseMediatype + "protocol-handler"
}
