/* For license and copyright information please see the LEGAL file in the code repository */

package errors_errs

import (
	"memar/computer/datatype"
	error_p "memar/process/error/protocol"
)

var DuplicateIdentifier duplicateIdentifier

type duplicateIdentifier struct {
	datatype.DataType
}

func (self *duplicateIdentifier) Init() (err error_p.Error) {
	// CANn't import `errors` package here due to import cycle problems.
	// err = Register(dt)
	return
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (self *duplicateIdentifier) MediaType() string {
	return domainBaseMediatype + "duplicate_identifier"
}

// This condition will just be true in the dev phase.
// panic("Error must have valid ID to save it in platform errors pools. Initialize inner e.MediaType.Init() first if use memar/service package.")

// This condition will just be true in the dev phase.
// panic("Error id exist and used for other Error. Check it now for bad media-type set or collision occurred" +
// "\nExiting error >> " + e.poolByID[errID].ToString() +
// "\nNew error >> " + err.ToString())
