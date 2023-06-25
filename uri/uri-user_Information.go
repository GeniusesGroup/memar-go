/* For license and copyright information please see the LEGAL file in the code repository */

package uri

import (
	"libgo/protocol"
	"libgo/utf8"
)

// UserInformation store user information part of an URI.
// https://datatracker.ietf.org/doc/html/rfc3986#section-3.2.1
type UserInformation struct {
	userinfo string //
	username string //
	password string //
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (u *UserInformation) Init(ui string) (err protocol.Error) {
	u.userinfo = ui
	u.username, u.password, _ = utf8.CutByte(ui, sign_Colon)
	return
}
func (u *UserInformation) Reinit() (err protocol.Error) {
	u.userinfo = ""
	u.username = ""
	u.password = ""
	return
}
func (u *UserInformation) Deinit() (err protocol.Error) {
	return
}

func (u *UserInformation) Userinfo() string { return u.userinfo }
func (u *UserInformation) Username() string { return u.username }
func (u *UserInformation) Password() string { return u.password }

func (u *UserInformation) SetUserinfo(ui string) { u.userinfo = ui }
func (u *UserInformation) SetUsername(un string) { u.username = un }
func (u *UserInformation) SetPassword(p string)  { u.password = p }
