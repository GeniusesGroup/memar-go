/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"libgo/detail"
	"libgo/mediatype"
	"libgo/protocol"
)

var (
	Package_MediaType mediaType
)

func init() {
	Package_MediaType.Init("domain/libgo.scm.geniuses.group; package=gp")
}

type mediaType struct {
	detail.Details
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
func (m *mediaType) Fields() []protocol.DataType         { return nil }
func (m *mediaType) Methods() []protocol.DataType_Method { return nil }
