/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	"crypto/tls"

	etime "../../earth-time"
	er "../../error"
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

func (idn *idn) getAccessToken(username, password string) (err *er.Error) {
	var req = idnGetTokenReq{
		Scope:     idnScope,
		Username:  username,
		Password:  password,
		GrantType: idnTokenType,
	}
	err = idn.idnGetToken(&req)
	return
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

func (idn *idn) sendHTTPRequest(req []byte) (res []byte, err *er.Error) {
	var goErr error
	var tlsConn *tls.Conn
	var tlsConf = tls.Config{
		ServerName: idnDomain,
	}
	tlsConn, goErr = tls.Dial("tcp", idnDomainPort, &tlsConf)
	if goErr != nil {
		err = ErrNoConnection
		return
	}

	_, goErr = tlsConn.Write(req)
	if goErr != nil {
		err = ErrInternalError
		return
	}

	res = make([]byte, 4096)
	var readSize int
	readSize, goErr = tlsConn.Read(res)
	if err != nil {
		err = ErrInternalError
		return
	}

	return res[:readSize], nil
}
