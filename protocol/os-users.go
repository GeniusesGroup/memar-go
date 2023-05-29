/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// OperatingSystem_User introduce all data about an applications
type OperatingSystem_User interface {
	ActiveUser() User
	Users() []User
}

type User interface {
	ID() UserID
	SocietyID() SocietyID
}
