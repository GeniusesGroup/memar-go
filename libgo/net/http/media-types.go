/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"memar/computer/datatype"
	datatype_p "memar/computer/datatype/protocol"
)

type mediaType struct {
	datatype.DataType
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (m *mediaType) MediaType() string { return "application/http" }

//memar:impl memar/protocol.File_Extension
func (m *mediaType) FileExtension() string { return "http" }

//memar:impl memar/datatype/protocol.DataType_Details
func (m *mediaType) LifeCycle() datatype_p.LifeCycle { return datatype_p.LifeCycle_PreAlpha }
func (m *mediaType) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (m *mediaType) IssueDate() string                    { return "" }
func (m *mediaType) ExpiryDate() string                   { return "" }
func (m *mediaType) ExpireInFavorOf() datatype_p.DataType { return nil }
