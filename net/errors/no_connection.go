/* For license and copyright information please see the LEGAL file in the code repository */

package errs

var (
	NoConnection errNoConnection
)

type (
	errNoConnection struct{}
)

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (dt *errNoConnection) MediaType() string {
	return domainBaseMediatype + "no-connection"
}
