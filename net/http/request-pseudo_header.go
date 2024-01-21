/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"memar/net/uri"
	"memar/protocol"
)

type PseudoHeader_Request struct {
	method  string
	U       uri.URI // Exported field to let consumers use other methods that protocol.HTTPHeader
	version string
}

//memar:impl memar/protocol.ObjectLifeCycle
func (r *PseudoHeader_Request) Init() (err protocol.Error) {
	// err = r.U.Init()
	return
}
func (r *PseudoHeader_Request) Reinit() (err protocol.Error) {
	r.method = ""
	r.version = ""
	err = r.U.Reinit()
	return
}
func (r *PseudoHeader_Request) Deinit() (err protocol.Error) {
	err = r.U.Deinit()
	return
}

//memar:impl memar/protocol.HTTP_PseudoHeader_Request
func (r *PseudoHeader_Request) Method() string            { return r.method }
func (r *PseudoHeader_Request) URI() protocol.URI         { return &r.U }
func (r *PseudoHeader_Request) Version() string           { return r.version }
func (r *PseudoHeader_Request) SetMethod(method string)   { r.method = method }
func (r *PseudoHeader_Request) SetVersion(version string) { r.version = version }
