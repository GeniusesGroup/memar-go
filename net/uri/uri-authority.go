/* For license and copyright information please see the LEGAL file in the code repository */

package uri

import (
	"strings"

	"memar/codec/string/utf8"
	error_p "memar/process/error/protocol"
)

// AU store authority part of an URI.
type AU struct {
	authority       string // host [ ":" port ]
	UserInformation        //
	host            string // host without port if any exist in authority
	port            string //
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (a *AU) Init(au string) (err error_p.Error) {
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
func (a *AU) Reinit() (err error_p.Error) {
	a.authority = ""
	a.UserInformation.Reinit()
	a.host = ""
	a.port = ""
	return
}
func (a *AU) Deinit() (err error_p.Error) {
	return
}

func (a *AU) Authority() string { return a.authority }
func (a *AU) Host() string      { return a.host }
func (a *AU) Port() string      { return a.port }

func (a *AU) SetAuthority(au string) { a.authority = au }
func (a *AU) SetHost(h string)       { a.host = h }
func (a *AU) SetPort(p string)       { a.port = p }
