/* For license and copyright information please see the LEGAL file in the code repository */

package datatype

import (
	datatype_p "memar/computer/datatype/protocol"
)

// DataType embed to provide some methods when not need to implement by others.
// Usually JUST embed `Detail` and let the capsule implement below methods.
type DataType struct {
	Identifiers
	Detail
}

//memar:impl memar/datatype/protocol.Field_LifeCycle
func (self *DataType) LifeCycle() datatype_p.LifeCycle { return datatype_p.LifeCycle_Unset }

//memar:impl memar/datatype/protocol.Details
func (self *DataType) ReferenceURI() string                 { return "" }
func (self *DataType) IssueDate() string                    { return "" }
func (self *DataType) ExpiryDate() string                   { return "" }
func (self *DataType) ExpireInFavorOf() datatype_p.DataType { return nil }
