/* For license and copyright information please see the LEGAL file in the code repository */

package utc_epoch

import (
	string_p "memar/codec/string/protocol"
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
	time_p "memar/time/protocol"
)

var DayInWeek = day_in_week{}

// It is use for any repetitive time that need to point from a day
// without need to indicate specific day, e.g. `8a.m`, `5p.m`
type day_in_week struct{}

//memar:impl memar/datatype/protocol.Field_ID
func (self *day_in_week) DataTypeID() datatype_p.ID { return 0 }

//memar:impl memar/datatype/protocol.Identifiers
func (self *day_in_week) DataTypeID_Base64() string { return "" }

//memar:impl memar/datatype/protocol.Detail
func (self *day_in_week) DevActionNote() string  { return "" }
func (self *day_in_week) Domain() string         { return "" }
func (self *day_in_week) Overview() string       { return "" }
func (self *day_in_week) Summary() string        { return "" }
func (self *day_in_week) TAGS() []string         { return nil }
func (self *day_in_week) UserActionNote() string { return "" }

//memar:impl memar/datatype/protocol.Quiddity
func (self *day_in_week) Abbreviation() string { return "" }
func (self *day_in_week) Aliases() []string    { return nil }
func (self *day_in_week) Name() string         { return "" }

//memar:impl memar/datatype/protocol.DataType_Details
func (self *day_in_week) ExpireInFavorOf() datatype_p.DataType { return nil }
func (self *day_in_week) ExpiryDate() string                   { return "" }
func (self *day_in_week) IssueDate() string                    { return "" }
func (self *day_in_week) ReferenceURI() string                 { return "" }

//memar:impl memar/datatype/protocol.Field_LifeCycle
func (self *day_in_week) LifeCycle() datatype_p.LifeCycle { return datatype_p.LifeCycle_PreAlpha }

//memar:impl memar/codec/string/protocol.Stringer_To
func (self *day_in_week) ToString() (str string_p.String, err error_p.Error) {
	return
}

func (self *day_in_week) Equivalence(with time_p.Epoch) bool {
	return self.DataTypeID() == with.DataTypeID()
}
