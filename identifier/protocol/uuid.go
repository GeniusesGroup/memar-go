/* For license and copyright information please see the LEGAL file in the code repository */

package uuid_p

type UUID_Hash interface {
	UUID() [32]byte // Hash of a record data
	ID() ID         // first 64bit of UUID
}
