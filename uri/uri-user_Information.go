/* For license and copyright information please see the LEGAL file in the code repository */

package uri

import (
	"strings"
)

// UserInformation store user information part of an URI.
// https://datatracker.ietf.org/doc/html/rfc3986#section-3.2.1
type UserInformation struct {
	userinfo string //
	username string //
	password string //
}

func (u *UserInformation) Init(ui string) {
	u.userinfo = ui
	var i = strings.IndexByte(ui, Colon)
	if i >= 0 {
		u.username = ui[:i]
		u.password = ui[i+1:]
	} else {
		u.userinfo = ui
	}
}
func (u *UserInformation) Reinit() {
	u.userinfo = ""
	u.username = ""
	u.password = ""
}
func (u *UserInformation) Deinit() {}

func (u *UserInformation) Userinfo() string { return u.userinfo }
func (u *UserInformation) Username() string { return u.username }
func (u *UserInformation) Password() string { return u.password }

func (u *UserInformation) SetUserinfo(ui string) { u.userinfo = ui }
func (u *UserInformation) SetUsername(un string) { u.username = un }
func (u *UserInformation) SetPassword(p string)  { u.password = p }
