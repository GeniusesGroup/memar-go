/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	error_p "memar/process/error/protocol"
	"memar/net/uri"
	string_p "memar/codec/string/protocol"
)

type PseudoHeader_Request[STR string_p.String] struct {
	method       STR
	uri.URI[STR] // Exported field to let consumers use other methods that uri_p.URI
	version      STR
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (r *PseudoHeader_Request[STR]) Init() (err error_p.Error) {
	// err = r.U.Init()
	return
}
func (r *PseudoHeader_Request[STR]) Reinit() (err error_p.Error) {
	// err = r.method.Reinit()
	// err = r.version.Reinit()
	// err = r.URI.Reinit()
	return
}
func (r *PseudoHeader_Request[STR]) Deinit() (err error_p.Error) {
	err = r.URI.Deinit()
	return
}

//memar:impl memar/net/http/protocol.PseudoHeader_Request
func (r *PseudoHeader_Request[STR]) Method() string_p.String  { return r.method }
func (r *PseudoHeader_Request[STR]) Version() string_p.String { return r.version }

func (r *PseudoHeader_Request[STR]) SetMethod(method STR)   { r.method = method }
func (r *PseudoHeader_Request[STR]) SetVersion(version STR) { r.version = version }
