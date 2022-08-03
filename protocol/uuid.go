/* For license and copyright information please see LEGAL file in repository */

package protocol

type ID uint64

type UUIDHash interface {
	UUID() [32]byte // Hash of a record data
	ID() ID         // first 64bit of UUID

	IDasString() string // Base64 of ID

	Stringer // Base64 of UUID
}

type UUID interface {
	UUID() [16]byte
	ExistenceTime() Time
	ID() [4]byte

	Stringer // Base64 of UUID
}
