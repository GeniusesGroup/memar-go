/* For license and copyright information please see the LEGAL file in the code repository */

package uri

import (
	"strings"

	"libgo/protocol"
	"libgo/utf8"
)

type AU = Authority

// Authority store authority part of an URI.
type Authority struct {
	authority       string // host [ ":" port ]
	UserInformation        //
	host            string // host without port if any exist in authority
	port            string //
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (a *Authority) Init(au string) (err protocol.Error) {
	if au == "" {
		return
	}

	a.authority = au

	var i = strings.IndexByte(au, sign_AtSign)
	if i > 0 {
		a.UserInformation.Init(au[:i])
		au = au[:i+1]
	}
	// TODO::: change below to respect RFC to respect IPv6 address
	a.host, a.port, _ = utf8.CutByte(au, sign_Colon)
	return
}
func (a *Authority) Reinit() (err protocol.Error) {
	a.authority = ""
	a.UserInformation.Reinit()
	a.host = ""
	a.port = ""
	return
}
func (a *Authority) Deinit() (err protocol.Error) {
	return
}

func (a *Authority) Authority() string { return a.authority }
func (a *Authority) Host() string      { return a.host }
func (a *Authority) Port() string      { return a.port }

func (a *Authority) SetAuthority(au string) { a.authority = au }
func (a *Authority) SetHost(h string)       { a.host = h }
func (a *Authority) SetPort(p string)       { a.port = p }
