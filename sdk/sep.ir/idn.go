/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	"crypto/tls"

	"../../achaemenid"
	etime "../../earth-time"
	er "../../error"
	"../../http"
	"../../log"
)

const (
	idnDomain     = "idn.seppay.ir"
	idnDomainPort = "idn.seppay.ir:443"

	idnTokenPath    = "/connect/token"
	idnUserInfoPath = "/connect/userinfo"

	idnTokenType                = "password"
	idnScope                    = "SepCentralPcPos openid"
	idnAuthorizationBasicHeader = "Basic cm9jbGllbnQ6c2VjcmV0" // username:roclient		password:secret
)

type idn struct {
	AccessToken string
	TokenType   string // "token_type" : "Bearer",
	ExpiresIn   etime.Time
}

func (idn *idn) checkAccessToken(username, password string) (err *er.Error) {
	if idn.ExpiresIn <= etime.Now() {
		var req = idnGetTokenReq{
			Scope:     idnScope,
			Username:  username,
			Password:  password,
			GrantType: idnTokenType,
		}
		err = idn.idnGetToken(&req)
	}
	return
}

func (idn *idn) doHTTP(httpReq *http.Request) (httpRes *http.Response, err *er.Error) {
	var goErr error
	var tlsConn *tls.Conn
	var tlsConf = tls.Config{
		ServerName: idnDomain,
	}
	tlsConn, goErr = tls.Dial("tcp", idnDomainPort, &tlsConf)
	if goErr != nil {
		err = achaemenid.ErrNoConnection
		return
	}

	var req []byte = httpReq.Marshal()
	_, goErr = tlsConn.Write(req)
	if goErr != nil {
		err = achaemenid.ErrSendRequest
		return
	}

	var res = make([]byte, 2048)
	var readSize int
	readSize, goErr = tlsConn.Read(res)
	if err != nil {
		err = achaemenid.ErrReceiveResponse
		return
	}

	httpRes = http.MakeNewResponse()
	err = httpRes.UnMarshal(res)
	if err != nil {
		return
	}

	var contentLength = httpRes.Header.GetContentLength()
	// TODO::: check content length not
	if contentLength > 0 && len(httpRes.Body) == 0 {
		httpRes.Body = make([]byte, contentLength)
		readSize, goErr = tlsConn.Read(httpRes.Body)
		if goErr != nil {
			return
		}
		if readSize != int(contentLength) {
			// err =
			// return
		}
	}

	if log.DeepDebugMode {
		log.DeepDebug("sep.ir - Send msg to:::", httpReq.URI.Raw, httpReq.Header, string(httpReq.Body))
		log.DeepDebug("sep.ir - Received msg from:::", httpRes.ReasonPhrase, httpRes.Header, string(httpRes.Body))
	}
	return
}
