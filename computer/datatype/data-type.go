/* For license and copyright information please see the LEGAL file in the code repository */

package datatype

import (
	"memar/protocol"
)

// DataType embed to provide some methods when not need to implement by others!
type DataType struct{}

//memar:impl memar/protocol.DataType_Details
func (dt *DataType) Status() protocol.SoftwareStatus    { return protocol.Software_Unset }
func (dt *DataType) ReferenceURI() string               { return "" }
func (dt *DataType) IssueDate() protocol.Time           { return nil }
func (dt *DataType) ExpiryDate() protocol.Time          { return nil }
func (dt *DataType) ExpireInFavorOf() protocol.DataType { return nil }

//memar:impl memar/protocol.Detail
func (dt *DataType) Domain() string   { return "" }
func (dt *DataType) Summary() string  { return "" }
func (dt *DataType) Overview() string { return "" }
func (dt *DataType) UserNote() string { return "" }
func (dt *DataType) DevNote() string  { return "" }
func (dt *DataType) TAGS() []string   { return []string{} }

//memar:impl memar/protocol.Quiddity
func (dt *DataType) Name() string         { return "" }
func (dt *DataType) Abbreviation() string { return "" }
func (dt *DataType) Aliases() []string    { return []string{} }
