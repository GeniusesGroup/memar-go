/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	"../../achaemenid"
	etime "../../earth-time"
	er "../../error"
	"../../http"
	"../../json"
	"../../log"
)

type idnGetTokenReq struct {
	GrantType string
	Username  string
	Password  string
	Scope     string
}

type idnGetTokenRes struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func (idn *idn) idnGetToken(req *idnGetTokenReq) (err *er.Error) {
	var httpReq = http.MakeNewRequest()
	httpReq.Method = http.MethodPOST
	httpReq.URI.Path = idnTokenPath
	httpReq.Version = http.VersionHTTP11

	httpReq.Header.Set(http.HeaderKeyHost, idnDomain)
	httpReq.Header.Set(http.HeaderKeyAcceptContent, "application/json")
	httpReq.Header.Set(http.HeaderKeyContentType, "application/x-www-form-urlencoded")
	httpReq.Header.Set(http.HeaderKeyAuthorization, idnAuthorizationBasicHeader)

	httpReq.Body = req.formEncoder()

	// Set some other header data
	httpReq.SetContentLength()
	httpReq.Header.Set(http.HeaderKeyServer, http.DefaultServer)

	var httpRes *http.Response
	httpRes, err = idn.doHTTP(httpReq)
	if err != nil {
		if err.Equal(achaemenid.ErrReceiveResponse) {
			// TODO::: check order
		}
		return
	}

	switch httpRes.StatusCode {
	case http.StatusBadRequestCode:
		return achaemenid.ErrBadRequest
	case http.StatusUnauthorizedCode:
		log.Warn("Authorization failed with sep services! Double check account username & password")
	case http.StatusInternalServerErrorCode:
		return achaemenid.ErrInternalError
	}

	switch httpRes.Header.GetContentType().SubType {
	case http.ContentTypeMimeSubTypeHTML:
		err = ErrPOSAuthenticationError
		return
	case http.ContentTypeMimeSubTypeJSON:
		var res = idnGetTokenRes{}
		err = res.jsonDecoder(httpRes.Body)
		if err != nil {
			return achaemenid.ErrBadResponse
		}
		idn.AccessToken = res.AccessToken
		idn.ExpiresIn = etime.Now().AddDuration(etime.Duration(res.ExpiresIn))
		idn.TokenType = res.TokenType
	}
	return
}

func (req *idnGetTokenReq) formEncoder() (buf []byte) {
	const fixed = "grant_type=&username=&password=&scope="
	buf = make([]byte, 0, len(fixed)+len(req.GrantType)+len(req.Username)+len(req.Password)+len(req.Scope))
	buf = append(buf, "grant_type="...)
	buf = append(buf, req.GrantType...)
	buf = append(buf, "&username="...)
	buf = append(buf, req.Username...)
	buf = append(buf, "&password="...)
	buf = append(buf, req.Password...)
	buf = append(buf, "&scope="...)
	buf = append(buf, req.Scope...)
	return
}

func (res *idnGetTokenRes) jsonDecoder(buf []byte) (err *er.Error) {
	var decoder = json.Decoder{
		Buf: buf,
	}
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "":
			return
		case "access_token":
			res.AccessToken, err = decoder.DecodeString()
		case "token_type":
			res.TokenType, err = decoder.DecodeString()
		case "expires_in":
			res.ExpiresIn, err = decoder.DecodeInt64()
		case "refresh_token":
			res.RefreshToken, err = decoder.DecodeString()
		default:
			err = decoder.NotFoundKey()
		}

		if len(decoder.Buf) < 3 {
			return
		}
	}
	return
}
