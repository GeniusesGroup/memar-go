/* For license and copyright information please see the LEGAL file in the code repository */

package achaemenid

import (
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
	mediatype.MT
}

//memar:impl memar/protocol.ObjectLifeCycle
func (dt *domain_) Init() (err protocol.Error) {
	err = dt.MT.Init("domain/memar.scm.geniuses.group; package=achaemenid")
	return
}

//memar:impl memar/protocol.DataType_Details
func (dt *domain_) Status() protocol.SoftwareStatus    { return protocol.Software_PreAlpha }
func (dt *domain_) ReferenceURI() string               { return "" }
func (dt *domain_) IssueDate() protocol.Time           { return nil }
func (dt *domain_) ExpiryDate() protocol.Time          { return nil }
func (dt *domain_) ExpireInFavorOf() protocol.DataType { return nil }
