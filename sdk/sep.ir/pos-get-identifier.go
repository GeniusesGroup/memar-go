/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	"../../achaemenid"
	er "../../error"
	"../../http"
	"../../json"
	"../../log"
)

type posGetIdentifierRes struct {
	IsSuccess        bool
	ErrorCode        int64
	ErrorDescription string

	Identifier string
}

func (pos *POS) posGetIdentifier() (Identifier string, err *er.Error) {
	err = pos.IDN.checkAccessToken(pos.Username, pos.Password)
	if err != nil {
		return
	}

	var httpReq = http.MakeNewRequest()
	httpReq.Method = http.MethodPOST
	httpReq.URI.Path = posPathReceiveIdentifier
	httpReq.Version = http.VersionHTTP11

	httpReq.Header.Set(http.HeaderKeyHost, posDomain)
	httpReq.Header.Set(http.HeaderKeyAcceptContent, "application/json")
	httpReq.Header.Set(http.HeaderKeyContentType, "application/json")
	httpReq.Header.Set(http.HeaderKeyAuthorization, pos.IDN.TokenType+" "+pos.IDN.AccessToken)

	// Set some other header data
	httpReq.SetContentLength()
	httpReq.Header.Set(http.HeaderKeyServer, http.DefaultServer)

	var httpRes *http.Response
	httpRes, err = pos.doHTTP(httpReq)
	if err != nil {
		if err.Equal(achaemenid.ErrReceiveResponse) {
			// TODO::: check order
		}
		return
	}

	switch httpRes.StatusCode {
	case http.StatusBadRequestCode:
		err = achaemenid.ErrBadRequest
		return
	case http.StatusUnauthorizedCode:
		log.Warn("Authorization failed with sep services! Double check account username & password")
	case http.StatusInternalServerErrorCode:
		err = achaemenid.ErrInternalError
		return
	}

	switch httpRes.Header.GetContentType().SubType {
	case http.ContentTypeMimeSubTypeHTML:
		err = ErrPOSInternalError
		return
	case http.ContentTypeMimeSubTypeJSON:
		var res = posGetIdentifierRes{}
		err = res.jsonDecoder(httpRes.Body)
		if err != nil {
			err = achaemenid.ErrBadResponse
			return
		}
		if !res.IsSuccess && res.ErrorCode == 0 {
			err = ErrPOSInternalError
			return
		}
		err = getErrorByResCode(res.ErrorCode)
		Identifier = res.Identifier
	}
	return
}

func (res *posGetIdentifierRes) jsonDecoder(buf []byte) (err *er.Error) {
	var decoder = json.DecoderUnsafe{
		Buf: buf,
	}
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "":
			return
		case "IsSuccess":
			res.IsSuccess, err = decoder.DecodeBool()
		case "ErrorCode":
			res.ErrorCode, err = decoder.DecodeInt64()
		case "ErrorDescription":
			res.ErrorDescription, err = decoder.DecodeString()

		case "Data":
			continue
		case "Identifier":
			res.Identifier, err = decoder.DecodeString()

		default:
			err = decoder.NotFoundKey()
		}

		if len(decoder.Buf) < 3 {
			return
		}
	}
	return
}
