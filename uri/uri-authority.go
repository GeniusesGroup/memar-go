/* For license and copyright information please see the LEGAL file in the code repository */

package uri

import "strings"

type AU = Authority

// Authority store authority part of an URI.
type Authority struct {
	authority       string // host [ ":" port ]
	UserInformation        //
	host            string // host without port if any exist in authority
	port            string //
}

func (a *Authority) Init(au string) {
	if au == "" {
		return
	}

	a.authority = au

	var i = strings.IndexByte(au, AtSign)
	if i > 0 {
		a.UserInformation.Init(au[:i])
		au = au[:i+1]
	}
	// TODO::: change below to respect RFC
	i = strings.IndexByte(au, Colon)
	if i > 0 {
		a.host = au[:i]
		a.port = au[i+1:]
	} else {
		a.host = au
	}
}
func (a *Authority) Reinit() {
	a.authority = ""
	a.UserInformation.Reinit()
	a.host = ""
	a.port = ""
}
func (a *Authority) Deinit() {}

func (a *Authority) Authority() string { return a.authority }
func (a *Authority) Host() string      { return a.host }
func (a *Authority) Port() string      { return a.port }

func (a *Authority) SetAuthority(au string) { a.authority = au }
func (a *Authority) SetHost(h string)       { a.host = h }
func (a *Authority) SetPort(p string)       { a.port = p }
