/* For license and copyright information please see the LEGAL file in the code repository */

package log

import (
	"memar/datatype"
	"memar/mediatype"
	"memar/protocol"
)

var (
	domain domain_
)

func init() {
	domain.Init()
}

type domain_ struct {
	datatype.DataType
	mediatype.MT
}

//memar:impl memar/protocol.ObjectLifeCycle
func (dt *domain_) Init() (err protocol.Error) {
	err = dt.MT.Init("domain/memar.scm.geniuses.group; package=log")
	return
}

//memar:impl memar/protocol.DataType_Details
func (dt *domain_) Status() protocol.SoftwareStatus    { return protocol.Software_PreAlpha }
func (dt *domain_) ReferenceURI() string               { return "" }
func (dt *domain_) IssueDate() protocol.Time           { return nil }
func (dt *domain_) ExpiryDate() protocol.Time          { return nil }
func (dt *domain_) ExpireInFavorOf() protocol.DataType { return nil }
