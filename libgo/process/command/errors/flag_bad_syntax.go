/* For license and copyright information please see the LEGAL file in the code repository */

package command_errs

import (
	"memar/computer/datatype"
	error_p "memar/process/error/protocol"
)

var FlagBadSyntax flagBadSyntax

type flagBadSyntax struct {
	datatype.DataType
}

func (dt *flagBadSyntax) Init() (err error_p.Error) {
	// err = dt.Err.Init("domain/memar.scm.geniuses.group; package=command; type=error; name=flag_bad_syntax")
	// if err != nil {
	// 	return
	// }
	// err = errors.Register(dt)
	return
}
