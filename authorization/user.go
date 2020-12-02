/* For license and copyright information please see LEGAL file in repository */

package authorization

import er "../error"

// UserType indicate connection user type
type UserType uint8

// User Type
const (
	UserTypeUnset      UserType = 0b00000000 // None
	UserTypeGuest      UserType = 0b00000001
	UserTypeRegistered UserType = 0b00000010 // To not indicate user type
	UserTypePerson     UserType = 0b00000100
	UserTypeOrg        UserType = 0b00001000
	UserTypeApp        UserType = 0b00010000
	UserTypeAI         UserType = 0b00100000 // Robots
	UserTypeAll        UserType = 0b11111111
)

// Set given user to given UserType!
// User logical operator OR| to add many types together e.g. ut.Set(UserTypePerson|UserTypeAI)
// User logical operator XOR^ to remove a UserType from base e.g. ut.Set(UserTypeAll^UserTypeAI)
// User logical operator NOT^ to reverse a UserType to accept all expect it self! e.g. ut.Set(^UserTypeAI)
func (ut UserType) Set(user UserType) {
	ut = user
}

// Check check given user type exist in given UserType!
func (ut UserType) Check(user UserType) (err *er.Error) {
	if user&ut != user {
		err = ErrUserNotAllow
	}
	return
}

// CheckReverse check given users type exist in given UserType!
func (ut UserType) CheckReverse(users UserType) (err *er.Error) {
	if ut&users != ut {
		err = ErrUserNotAllow
	}
	return
}
