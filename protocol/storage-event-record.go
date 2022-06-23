/* For license and copyright information please see LEGAL file in repository */

package protocol

type RecordEvent interface {
	Event

	MediaTypeID() uint64
	ID() [16]byte
	CRUD() CRUD
	VersionOffset() uint64
	Version() Time
}
