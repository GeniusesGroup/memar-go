/* For license and copyright information please see the LEGAL file in the code repository */

package unix

import (
	string_p "memar/codec/string/protocol"
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
	time_p "memar/time/protocol"
)

var Epoch = epoch{}

type epoch struct{}

//memar:impl memar/datatype/protocol.Field_ID
func (self *epoch) DataTypeID() datatype_p.ID { return 0 }

//memar:impl memar/datatype/protocol.Identifiers
func (self *epoch) DataTypeID_Base64() string { return "" }

//memar:impl memar/datatype/protocol.Detail
func (self *epoch) DevActionNote() string  { return "" }
func (self *epoch) Domain() string         { return "" }
func (self *epoch) Overview() string       { return "" }
func (self *epoch) Summary() string        { return "" }
func (self *epoch) TAGS() []string         { return nil }
func (self *epoch) UserActionNote() string { return "" }

//memar:impl memar/datatype/protocol.Quiddity
func (self *epoch) Abbreviation() string { return "" }
func (self *epoch) Aliases() []string    { return nil }
func (self *epoch) Name() string         { return "" }

//memar:impl memar/datatype/protocol.DataType_Details
func (self *epoch) ExpireInFavorOf() datatype_p.DataType { return nil }
func (self *epoch) ExpiryDate() string                   { return "" }
func (self *epoch) IssueDate() string                    { return "" }
func (self *epoch) ReferenceURI() string                 { return "" }

//memar:impl memar/datatype/protocol.Field_LifeCycle
func (self *epoch) LifeCycle() datatype_p.LifeCycle { return datatype_p.LifeCycle_PreAlpha }

//memar:impl memar/codec/string/protocol.Stringer_To
func (self *epoch) ToString() (str string_p.String, err error_p.Error) {
	return
}

func (self *epoch) Equivalence(with time_p.Epoch) bool { return self.DataTypeID() == with.DataTypeID() }
