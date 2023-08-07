/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	er "memar/error"
	"memar/protocol"
)

var ErrDuplicateIdentifier errDuplicateIdentifier

type errDuplicateIdentifier struct{ er.Err }

func (dt *errDuplicateIdentifier) Init() (err protocol.Error) {
	err = dt.Err.Init("domain/memar.scm.geniuses.group; package=errors; type=error; name=duplicate_identifier")
	if err != nil {
		return
	}
	err = protocol.App.RegisterError(dt)
	return
}

// This condition will just be true in the dev phase.
// panic("Error must have valid ID to save it in platform errors pools. Initialize inner e.MediaType.Init() first if use memar/service package.")

// This condition will just be true in the dev phase.
// panic("Error id exist and used for other Error. Check it now for bad media-type set or collision occurred" +
// "\nExiting error >> " + e.poolByID[errID].ToString() +
// "\nNew error >> " + err.ToString())
