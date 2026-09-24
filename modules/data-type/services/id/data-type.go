/* For license and copyright information please see the LEGAL file in the code repository */

package datatype

import (
	datatype_p "memar/computer/datatype/protocol"
	uuid "memar/identifier/32byte"
	error_p "memar/process/error/protocol"
)

// DataType embed to provide some methods when not need to implement by others!
type DataType struct {
	uuid.Generated
}

func (dt *DataType) Init(mediatype string) (err error_p.Error) {
	dt.Generated.NewHashString(mediatype)
	return
}

//memar:impl memar/datatype/protocol.DataType_ID
func (dt *DataType) DataTypeID() datatype_p.ID { return dt.Generated.ID() }
func (dt *DataType) DataTypeID_Base64() string { return "" }

//memar:impl memar/datatype/protocol.DataType_Details
func (dt *DataType) LifeCycle() datatype_p.LifeCycle      { return datatype_p.LifeCycle_Unset }
func (dt *DataType) ReferenceURI() string                 { return "" }
func (dt *DataType) IssueDate() string                    { return "" }
func (dt *DataType) ExpiryDate() string                   { return "" }
func (dt *DataType) ExpireInFavorOf() datatype_p.DataType { return nil }

//memar:impl memar/protocol.Detail
func (dt *DataType) Domain() string         { return "" }
func (dt *DataType) Summary() string        { return "" }
func (dt *DataType) Overview() string       { return "" }
func (dt *DataType) UserActionNote() string { return "" }
func (dt *DataType) DevActionNote() string  { return "" }
func (dt *DataType) TAGS() []string         { return []string{} }

//memar:impl memar/protocol.Quiddity
func (dt *DataType) Name() string         { return "" }
func (dt *DataType) Abbreviation() string { return "" }
func (dt *DataType) Aliases() []string    { return []string{} }
