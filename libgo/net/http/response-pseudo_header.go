/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	errs "memar/libgo/net/http/errors"
	"memar/math/integer"
	error_p "memar/process/error/protocol"
)

type PseudoHeader_Response struct {
	version      string
	statusCode   string
	reasonPhrase string
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (r *PseudoHeader_Response) Init() (err error_p.Error) {
	return
}
func (r *PseudoHeader_Response) Reinit() (err error_p.Error) {
	// r.version.Reinit()
	// r.statusCode.Reinit()
	// r.reasonPhrase.Reinit()
	return
}
func (r *PseudoHeader_Response) Deinit() (err error_p.Error) {
	return
}

//memar:impl memar/libgo/net/http/protocol.PseudoHeader_Response
func (r *PseudoHeader_Response) Version() string      { return r.version }
func (r *PseudoHeader_Response) StatusCode() string   { return r.statusCode }
func (r *PseudoHeader_Response) ReasonPhrase() string { return r.reasonPhrase }

func (r *PseudoHeader_Response) SetVersion(version string) { r.version = version }
func (r *PseudoHeader_Response) SetStatus(code, phrase string) {
	r.statusCode = code
	r.reasonPhrase = phrase
}

// GetStatusCode get status code as uit16
func (r *PseudoHeader_Response) GetStatusCode() (code integer.U16, err error_p.Error) {
	// TODO::: don't use strconv for such simple task
	err = code.FromString(r.StatusCode())
	// var c, goErr = strconv.ParseUint(r.StatusCode(), 10, 16)
	if err != nil {
		return 0, &errs.ErrParseStatusCode
	}
	return
}

func (r *Response) SetStatusByError(err error_p.Error) {
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
