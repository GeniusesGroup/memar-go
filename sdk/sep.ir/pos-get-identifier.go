/* For license and copyright information please see LEGAL file in repository */

package sep

import (
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

	var serverReq []byte = httpReq.Marshal()
	var serverRes []byte
	serverRes, err = pos.sendHTTPRequest(serverReq)
	if err != nil {
		return
	}

	if log.DevMode {
		log.Debug("sep.ir - Send msg to /v1/PcPosTransaction/ReciveIdentifier ::\n", string(serverReq))
		log.Debug("sep.ir - Received msg from /v1/PcPosTransaction/ReciveIdentifier::\n", string(serverRes))
	}

	var httpRes = http.MakeNewResponse()
	err = httpRes.UnMarshal(serverRes)
	if err != nil {
		return
	}

	switch httpRes.StatusCode {
	case http.StatusBadRequestCode:
		err = ErrBadRequest
		return
	case http.StatusUnauthorizedCode:
		log.Warn("Authorization failed with sep services! Double check account username & password")
	case http.StatusInternalServerErrorCode:
		err = ErrInternalError
		return
	}

	var res = posGetIdentifierRes{}
	err = res.jsonDecoder(httpRes.Body)
	if err != nil {
		err = ErrBadResponse
		return
	}
	err = getErrorByResCode(res.ErrorCode)
	Identifier = res.Identifier
	return
}

func (res *posGetIdentifierRes) jsonDecoder(buf []byte) (err *er.Error) {
	var decoder = json.DecoderUnsafe{
		Buf: buf,
	}
	var keyName string
	for len(decoder.Buf) > 2 {
		keyName = decoder.DecodeKey()
		switch keyName {
		case "":
			return
		case "IsSuccess":
			res.IsSuccess, err = decoder.DecodeBool()
			if err != nil {
				return
			}
		case "ErrorCode":
			res.ErrorCode, err = decoder.DecodeInt64()
			if err != nil {
				return
			}
		case "ErrorDescription":
			res.ErrorDescription, err = decoder.DecodeString()
			if err != nil {
				return
			}

		case "Data":
			continue
		case "Identifier":
			res.Identifier, err = decoder.DecodeString()
			if err != nil {
				return
			}

		default:
			err = decoder.NotFoundKey()
			if err != nil {
				return
			}
		}
	}
	return
}
