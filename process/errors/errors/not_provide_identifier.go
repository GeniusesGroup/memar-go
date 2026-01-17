/* For license and copyright information please see the LEGAL file in the code repository */

package errors_errs

import (
	"memar/computer/datatype"
	error_p "memar/process/error/protocol"
)

var NotProvideIdentifier notProvideIdentifier

type notProvideIdentifier struct {
	datatype.DataType
}

func (self *notProvideIdentifier) Init() (err error_p.Error) {
	// CANn't import `errors` package here due to import cycle problems.
	// err = errors.Register(dt)
	return
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (self *notProvideIdentifier) MediaType() string {
	return domainBaseMediatype + "not_provide_identifier"
}
