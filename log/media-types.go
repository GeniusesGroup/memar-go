/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"memar/datatype"
	"memar/mediatype"
	"memar/protocol"
)

var (
	DT domainType
)

func init() {
	DT.Init()
}

type domainType struct {
	datatype.DataType
	mediatype.MT
}

//memar:impl memar/protocol.ObjectLifeCycle
func (dt *domainType) Init() (err protocol.Error) {
	err = dt.MT.Init("domain/memar.scm.geniuses.group; package=log")
	return
}

//memar:impl memar/protocol.DataType_Details
func (dt *domainType) Status() protocol.SoftwareStatus    { return protocol.Software_PreAlpha }
func (dt *domainType) ReferenceURI() string               { return "" }
func (dt *domainType) IssueDate() protocol.Time           { return nil }
func (dt *domainType) ExpiryDate() protocol.Time          { return nil }
func (dt *domainType) ExpireInFavorOf() protocol.DataType { return nil }
