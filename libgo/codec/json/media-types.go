/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	"memar/computer/datatype"
	datatype_p "memar/computer/datatype/protocol"
)

var DT domainType

type domainType struct {
	datatype.DataType
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (s *domainType) MediaType() string { return "application/json" }

//memar:impl memar/storage/protocol.FileExtension
func (s *domainType) FileExtension() string { return "json" }

//memar:impl memar/datatype/protocol.DataType_Details
func (s *domainType) LifeCycle() datatype_p.LifeCycle      { return datatype_p.LifeCycle_StableRelease }
func (s *domainType) ReferenceURI() string                 { return "" }
func (s *domainType) IssueDate() string                    { return "" }
func (s *domainType) ExpiryDate() string                   { return "" }
func (s *domainType) ExpireInFavorOf() datatype_p.DataType { return nil }
