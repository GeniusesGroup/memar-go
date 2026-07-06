/* For license and copyright information please see the LEGAL file in the code repository */

package uri

import (
	"strings"

	error_p "memar/process/error/protocol"
	string_p "memar/codec/string/protocol"
	"memar/codec/string/utf8"
)

// AU store authority part of an URI.
type AU[STR string_p.String] struct {
	authority       STR // host [ ":" port ]
	UserInformation     //
	host            STR // host without port if any exist in authority
	port            STR //
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (a *AU[STR]) Init(au string) (err error_p.Error) {
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
func (a *AU[STR]) Reinit() (err error_p.Error) {
	a.authority = ""
	a.UserInformation.Reinit()
	a.host = ""
	a.port = ""
	return
}
func (a *AU[STR]) Deinit() (err error_p.Error) {
	return
}

func (a *AU[STR]) Authority() string_p.String { return a.authority }
func (a *AU[STR]) Host() string_p.String      { return a.host }
func (a *AU[STR]) Port() string_p.String      { return a.port }

func (a *AU[STR]) SetAuthority(au string) { a.authority = au }
func (a *AU[STR]) SetHost(h string)       { a.host = h }
func (a *AU[STR]) SetPort(p string)       { a.port = p }
