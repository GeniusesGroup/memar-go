/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type UserUUID = UUID

type UserID interface {
	UUID() UserUUID

	// Below fields extract from(part of) above UUID
	ExistenceTime() Time
	Type() UserType
	ID() [3]byte
}

// UserType indicate connection user type
// Set desire user by:
// User logical operator OR| to add many types together e.g. (UserType_Person|UserType_Thing)
// User logical operator XOR^ to remove a UserType from base e.g. (UserType_All^UserType_Guest)
// User logical operator NOT^ to reverse a UserType to accept all expect it self! e.g. (^UserType_Guest)
type UserType byte

// User Type
const (
	UserType_Unset   UserType = 0b00000000 //  Don't know anything about user yet e.g. connection not fully established yet.
	UserType_Guest   UserType = 0b00000001 // (1 << 0) Unknown user type
	UserType_Person  UserType = 0b00000010 // (1 << 1) a human being
	UserType_Thing   UserType = 0b00000100 // (1 << 2) any device or robot with specific intelligence - AI
	UserType_Org     UserType = 0b00001000 // (1 << 3) other applications
	UserType_App     UserType = 0b00010000 // (1 << 4) same app with difference instance(node)
	UserType_Society UserType = 0b00100000 // (1 << 5)
	// UserType_...      UserType = 0b01000000 // (1 << 6)
	// UserType_...      UserType = 0b10000000 // (1 << 7)
	UserType_Known UserType = 0b11111110 // To not indicate user type but not guest one
	UserType_All   UserType = 0b11111111
)
