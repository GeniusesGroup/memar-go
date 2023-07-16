/* For license and copyright information please see the LEGAL file in the code repository */

package datatype

import (
	"libgo/detail"
	"libgo/protocol"
)

type DataType struct {
	detail.Details
}

func (dt *DataType) Init() (err protocol.Error) {
	return
}

//libgo:impl libgo/protocol.DataType_Details
func (dt *DataType) Status() protocol.SoftwareStatus    { return protocol.Software_Unset }
func (dt *DataType) ReferenceURI() string               { return "" }
func (dt *DataType) IssueDate() protocol.Time           { return nil }
func (dt *DataType) ExpiryDate() protocol.Time          { return nil }
func (dt *DataType) ExpireInFavorOf() protocol.DataType { return nil }
