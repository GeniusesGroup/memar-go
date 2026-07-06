/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	error_p "memar/process/error/protocol"
	"memar/math/integer"
	errs "memar/net/http/errors"
	string_p "memar/codec/string/protocol"
)

type PseudoHeader_Response[STR string_p.String] struct {
	version      STR
	statusCode   STR
	reasonPhrase STR
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (r *PseudoHeader_Response[STR]) Init() (err error_p.Error) {
	return
}
func (r *PseudoHeader_Response[STR]) Reinit() (err error_p.Error) {
	// r.version.Reinit()
	// r.statusCode.Reinit()
	// r.reasonPhrase.Reinit()
	return
}
func (r *PseudoHeader_Response[STR]) Deinit() (err error_p.Error) {
	return
}

//memar:impl memar/net/http/protocol.PseudoHeader_Response
func (r *PseudoHeader_Response[STR]) Version() string_p.String      { return r.version }
func (r *PseudoHeader_Response[STR]) StatusCode() string_p.String   { return r.statusCode }
func (r *PseudoHeader_Response[STR]) ReasonPhrase() string_p.String { return r.reasonPhrase }

func (r *PseudoHeader_Response[STR]) SetVersion(version STR) { r.version = version }
func (r *PseudoHeader_Response[STR]) SetStatus(code, phrase STR) {
	r.statusCode = code
	r.reasonPhrase = phrase
}

// GetStatusCode get status code as uit16
func (r *PseudoHeader_Response[STR]) GetStatusCode() (code integer.U16, err error_p.Error) {
	// TODO::: don't use strconv for such simple task
	err = code.FromString(r.StatusCode())
	// var c, goErr = strconv.ParseUint(r.StatusCode(), 10, 16)
	if err != nil {
		return 0, &errs.ErrParseStatusCode
	}
	return
}

func (r *Response[STR]) SetStatusByError(err error_p.Error) {
	if err != nil {
		switch {
		case err.Equal(&errs.ErrParseHeaderTooLarge):
			r.SetStatus(StatusHeaderFieldsTooLargeCode, StatusHeaderFieldsTooLargePhrase)
		case err.Equal(&errs.UnsupportedTransferEncoding):
			r.SetStatus(StatusNotImplementedCode, StatusNotImplementedPhrase)
		default:
			r.SetStatus(StatusBadRequestCode, StatusBadRequestPhrase)
		}
	}
}
