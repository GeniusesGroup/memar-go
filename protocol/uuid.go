/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type ID uint64

type UUID [16]byte

type UUID_Time interface {
	UUID() UUID
	ExistenceTime() Time
}

type UUID_Hash interface {
	UUID() [32]byte // Hash of a record data
	ID() ID         // first 64bit of UUID
}
