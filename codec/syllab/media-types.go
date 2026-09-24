/* For license and copyright information please see the LEGAL file in the code repository */

package syllab

import (
	"memar/computer/datatype"
	datatype_p "memar/computer/datatype/protocol"
)

var MediaType mediaType

type mediaType struct {
	datatype.DataType
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (s *mediaType) MediaType() string { return "application/syllab" }

//memar:impl memar/protocol.FileExtension
func (s *mediaType) FileExtension() string { return "syllab" }

//memar:impl memar/datatype/protocol.DataType_Details
func (s *mediaType) LifeCycle() datatype_p.LifeCycle      { return datatype_p.LifeCycle_Unset }
func (s *mediaType) ReferenceURI() string                 { return "" }
func (s *mediaType) IssueDate() string                    { return "" }
func (s *mediaType) ExpiryDate() string                   { return "" }
func (s *mediaType) ExpireInFavorOf() datatype_p.DataType { return nil }
