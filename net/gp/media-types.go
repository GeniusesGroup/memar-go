/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"memar/datatype"
	"memar/mediatype"
	"memar/protocol"
)

var (
	Package_MediaType mediaType
)

func init() {
	Package_MediaType.Init("domain/memar.scm.geniuses.group; package=gp")
}

type mediaType struct {
	datatype.DataType
	mediatype.MT
}

//memar:impl memar/protocol.DataType_Details
func (m *mediaType) Status() protocol.SoftwareStatus { return protocol.Software_PreAlpha }
func (m *mediaType) ReferenceURI() string {
	return ""
}
func (m *mediaType) IssueDate() protocol.Time           { return nil }
func (m *mediaType) ExpiryDate() protocol.Time          { return nil }
func (m *mediaType) ExpireInFavorOf() protocol.DataType { return nil }
