/* For license and copyright information please see the LEGAL file in the code repository */

package file

import (
	container_p "memar/computer/adt/container/protocol"
	"memar/computer/datatype"
	datatype_p "memar/computer/datatype/protocol"
)

var MediaType mediaType

type mediaType struct {
	datatype.DataType
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (m *mediaType) MediaType() string { return "application/uri" }

//memar:impl memar/storage/protocol.FileExtension
func (m *mediaType) FileExtension() string { return "uri" }

//memar:impl memar/datatype/protocol.DataType_Details
func (m *mediaType) LifeCycle() datatype_p.LifeCycle { return datatype_p.LifeCycle_PreAlpha }
func (m *mediaType) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/uri"
}
func (m *mediaType) IssueDate() string                    { return "" }
func (m *mediaType) ExpiryDate() string                   { return "" }
func (m *mediaType) ExpireInFavorOf() datatype_p.DataType { return nil }

func (m *mediaType) MinLength() container_p.NumberOfElement { return -1 }
func (m *mediaType) MaxLength() container_p.NumberOfElement { return -1 }
func (m *mediaType) Length() container_p.NumberOfElement    { return -1 }
