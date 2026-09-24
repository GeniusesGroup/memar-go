/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	"memar/computer/datatype"
	datatype_p "memar/computer/datatype/protocol"
	er "memar/process/error"
	error_p "memar/process/error/protocol"
	"memar/process/errors"
)

var GuestMaxReached guestMaxReached

type guestMaxReached struct {
	datatype.DataType
}

func (self *guestMaxReached) Init() (err error_p.Error) {
	err = errors.Register(self)
	return
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (self *guestMaxReached) MediaType() string { return domainBaseMediatype + "guest_max_reached" }

//memar:impl memar/process/error/protocol.Internal
func (self *guestMaxReached) Internal() bool { return false }

//memar:impl memar/datatype/protocol.Identifiers
func (self *guestMaxReached) DataTypeID() datatype_p.ID { return 0 }
func (self *guestMaxReached) DataTypeID_Base64() string { return "" }

//memar:impl memar/datatype/protocol.DataType_Details
func (self *guestMaxReached) LifeCycle() datatype_p.LifeCycle      { return datatype_p.LifeCycle_Unset }
func (self *guestMaxReached) ReferenceURI() string                 { return "" }
func (self *guestMaxReached) IssueDate() string                    { return "" }
func (self *guestMaxReached) ExpiryDate() string                   { return "" }
func (self *guestMaxReached) ExpireInFavorOf() datatype_p.DataType { return nil }

//memar:impl memar/math/logic/protocol.Equivalence
func (self *guestMaxReached) Equal(with error_p.Error) bool { return er.IsEqual(self, with) }

//memar:impl memar/computer/adt/protocol.Empty
func (self *guestMaxReached) IsEmpty() bool { return self.IsNil() && self.IsNull() }

//memar:impl memar/computer/adt/protocol.Nil
func (self *guestMaxReached) IsNil() bool { return self == nil }

//memar:impl memar/computer/adt/protocol.Null
func (self *guestMaxReached) IsNull() bool { return self.DataTypeID() == 0 }
