/* For license and copyright information please see LEGAL file in repository */

package protocol

type UserID interface {
	UUID() [16]byte
	ExistenceTime() TimeUnixMilli
	Type() UserType
	ID() uint64
}

// UserType indicate connection user type
// Set desire user by:
// User logical operator OR| to add many types together e.g. (UserType_Person|UserType_Thing)
// User logical operator XOR^ to remove a UserType from base e.g. (UserType_All^UserType_Guest)
// User logical operator NOT^ to reverse a UserType to accept all expect it self! e.g. (^UserType_Guest)
type UserType uint8

// User Type
const (
	UserType_Unset   UserType = 0b00000000 // Don't know anything about user yet e.g. connection not fully established yet.
	UserType_Guest   UserType = 0b00000001 // Unknown user type
	UserType_Person  UserType = 0b00000010 // a human being
	UserType_Thing   UserType = 0b00000100 // any device or robot with specific intelligence - AI
	UserType_Org     UserType = 0b00001000
	UserType_App     UserType = 0b00010000 // same app with difference instance
	UserType_Society UserType = 0b00100000 //
	// UserType_...      UserType = 0b00100000
	// UserType_...      UserType = 0b01000000
	// UserType_...      UserType = 0b10000000
	UserType_Known UserType = 0b11111110 // To not indicate user type but not guest one
	UserType_All   UserType = 0b11111111
)
