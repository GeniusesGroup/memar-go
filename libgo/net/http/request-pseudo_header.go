/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"memar/net/uri"
	error_p "memar/process/error/protocol"
)

type PseudoHeader_Request struct {
	method  string
	uri.URI // Exported field to let consumers use other methods that uri_p.URI
	version string
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (r *PseudoHeader_Request) Init() (err error_p.Error) {
	// err = r.U.Init()
	return
}
func (r *PseudoHeader_Request) Reinit() (err error_p.Error) {
	// err = r.method.Reinit()
	// err = r.version.Reinit()
	// err = r.URI.Reinit()
	return
}
func (r *PseudoHeader_Request) Deinit() (err error_p.Error) {
	err = r.URI.Deinit()
	return
}

//memar:impl memar/libgo/net/http/protocol.PseudoHeader_Request
func (r *PseudoHeader_Request) Method() string  { return r.method }
func (r *PseudoHeader_Request) Version() string { return r.version }

func (r *PseudoHeader_Request) SetMethod(method string)   { r.method = method }
func (r *PseudoHeader_Request) SetVersion(version string) { r.version = version }
