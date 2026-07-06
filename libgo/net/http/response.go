/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"strconv"

	"memar/errors"
	errs "memar/net/http/errors"
	"memar/protocol"
)

// Response is represent response protocol structure.
// https://tools.ietf.org/html/rfc2616#section-6
type Response struct {
	PseudoHeader_Response
	Header
	body
}

//memar:impl memar/protocol.ObjectLifeCycle
func (r *Response) Init() (err protocol.Error) {
	err = r.PseudoHeader_Response.Init()
	if err != nil {
		return
	}
	err = r.Header.Init()
	if err != nil {
		return
	}
	err = r.body.Init()
	return
}
func (r *Response) Reinit() (err protocol.Error) {
	err = r.PseudoHeader_Response.Reinit()
	if err != nil {
		return
	}
	err = r.Header.Reinit()
	if err != nil {
		return
	}
	err = r.body.Reinit()
	return
}
func (r *Response) Deinit() (err protocol.Error) {
	err = r.PseudoHeader_Response.Deinit()
	if err != nil {
		return
	}
	err = r.Header.Deinit()
	if err != nil {
		return
	}
	err = r.body.Deinit()
	return
}

// GetStatusCode get status code as uit16
func (r *Response) GetStatusCode() (code uint16, err protocol.Error) {
	// TODO::: don't use strconv for such simple task
	var c, goErr = strconv.ParseUint(r.StatusCode(), 10, 16)
	if goErr != nil {
		return 0, &errs.ErrParseStatusCode
	}
	return uint16(c), nil
}

// GetError return related protocol.Error in header of the Response
func (r *Response) GetError() (err protocol.Error) {
	var errIDString = r.Header_Get(HeaderKey_ErrorID)
	var errID, _ = strconv.ParseUint(errIDString, 10, 64)
	if errID == 0 {
		return
	}
	err = errors.GetByID(protocol.ID(errID))
	return
}

// SetError set given protocol.Error to header of the response
func (r *Response) SetError(err protocol.Error) {
	r.Header_Set(HeaderKey_ErrorID, err.IDasString())
}

// Redirect set given status and target location to the response
// httpRes.Redirect(http.StatusMovedPermanentlyCode, http.StatusMovedPermanentlyPhrase, "http://www.google.com/")
func (r *Response) Redirect(code, phrase string, target string) {
	r.SetStatus(code, phrase)
	r.Header_Set(HeaderKey_Location, target)
}

func (r *Response) SetStatusByError(err protocol.Error) {
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
