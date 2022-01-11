/* For license and copyright information please see LEGAL file in repository */

package user

import "../protocol"

// CheckUserType check and return error if user not exist in base users
func CheckUserType(baseUsers, user protocol.UserType) (exist bool) {
	return user&baseUsers == user
	// err = ErrUserNotAllow
}
