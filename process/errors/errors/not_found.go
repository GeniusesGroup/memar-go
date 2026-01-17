/* For license and copyright information please see the LEGAL file in the code repository */

package errors_errs

import (
	"memar/computer/datatype"
	error_p "memar/process/error/protocol"
)

var (
	ErrNotFound errNotFound
)

type errNotFound struct {
	datatype.DataType
}

func (dt *errNotFound) Init() (err error_p.Error) {
	// CANn't import `errors` package here due to import cycle problems.
	// err = Register(dt)
	return
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (dt *errNotFound) MediaType() string {
	return domainBaseMediatype + "not_found"
}
