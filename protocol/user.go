/* For license and copyright information please see LEGAL file in repository */

package protocol

type UserID interface {
	UUID() [16]byte
	ExistenceTime() int64 // Unix milli second time
	Type() UserType
	ID() uint64
}

// UserType indicate connection user type
// Set desire user by:
// User logical operator OR| to add many types together e.g. ut.Set(UserTypePerson|UserTypeThing)
// User logical operator XOR^ to remove a UserType from base e.g. ut.Set(UserTypeAll^UserTypeThing)
// User logical operator NOT^ to reverse a UserType to accept all expect it self! e.g. ut.Set(^UserTypeThing)
type UserType uint8

// User Type
const (
	UserTypeUnset      UserType = 0b00000000 // None - usually use for unknown users like guest one
	UserTypePerson     UserType = 0b00000001 // a human being
	UserTypeThing      UserType = 0b00000010 // any device or robot with specific intelligence - AI
	UserTypeOrg        UserType = 0b00000100
	UserTypeApp        UserType = 0b00001000 // same app with difference instance
	UserTypeSociety    UserType = 0b00010000 // 
	// UserType...      UserType = 0b00100000
	// UserType...      UserType = 0b01000000
	// UserType...      UserType = 0b10000000
	// UserTypeRegistered UserType = 0b11111110 // To not indicate user type but not guest one
	UserTypeAll        UserType = 0b11111111
)
