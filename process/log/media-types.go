/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"memar/datatype"
	datatype_p "memar/datatype/protocol"
)

var DT domainType

type domainType struct {
	datatype.DataType
}

//memar:impl memar/protocol.MediaType
func (dt *domainType) MediaType() string { return "domain/memar.scm.geniuses.group; package=log" }

//memar:impl memar/datatype/protocol.DataType_Details
func (dt *domainType) LifeCycle() datatype_p.LifeCycle      { return datatype_p.LifeCycle_PreAlpha }
func (dt *domainType) ReferenceURI() string                 { return "" }
func (dt *domainType) IssueDate() string                    { return "" }
func (dt *domainType) ExpiryDate() string                   { return "" }
func (dt *domainType) ExpireInFavorOf() datatype_p.DataType { return nil }
