/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"memar/protocol"
)

type PseudoHeader_Response struct {
	version      string
	statusCode   string
	reasonPhrase string
}

//memar:impl memar/protocol.ObjectLifeCycle
func (r *PseudoHeader_Response) Init() (err protocol.Error) {
	return
}
func (r *PseudoHeader_Response) Reinit() (err protocol.Error) {
	r.version = ""
	r.statusCode = ""
	r.reasonPhrase = ""
	return
}
func (r *PseudoHeader_Response) Deinit() (err protocol.Error) {
	return
}

//memar:impl memar/protocol.HTTP_PseudoHeader_Response
func (r *PseudoHeader_Response) Version() string           { return r.version }
func (r *PseudoHeader_Response) StatusCode() string        { return r.statusCode }
func (r *PseudoHeader_Response) ReasonPhrase() string      { return r.reasonPhrase }
func (r *PseudoHeader_Response) SetVersion(version string) { r.version = version }
func (r *PseudoHeader_Response) SetStatus(code, phrase string) {
	r.statusCode = code
	r.reasonPhrase = phrase
}
