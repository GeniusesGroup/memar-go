/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"memar/datatype"
	"memar/protocol"
)

var DT domainType

type domainType struct {
	datatype.DataType
}

//memar:impl memar/protocol.MediaType
func (dt *domainType) MediaType() string { return "domain/memar.scm.geniuses.group; package=log" }

//memar:impl memar/protocol.DataType_Details
func (dt *domainType) Status() protocol.SoftwareStatus    { return protocol.Software_PreAlpha }
func (dt *domainType) ReferenceURI() string               { return "" }
func (dt *domainType) IssueDate() string                  { return "" }
func (dt *domainType) ExpiryDate() string                 { return "" }
func (dt *domainType) ExpireInFavorOf() protocol.DataType { return nil }
