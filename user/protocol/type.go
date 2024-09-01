/* For license and copyright information please see the LEGAL file in the code repository */

package user_p

type Field_UserType interface {
	UserType() Type
}

// Type indicate connection user type
// Set desire user by:
// User logical operator OR| to add many types together e.g. (Type_Person|Type_Thing)
// User logical operator XOR^ to remove a Type from base e.g. (Type_All^Type_Guest)
// User logical operator NOT^ to reverse a Type to accept all expect it self! e.g. (^Type_Guest)
type Type byte

// User Type
const (
	Type_Unset   Type = 0b00000000 //  Don't know anything about user yet e.g. connection not fully established yet.
	Type_Guest   Type = 0b00000001 // (1 << 0) Unknown user type
	Type_Person  Type = 0b00000010 // (1 << 1) a human being
	Type_Thing   Type = 0b00000100 // (1 << 2) any device or robot with specific intelligence - AI
	Type_Org     Type = 0b00001000 // (1 << 3) other applications
	Type_App     Type = 0b00010000 // (1 << 4) same app with difference instance(node)
	Type_Society Type = 0b00100000 // (1 << 5)
	// Type_...      Type = 0b01000000 // (1 << 6)
	// Type_...      Type = 0b10000000 // (1 << 7)
	Type_Known Type = 0b11111110 // To not indicate user type but not guest one
	Type_All   Type = 0b11111111
)
