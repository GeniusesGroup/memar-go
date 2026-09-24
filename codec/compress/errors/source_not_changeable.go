/* For license and copyright information please see the LEGAL file in the code repository */

package errs

import (
	"memar/computer/datatype"
	er "memar/process/error"
	error_p "memar/process/error/protocol"
)

var SourceNotChangeable sourceNotChangeable

type sourceNotChangeable struct{ datatype.DataType }

func (self *sourceNotChangeable) Init() (err error_p.Error) {
	// err = dt.Init(domainBaseMediatype + "source-not-changeable")
	// if err != nil {
	// 	return
	// }
	// err = errors.Register(dt)
	return
}

func (self *sourceNotChangeable) Equivalence(with error_p.Error) bool { return er.IsEqual(self, with) }
