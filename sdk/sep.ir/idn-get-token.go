/* For license and copyright information please see LEGAL file in repository */

package sep

import (
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

	var serverReq []byte = httpReq.Marshal()
	var serverRes []byte
	serverRes, err = idn.sendHTTPRequest(serverReq)
	if err != nil {
		return
	}

	if log.DevMode {
		log.Debug("sep.ir - Send msg to idn.seppay.ir::\n", string(serverReq))
		log.Debug("sep.ir - Received msg from idn.seppay.ir::\n", string(serverRes))
	}

	var httpRes = http.MakeNewResponse()
	err = httpRes.UnMarshal(serverRes)
	if err != nil {
		return
	}

	switch httpRes.StatusCode {
	case http.StatusBadRequestCode:
		return ErrBadRequest
	case http.StatusUnauthorizedCode:
		log.Warn("Authorization failed with sep services! Double check account username & password")
	case http.StatusInternalServerErrorCode:
		return ErrInternalError
	}

	var res = idnGetTokenRes{}
	err = res.jsonDecoder(httpRes.Body)
	if err != nil {
		return ErrBadResponse
	}
	idn.AccessToken = res.AccessToken
	idn.ExpiresIn = etime.Now().AddDuration(etime.Duration(res.ExpiresIn))
	idn.TokenType = res.TokenType
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
	var keyName string
	for len(decoder.Buf) > 2 {
		keyName = decoder.DecodeKey()
		switch keyName {
		case "":
			return
		case "access_token":
			res.AccessToken, err = decoder.DecodeString()
			if err != nil {
				return
			}
		case "token_type":
			res.TokenType, err = decoder.DecodeString()
			if err != nil {
				return
			}
		case "expires_in":
			res.ExpiresIn, err = decoder.DecodeInt64()
			if err != nil {
				return
			}
		case "refresh_token":
			res.RefreshToken, err = decoder.DecodeString()
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
