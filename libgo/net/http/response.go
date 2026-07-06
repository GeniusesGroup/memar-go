/* For license and copyright information please see the LEGAL file in the code repository */

package http

import (
	"strconv"

	"memar/computer/datatype"
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
	"memar/process/errors"
	string_p "memar/codec/string/protocol"
)

// Response is represent response protocol structure.
// https://tools.ietf.org/html/rfc2616#section-6
type Response[STR string_p.String] struct {
	PseudoHeader_Response[STR]
	Header[STR]
	body

	datatype.DataType
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (r *Response[STR]) Init() (err error_p.Error) {
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
func (r *Response[STR]) Reinit() (err error_p.Error) {
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
func (r *Response[STR]) Deinit() (err error_p.Error) {
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

// GetError return related error_p.Error in header of the Response
func (r *Response[STR]) GetError() (err error_p.Error) {
	var errIDString = r.Header_Get(Key_ErrorID)
	var errID, _ = strconv.ParseUint(errIDString, 10, 64)
	if errID == 0 {
		return
	}
	err = errors.GetByID(datatype_p.ID(errID))
	return
}

// SetError set given error_p.Error to header of the response
func (r *Response[STR]) SetError(err error_p.Error) {
	r.Header_Set(Key_ErrorID, err.DataTypeID_Base64())
}

// Redirect set given status and target location to the response
// httpRes.Redirect(http.StatusMovedPermanentlyCode, http.StatusMovedPermanentlyPhrase, "http://www.google.com/")
func (r *Response[STR]) Redirect(code, phrase, target STR) {
	r.SetStatus(code, phrase)
	r.Header_Set(Key_Location, target)
}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (r *Response[STR]) MediaType() string { return "application/http; response" }

//memar:impl memar/protocol.FileExtension
func (r *Response[STR]) FileExtension() string { return "res.http" }

//memar:impl memar/datatype/protocol.DataType_Details
func (r *Response[STR]) LifeCycle() datatype_p.LifeCycle { return datatype_p.LifeCycle_PreAlpha }
func (r *Response[STR]) ReferenceURI() string {
	return "https://www.iana.org/assignments/media-types/application/http"
}
func (r *Response[STR]) IssueDate() string                    { return "" }
func (r *Response[STR]) ExpiryDate() string                   { return "" }
func (r *Response[STR]) ExpireInFavorOf() datatype_p.DataType { return nil }
