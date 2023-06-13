/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"libgo/detail"
	"libgo/mediatype"
	"libgo/protocol"
)

var (
	DefaultEvent_MediaType mediaType
)

func init() {
	DefaultEvent_MediaType.Init("domain/libgo.scm.geniuses.group; package=log; type=event; name=default")
}

type mediaType struct {
	detail.DS
	mediatype.MT
}

//libgo:impl libgo/protocol.MediaType
func (m *mediaType) FileExtension() string           { return "" }
func (m *mediaType) Status() protocol.SoftwareStatus { return protocol.Software_PreAlpha }
func (m *mediaType) ReferenceURI() string {
	return ""
}
func (m *mediaType) IssueDate() protocol.Time            { return nil }
func (m *mediaType) ExpiryDate() protocol.Time           { return nil }
func (m *mediaType) ExpireInFavorOf() protocol.MediaType { return nil }

//libgo:impl libgo/protocol.Object
func (m *mediaType) Fields() []protocol.Object_Member_Field   { return nil }
func (m *mediaType) Methods() []protocol.Object_Member_Method { return nil }
