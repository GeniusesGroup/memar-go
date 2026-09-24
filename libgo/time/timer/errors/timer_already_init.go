/* For license and copyright information please see the LEGAL file in the code repository */

package timer_errs

import (
	"memar/computer/datatype"
	error_p "memar/process/error/protocol"
	"memar/process/errors"
)

var TimerAlreadyInit timerAlreadyInit

type timerAlreadyInit struct{ datatype.DataType }

func (dt *timerAlreadyInit) Init() (err error_p.Error) {
	err = dt.Err.Init(dt.MediaType())
	if err != nil {
		return
	}
	err = errors.Register(dt)
	return
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (dt *timerAlreadyInit) MediaType() string {
	return domainBaseMediatype + "already_initialized"
}
