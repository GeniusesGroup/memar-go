/* For license and copyright information please see the LEGAL file in the code repository */

package gp

import (
	"memar/computer/datatype"
	datatype_p "memar/computer/datatype/protocol"
)

var (
	domain domainType
)

type domainType struct {
	datatype.DataType
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (dt *domainType) MediaType() string { return "domain/memar.scm.geniuses.group; package=gp" }

//memar:impl memar/datatype/protocol.DataType_Details
func (dt *domainType) LifeCycle() datatype_p.LifeCycle      { return datatype_p.LifeCycle_PreAlpha }
func (dt *domainType) ReferenceURI() string                 { return "" }
func (dt *domainType) IssueDate() string                    { return "" }
func (dt *domainType) ExpiryDate() string                   { return "" }
func (dt *domainType) ExpireInFavorOf() datatype_p.DataType { return nil }
